package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type User struct{
	userId int "用户ID唯一"
	passwd string "密码"
	balance int "账号余额"
	power int "用户权限，0为普通用户，1为银行职员"
	lock Lock "读写锁，保障并发数据安全"
}

func New(passwd string, userId, balance, power int) User {
	var user User
	user.userId = userId
	user.passwd = passwd
	user.balance = balance
	user.power = power
	return user
}

func NewPointer(passwd string, userId, balance, power int) *User {
	var user User
	user.userId = userId
	user.passwd = passwd
	user.balance = balance
	user.power = power
	return &user
}
func (user *User)GetBalance() int {
	return user.balance
}
func (user *User) IsHavePower() bool {
	if user.power == 1 {
		return true
	} else {
		return false
	}
}

func (user *User) GetUserId() int {
	return user.userId
}

func (user *User) OpenAccount(passwd string, power int) (int, error) {
	if !user.IsHavePower() {
		return -1, errors.New("do not have power to open account")
	} else {
		stmt, err := Db.Prepare("INSERT INTO User (passwd, balance, power) values (?, ?, ?)")
		defer stmt.Close()
		if err != nil {
			log.Println(err)
			return -1, err
		}
		res, err := stmt.Exec(passwd, 0, power)
		if err != nil {
			log.Println(err)
			return -1, err
		}
		userId, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return -2, err
		}
		return int(userId), nil
	}
}

func (source *User) Transfer(destinationId, money int) (bool, error) {
	var balance int
	dLock, err := GlobalBank.LockRead(strconv.Itoa(destinationId))
	dLock.lock.Lock()
	defer dLock.lock.Unlock()
	if err != nil {
		log.Println("获取目的用户锁失败")
		return false, err
	}
	err = Db.QueryRow("Select balance From User Where userId = ?", destinationId).Scan(&balance)
	if err != nil {
		log.Println("no this User")
		return false, err
	}
	destination := &User{
		userId: destinationId,
		balance: balance,
		lock: dLock,
	}
	var SourceBalance int
	err = Db.QueryRow("Select balance From User Where userId = ?", source.GetUserId()).Scan(&SourceBalance)
	source.balance = SourceBalance
	if (source.balance - money) >= 0 {
		source.balance -= money
		destination.balance += money
		tx, err := Db.Begin()
		if err != nil {
			log.Println("开启事务失败")
			return false, err
		}
		stmt, err := tx.Prepare("Update User Set balance = ? Where userId = ?")
		if err != nil {
			log.Println("创建预处理语句失败")
			tx.Rollback()
			return false, err
		}
		_, err = stmt.Exec(source.balance, source.GetUserId())
		if err != nil {
			log.Println("修改源用户余额失败")
			tx.Rollback()
			return false, err
		}
		_, err = stmt.Exec(destination.balance, destination.GetUserId())
		if err != nil {
			log.Println("修改目的用户余额失败")
			tx.Rollback()
			return false, err
		}
		err = tx.Commit()
		if err != nil {
			log.Println("事务提交失败")
			tx.Rollback()
			return false, err
		}
		return true, nil
	}
	log.Println("源用户余额不足")
	return false, errors.New("源用户余额不足")
}

func (user *User) Balance() error {
	userLock, err := GlobalBank.LockRead(strconv.Itoa(user.GetUserId()))
	if err != nil {
		log.Println(err)
		return errors.New("获取用户锁失败")
	}
	user.lock = userLock
	user.lock.lock.RLock()
	defer user.lock.lock.RUnlock()
	var balance int
	err = Db.QueryRow("Select balance From User Where userId = ?", user.GetUserId()).Scan(&balance)
	if err != nil {
		log.Println("no this user or get user balance error")
		return err
	}
	user.balance = balance
	return nil
}

func (user *User) DeleteAccount(destinationId string) error {
	if user.power == 0 {
		log.Println("not have enough power")
		return errors.New("not have enough power")
	}
	dLock, err := GlobalBank.LockRead(destinationId)
	if err != nil {
		log.Println("获取用户锁失败")
		return err
	}
	dLock.lock.Lock()
	defer dLock.lock.Unlock()
	tx, err := Db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	stmt, err := tx.Prepare("Delete From User Where userId = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	userId, _ := strconv.Atoi(destinationId)
	_, err = stmt.Exec(userId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	return nil
}
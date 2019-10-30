package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct{
	userId int "用户ID唯一"
	UserName string "用户名"
	passwd string "密码"
	balance int "账号余额"
	power int "用户权限，0为普通用户，1为银行职员"
	lock Lock "读写锁，保障并发数据安全"
}

func New(UserName, passwd string, userId, balance, power int) User{
	var user User
	user.userId = userId
	user.UserName = UserName
	user.passwd = passwd
	user.balance = balance
	user.power = power
	return user
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
func (user *User) OpenAccount(username, passwd string, power int) (int64, error) {
	if !user.IsHavePower() {
		return -1, errors.New("do not have power to open account")
	} else {
		stmt, err := Db.Prepare("INSERT INTO User (UserName, passwd, balance, power) values (?, ?, ?, ?)")
		if err != nil {
			log.Println(err)
			return -1, err
		}
		res,err := stmt.Exec(username, passwd, 0, power)
		if err != nil {
			log.Println(err)
			return -1, err
		}
		userId, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return -2, err
		}
		return userId, nil
	}
} 
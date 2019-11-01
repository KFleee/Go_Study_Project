package main

import (
	"net/http"
	"log"
	"strconv"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func OpenAccount(w http.ResponseWriter, r *http.Request){
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
    }
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	UserName := formData["UserName"].(string)
	passwd := formData["Passwd"].(string)
	Power := formData["Power"].(string)
	OperatorId := formData["OperatorId"].(string)
	operatorId, _ := strconv.Atoi(OperatorId)
	power, _ := strconv.Atoi(Power)
	rows, err := Db.Query("SELECT userId, power FROM User Where userId = ?", operatorId)
	defer rows.Close()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
	}
	var operator User
	for rows.Next() {
		var o_userId, o_power int
		_ = rows.Scan(&o_userId, &o_power)
		operator = New("", "", o_userId, 0, o_power)
	}
	userId, err := operator.OpenAccount(UserName, passwd, power)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
	}
	w.Write([]byte("open account success"))
	log.Println("new User Id = ", userId)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
    }
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	SourceId := formData["SourceId"].(string)
	DestinationId := formData["DestinationId"].(string)
	Money := formData["Money"].(string)
	SourceUserId, _ := strconv.Atoi(SourceId)
	DestinationUserId, _ := strconv.Atoi(DestinationId)
	money, _ := strconv.Atoi(Money)
	var SourceUserName string
	err = Db.QueryRow("Select UserName From User Where userId = ?", SourceUserId).Scan(&SourceUserName)
	if err != nil {
		w.Write([]byte("no this source User"))
		log.Panicln("no this User")
	}
	source := NewPointer(SourceUserName, "", SourceUserId, 0, 0)
	sLock, err := GlobalBank.LockRead(SourceId)
	if err != nil {
		w.Write([]byte("获取源用户锁失败"))
		log.Panicln("获取源用户锁失败")
	}
	source.lock = sLock
	source.lock.lock.Lock()
	defer source.lock.lock.Unlock()
	if ok, err := source.Transfer(DestinationUserId, money); !ok {
		w.Write([]byte("transfer money erro"))
		log.Panicln(err)
	}
	w.Write([]byte("transfer money success"))
}

func Balance(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
    }
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	UserId := formData["UserId"].(string)
	userId, _ := strconv.Atoi(UserId)
	user := NewPointer("", "", userId, 0, 0)
	err = user.Balance()
	if err != nil {
		w.Write([]byte("获取用户余额失败"))
		log.Panicln(err)
	}
	w.Write([]byte("User balance = " + strconv.Itoa(user.GetBalance())))
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
    }
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	OperatorIdString := formData["OperatorId"].(string)
	DestinationIdString := formData["DestinationId"].(string)
	OperatorId, _ := strconv.Atoi(OperatorIdString)
	var power int
	err = Db.QueryRow("Select power From User Where userId = ?", OperatorId).Scan(&power)
	if err != nil {
		w.Write([]byte("该银行员工不存在"))
		log.Panicln("该银行员工不存在")
	}
	Operator := NewPointer("", "", OperatorId, 0, power)
	err = Operator.DeleteAccount(DestinationIdString)
	if err != nil {
		w.Write([]byte("删除用户账户失败"))
		log.Panicln(err)
	}
	w.Write([]byte("Delete User Account Success"))
}
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
	log.Println("new User Id = ", userId)
	w.Write([]byte("open account success"))
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
	tsource := New(SourceUserName, "", SourceUserId, 0, 0)
	sLock, err := GlobalBank.LockRead(SourceId)
	if err != nil {
		log.Panicln("获取源用户锁失败")
		w.Write([]byte("获取源用户锁失败"))
	}
	tsource.lock = sLock
	source := &tsource
	source.lock.lock.Lock()
	defer source.lock.lock.Unlock()
	if ok, err := source.Transfer(DestinationUserId, money); !ok {
		w.Write([]byte("transfer money erro"))
		log.Panicln(err)
	}
	w.Write([]byte("transfer money success"))
}
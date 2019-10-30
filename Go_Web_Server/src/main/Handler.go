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
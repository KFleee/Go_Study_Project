package main

import (
	"net/http"
	"log"
	"strconv"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)
func Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			GlobalSessions.SessionDestroy(w, r)
			log.Println("Destroy Session")
			return
		}
	}()
	logUrl(r, "Login")
	sess, err := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if err != nil {
		log.Println("Open Session error", err)
	} else {
		userId := sess.Get("userId")
		power := sess.Get("power")
		if userId == nil || power == nil {
			log.Println("need passwd to login")
		} else {
			w.Write([]byte(strconv.Itoa(power.(int))))
			return
		}
	}
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	UserId := formData["UserId"].(string)
	passwd := formData["Passwd"].(string)
	log.Println(UserId, passwd)
	var userId, power int
	err = Db.QueryRow("Select userId, power From User Where userId = ? and passwd = ?", UserId, passwd).Scan(&userId, &power)
	if err != nil {
		w.Write([]byte("no this User or UserId and Passwd error"))
		log.Panicln(err)
	}
	err = sess.Set("userId", userId)
	if err != nil {
		w.Write([]byte("Login failure"))
		log.Panicln("Insert userId error")
	}
	err = sess.Set("power", power)
	if err != nil {
		w.Write([]byte("Login failure"))
		log.Panicln("Insert power error")
	}
	w.Write([]byte(strconv.Itoa(power)))
}

func Logout(w http.ResponseWriter, r *http.Request){
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	logUrl(r, "Logout")
	GlobalSessions.SessionDestroy(w, r)
	w.Write([]byte("Logout success"))
}

func OpenAccount(w http.ResponseWriter, r *http.Request){
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	logUrl(r, "OpenAccount")
	sess, err := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if err != nil {
		w.Write([]byte("need Login again"))
		log.Panicln("Open Session error", err)
	}
	operatorId := sess.Get("userId")
	if operatorId == nil {
		w.Write([]byte("need Login again"))
		GlobalSessions.SessionDestroy(w, r)
		log.Panicln("get session context error")
	}
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	passwd := formData["Passwd"].(string)
	Power := formData["Power"].(string)
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
		operator = New("", o_userId, 0, o_power)
	}
	userId, err := operator.OpenAccount(passwd, power)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Panicln(err)
	}
	w.Write([]byte("open account success userId = " + strconv.Itoa(userId)))
	log.Println("new User Id = ", userId)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	logUrl(r, "Transfer")
	sess, err := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if err != nil {
		w.Write([]byte("need Login again"))
		log.Panicln("Open Session error", err)
	}
	SourceUserId := sess.Get("userId")
	if SourceUserId == nil {
		w.Write([]byte("need Login again"))
		GlobalSessions.SessionDestroy(w, r)
		log.Panicln("get session context error")
	}
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	DestinationId := formData["DestinationId"].(string)
	Money := formData["Money"].(string)
	DestinationUserId, _ := strconv.Atoi(DestinationId)
	money, _ := strconv.Atoi(Money)
	if DestinationUserId == SourceUserId.(int) {
		w.Write([]byte("can not transfer money to myself"))
		log.Panicln("can not transfer money to myself")
	}
	var tempUserId int
	err = Db.QueryRow("Select userId From User Where userId = ?", SourceUserId).Scan(&tempUserId)
	if err != nil {
		w.Write([]byte("no this source User"))
		log.Panicln("no this User")
	}
	source := NewPointer("", SourceUserId.(int), 0, 0)
	sLock, err := GlobalBank.LockRead(strconv.Itoa(SourceUserId.(int)))
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
	logUrl(r, "Balance")
	sess, err := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if err != nil {
		w.Write([]byte("need Login again"))
		log.Panicln("Open Session error", err)
	}
	userId := sess.Get("userId")
	if userId == nil {
		w.Write([]byte("need Login again"))
		GlobalSessions.SessionDestroy(w, r)
		log.Panicln("get session context error")
	}
	user := NewPointer("", userId.(int), 0, 0)
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
	logUrl(r, "DeleteAccount")
	sess, err := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if err != nil {
		w.Write([]byte("need Login again"))
		log.Panicln("Open Session error", err)
	}
	OperatorId := sess.Get("userId")
	if OperatorId == nil {
		w.Write([]byte("need Login again"))
		GlobalSessions.SessionDestroy(w, r)
		log.Panicln("get session context error")
	}
	formData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&formData)
	DestinationIdString := formData["DestinationId"].(string)
	var power int
	err = Db.QueryRow("Select power From User Where userId = ?", OperatorId).Scan(&power)
	if err != nil {
		w.Write([]byte("该银行员工不存在"))
		log.Panicln("该银行员工不存在")
	}
	Operator := NewPointer("", OperatorId.(int), 0, power)
	err = Operator.DeleteAccount(DestinationIdString)
	if err != nil {
		w.Write([]byte("删除用户账户失败"))
		log.Panicln(err)
	}
	w.Write([]byte("Delete User Account Success"))
}

func logUrl(r *http.Request, methodName string) {
	requestMethod := r.Method
	requestUrl := r.RemoteAddr
	log.Println("requestMethod = " + methodName + " : " + "Method = " + requestMethod + "----------" + " From " + requestUrl)
}
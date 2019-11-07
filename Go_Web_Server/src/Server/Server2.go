package main

import (
	"net/http"
	 "github.com/astaxie/beego/session"
	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"
	 "log"
)

var GlobalSessions *session.Manager
var Db *sql.DB
var GlobalBank *Bank
const(
	MaxLifeTime int64 = 120
	GcLifeTime int64 = 120
)
func init() {
    sessionConfig := &session.ManagerConfig{
	    CookieName:"gosessionid", 
	    EnableSetCookie: true, 
	    Gclifetime:3600,
	    Maxlifetime: 3600, 
	    Secure: false,
	    CookieLifeTime: 3600,
	    ProviderConfig: "./tmp",
    }
    GlobalSessions, _ = session.NewManager("memory", sessionConfig)
    GlobalBank = &Bank{
    	userLock: make(map[string]Lock),
    	MaxLifeTime: MaxLifeTime,
    	GcLifeTime: GcLifeTime,
    }
    var err error
    Db, err = sql.Open("mysql", "root:LJH787807080886@/Bank")
    if err != nil {
    	log.Fatalln(err)
    }
    go GlobalSessions.GC()
    go GlobalBank.Gc()
}

func main() {
	log.Println("Server Start....")
	mux := http.NewServeMux()
	mux.HandleFunc("/Login", Login)
	mux.HandleFunc("/Logout", Logout)
	mux.HandleFunc("/OpenAccount", OpenAccount)
	mux.HandleFunc("/Transfer", Transfer)
	mux.HandleFunc("/Balance", Balance)
	mux.HandleFunc("/DeleteAccount", DeleteAccount)
	http.ListenAndServe("0.0.0.0:1234", mux)
}
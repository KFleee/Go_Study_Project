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
	MaxLifeTime int64 = 1800
	GcLifeTime int64 = 300
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

func login(w http.ResponseWriter, r *http.Request) {
    sess, _ := GlobalSessions.SessionStart(w, r)
    defer sess.SessionRelease(w)
    username := sess.Get("username")
    if username == nil {
    	sess.Set("username", "saljdhasjkdha")
    	w.Write([]byte("no session"))
    	return
    }else {
    	w.Write([]byte(username.(string)))
    	return
    }
}

func main() {
	http.HandleFunc("/", OpenAccount)
	http.ListenAndServe(":1234", nil)
}
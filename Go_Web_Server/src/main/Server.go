package main

import (
//	"log"
//	"io"
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
	"encoding/json"
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
)

type A struct {
	Foo string
	Id int
}
var store = sessions.NewCookieStore([]byte("session-test"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
    // Get a session. Get() always returns a session, even if empty.
    session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set some session values.
    session.Values["foo"] = "bssar"
    session.Values[42] = 53
    // Save it before we write to the response/return from the handler.
    err = session.Save(r, w)
    w.Write([]byte("hello session"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
//	temp := struct{
//		foo string
//		id int
//	} {
//		session.Values["foo"].(string),
//		session.Values[42].(int),
//	}
	var temp, d A
	temp.Foo = session.Values["foo"].(string)
	temp.Id = session.Values[42].(int)
//	temp.Foo = "gjhg"
//	temp.Id = 6
	respone, _ := json.Marshal(temp)
	_ = json.Unmarshal(respone, &d)
	fmt.Println(temp)
	fmt.Println(d)
	w.Header().Set("Content-Type","application/json")
	w.Write(respone)
}
//func main() {
//	http.HandleFunc("/", MyHandler)
//	http.HandleFunc("/getSession", GetHandler)
//    http.ListenAndServe(":12345", nil)
//}
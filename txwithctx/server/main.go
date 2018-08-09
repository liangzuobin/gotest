package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const (
	driver = "mysql"
	dsn    = "user:pwd@tcp(localdb:3306)/database" // your dsn here

	userID int64 = 209008
)

var conn *dbr.Connection

func main() {
	var err error
	conn, err = dbr.Open(driver, dsn, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	handler := http.NewServeMux()
	handler.HandleFunc("/tx", wrapper(foo))
	handler.HandleFunc("/txpanic", wrapper(bar))
	handler.HandleFunc("/txwithctxpanic", wrapper(baz))

	svr := &http.Server{
		Addr:    ":3336",
		Handler: handler,
	}

	go func() {
		log.Fatal(svr.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)
	<-ch
	defer close(ch)

	if svr != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := svr.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
		log.Println("server shutdown gracefully")
	}
}

func wrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered from panic, r: %v", r)
			}
		}()
		fn(w, r)
	}
}

// try to lock
func foo(w http.ResponseWriter, _ *http.Request) {
	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tx.RollbackUnlessCommitted()

	id, err := tx.SelectBySql("select id from user where id = ? for update", userID).ReturnInt64()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id != userID {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"ok"}`)
}

// panic after lock
func bar(w http.ResponseWriter, _ *http.Request) {
	log.Println("bar")
	defer func() {
		log.Println("bar done")
	}()

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// no defer rollback

	_, err = tx.SelectBySql("select id from user where id = ? for update", userID).ReturnInt64()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	panic(errors.New("panic here"))
}

// panic after lock, but tx has ctx
func baz(w http.ResponseWriter, _ *http.Request) {
	log.Println("baz")
	defer func() {
		log.Println("baz done")
	}()

	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess := conn.NewSession(nil)
	tx, err := sess.BeginTx(ctx, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// no defer rollback

	_, err = tx.SelectBySql("select id from user where id = ? for update", userID).ReturnInt64()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	panic(errors.New("panic here"))
}

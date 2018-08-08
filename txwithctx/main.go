package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driver = "mysql"
	dsn    = "zhimaa:zhimaa@tcp(localdb:3306)/fan"
)

var db *sql.DB

func init() {
}

func main() {
	var err error
	db, err = sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(40)

	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// baz()

	// panic without ctx
	// tx, err := db.Begin()
	// if err != nil {
	// 	panic(err)
	// }

	// r := tx.QueryRow("select id, username from user where id = ? for update", 209008)
	// var (
	// 	id       uint64
	// 	username string
	// )
	// if err = r.Scan(&id, &username); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("user = %v, %v", id, username)

	// // time.Sleep(20 * time.Second)

	// // panic here
	// panic(errors.New("panic here"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	// _ = tx.QueryRow("select id from user where id = ? for update", 209008)

	// _, err = tx.Exec("update user set balance = balance+1 where id = ?", 209008)
	// if err != nil {
	// 	panic(err)
	// }

	r := tx.QueryRow("select balance from user where id = ? for update", 209008)
	var b float32
	if err = r.Scan(&b); err != nil {
		panic(err)
	}
	fmt.Printf("balance = %v", b)

	// if err := tx.Commit(); err != nil {
	// 	panic(err)
	// }

	time.Sleep(15 * time.Second)

	if err := tx.Commit(); err != nil {
		fmt.Printf("commit failed, err: %v", err)
	}

	// panic(errors.New("panic here"))

}

func foo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	r := tx.QueryRow("select id, username from user where id = ? for update", 209008)
	var id uint64
	var username string
	if err = r.Scan(&id, &username); err != nil {
		panic(err)
	}
	fmt.Printf("user = %v, %v", id, username)

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

func bar() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	r := tx.QueryRow("select id, username from user where id = ? for update", 209008)
	var id uint64
	var username string
	if err = r.Scan(&id, &username); err != nil {
		panic(err)
	}
	fmt.Printf("user = %v, %v", id, username)

	// panic here
	// panic(errors.New("panic here"))

	// no panic but not commit
}

func baz() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover from panic")
			}
		}()

		bar()

	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		timer := time.Now()
		defer func() {
			fmt.Printf("time cost %d", time.Since(timer)/time.Second)
		}()

		// delaV
		time.Sleep(3 * time.Second)

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover from panic")
			}
		}()

		foo()
	}(&wg)

	wg.Wait()
}

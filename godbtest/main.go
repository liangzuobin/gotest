package main

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	// mysql driver
	"github.com/gocraft/dbr"
)

const (
	driver = "mysql"
	dsn    = "you:me@tcp(127.0.0.1:3306)/amores?parseTime=true"
)

type nested struct {
	Key   string `db:"some_key" json:"some_key"`
	Value uint   `db:"some_value" json:"some_value"`
}

type outter struct {
	ID    uint   `db:"id"`
	Name  string `db:"name"`
	Extra nested `db:"extra"`
}

func main() {
	conn, err := dbr.Open(driver, dsn, nil)
	if err != nil {
		log.Fatal(err)
	}
	sess := conn.NewSession(nil)
	foo(sess)
	// bar(sess)
}

func foo(sess *dbr.Session) {
	{
		o := outter{
			Name: "name_1",
			Extra: nested{
				Key:   "key_1",
				Value: 1,
			},
		}
		if _, err := sess.InsertInto("gostruct").Columns("name", "extra").Record(&o).Exec(); err != nil {
			log.Println("try insert")
			log.Fatal(err)
		}
	}
	{
		o := outter{}
		if err := sess.Select("*").From("gostruct").Where("id = 1").LoadStruct(&o); err != nil {
			log.Println("try select")
			log.Fatal(err)
		}
	}
	log.Println("done.")
}

func bar(sess *dbr.Session) {
	{
		o := outter{
			Name: "name_1",
			Extra: nested{
				Key:   "key_1",
				Value: 1,
			},
		}
		if _, err := sess.InsertInto("gostructplain").Columns("name", "some_key", "some_value").Record(&o).Exec(); err != nil {
			log.Println("try insert")
			log.Fatal(err)
		}
	}
	{
		o := outter{}
		if err := sess.Select("*").From("gostructplain").Where("id = 1").LoadStruct(&o); err != nil {
			log.Println("try select")
			log.Fatal(err)
		} else {
			log.Printf("outer.Extra = %v \n", o.Extra)
		}
	}
	{
		str := `{"some_key":"key_1","some_value":1}`
		n := nested{}
		if err := json.Unmarshal([]byte(str), &n); err != nil {
			log.Println("try unmarshal")
			log.Fatal(err)
		} else {
			log.Printf("n = %v", n)
		}

		o := outter{}
		if err := sess.Select("*").From("gostruct").Where("id = 1").LoadStruct(&o); err != nil {
			log.Println("try select")
			log.Fatal(err)
		} else {
			log.Printf("outer.Extra = %v \n", o.Extra)
		}
	}
	log.Println("done.")
}

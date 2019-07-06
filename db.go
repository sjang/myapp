package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Dept2 struct {
	DeptNo   int64  `db:"dept_no"`
	DeptCode string `db:"dept_code"`
	DeptName string `db:"dept_name"`
}

func main() {
	db, err := sqlx.Connect("mysql", "testuser:testuserpw@(localhost:3306)/testuser")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var dept2Args [2]string
	dept2Args[0] = "code1"

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO dept2 (dept_code, dept_name) VALUES (?, ?)", "code1", "name1")
	tx.MustExec("INSERT INTO dept2 (dept_code, dept_name) VALUES (?, ?)", "code2", "name2")
	tx.Commit()

	d := Dept2{DeptCode: "struct_code", DeptName: "struct_name"}
	r, e := db.NamedExec("INSERT INTO dept2 (dept_code, dept_name) VALUES (:dept_code, :dept_name)", d)

	log.Print(r)
	log.Print(e)

	dept2 := []Dept2{}
	db.Select(&dept2, "SELECT * FROM dept2 ORDER BY dept_no DESC")

	log.Print(dept2)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return
}

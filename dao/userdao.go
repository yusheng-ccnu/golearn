package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"webdemo/model"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@/goweb")
	checkErr(err)
	return db, err
}

func close(db *sql.DB) {
	db.Close()
}

func Save(u model.User) int64 {
	db, err := connect()

	stmt, err := db.Prepare("insert user set name=?,password=?,age=?,sex=?")
	checkErr(err)

	res, err := stmt.Exec(u.Name, u.Password, u.Age, u.Sex)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

func Update(u model.User) int64 {
	db, err := connect()
	checkErr(err)

	stmt, err := db.Prepare("update user set name=?,password=?,age=?,sex=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(u.Name, u.Password, u.Age, u.Sex, u.Id)
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	return id
}

func Delete(uid int64) int64 {
	db, err := connect()
	checkErr(err)

	stmt, err := db.Prepare("delete from user where id=?")
	checkErr(err)

	res, err := stmt.Exec(uid)
	checkErr(err)

	id, err := res.RowsAffected()

	return id
}

func Query(id int64) model.User {
	db, err := connect()

	rows, err := db.Query("select * from user where id=?", id)
	checkErr(err)

	var u model.User
	for rows.Next() {
		var id int
		var name string
		var password string
		var age int
		var sex int

		rows.Scan(&id, &name, &password, &age, &sex)

		u.Name = name
		u.Password = password
		u.Age = age
		if sex == 1 {
			u.Sex = true
		} else {
			u.Sex = false
		}
		u.Id = id
	}
	return u

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DBconf struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

func InitDB(d DBconf) error {
	var err error
	dsn := `%s:%s@tcp(%s:3306)/%s`
	db, err = sql.Open("mysql", fmt.Sprintf(dsn, d.User, d.Password, d.Host, d.DB))
	if err != nil {
		fmt.Printf("Failed to open connection to mysql, error:%v\n", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to mysql, error:%v\n", err)
		return err
	}
	fmt.Println("connected to mysql")
	db.SetMaxOpenConns(1000)
	return nil
}

func DBConn() *sql.DB {
	return db
}

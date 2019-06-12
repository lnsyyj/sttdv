package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MariaDBInfo struct {
	MariaHostIP			string
	MariaPort			string
	MariaDatabase		string
	MariaTableName		string
	MariaUserName		string
	MariaUserPassword	string
}

func ConnectionMariadb(mariaDBInfo *MariaDBInfo) *sql.DB {
	//   db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	db, err := sql.Open("mysql", mariaDBInfo.MariaUserName+":"+mariaDBInfo.MariaUserPassword+"@"+"tcp"+"("+mariaDBInfo.MariaHostIP+":"+mariaDBInfo.MariaPort+")"+"/"+mariaDBInfo.MariaDatabase)
	if err != nil {
		log.Println("[ERROR] [ConnectionMariadb] : ", err)
		return nil
	}
	return db
}

func CloseConnectionMariadb(db *sql.DB) bool {
	err := db.Close()
	if err != nil {
		log.Println("[ERROR] [CloseConnectionMariadb] : ", err)
		return false
	}
	return true
}

func InsertMariadb(db *sql.DB, sqlStatement string) {
	stmtIns, err := db.Prepare(sqlStatement) // ? = placeholder
	if err != nil {
		log.Println("[ERROR] [InsertMariadb] : ", err)
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	ret, err := stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ret)
}
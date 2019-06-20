package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDBInfo struct {
	MariaHostIP			string
	MariaPort			string
	MariaDatabaseName	string
	MariaTableName		string
	MariaUserName		string
	MariaUserPassword	string
}

func (mi *MariaDBInfo) CheckParameter() bool {
	if mi.MariaHostIP == "" {
		fmt.Println("MariaHostIP parameter is empty ")
		return false
	} else if mi.MariaPort == "" {
		fmt.Println("MariaPort parameter is empty ")
		return false
	} else if mi.MariaDatabaseName == "" {
		fmt.Println("MariaDatabaseName parameter is empty ")
		return false
	} else if mi.MariaTableName == "" {
		fmt.Println("MariaTableName parameter is empty ")
		return false
	} else if mi.MariaUserName == "" {
		fmt.Println("MariaUserName parameter is empty ")
		return false
	} else if mi.MariaUserPassword == "" {
		fmt.Println("MariaUserPassword parameter is empty ")
		return false
	} else {
		return true
	}
}

func ConnectionMariadb(mariaDBInfo *MariaDBInfo) *sql.DB {
	//   db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	db, err := sql.Open("mysql", mariaDBInfo.MariaUserName+":"+mariaDBInfo.MariaUserPassword+"@"+"tcp"+"("+mariaDBInfo.MariaHostIP+":"+mariaDBInfo.MariaPort+")"+"/"+mariaDBInfo.MariaDatabaseName)
	if err != nil {
		panic("[ERROR] [ConnectionMariadb] [Open] : " + err.Error())
	}
	return db
}

func CloseConnectionMariadb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic("[ERROR] [CloseConnectionMariadb] [Close] : " + err.Error())
	}
}

func InsertMariadb(db *sql.DB, sqlStatement string) {
	stmtIns, err := db.Prepare(sqlStatement) // ? = placeholder
	if err != nil {
		fmt.Println(sqlStatement)
		panic("[ERROR] [InsertMariadb] [Prepare] : " + err.Error())
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec()
	if err != nil {
		panic("[ERROR] [InsertMariadb] [Exec] : " + err.Error())
	}
}
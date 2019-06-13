package dbs

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDBInfo struct {
	MariaHostIP			string
	MariaPort			string
	MariaDatabase		string
	MariaTableName		string
	MariaUserName		string
	MariaUserPassword	string
}

func (mdbi *MariaDBInfo) CheckParameterValid() {
	if mdbi.MariaHostIP == "" {
		panic("[ERROR] [CheckParameterValid] : MariaHostIP is null")
	}
}

func ConnectionMariadb(mariaDBInfo *MariaDBInfo) *sql.DB {
	//   db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	db, err := sql.Open("mysql", mariaDBInfo.MariaUserName+":"+mariaDBInfo.MariaUserPassword+"@"+"tcp"+"("+mariaDBInfo.MariaHostIP+":"+mariaDBInfo.MariaPort+")"+"/"+mariaDBInfo.MariaDatabase)
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
		panic("[ERROR] [InsertMariadb] [Prepare] : " + err.Error())
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec()
	if err != nil {
		panic("[ERROR] [InsertMariadb] [Exec] : " + err.Error())
	}
}
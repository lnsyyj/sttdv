package main

import (
	"flag"
	"fmt"
	"github.com/lnsyyj/sttdv/analysis"
	. "github.com/lnsyyj/sttdv/commons"
	"github.com/lnsyyj/sttdv/dbs"
)

func main() {

	var (
		visualizationType = flag.String("visualizationType", "", "Currently supports parsing: vdbench")
		logPath	= flag.String("logPath", "", "Specify the log file to be parsed, absolute or relative path")
		outputInterval	= flag.Int("outputinterval", -1, "Specify vdbench data interval")
		mariaDBHostIP = flag.String("mariaDBHostIP", "", "Specify mariaDB IP address")
		mariaDBPort = flag.String("mariaDBPort", "3306", "Specify mariaDB port")
		mariaDBDatabase = flag.String("mariaDBDatabase", "", "Specify the mariaDB database name")
		mariaDBTableName = flag.String("mariaDBTableName", "", "Specify mariaDB table name")
		mariaDBUserName = flag.String("mariaDBUserName", "", "Specify mariaDB username")
		mariaDBUserPassword = flag.String("mariaDBUserPassword", "", "Specify mariaDB password")
		TestCase = flag.String("TestCase", "", "Specify TestCaseName")
		ClientNumber = flag.String("ClientNumber", "", "Specify TestCaseName")
	)
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}
	fmt.Println(visualizationType, logPath, outputInterval, mariaDBDatabase, mariaDBHostIP, mariaDBPort, mariaDBTableName, mariaDBUserName, mariaDBUserPassword, TestCase, ClientNumber)
	mariaDBInfo := dbs.MariaDBInfo{}					// MariaDB Info
	summaryFileSystemCombination := analysis.SummaryFileSystemCombination{}	// Data Combination

	mariaDBInfo.MariaHostIP = *mariaDBHostIP
	mariaDBInfo.MariaPort = *mariaDBPort
	mariaDBInfo.MariaDatabase = *mariaDBDatabase
	mariaDBInfo.MariaTableName = *mariaDBTableName
	mariaDBInfo.MariaUserName = *mariaDBUserName
	mariaDBInfo.MariaUserPassword = *mariaDBUserPassword
	summaryFileSystemCombination.TestCase = *TestCase
	summaryFileSystemCombination.ClientNumber = *ClientNumber

	fmt.Println(visualizationType, logPath, outputInterval, mariaDBInfo.MariaHostIP, mariaDBInfo.MariaDatabase, mariaDBInfo.MariaTableName, mariaDBInfo.MariaUserName, mariaDBInfo.MariaUserPassword)

	CheckParameter(&mariaDBInfo)
	CheckParameter(&summaryFileSystemCombination)
	InitDateTimeInterval(logPath, &summaryFileSystemCombination, outputInterval)
	InitData(logPath, &summaryFileSystemCombination)
	analysis.AssemblingTime(&summaryFileSystemCombination)

	db := dbs.ConnectionMariadb(&mariaDBInfo)
	analysis.InsertFilesystemData(db, &mariaDBInfo, &summaryFileSystemCombination)
	dbs.CloseConnectionMariadb(db)
}
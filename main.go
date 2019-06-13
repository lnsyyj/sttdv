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
		startData	= flag.String("startData", "", "Specify test start date")
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
	fmt.Println(visualizationType, logPath, startData, outputInterval, mariaDBDatabase, mariaDBHostIP, mariaDBPort, mariaDBTableName, mariaDBUserName, mariaDBUserPassword, TestCase, ClientNumber)
	mariaDBInfo := dbs.MariaDBInfo{}					// MariaDB Info
	summaryFileSystemCombination := analysis.SummaryFileSystemCombination{}	// Data Combination

	mariaDBInfo.MariaHostIP = *mariaDBHostIP
	//mariaDBInfo.MariaPort =
	//mariaDBInfo.MariaDatabase =
	//mariaDBInfo.MariaTableName =
	//mariaDBInfo.MariaUserName =
	//mariaDBInfo.MariaUserPassword =
	//summaryFileSystemCombination.TestCase =
	//summaryFileSystemCombination.ClientNumber =

	fmt.Println(visualizationType, logPath, startData, outputInterval, mariaDBInfo.MariaHostIP, mariaDBInfo.MariaDatabase, mariaDBInfo.MariaTableName, mariaDBInfo.MariaUserName, mariaDBInfo.MariaUserPassword)
	// ******************************Test
	//*logPath = "D:\\SourceCode\\GitHub\\Golang\\src\\github.com\\lnsyyj\\sttdv\\863.log"
	////*logPath = "E:\\summary.html"
	//mariaDBInfo.MariaHostIP = "10.121.9.23"
	//mariaDBInfo.MariaPort = "3306"
	//mariaDBInfo.MariaDatabase = "cephtest"
	//mariaDBInfo.MariaTableName = "vdbench_filesystem"
	//mariaDBInfo.MariaUserName = "root"
	//mariaDBInfo.MariaUserPassword = "1234567890"
	//*outputInterval = 1
	//summaryFileSystemCombination.TestCase = "yu"
	//summaryFileSystemCombination.ClientNumber = "jiang"
	// ******************************Test

	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}
	CheckParameter(&mariaDBInfo)
	CheckParameter(&summaryFileSystemCombination)
	InitDateTimeInterval(logPath, &summaryFileSystemCombination, outputInterval)
	InitData(logPath, &summaryFileSystemCombination)
	analysis.AssemblingTime(&summaryFileSystemCombination)
	//fmt.Println(summaryFileSystemCombination)


	db := dbs.ConnectionMariadb(&mariaDBInfo)
	analysis.InsertFilesystemData(db, &mariaDBInfo, &summaryFileSystemCombination)
	dbs.CloseConnectionMariadb(db)
}
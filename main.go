package main

import (
	"flag"
	"fmt"
	"github.com/lnsyyj/sttdv/analysis"
	"github.com/lnsyyj/sttdv/com"
	"github.com/lnsyyj/sttdv/comst"
	"github.com/lnsyyj/sttdv/dbs"
)

func main() {

	var (
		visualizationType = flag.String("visualizationType", "", "Currently supports parsing: VdbenchFSStaticData, VdbenchFSDynamicData")

		mariaDBHostIP = flag.String("mariaDBHostIP", "", "Specify mariaDB IP address")
		mariaDBPort = flag.String("mariaDBPort", "3306", "Specify mariaDB port")
		mariaDBDatabaseName = flag.String("mariaDBDatabaseName", "", "Specify the mariaDB database name")
		mariaDBTableName = flag.String("mariaDBTableName", "", "Specify mariaDB table name")
		mariaDBUserName = flag.String("mariaDBUserName", "", "Specify mariaDB username")
		mariaDBUserPassword = flag.String("mariaDBUserPassword", "", "Specify mariaDB password")

		logPath	= flag.String("logPath", "", "Specify the log file to be parsed, absolute or relative path")
		outputInterval	= flag.String("outputInterval", "-1", "Specify vdbench data interval")
		testCase = flag.String("testCase", "", "Specify TestCaseName")
		clientNumber = flag.String("clientNumber", "", "Specify TestCaseName")
	)
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	// Set MariaDB Info
	mariaDBInfo := dbs.MariaDBInfo{}
	comst.SetMariaDBInfo(&mariaDBInfo, mariaDBHostIP, mariaDBPort, mariaDBDatabaseName, mariaDBTableName, mariaDBUserName, mariaDBUserPassword)
	if com.CheckParameterValid(&mariaDBInfo) == false {
		return
	}

	// Set extra Info
	extraInfo := comst.ExtraInfo{}
	comst.SetExtraInfo(&extraInfo, logPath, outputInterval, testCase, clientNumber)

	switch *visualizationType {
	case "VdbenchFSStaticData":
		// read file
		com.StaticDataProcess(&analysis.SummaryFileSystemInfo{}, &mariaDBInfo, &extraInfo)
		fmt.Println("VdbenchFSStaticData")
	case "VdbenchFSDynamicData":
		fmt.Println("VdbenchFSDynamicData")
	default:
		fmt.Println("Please specify the visualizationType !")
		flag.PrintDefaults()
	}

}
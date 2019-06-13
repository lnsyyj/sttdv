package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lnsyyj/sttdv/analysis"
	. "github.com/lnsyyj/sttdv/commons"
	"github.com/lnsyyj/sttdv/dbs"
	"os"
)

func InitDataTimeInterval(logPath *string, summaryData *analysis.SummaryData, outputInterval *int) {
	file, err := os.Open(*logPath)
	lineInfo := ""
	defer file.Close()
	Check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineInfo = scanner.Text()
		if summaryData.Data != "" && summaryData.Time != "" {
			break
		}
		analysis.AnalysisFirstData(lineInfo, summaryData)
		analysis.AnalysisFirstTime(lineInfo, summaryData)
	}
	summaryData.Conversion = summaryData.Data + " " + summaryData.Time
	summaryData.OutputInterval = *outputInterval
}

func CheckParameter(cp ICheckParameterValid) {
	cp.CheckParameterValid()
}

func main() {

	mariaDBInfo := dbs.MariaDBInfo{}
	summaryInfo := analysis.SummaryFileSystemInfo{}
	summaryData := analysis.SummaryData{}
	summaryFileSystemCombination := analysis.SummaryFileSystemCombination{}

	var (
		visualizationType = flag.String("visualizationType", "", "Currently supports parsing: vdbench")
		logPath	= flag.String("logPath", "", "Specify the log file to be parsed, absolute or relative path")
		startData	= flag.String("startData", "", "Specify test start date")
		outputInterval	= flag.Int("outputinterval", -1, "Specify vdbench data interval")
	)
	mariaDBInfo.MariaHostIP = *flag.String("mariaDBHostIP", "", "Specify mariaDB IP address")
	mariaDBInfo.MariaPort =  *flag.String("mariaDBPort", "3306", "Specify mariaDB port")
	mariaDBInfo.MariaDatabase = *flag.String("mariaDBDatabase", "", "Specify the mariaDB database name")
	mariaDBInfo.MariaTableName = *flag.String("mariaDBTableName", "", "Specify mariaDB table name")
	mariaDBInfo.MariaUserName = *flag.String("mariaDBUserName", "", "Specify mariaDB username")
	mariaDBInfo.MariaUserPassword = *flag.String("mariaDBUserPassword", "", "Specify mariaDB password")
	summaryFileSystemCombination.TestCase = *flag.String("TestCase", "", "Specify TestCaseName")
	summaryFileSystemCombination.ClientNumber = *flag.String("client_number", "", "Specify TestCaseName")
	flag.Parse()

	// Test
	*logPath = "D:\\SourceCode\\GitHub\\Golang\\src\\github.com\\lnsyyj\\sttdv\\863.log"
	//*logPath = "E:\\summary.html"
	mariaDBInfo.MariaHostIP = "10.121.9.23"
	mariaDBInfo.MariaPort = "3306"
	mariaDBInfo.MariaDatabase = "cephtest"
	mariaDBInfo.MariaTableName = "vdbench_filesystem"
	mariaDBInfo.MariaUserName = "root"
	mariaDBInfo.MariaUserPassword = "1234567890"
	*outputInterval = 1
	summaryFileSystemCombination.TestCase = "yu"
	summaryFileSystemCombination.ClientNumber = "jiang"

	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}
	CheckParameter(&mariaDBInfo)
	CheckParameter(&summaryFileSystemCombination)

	fmt.Println(visualizationType, logPath, startData, outputInterval, mariaDBInfo.MariaHostIP, mariaDBInfo.MariaDatabase, mariaDBInfo.MariaTableName, mariaDBInfo.MariaUserName, mariaDBInfo.MariaUserPassword)
	//fmt.Println(visualizationType, logPath, mysqlDatabase, mysqlTableName, mysqlHostIP, mysqlUserName, mysqlUserPassword)

	InitDataTimeInterval(logPath, &summaryData, outputInterval)
	summaryFileSystemCombination.SD = summaryData

	fmt.Println(summaryData)
	file, err := os.Open(*logPath)
	lineInfo := ""
	defer file.Close()
	Check(err)
	scanner := bufio.NewScanner(file)

	//fmt.Println(summaryData.Conversion)

	var i = 1
	var result bool
	for scanner.Scan() {
		lineInfo = scanner.Text()
		//fmt.Println(lineInfo)
		summaryInfo = analysis.AnalysisSummaryInfo(lineInfo)

		if result == false {
			continue
		}
		summaryFileSystemCombination.SFSI = append(summaryFileSystemCombination.SFSI, summaryInfo)
		//fmt.Println(i, summaryFileSystemCombination.SFSI)
		i++
		//fmt.Println(summaryData)
	}
	err = scanner.Err()
	Check(err)
	analysis.AssemblingTime(&summaryFileSystemCombination)
	db := dbs.ConnectionMariadb(&mariaDBInfo)
	analysis.InsertFilesystemData(db, &mariaDBInfo, &summaryFileSystemCombination)
	dbs.CloseConnectionMariadb(db)
}
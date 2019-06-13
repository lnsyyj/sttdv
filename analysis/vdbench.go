package analysis

import (
	"database/sql"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/lnsyyj/sttdv/dbs"
	"regexp"
	"strconv"
	"time"
)

type SummaryFileSystemInfo struct {
	//Id							int64
	DateTime					string
	OutputInterval				int
	ReqstdOpsRate				string
	ReqstdOpsResp				string
	CpuTotal					string
	CpuSys						string
	ReadPct						string
	ReadRate					string
	ReadResp					string
	WriteRate					string
	WriteResp					string
	MbSecRead					string
	MbSecWrite					string
	MbSecTotal					string
	XferSize					string
	MkdirRate					string
	MkdirResp					string
	RmdirRate					string
	RmdirResp					string
	CreateRate					string
	CreateResp					string
	OpenRate					string
	OpenResp					string
	CloseRate					string
	CloseResp					string
	DeleteRate					string
	DeleteResp					string
}

type SummaryData struct {
	Data						string
	Time						string
	Conversion					string
	OutputInterval 				int
}

type SummaryFileSystemCombination struct {
	SD				SummaryData
	SFSI			[]SummaryFileSystemInfo
	OperationTableDate			string
	TestCase					string
	ClientNumber				string
}

func (sfsc *SummaryFileSystemCombination) CheckParameterValid() {
	if sfsc.TestCase == "" {
		panic("[ERROR] [CheckParameterValid] : TestCase is null")
	}
	if sfsc.ClientNumber == "" {
		panic("[ERROR] [CheckParameterValid] : ClientNumber is null")
	}
}

func AnalysisFirstData(lineInfo string, sfsc *SummaryFileSystemCombination) {
	var timeConversion string
	// Jun 08, 2019  interval        i/o   MB/sec   bytes   read     resp     read    write     resp     resp queue  cpu%  cpu%
	re := regexp.MustCompile(`(\w+\s\d+,\s\d+).*`)
	match := re.FindStringSubmatch(lineInfo)
	if len(match) > 1 {
		t, err := dateparse.ParseAny(match[1])
		if err != nil {
			panic("[ERROR] [AnalysisFirstData] : " + err.Error())
		}
		timeConversion = t.String()
	}

	re = regexp.MustCompile(`(\d+-\d+-\d+).*`)
	match = re.FindStringSubmatch(timeConversion)
	if len(match) > 1 {
		sfsc.SD.Data = match[1]
	}
}

func AnalysisFirstTime(lineInfo string, sfsc *SummaryFileSystemCombination) {
	re := regexp.MustCompile(`([0-9][0-9]\:[0-9][0-9]\:[0-9][0-9])\.[0-9]+[\s]+[0-9]+`)
	match := re.FindStringSubmatch(lineInfo)
	if len(match) > 1 {
		sfsc.SD.Time = match[1]
	}
}

func AnalysisResult(regular string, lineInfo string) string {
	re := regexp.MustCompile(regular)
	match := re.FindStringSubmatch(lineInfo)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func AssemblingTime(summaryFileSystemCombination *SummaryFileSystemCombination) {
	t, err := dateparse.ParseLocal(summaryFileSystemCombination.SD.Conversion)
	if err != nil {
		panic("[ERROR] [AssemblingTime] : " + err.Error())
	}
	fmt.Println(t)
	for key, _ := range summaryFileSystemCombination.SFSI {
		m, _ := time.ParseDuration(strconv.Itoa(summaryFileSystemCombination.SD.OutputInterval * key) + "s")
		summaryFileSystemCombination.SFSI[key].DateTime = t.Add(m).Format("2006-01-02 15:04:05")
	}
}

func StringTOInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic("[ERROR] [StringTOInt] : " + err.Error())
	}
	return result
}

func AnalysisSummaryInfo(lineInfo string) SummaryFileSystemInfo {
	summaryInfo := SummaryFileSystemInfo{}
	// 17:17:20.077            1   39.1 143.61  24.5 7.77   0.0    0.0  0.000   39.0 143.61  0.00 39.00  39.00 1048576   0.0  0.000   0.0  0.000   0.0  0.000  18.0  6.705   1.0 77.087   0.0  0.000
	// 	Outputinterval
	outputinterval := `\d+\:\d+\:\d+\.\d+[\s]+(\d+).*`
	result := AnalysisResult(outputinterval, lineInfo)
	if result == "" {
		return summaryInfo
	}
	summaryInfo.OutputInterval =  StringTOInt(result)

	// 	ReqstdOpsRate
	reqstdOpsRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(reqstdOpsRate, lineInfo)
	summaryInfo.ReqstdOpsRate = result

	//	ReqstdOpsResp
	reqstdOpsResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+(\d+.\d+).*`
	result = AnalysisResult(reqstdOpsResp, lineInfo)
	summaryInfo.ReqstdOpsResp = result

	//	CpuTotal
	cpuTotal := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(cpuTotal, lineInfo)
	summaryInfo.CpuTotal = result

	//	CpuSys
	cpuSys := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(cpuSys, lineInfo)
	summaryInfo.CpuSys = result

	//	ReadPct
	readPct := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(readPct, lineInfo)
	summaryInfo.ReadPct = result

	//	ReadRate
	readRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(readRate, lineInfo)
	summaryInfo.ReadRate = result

	//	ReadResp
	readResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(readResp, lineInfo)
	summaryInfo.ReadResp = result

	//	WriteRate
	writeRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(writeRate, lineInfo)
	summaryInfo.WriteRate = result

	// WriteResp
	writeResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(writeResp, lineInfo)
	summaryInfo.WriteResp = result

	//	MbSecRead
	mbSecRead := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(mbSecRead, lineInfo)
	summaryInfo.MbSecRead = result

	//	MbSecWrite
	mbSecWrite := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(mbSecWrite, lineInfo)
	summaryInfo.MbSecWrite = result

	//	MbSecTotal
	mbSecTotal := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(mbSecTotal, lineInfo)
	summaryInfo.MbSecTotal = result

	//	XferSize
	xferSize := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+).*`
	result = AnalysisResult(xferSize, lineInfo)
	summaryInfo.XferSize = result

	//	MkdirRate
	mkdirRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(mkdirRate, lineInfo)
	summaryInfo.MkdirRate = result

	//	MkdirResp
	mkdirResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(mkdirResp, lineInfo)
	summaryInfo.MkdirResp = result

	//	RmdirRate
	rmdirRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(rmdirRate, lineInfo)
	summaryInfo.RmdirRate = result

	//	RmdirResp
	rmdirResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(rmdirResp, lineInfo)
	summaryInfo.RmdirResp = result

	//	CreateRate
	createRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(createRate, lineInfo)
	summaryInfo.CreateRate = result

	//	CreateResp
	createResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(createResp, lineInfo)
	summaryInfo.CreateResp = result

	//	OpenRate
	openRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(openRate, lineInfo)
	summaryInfo.OpenRate = result

	//	OpenResp
	openResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(openResp, lineInfo)
	summaryInfo.OpenResp = result

	//	CloseRate
	closeRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(closeRate, lineInfo)
	summaryInfo.CloseRate = result

	//	CloseResp
	closeResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(closeResp, lineInfo)
	summaryInfo.CloseResp = result

	//	DeleteRate
	deleteRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(deleteRate, lineInfo)
	summaryInfo.DeleteRate = result

	//	DeleteResp
	deleteResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result = AnalysisResult(deleteResp, lineInfo)
	summaryInfo.DeleteResp = result

	return summaryInfo
}

func InsertFilesystemData(db *sql.DB, mariaDBInfo *dbs.MariaDBInfo, sfsc *SummaryFileSystemCombination) {
	var sqlStatement string
	for key, _ := range sfsc.SFSI {
		sqlStatement = "INSERT INTO " + mariaDBInfo.MariaTableName + "(Id, DateTime, OutputInterval, ReqstdOpsRate, ReqstdOpsResp, CpuTotal, CpuSys, ReadPct, ReadRate, ReadResp, WriteRate, WriteResp, MbSecRead, MbSecWrite, MbSecTotal, XferSize, MkdirRate, " +
			"MkdirResp, RmdirRate, RmdirResp, CreateRate, CreateResp, OpenRate, OpenResp, CloseRate, CloseResp, DeleteRate, DeleteResp, OperationTableDate, TestCase, ClientNumber)" + " VALUES " +
			"(" + "NULL, " + "\"" + sfsc.SFSI[key].DateTime+ "\", " + strconv.Itoa(sfsc.SFSI[key].OutputInterval) + ", " +	sfsc.SFSI[key].ReqstdOpsRate + ", " + sfsc.SFSI[key].ReqstdOpsResp + ", " + sfsc.SFSI[key].CpuTotal + ", " + sfsc.SFSI[key].CpuSys +
			", " + sfsc.SFSI[key].ReadPct + ", " + sfsc.SFSI[key].ReadRate + ", " + sfsc.SFSI[key].ReadResp + ", " + sfsc.SFSI[key].WriteRate + ", " + sfsc.SFSI[key].WriteResp + ", " + sfsc.SFSI[key].MbSecRead + ", " + sfsc.SFSI[key].MbSecWrite + ", " +
			sfsc.SFSI[key].MbSecTotal + ", " + sfsc.SFSI[key].XferSize + ", " + sfsc.SFSI[key].MkdirRate + ", " + sfsc.SFSI[key].MkdirResp + ", " +	sfsc.SFSI[key].RmdirRate + ", " + sfsc.SFSI[key].RmdirResp + ", " + sfsc.SFSI[key].CreateRate +
			", " + sfsc.SFSI[key].CreateResp + ", " + sfsc.SFSI[key].OpenRate + ", " + sfsc.SFSI[key].OpenResp + ", " + sfsc.SFSI[key].CloseRate + ", " + sfsc.SFSI[key].CloseResp + ", " +	sfsc.SFSI[key].DeleteRate + ", " + sfsc.SFSI[key].DeleteResp +
			", \"" + time.Now().Format("2006-01-02 15:04:05") + "\", \"" + sfsc.TestCase + "\", \"" +  sfsc.ClientNumber + "\")"
		dbs.InsertMariadb(db, sqlStatement)
	}
}
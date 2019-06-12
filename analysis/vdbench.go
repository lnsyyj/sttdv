package analysis

import (
	"database/sql"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/lnsyyj/SSTDV/DBs"
	"regexp"
	"strconv"
	"time"
)

type SummaryFileSystemInfo struct {
	//Id							int64
	DateTime					string
	OutputInterval				int
	ReqstdOpsRate				float64
	ReqstdOpsResp				float64
	CpuTotal					float64
	CpuSys						float64
	ReadPct						float64
	ReadRate					float64
	ReadResp					float64
	WriteRate					float64
	WriteResp					float64
	MbSecRead					float64
	MbSecWrite					float64
	MbSecTotal					float64
	XferSize					float64
	MkdirRate					float64
	MkdirResp					float64
	RmdirRate					float64
	RmdirResp					float64
	CreateRate					float64
	CreateResp					float64
	OpenRate					float64
	OpenResp					float64
	CloseRate					float64
	CloseResp					float64
	DeleteRate					float64
	DeleteResp					float64
	OperationTableDate			string
	TestCase					string
	ClientNumber				string
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
	TestCase		string
	ClientNumber	string
}

func AnalysisFirstData(lineInfo string, summaryData *SummaryData) {
	var timeConversion string
	// Jun 08, 2019  interval        i/o   MB/sec   bytes   read     resp     read    write     resp     resp queue  cpu%  cpu%
	re := regexp.MustCompile(`(\w+\s\d+,\s\d+).*`)
	match := re.FindStringSubmatch(lineInfo)
	if len(match) > 1 {
		t, err := dateparse.ParseAny(match[1])
		if err != nil {
			panic("[ERROR] [AnalysisDataTime] : ERROR")
		}
		timeConversion = t.String()
	}

	re = regexp.MustCompile(`(\d+-\d+-\d+).*`)
	match = re.FindStringSubmatch(timeConversion)
	if len(match) > 1 {
		summaryData.Data = match[1]
	}
}

func AnalysisFirstTime(lineInfo string, summaryData *SummaryData) {
	re := regexp.MustCompile(`([0-9][0-9]\:[0-9][0-9]\:[0-9][0-9])\.[0-9]+[\s]+[0-9]+`)
	match := re.FindStringSubmatch(lineInfo)
	if len(match) > 1 {
		summaryData.Time = match[1]
	}
}

func AnalysisResult(regular string, lineInfo string) (float64, bool) {
	re := regexp.MustCompile(regular)
	match := re.FindStringSubmatch(lineInfo)
	if len(match) > 1 {
		result,err := strconv.ParseFloat(match[1],64)
		if err != nil {
			panic("[ERROR] [AnalysisSummaryInfo] [Iorate] [strconv.ParseFloat]")
		}
		return result, true
	} else {
		return -1, false
	}
}

func AssemblingTime(summaryFileSystemCombination *SummaryFileSystemCombination) {
	t, err := dateparse.ParseLocal(summaryFileSystemCombination.SD.Conversion)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(t)
	for key, _ := range summaryFileSystemCombination.SFSI {
		m, _ := time.ParseDuration(strconv.Itoa(summaryFileSystemCombination.SD.OutputInterval * key) + "s")
		summaryFileSystemCombination.SFSI[key].DateTime = t.Add(m).Format("2006-01-02 15:04:05")
		//fmt.Println(val.Datetime)
	}
	fmt.Println(summaryFileSystemCombination.SFSI)
}

func AnalysisSummaryInfo(lineInfo string) (SummaryFileSystemInfo, bool) {
	summaryInfo := SummaryFileSystemInfo{}

	// 17:17:20.077            1   39.1 143.61  24.5 7.77   0.0    0.0  0.000   39.0 143.61  0.00 39.00  39.00 1048576   0.0  0.000   0.0  0.000   0.0  0.000  18.0  6.705   1.0 77.087   0.0  0.000
	// 	Outputinterval				int64
	outputinterval := `\d+\:\d+\:\d+\.\d+[\s]+(\d+).*`
	result, err := AnalysisResult(outputinterval, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.OutputInterval = int(result)

	// 	ReqstdOpsRate				float64
	reqstdOpsRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(reqstdOpsRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.ReqstdOpsRate = result

	//	ReqstdOpsResp				float64
	reqstdOpsResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+(\d+.\d+).*`
	result, err = AnalysisResult(reqstdOpsResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.ReqstdOpsResp = result

	//	CpuTotal					float64
	cpuTotal := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(cpuTotal, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.CpuTotal = result

	//	CpuSys						float64
	cpuSys := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(cpuSys, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.CpuSys = result

	//	ReadPct						float64
	readPct := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(readPct, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.ReadPct = result

	//	ReadRate					float64
	readRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(readRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.ReadRate = result

	//	ReadResp					float64
	readResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(readResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.ReadResp = result

	//	WriteRate					float64
	writeRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(writeRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.WriteRate = result

	// WriteResp					float64
	writeResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(writeResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.WriteResp = result
	//	MbSecRead					float64
	mbSecRead := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(mbSecRead, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.MbSecRead = result
	//	MbSecWrite					float64
	mbSecWrite := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(mbSecWrite, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.MbSecWrite = result
	//	MbSecTotal					float64
	mbSecTotal := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(mbSecTotal, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.MbSecTotal = result
	//	XferSize					float64
	xferSize := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+).*`
	result, err = AnalysisResult(xferSize, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.XferSize = result
	//	MkdirRate					float64
	mkdirRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(mkdirRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.MkdirRate = result
	//	MkdirResp					float64
	mkdirResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(mkdirResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.MkdirResp = result
	//	RmdirRate					float64
	rmdirRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(rmdirRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.RmdirRate = result
	//	RmdirResp					float64
	rmdirResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(rmdirResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.RmdirResp = result
	//	CreateRate					float64
	createRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(createRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.CreateRate = result
	//	CreateResp					float64
	createResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(createResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.CreateResp = result
	//	OpenRate					float64
	openRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(openRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.OpenRate = result
	//	OpenResp					float64
	openResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(openResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.OpenResp = result
	//	CloseRate					float64
	closeRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(closeRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.CloseRate = result
	//	CloseResp					float64
	closeResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(closeResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.CloseResp = result
	//	DeleteRate					float64
	deleteRate := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(deleteRate, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.DeleteRate = result
	//	DeleteResp					float64
	deleteResp := `\d+\:\d+\:\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+\d+\.\d+[\s]+(\d+\.\d+).*`
	result, err = AnalysisResult(deleteResp, lineInfo)
	if err == false {
		return summaryInfo, false
	}
	summaryInfo.DeleteResp = result

	return summaryInfo, true
}

func InsertFilesystemData(db *sql.DB, mariaDBInfo *DBs.MariaDBInfo, sfsc *SummaryFileSystemCombination) {
	var sqlStatement string
	for key, _ := range sfsc.SFSI {
		sqlStatement = "INSERT INTO " + mariaDBInfo.MariaTableName + "(Id, DateTime, OutputInterval, ReqstdOpsRate, ReqstdOpsResp, CpuTotal, CpuSys, ReadPct, ReadRate, ReadResp, " +
			"WriteRate, WriteResp, MbSecRead, MbSecWrite, MbSecTotal, XferSize, MkdirRate, MkdirResp, RmdirRate, RmdirResp, CreateRate, CreateResp, OpenRate, OpenResp, CloseRate, " +
			"CloseResp, DeleteRate, DeleteResp, OperationTableDate, TestCase, ClientNumber)" + " VALUES " + "(" + "NULL, " + "\"" + sfsc.SFSI[key].DateTime+ "\"" + ", " + strconv.Itoa(sfsc.SFSI[key].OutputInterval) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].ReqstdOpsRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].ReqstdOpsResp) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].CpuTotal) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].CpuSys) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].ReadPct) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].ReadRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].ReadResp) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].WriteRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].WriteResp) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].MbSecRead) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].MbSecWrite) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].MbSecTotal) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].XferSize) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].MkdirRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].MkdirResp) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].RmdirRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].RmdirResp) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].CreateRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].CreateResp) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].OpenRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].OpenResp) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].CloseRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].CloseResp) + ", " +
			fmt.Sprintf("%f", sfsc.SFSI[key].DeleteRate) + ", " + fmt.Sprintf("%f", sfsc.SFSI[key].DeleteResp) + ", " + "\"" + time.Now().Format("2006-01-02 15:04:05") + "\"" + ", " + "\"" + sfsc.TestCase + "\"" + ", " + "\"" +  sfsc.ClientNumber + "\"" + ")"
			fmt.Println(sqlStatement)
			DBs.InsertMariadb(db, sqlStatement)
	}
	//fmt.Println(sqlStatement)
}
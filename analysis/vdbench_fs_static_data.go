package analysis

import (
	"bufio"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/lnsyyj/sttdv/comst"
	"github.com/lnsyyj/sttdv/dbs"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (ssd *FSStaticSummaryData) CheckParameter(extraInfo *comst.ExtraInfo) {
	outputInterval, err := strconv.Atoi(extraInfo.OutputInterval)
	if err != nil {
		fmt.Println("Parsing OutputInterval failed")
		return
	}
	if outputInterval <= 0 {
		fmt.Println("Please check the outputInterval parameter")
		return
	}
	if extraInfo.LogPath == "" {
		fmt.Println("Please check the file path")
		return
	}
}

func (ssd *FSStaticSummaryData) Process(mi *dbs.MariaDBInfo, extraInfo *comst.ExtraInfo) {
	file, err := os.Open(extraInfo.LogPath)
	if err != nil {
		panic("Read File Failed")
	}
	defer file.Close()

	fileInfo := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileInfo = append(fileInfo, scanner.Text())
	}

	// Parsing First Data

	//sfsiFirstDate := SummaryFileSystemInfoFirstDate{}
	ssd.SummaryFirstDate.Date, ssd.SummaryFirstDate.Time = ssd.ParsingFirstData(fileInfo)
	fmt.Println(ssd.SummaryFirstDate.Date, ssd.SummaryFirstDate.Time)

	// Parsing data
	//summaryFileSystemInfo := []SummaryFileSystemInfo{}
	for _, val := range fileInfo {
		re := ssd.ParsingData(val)
		if re.OutputInterval == "" {
			continue
		} else {
			ssd.SummaryInfo = append(ssd.SummaryInfo, re)
		}
	}

	// Assembly time
	ssd.AssemblyDateTime(extraInfo, ssd.SummaryFirstDate, ssd.SummaryInfo)

	//
	for _, val := range ssd.SummaryInfo {
		fmt.Println(val)
	}
	ssd.InsertDBSummaryFileSystemInfo(extraInfo, mi, ssd.SummaryInfo)
}

func (ssd *FSStaticSummaryData) ParsingFirstData(str []string) (string, string) {
	var reDate, reTime string
	for _, val := range str {
		// Jun 08, 2019  interval        i/o   MB/sec   bytes   read     resp     read    write     resp     resp queue  cpu%  cpu%
		reDateRegularTemp := regexp.MustCompile(`(\w+\s\d+,\s\d+).*`)
		reDateMatchTemp := reDateRegularTemp.FindStringSubmatch(val)
		if len(reDateMatchTemp) > 1 {
			reDate = reDateMatchTemp[1]
		}
		reTimeRegularTemp := regexp.MustCompile(`([0-9][0-9]\:[0-9][0-9]\:[0-9][0-9])\.[0-9]+[\s]+[0-9]+`)
		reTimeMatchTemp := reTimeRegularTemp.FindStringSubmatch(val)
		if len(reTimeMatchTemp) > 1 {
			reTime = reTimeMatchTemp[1]
		}
		if reDate != "" && reTime != "" {
			break
		}
	}
	return reDate, reTime
}

func (ssd *FSStaticSummaryData) InsertDBSummaryFileSystemInfo(extraInfo *comst.ExtraInfo, mi *dbs.MariaDBInfo, sfsi []SummaryFileSystemInfo) {
	db := dbs.ConnectionMariadb(mi)

	nowDateTime := time.Now().Format("2006-01-02 15:04:05")
	for _, val := range sfsi {
		sqlStatement := "INSERT INTO " + mi.MariaTableName + "(Id, DateTime, OutputInterval, ReqstdOpsRate, ReqstdOpsResp, CpuTotal, CpuSys, ReadPct, ReadRate, ReadResp, " +
			"WriteRate, WriteResp, MbSecRead, MbSecWrite, MbSecTotal, XferSize, MkdirRate, MkdirResp, RmdirRate, RmdirResp, " +
			"CreateRate, CreateResp, OpenRate, OpenResp, CloseRate, CloseResp, DeleteRate, DeleteResp, " +
			"OperationTableDate, TestCase, ClientNumber)" + " VALUES " +
			"(" + "NULL, " + "\"" + val.DateTime+ "\", " + val.OutputInterval + ", " +	val.ReqstdOpsRate + ", " + val.ReqstdOpsResp + ", " + val.CpuTotal + ", " + val.CpuSys + ", " + val.ReadPct + ", " + val.ReadRate + ", " + val.ReadResp + ", " +
			val.WriteRate + ", " + val.WriteResp + ", " + val.MbSecRead + ", " + val.MbSecWrite + ", " + val.MbSecTotal + ", " + val.XferSize + ", " + val.MkdirRate + ", " + val.MkdirResp + ", " +	val.RmdirRate + ", " + val.RmdirResp + 	", " +
			val.CreateRate + ", " + val.CreateResp + ", " + val.OpenRate + ", " + val.OpenResp + ", " + val.CloseRate + ", " + val.CloseResp + ", " +	val.DeleteRate + ", " + val.DeleteResp + ", \"" +
			nowDateTime + "\", \"" + extraInfo.TestCase + "\", \"" +  extraInfo.ClientNumber + "\")"
		dbs.InsertMariadb(db, sqlStatement)
	}

	dbs.CloseConnectionMariadb(db)
}

func (ssd *FSStaticSummaryData) AssemblyDateTime(extraInfo *comst.ExtraInfo, sfsifd SummaryFileSystemInfoFirstDate, sfsi []SummaryFileSystemInfo) {
	t, err := dateparse.ParseLocal(sfsifd.Date + " " + sfsifd.Time)
	if err != nil {
		panic("[ERROR] [AssemblingTime] : " + err.Error())
	}
	fmt.Println(t)
	for key, _ := range sfsi {
		m, _ := time.ParseDuration(strconv.Itoa(comst.StringToInt(extraInfo.OutputInterval) * key)+ "s")
		sfsi[key].DateTime = t.Add(m).Format("2006-01-02 15:04:05")
	}
}

func (ssd *FSStaticSummaryData) ParsingData(str string) SummaryFileSystemInfo {
	//Jun 08, 2019 .Interval. .ReqstdOps... ...cpu%...  read ....read..... ....write.... ..mb/sec... mb/sec .xfer.. ...mkdir.... ...rmdir.... ...create... ....open.... ...close.... ...delete...
	//                          rate   resp total  sys   pct   rate   resp   rate   resp  read write  total    size  rate   resp  rate   resp  rate   resp  rate   resp  rate   resp  rate   resp
	// 17:17:20.077       1     39.1 143.61  24.5 7.77   0.0    0.0  0.000   39.0 143.61  0.00 39.00  39.00 1048576   0.0  0.000   0.0  0.000   0.0  0.000  18.0  6.705   1.0 77.087   0.0  0.000
	sfsi := SummaryFileSystemInfo{}
	outputIntervalRegularTemp := `^\d+\:\d+\:\d+\.\d+[\s]+(\d+).*`
	re := regexp.MustCompile(outputIntervalRegularTemp)
	match := re.FindStringSubmatch(str)
	if len(match) > 1 {
		sfsi.OutputInterval = match[1]
	} else {
		return sfsi
	}

	dataLineRegularTemp := `\d+\:\d+\:\d+\.\d+[\s]+\d+(.*)`
	re = regexp.MustCompile(dataLineRegularTemp)
	match = re.FindStringSubmatch(str)
	if len(match) > 1 {
		strarray := strings.Fields(strings.TrimSpace(match[1]))
		sfsi.ReqstdOpsRate = strarray[0]
		sfsi.ReqstdOpsResp = strarray[1]
		sfsi.CpuTotal = strarray[2]
		sfsi.CpuSys = strarray[3]
		sfsi.ReadPct = strarray[4]
		sfsi.ReadRate = strarray[5]
		sfsi.ReadResp = strarray[6]
		sfsi.WriteRate = strarray[7]
		sfsi.WriteResp = strarray[8]
		sfsi.MbSecRead = strarray[9]
		sfsi.MbSecWrite = strarray[10]
		sfsi.MbSecTotal = strarray[11]
		sfsi.XferSize = strarray[12]
		sfsi.MkdirRate = strarray[13]
		sfsi.MkdirResp = strarray[14]
		sfsi.RmdirRate = strarray[15]
		sfsi.RmdirResp = strarray[16]
		sfsi.CreateRate = strarray[17]
		sfsi.CreateResp = strarray[18]
		sfsi.OpenRate = strarray[19]
		sfsi.OpenResp = strarray[20]
		sfsi.CloseRate = strarray[21]
		sfsi.CloseResp = strarray[22]
		sfsi.DeleteRate = strarray[23]
		sfsi.DeleteResp = strarray[24]
	} else {
		sfsi.OutputInterval = ""
		return sfsi
	}

	fmt.Println(sfsi)
	return sfsi
}
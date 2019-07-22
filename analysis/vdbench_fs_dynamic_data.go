package analysis

import (
	"bufio"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/lnsyyj/sttdv/comst"
	"github.com/lnsyyj/sttdv/dbs"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (dsd *FSDynamicSummaryData) Init() {

}

func (dsd *FSDynamicSummaryData) CheckParameter(extraInfo *comst.ExtraInfo) {
	outputInterval, err := strconv.Atoi(extraInfo.OutputInterval)
	if err != nil {
		fmt.Println("Parsing OutputInterval failed")
		return
	}
	if outputInterval <= 0 {
		fmt.Println("Please check the outputInterval parameter")
		return
	}
	if extraInfo.ToolPath == "" {
		fmt.Println("Please check the toolPath parameter")
		return
	}
}

func CheckPretreatment(str string) bool {
	reDateRegularTemp := regexp.MustCompile(`^\d+\:\d+\:\d+\.\d+[\s]+Miscellaneous statistics.*`)
	reDateMatchTemp := reDateRegularTemp.FindStringSubmatch(str)
	if len(reDateMatchTemp) > 1 {
		return false
	}
	return true
}

func (dsd *FSDynamicSummaryData) Process(mi *dbs.MariaDBInfo, extraInfo *comst.ExtraInfo) {
	cmdArgs := strings.Fields(extraInfo.ToolPath)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()

	pretreatment := true

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		sfsi := dsd.ParsingData(line)
		
		if pretreatment == true {
			if CheckPretreatment(line) {
				sfsi.Pretreatment = ""
			} else {
				pretreatment = false
				sfsi.Pretreatment = "no"
			}
		}
		
		fmt.Print(line)
		//fmt.Println(line)
		if dsd.SummaryFirstDate.Date == "" {
			dsd.SummaryFirstDate.Date = dsd.ParsingFirstDate(line)
		}
		if dsd.SummaryFirstDate.Time == "" {
			dsd.SummaryFirstDate.Time = dsd.ParsingFirstTime(line)
		}

		if sfsi.OutputInterval != "" {
			sfsi.DateTime = dsd.CalculationTime(extraInfo, &sfsi)
			dsd.InsertDBSummaryFileSystemInfo(extraInfo, mi, sfsi)
			//dsd.SummaryInfo = append(dsd.SummaryInfo, sfsi)
		}

	}
	for key, val := range dsd.SummaryInfo {
		fmt.Println(key, val)
	}

	cmd.Wait()
}

func (dsd *FSDynamicSummaryData) InsertDBSummaryFileSystemInfo(extraInfo *comst.ExtraInfo, mi *dbs.MariaDBInfo, sfsi SummaryFileSystemInfo) {
	db := dbs.ConnectionMariadb(mi)

	nowDateTime := time.Now().Format("2006-01-02 15:04:05")
	sqlStatement := "INSERT INTO " + mi.MariaTableName + "(Id, DateTime, OutputInterval, ReqstdOpsRate, ReqstdOpsResp, CpuTotal, CpuSys, ReadPct, ReadRate, ReadResp, " +
		"WriteRate, WriteResp, MbSecRead, MbSecWrite, MbSecTotal, XferSize, MkdirRate, MkdirResp, RmdirRate, RmdirResp, " +
		"CreateRate, CreateResp, OpenRate, OpenResp, CloseRate, CloseResp, DeleteRate, DeleteResp, " +
		"OperationTableDate, TestCase, ClientNumber, Pretreatment)" + " VALUES " +
		"(" + "NULL, " + "\"" + sfsi.DateTime+ "\", " + sfsi.OutputInterval + ", " +	sfsi.ReqstdOpsRate + ", " + sfsi.ReqstdOpsResp + ", " + sfsi.CpuTotal + ", " + sfsi.CpuSys + ", " + sfsi.ReadPct + ", " + sfsi.ReadRate + ", " + sfsi.ReadResp + ", " +
		sfsi.WriteRate + ", " + sfsi.WriteResp + ", " + sfsi.MbSecRead + ", " + sfsi.MbSecWrite + ", " + sfsi.MbSecTotal + ", " + sfsi.XferSize + ", " + sfsi.MkdirRate + ", " + sfsi.MkdirResp + ", " +	sfsi.RmdirRate + ", " + sfsi.RmdirResp + 	", " +
		sfsi.CreateRate + ", " + sfsi.CreateResp + ", " + sfsi.OpenRate + ", " + sfsi.OpenResp + ", " + sfsi.CloseRate + ", " + sfsi.CloseResp + ", " +	sfsi.DeleteRate + ", " + sfsi.DeleteResp + ", \"" +
		nowDateTime + "\", \"" + extraInfo.TestCase + "\", \"" +  extraInfo.ClientNumber + sfsi.Pretreatment + "\")"
	dbs.InsertMariadb(db, sqlStatement)

	dbs.CloseConnectionMariadb(db)
}

func (dsd *FSDynamicSummaryData) CalculationTime(extraInfo *comst.ExtraInfo, sfsi *SummaryFileSystemInfo) string {
	var reDateTime string
	t, err := dateparse.ParseLocal(dsd.SummaryFirstDate.Date + " " + dsd.SummaryFirstDate.Time)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	m, _ := time.ParseDuration(strconv.Itoa(comst.StringToInt(extraInfo.OutputInterval) * (comst.StringToInt(sfsi.OutputInterval) - 1))+ "s")
	reDateTime = t.Add(m).Format("2006-01-02 15:04:05")
	return	reDateTime
}

func (dsd *FSDynamicSummaryData) ParsingFirstDate(str string) string {
	var reDate string
	reDateRegularTemp := regexp.MustCompile(`(\w+\s\d+,\s\d+).*`)
	reDateMatchTemp := reDateRegularTemp.FindStringSubmatch(str)
	if len(reDateMatchTemp) > 1 {
		reDate = reDateMatchTemp[1]
	}
	return reDate
}

func (dsd *FSDynamicSummaryData) ParsingFirstTime(str string) string {
	var reTime string
	reTimeRegularTemp := regexp.MustCompile(`([0-9][0-9]\:[0-9][0-9]\:[0-9][0-9])\.[0-9]+[\s]+[0-9]+`)
	reTimeMatchTemp := reTimeRegularTemp.FindStringSubmatch(str)
	if len(reTimeMatchTemp) > 1 {
		reTime = reTimeMatchTemp[1]
	}
	return reTime
}

func (dsd *FSDynamicSummaryData) ParsingData(str string) SummaryFileSystemInfo {

	sfsi := SummaryFileSystemInfo{}
	outputIntervalRegularTemp := `^\d+\:\d+\:\d+\.\d+[\s]+(\d+)[\s]+.*`
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
	return  sfsi
}
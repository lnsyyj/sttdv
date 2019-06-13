package commons

import (
	"bufio"
	"github.com/lnsyyj/sttdv/analysis"
	"os"
)

type ReaderData interface {

}

type ICheckParameterValid interface {
	CheckParameterValid()
}

func InitDateTimeInterval(logPath *string, sfsc *analysis.SummaryFileSystemCombination, outputInterval *int) {
	file, err := os.Open(*logPath)
	lineInfo := ""
	defer file.Close()
	Check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineInfo = scanner.Text()
		if sfsc.SD.Data != "" && sfsc.SD.Time != "" {
			break
		}
		analysis.AnalysisFirstData(lineInfo, sfsc)
		analysis.AnalysisFirstTime(lineInfo, sfsc)
	}
	err = scanner.Err()
	Check(err)
	sfsc.SD.Conversion = sfsc.SD.Data + " " + sfsc.SD.Time
	sfsc.SD.OutputInterval = *outputInterval
}

func CheckParameter(cp ICheckParameterValid) {
	cp.CheckParameterValid()
}

func InitData(logPath *string, sfsc *analysis.SummaryFileSystemCombination) {
	summaryInfo := analysis.SummaryFileSystemInfo{}		// Data Info
	file, err := os.Open(*logPath)
	lineInfo := ""
	defer file.Close()
	Check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInfo = scanner.Text()
		summaryInfo = analysis.AnalysisSummaryInfo(lineInfo)
		if summaryInfo.OutputInterval == 0 {
			continue
		}
		sfsc.SFSI = append(sfsc.SFSI, summaryInfo)
	}
	err = scanner.Err()
	Check(err)
}
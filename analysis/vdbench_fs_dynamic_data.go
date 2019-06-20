package analysis

import (
	"bufio"
	"fmt"
	"github.com/lnsyyj/sttdv/comst"
	"github.com/lnsyyj/sttdv/dbs"
	"io"
	"os/exec"

	"strings"

)

func (dsd *FSDynamicSummaryData) CheckParameter(extraInfo *comst.ExtraInfo)  {
	if extraInfo.ToolPath == "" {
		fmt.Println("Please check the toolPath parameter")
		return
	}
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

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
}
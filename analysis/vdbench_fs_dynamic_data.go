package analysis

import (
	"fmt"
	"github.com/lnsyyj/sttdv/comst"
	"github.com/lnsyyj/sttdv/dbs"
)

func (dsd *FSDynamicSummaryData) CheckParameter(extraInfo *comst.ExtraInfo)  {
	if extraInfo.ToolPath == "" {
		fmt.Println("Please check the toolPath parameter")
		return
	}
}

func (dsd *FSDynamicSummaryData) Process(mi *dbs.MariaDBInfo, extraInfo *comst.ExtraInfo) {


}
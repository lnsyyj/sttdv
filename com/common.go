package com

import (
	"github.com/lnsyyj/sttdv/comst"
	"github.com/lnsyyj/sttdv/dbs"
)

type IParsingData interface {
	ParsingData()
}

type ICheckParameterValid interface {
	CheckParameter()	bool
}

func CheckParameterValid(cpv ICheckParameterValid) bool {
	return cpv.CheckParameter()
}

type IStaticDataProcess interface {
	CheckParameter(*comst.ExtraInfo)
	Process(*dbs.MariaDBInfo, *comst.ExtraInfo)
}

func StaticDataProcess(sdp IStaticDataProcess, mi *dbs.MariaDBInfo, extraInfo *comst.ExtraInfo) {
	sdp.CheckParameter(extraInfo)
	sdp.Process(mi, extraInfo)
}

type IDynamicDataProcess interface {
	CheckParameter(*comst.ExtraInfo)
	Process(*dbs.MariaDBInfo, *comst.ExtraInfo)
}

func DynamicDataProcess(ddp IDynamicDataProcess, mi *dbs.MariaDBInfo, extraInfo *comst.ExtraInfo) {
	ddp.CheckParameter(extraInfo)
	ddp.Process(mi, extraInfo)
}
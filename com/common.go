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
	Init(*dbs.MariaDBInfo, *comst.ExtraInfo)
	CheckParameter(*comst.ExtraInfo)
	Process(*dbs.MariaDBInfo, *comst.ExtraInfo)
}

func StaticDataProcess(sdp IStaticDataProcess, mi *dbs.MariaDBInfo, extraInfo *comst.ExtraInfo) {
	sdp.Init(mi, extraInfo)
	sdp.CheckParameter(extraInfo)
	sdp.Process(mi, extraInfo)
}

type IDynamicDataProcess interface {
	Init()
}


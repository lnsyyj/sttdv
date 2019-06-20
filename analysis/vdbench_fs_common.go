package analysis

type SummaryFileSystemInfo struct {
	//Id							int64
	DateTime					string
	OutputInterval				string
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
	OperationTableDate			string
	TestCase					string
	ClientNumber				string
}

type SummaryFileSystemInfoFirstDate struct {
	Date						string
	Time						string
}

type FSStaticSummaryData struct {
	SummaryFirstDate	SummaryFileSystemInfoFirstDate
	SummaryInfo	[]SummaryFileSystemInfo
}

type FSDynamicSummaryData struct {
	SummaryInfo	[]SummaryFileSystemInfo
}
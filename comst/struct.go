package comst

import (
	"github.com/lnsyyj/sttdv/dbs"
	"strconv"
)

type ExtraInfo struct {
	ToolPath		string
	LogPath			string
	OutputInterval 	string
	TestCase		string
	ClientNumber	string
}

func SetMariaDBInfo(mariaDBInfo *dbs.MariaDBInfo, ip *string, port *string, dbName *string, tableName *string, user *string, passwd *string) {
	mariaDBInfo.MariaHostIP = *ip
	mariaDBInfo.MariaPort = *port
	mariaDBInfo.MariaDatabaseName = *dbName
	mariaDBInfo.MariaTableName = *tableName
	mariaDBInfo.MariaUserName = *user
	mariaDBInfo.MariaUserPassword = *passwd
}

func SetExtraInfo(extraInfo *ExtraInfo, lp *string, oi *string, tc *string, cn *string, tp *string) {
	extraInfo.LogPath = *lp
	extraInfo.OutputInterval = *oi
	extraInfo.TestCase = *tc
	extraInfo.ClientNumber = *cn
	extraInfo.ToolPath = *tp
}

func StringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic("StringToInt failure")
	}
	return result
}
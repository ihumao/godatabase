package cdatabase

import (
	"database/sql"
)

type TField struct {
	Index   int
	Name    string
	Variant interface{}
}

type TFields []*TField

type TTransaction struct {
	ptx      *sql.Tx
	sql_text string
	sql_args []interface{}
}

type TQuery struct {
	pConn    *TCDatabase
	fields   TFields
	rows     []*TFields
	rowCount int // 数据行数
	rowIndex int // 当前所在行
	sql_text string
	sql_args []interface{}
}

type TCDatabase struct {
	remark     string
	pDB        *sql.DB
	DriveName  string
	SourceName string
}

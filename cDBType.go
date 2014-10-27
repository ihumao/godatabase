package cdatabase

import (
	"database/sql"
)

//type TAllRows struct {
//	prows  *sql.Rows
//	Fields []*TField
//}

//type TField struct {
//	Index   int
//	Name    string
//	Variant interface{}
//}

type TFields map[string]interface{}

type TTransaction struct {
	ptx      *sql.Tx
	sql_text string
	sql_args []interface{}
}

//type TRow struct {
//	fields TFields
//}

type TQuery struct {
	pConn    *TCDatabase
	fields   TFields
	rows     []*TFields
	sql_text string
	sql_args []interface{}
}

type TCDatabase struct {
	remark     string
	pDB        *sql.DB
	DriveName  string
	SourceName string
}

package cdatabase

import (
	"database/sql"
	"fmt"
)

func (this *TTrans) SQL(sql string, args ...interface{}) {
	this.sql_text = sql
	this.sql_args = args
}

func (this *TTrans) ExecSQL(sql string, args ...interface{}) int64 {
	this.SQL(sql, args)
	return this.Execute()
}

func (this *TTrans) Execute() int64 {
	var err error
	var pStmt *sql.Stmt
	var sqlResult sql.Result
	var nResult int64
	if pStmt, err = this.pTx.Prepare(this.sql_text); err != nil {
		fmt.Printf("[mysql Execute 1] >%s\n", err.Error())
		return -1
	}
	defer pStmt.Close()
	if sqlResult, err = pStmt.Exec(this.sql_args...); err != nil {
		fmt.Printf("[mysql Execute 2] >%s\n", err.Error())
		return -1
	}
	if nResult, err = sqlResult.LastInsertId(); err != nil {
		fmt.Printf("[mysql Execute 3] >%s\n", err.Error())
		return -1
	}
	return nResult
}

// cdatabase
package cdatabase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewTCDatabase(driverName, sourceName string) (*TCDatabase, error) {
	var err error
	p := new(TCDatabase)
	p.remark = ``
	p.DriveName = driverName
	p.SourceName = sourceName
	p.pDB, err = sql.Open(driverName, sourceName)
	return p, err
}

func (this *TCDatabase) Connectioned() bool {
	if this.pDB == nil {
		var err error
		if this.pDB, err = sql.Open(this.DriveName, this.SourceName); err != nil {
			fmt.Printf("[TCDatabase Connection] >%s\n", err.Error())
			return false
		}
	}
	return true
}

func (this *TCDatabase) NewQuery() *TQuery {
	p := new(TQuery)
	p.pConn = this
	p.bInTrans = false
	return p
}

func (this *TCDatabase) NewTrans() *TQuery {
	var err error
	var pTx *sql.Tx
	p := new(TQuery)
	p.pConn = this
	if pTx, err = p.pConn.pDB.Begin(); err != nil || pTx == nil {
		return nil
	}
	p.pTx = pTx
	p.bInTrans = true
	return p
}

//func (this *TTransaction) Commit() error {
//	return this.ptx.Commit()
//}

//func (this *TTransaction) Rollback() error {
//	return this.ptx.Rollback()
//}

//func (this *TTransaction) SQL(sql string, args ...interface{}) {
//	this.sql_text = sql
//	this.sql_args = args
//}

//func (this *TTransaction) Execute() (int64, error) {
//	stmt_exec, err := this.ptx.Prepare(this.sql_text)
//	if err != nil {
//		fmt.Printf("[mysql execute 1] >%s\n", err.Error())
//		return -1, err
//	}
//	defer stmt_exec.Close()

//	exec_result, err := stmt_exec.Exec(this.sql_args...)
//	if err != nil {
//		fmt.Printf("[mysql execute 2] >%s\n", err.Error())
//		return -1, err
//	}
//	nLastInsertId, err := exec_result.LastInsertId()
//	if err != nil {
//		fmt.Printf("[mysql execute 3] >%s\n", err.Error())
//		return -1, err
//	}
//	return nLastInsertId, nil
//}

// dboperate
package cdatabase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewDBOperator(driverName, sourceName string) (*TCDatabase, error) {
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
		this.pDB, err = sql.Open(this.DriveName, this.SourceName)
		if err != nil {
			fmt.Printf("[dboperate Connection] >%s\n", err.Error())
			return false
		}
	}
	return true
}

func (this *TCDatabase) NewQuery() *TQuery {
	p := new(TQuery)
	p.pConn = this
	return p
}

func (this *TCDatabase) NewTransaction() (*TTransaction, error) {
	var err error
	p := new(TTransaction)
	p.ptx, err = this.pDB.Begin()
	return p, err
}

func (qry *TQuery) SQL(sql string, args ...interface{}) {
	qry.sql_text = sql
	qry.sql_args = args
}

func (qry *TQuery) Execute() (int64, error) {
	stmt_exec, err := qry.pConn.pDB.Prepare(qry.sql_text)
	if err != nil {
		fmt.Printf("[mysql execute 1] >%s\n", err.Error())
		return -1, err
	}
	defer stmt_exec.Close()

	exec_result, err := stmt_exec.Exec(qry.sql_args...)
	if err != nil {
		fmt.Printf("[mysql execute 2] >%s\n", err.Error())
		return -1, err
	}
	nLastInsertId, err := exec_result.LastInsertId()
	if err != nil {
		fmt.Printf("[mysql execute 3] >%s\n", err.Error())
		return -1, err
	}
	return nLastInsertId, nil
}

func (this *TQuery) newFileds() *TFields {
	p := new(TFields)
	for k, _ := range this.fields {
		(*p)[k] = nil
	}
	return p
}

func (this *TQuery) Open() error {
	var pStmt *sql.Stmt
	var pRows *sql.Rows
	var cols []string
	var err error
	if pStmt, err = this.pConn.pDB.Prepare(this.sql_text); err != nil {
		fmt.Printf("[mysql query 1] >%s", err.Error())
		return err
	}
	defer pStmt.Close()
	if pRows, err = pStmt.Query(this.sql_args...); err != nil {
		fmt.Printf("[mysql query 2] >%s", err.Error())
		return err
	}
	if cols, err = pRows.Columns(); err != nil {
		fmt.Printf("[mysql query 3] >%s", err.Error())
		return err
	}
	for _, v := range cols {
		this.fields[v] = nil
	}
	var arrValue []interface{}
	for pRows.Next() {
		arrValue = nil
		p := this.newFileds()
		for _, v := range *p {
			arrValue = append(arrValue, &v)
		}
		pRows.Scan(arrValue...)
		this.rows = append(this.rows, p)
	}
	return nil
}

func (this *TTransaction) Commit() error {
	return this.ptx.Commit()
}

func (this *TTransaction) Rollback() error {
	return this.ptx.Rollback()
}

func (this *TTransaction) SQL(sql string, args ...interface{}) {
	this.sql_text = sql
	this.sql_args = args
}

func (this *TTransaction) Execute() (int64, error) {
	stmt_exec, err := this.ptx.Prepare(this.sql_text)
	if err != nil {
		fmt.Printf("[mysql execute 1] >%s\n", err.Error())
		return -1, err
	}
	defer stmt_exec.Close()

	exec_result, err := stmt_exec.Exec(this.sql_args...)
	if err != nil {
		fmt.Printf("[mysql execute 2] >%s\n", err.Error())
		return -1, err
	}
	nLastInsertId, err := exec_result.LastInsertId()
	if err != nil {
		fmt.Printf("[mysql execute 3] >%s\n", err.Error())
		return -1, err
	}
	return nLastInsertId, nil
}

package cdatabase

import (
	"database/sql"
	"fmt"
)

func (this *TQuery) SQL(sql string, args ...interface{}) {
	this.sql_text = sql
	this.sql_args = args
}

func (this *TQuery) ExecSQL(sql string, args ...interface{}) int64 {
	this.SQL(sql, args)
	return this.Execute()
}

func (this *TQuery) Execute() int64 {
	var err error
	var pStmt *sql.Stmt
	var sqlResult sql.Result
	var nResult int64
	if pStmt, err = this.pConn.pDB.Prepare(this.sql_text); err != nil {
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

func (this *TQuery) newFields() TFields {
	var fields TFields
	//fmt.Println(this.fields)
	for _, v := range this.fields {
		field := TField{v.Index, v.Name, nil}
		fields = append(fields, &field)
	}
	return fields
}

func (this *TQuery) Query(sql string, args ...interface{}) error {
	this.SQL(sql, args)
	return this.Open()
}

func (this *TQuery) Open() error {
	var pStmt *sql.Stmt
	var pRows *sql.Rows
	var cols []string
	var err error
	this.rows = nil
	this.fields = nil
	this.rowCount = 0
	this.rowIndex = -1
	if pStmt, err = this.pConn.pDB.Prepare(this.sql_text); err != nil {
		fmt.Printf("[mysql Open 1] >%s\n", err.Error())
		return err
	}
	defer pStmt.Close()
	if pRows, err = pStmt.Query(this.sql_args...); err != nil {
		fmt.Printf("[mysql Open 2] >%s\n", err.Error())
		return err
	}
	defer pRows.Close()
	if cols, err = pRows.Columns(); err != nil {
		fmt.Printf("[mysql Open 3] >%s\n", err.Error())
		return err
	}
	for k, v := range cols {
		field := TField{k, v, nil}
		this.fields = append(this.fields, &field)
	}
	var nCount int = 0
	for pRows.Next() {
		nCount = nCount + 1
		var arrValue []interface{}
		fields := this.newFields()
		for _, v := range fields {
			arrValue = append(arrValue, &(v.Variant))
		}
		pRows.Scan(arrValue...)
		this.rows = append(this.rows, &fields)
	}
	this.rowCount = nCount
	this.rowIndex = -1
	//fmt.Println(nCount)
	//fmt.Println(this.rows)
	//for _, v := range this.rows {
	//	fmt.Println((*v)[0].Variant, (*v)[1].Variant, (*v)[2].Variant)
	//}
	return nil
}

func (this *TQuery) Next() bool {
	if this.rowCount > 0 && this.rowIndex < (this.rowCount-1) {
		this.rowIndex = this.rowIndex + 1
		return true
	}
	return false
}

func (this *TQuery) FieldByName(value string) *TField {
	if this.rowIndex >= 0 && this.rowIndex < this.rowCount {
		for _, v := range *this.rows[this.rowIndex] {
			if v.Name == value {
				return v
			}
		}
	}
	panic(`Field '` + value + `' not found`)
}

func (this *TQuery) FieldByIndex(value int) *TField {
	if this.rowIndex >= 0 && this.rowIndex < this.rowCount {
		if value >= 0 && value < len(*this.rows[this.rowIndex]) {
			return (*this.rows[this.rowIndex])[value]
		}
	}
	panic(fmt.Sprintf(`List index out of bounds (%d)`, value))
}

func (this *TQuery) First() {
	if this.rowCount < 1 {
		return
	}
	this.rowIndex = 0
}

func (this *TQuery) Last() {
	if this.rowCount < 1 {
		return
	}
	this.rowIndex = this.rowCount - 1
}

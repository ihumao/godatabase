package main

import (
	".."
	"fmt"
	//"time"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
	//pdb, _ := cdatabase.NewTCDatabase(`mysql`, `tigercat:3206.net@tcp(mail.zs310.com:8529)/test?charset=utf8`)
	pdb, _ := cdatabase.NewTCDatabase(`mysql`, `tigercat:3206.net@tcp(mail.zs310.com:8529)/lfsvr?charset=utf8`)
	pQuery := pdb.NewQuery()
	//pQuery.SQL(`SELECT * FROM issue_basic;`)
	pQuery.SQL(`SELECT * FROM order_basic LIMIT 10;`)
	pQuery.Open()
	for pQuery.Next() {
		//fmt.Println(pQuery.FieldByName(`r1`).AsString(),
		//	pQuery.FieldByName(`r2`).AsString(), pQuery.FieldByName(`r3`).AsString(),
		//	pQuery.FieldByName(`r4`).AsString())
		//fmt.Println(pQuery.FieldByName(`lf_issue`).AsString(),
		//	pQuery.FieldByName(`lottery_id`).AsString(), pQuery.FieldByName(`lottery_issue`).AsString(),
		//	pQuery.FieldByName(`r4`).AsString())
		fmt.Println(pQuery.FieldByIndex(0).AsString(),
			pQuery.FieldByIndex(1).AsString(), pQuery.FieldByIndex(2).AsString(),
			pQuery.FieldByIndex(3).AsString())
	}
	//time.Sleep(5 * time.Second)
	runtime.GC()
	fmt.Println(`aaaaaaaaaaaaaaaaaaaaaaaa`)
}

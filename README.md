# sqltransactioncheck
sqltransactioncheck 用于检测公司代码事务是否提交

开发调试
go run main.go ./testdata/trans

集成环境运行
/Users/xxxx/xxxx/golangci-lint-1.55.2/golangci-lint run ./testdata/trans


#demo


  tm := dbutil.NewTransaction()
	tm.SetCommitFlag(true)
	//tm.RollbackIfNotCommit()

 /Users/xxxx/sqltransactioncheck/testdata/trans/gorm.go:14:8: SQL Transaction(commit|rollback) was not submited warning
exit status 3

package trans_examples

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/macopad/sqltransactioncheck/pkg/analyzer/testdata/dbutil"
)

/*
var (
	db *gorm.DB
)

type TM struct {
	tx         *gorm.DB
	commitFlag bool
}

type TransactionManager interface {
	SetCommitFlag(flag bool)
	RollbackIfNotCommit()
	GetConn() *gorm.DB
}

func NewTransaction() TransactionManager {
	tm := &TM{}
	tm.tx = db.Begin()
	tm.commitFlag = false
	err := tm.tx.Error
	if err != nil {
		//
	}
	return tm
}

func (tm *TM) GetConn() *gorm.DB {
	return tm.tx
}

func (tm *TM) SetCommitFlag(flag bool) {
	tm.commitFlag = flag
}

func (tm *TM) RollbackIfNotCommit() {
	if tm.commitFlag == true {
		err := tm.tx.Commit().Error
		if err != nil {
			//
		}
	} else {
		err := tm.tx.Rollback().Error
		if err != nil {
			//
		}
	}
}*/

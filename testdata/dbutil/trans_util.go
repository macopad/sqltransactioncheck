package dbutil

import (
	"github.com/jinzhu/gorm"
)

type TransactionManager interface {
	SetCommitFlag(flag bool)
	RollbackIfNotCommit()
	GetConn() *gorm.DB
	SetConnOption(name string, value interface{}) TransactionManager
}

type TM struct {
	tx         *gorm.DB
	commitFlag bool
}

func (tm *TM) clone() *TM {

	ntm := &TM{}
	ntm.tx = tm.tx
	ntm.commitFlag = tm.commitFlag

	return ntm
}

func (tm *TM) SetCommitFlag(flag bool) {
	tm.commitFlag = flag
}

func (tm *TM) RollbackIfNotCommit() {
	if tm.commitFlag == true {
		tm.tx.Commit()
	} else {
		tm.tx.Rollback()
	}
}

func (tm *TM) GetConn() *gorm.DB {
	return tm.tx
}

func (tm *TM) SetConnOption(name string, value interface{}) TransactionManager {
	ntm := tm.clone()
	ntm.tx = ntm.tx.Set(name, value)
	return ntm
}

type TMWithoutTrans struct {
	tx *gorm.DB
}

func (tm *TMWithoutTrans) clone() *TMWithoutTrans {

	ntm := &TMWithoutTrans{}
	ntm.tx = tm.tx

	return ntm
}

func (tm *TMWithoutTrans) SetCommitFlag(flag bool) {
	panic("CANNOT use SetCommitFlag in TMWithoutTrans")
}

func (tm *TMWithoutTrans) RollbackIfNotCommit() {
	panic("CANNOT use RollbackIfNotCommit in TMWithoutTrans")
}

func (tm *TMWithoutTrans) GetConn() *gorm.DB {
	return tm.tx
}

func (tm *TMWithoutTrans) SetConnOption(name string, value interface{}) TransactionManager {
	ntm := tm.clone()
	ntm.tx = ntm.tx.Set(name, value)
	return ntm
}

type TMSlave struct {
	tx *gorm.DB
}

func (tm *TMSlave) clone() *TMSlave {

	ntm := &TMSlave{}
	ntm.tx = tm.tx

	return ntm
}

func (tm *TMSlave) SetCommitFlag(flag bool) {
	panic("CANNOT use SetCommitFlag in TMSlave")
}

func (tm *TMSlave) RollbackIfNotCommit() {
	panic("CANNOT use RollbackIfNotCommit in TMSlave")
}

func (tm *TMSlave) GetConn() *gorm.DB {
	return tm.tx
}

func (tm *TMSlave) SetConnOption(name string, value interface{}) TransactionManager {
	ntm := tm.clone()
	ntm.tx = ntm.tx.Set(name, value)
	return ntm
}

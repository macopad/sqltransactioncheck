package trans_examples

import (
	"testdata/trans_examples/testdata/dbutil"
)

func testTransaction1() {

	tm := dbutil.NewTransaction()
	tm.SetCommitFlag(true)
	//tm.RollbackIfNotCommit()

	/*
		var tm = dbutil.NewTransaction()

		err := tm.GetConn().Exec("select 1")

		if err != nil {
			tm.SetCommitFlag(false)
		}

		tm.SetCommitFlag(true)
		//defer tm.RollbackIfNotCommit()

		fmt.Printf("data:%+v", err)
	*/
}

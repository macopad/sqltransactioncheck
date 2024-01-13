package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

// NewAnalyzer returns a non-configurable analyzer that defaults to the defer-only mode.
// Deprecated, this will be removed in v1.0.0.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "sqltransactioncheck",
		Doc:  "Checks transaction are submit.",
		Run:  run,
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}
}

// Run implements the main analysis pass
func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {
		//定义是否开启事务
		transaction := false
		var tempPos ast.Node
		//定义是否检测到事务提交
		isClosed := false

		ast.Inspect(file, func(node ast.Node) bool {
			callExpr, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			// 检查函数调用是否是 `NewTransaction` 方法
			if isMethodCall(callExpr, "NewTransaction") {
				transaction = true
				//fmt.Printf("事务开启状态： %+v \n", transaction)

				tempPos = callExpr.Fun
				//pass.Reportf(tempPos.Pos(), "SQL Transaction start")
			}

			//检查事务是否有提交
			if checkTransactionClose(pass, callExpr) == true {
				isClosed = true
			}

			return true
		})

		if transaction == true && isClosed == false {
			pass.Reportf(tempPos.Pos(), "SQL Transaction(commit|rollback) was not submit warning")
		}
	}

	return nil, nil

}

func isMethodCall(callExpr *ast.CallExpr, methodName string) bool {

	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if ident, ok := selectorExpr.X.(*ast.Ident); ok && ident.Name == "dbutil" {
		if selectorExpr.Sel.Name == methodName {
			return true
		}
	}
	return false
}

func isTmMethodCall(callExpr *ast.CallExpr, methodName string) bool {
	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := selectorExpr.X.(*ast.Ident)
	if ok && (ident.Name == "tm" || ident.Name == "dbMgr") {
		if selectorExpr.Sel.Name == methodName {
			return true
		}
	}
	return false
}

func checkTransactionClose(pass *analysis.Pass, callExpr *ast.CallExpr) bool {
	if !isTransactionClosed(callExpr) {
		//pass.Reportf(callExpr.Pos(), "事务未正确关闭")
		return false
	}
	return true
}

func isTransactionClosed(callExpr *ast.CallExpr) bool {
	/*
		//排除other非tm函数
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return false
		}

		ident, ok := selectorExpr.X.(*ast.Ident)
		//如果不是tm或者dbMgr的类，直接忽略掉不检查
		notCheckObjList := map[string]bool{
			"tm":    true,
			"dbMgr": true,
		}

		//如果在map内，直接忽略
		if notCheckObjList[ident.Name] != true {
			return false
		}
	*/

	// 检查事务关闭的常见方式
	if isTmMethodCall(callExpr, "RollbackIfNotCommit") {
		return true
	}

	// 检查 defer 关闭事务
	//if isDeferTransactionClose(callExpr) {
	//	return true
	//}
	return false
}

/*
func isDeferTransactionClose(callExpr *ast.CallExpr) bool {
	deferStmt, ok := getParentDeferStmt(callExpr)
	if !ok {
		return false
	}

	// 检查 defer 语句中的函数调用是否是事务关闭的方法
	return isTransactionClosed(deferStmt.Call)
}

func getParentDeferStmt(callExpr *ast.CallExpr) (*ast.DeferStmt, bool) {
	for _, stmt := range getContainingStmts(callExpr) {
		if deferStmt, ok := stmt.(*ast.DeferStmt); ok {
			return deferStmt, true
		}
	}

	return nil, false
}

func getContainingStmts(node ast.Node) []ast.Stmt {
	var stmts []ast.Stmt

	for {
		node = node.Parent()
		if node == nil {
			break
		}

		if stmt, ok := node.(ast.Stmt); ok {
			stmts = append(stmts, stmt)
		}
	}

	return stmts
}*/

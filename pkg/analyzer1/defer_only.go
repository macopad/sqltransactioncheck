package analyzer1

//import (
//	"flag"
//	"fmt"
//	"go/ast"
//	"go/token"
//	"golang.org/x/tools/go/analysis"
//	"golang.org/x/tools/go/analysis/passes/buildssa"
//)
//
//type deferOnlyAnalyzer struct{}
//
//func NewDeferOnlyAnalyzer() *analysis.Analyzer {
//	analyzer := &deferOnlyAnalyzer{}
//	flags := flag.NewFlagSet("deferOnlyAnalyzer", flag.ExitOnError)
//	return newAnalyzer(analyzer.Run, flags)
//}
//
//// Run implements the main analysis pass
//func (a *deferOnlyAnalyzer) Run(pass *analysis.Pass) (interface{}, error) {
//
//	_, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
//	if !ok {
//		return nil, nil
//	}
//
//	for _, file := range pass.Files {
//		//暂定检测结构体
//		var checkTransactionObj = "&{dbutil TransactionManager}"
//		var checkTransactionSubmit = "RollbackIfNotCommit"
//
//		// 遍历源代码文件中的所有节点
//		ast.Inspect(file, func(n ast.Node) bool {
//			//遍历函数
//			if function, ok := n.(*ast.FuncDecl); ok {
//				//fmt.Printf("函数名: %s\n", function.Name.Name)
//
//				//定义事务是否开启
//				var transaction bool = false
//				//定义事务是否提交
//				var submit bool = false
//
//				// 遍历函数内部的语句
//				ast.Inspect(function.Body, func(n ast.Node) bool {
//
//					// 检查变量声明语句
//					if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
//						for _, spec := range decl.Specs {
//							if valueSpec, ok := spec.(*ast.ValueSpec); ok {
//								// 遍历变量声明中的每个变量
//								for _, ident := range valueSpec.Names {
//									// 临时变量
//									tempTransaction := fmt.Sprintf("%s", valueSpec.Type)
//									if tempTransaction == checkTransactionObj {
//										transaction = true
//										fmt.Sprintf("变量: %s, 类型: %s\n", ident.Name, valueSpec.Type)
//									}
//								}
//							}
//						}
//					}
//
//					// 检查函数调用表达式
//					if callExpr, ok := n.(*ast.CallExpr); ok {
//						selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
//						if !ok {
//							return false
//						}
//
//						if _, ok := selectorExpr.X.(*ast.Ident); ok {
//							// 打印函数调用的函数名称
//							//fmt.Printf("函数调用: %s\n", selectorExpr.Sel.Name)
//							if selectorExpr.Sel.Name == checkTransactionSubmit {
//								submit = true
//							}
//						}
//
//					}
//
//					return true
//				})
//
//				if transaction == true && submit == false {
//					pass.Reportf(function.Pos(), "SQL Transaction(commit|rollback) was not submit warning")
//				}
//
//			}
//
//			return true
//		})
//
//	}
//
//	return nil, nil
//}

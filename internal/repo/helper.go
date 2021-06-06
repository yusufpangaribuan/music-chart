package repository

import (
	"database/sql"
	"log"

	"github.com/lp/music-chart/connections/mysql"
)

type Stmt struct {
	Query     string
	ActionNum int
}

func BuildStmt(db mysql.DB, stmtList []Stmt, repoName string) (res map[int]*sql.Stmt) {
	if !isUniqueStmt(stmtList) {
		log.Fatalf("%s: stmt action must be unique", repoName)
	}

	res = make(map[int]*sql.Stmt)

	for _, val := range stmtList {

		stmt, err := db.Client.Prepare(val.Query)
		if err != nil {
			log.Fatal(err)
		}
		res[val.ActionNum] = stmt
	}
	return
}

func isUniqueStmt(stmt []Stmt) bool {
	flag := map[int]bool{}
	for _, val := range stmt {
		if flag[val.ActionNum] == true {
			return false
		}
		flag[val.ActionNum] = true
	}

	return true
}

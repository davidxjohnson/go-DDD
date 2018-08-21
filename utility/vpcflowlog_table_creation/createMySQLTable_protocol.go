package main

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var tablename = "aws_protocols"

func createTable(dbhandler *sql.DB, table string) *sql.Stmt {
	statement := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id                    INTEGER NOT NULL AUTO_INCREMENT, 
			protocol_integer      INTEGER NOT NULL,
                        protocol_keyword      VARCHAR(128) NOT NULL,
			protocol_description  VARCHAR(128) NOT NULL,
			inserted_timestamp    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)    
		)`, table)

	evaluatedStatement, err := dbhandler.Prepare(statement)
	
	if err != nil {
		fmt.Println(err)
	}
	
	return evaluatedStatement
}

func execStatement(dbhandler *sql.DB) sql.Result {
	statement := createTable(dbhandler, tablename)
	result, err := statement.Exec()
	
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func main(){
	dbconn,_ := sql.Open("mysql", os.Getenv("CONNSTRING"))
	dberr := dbconn.Ping()
	defer dbconn.Close()

	if dberr != nil {
		fmt.Println(dberr)
	}
	
	execStatement(dbconn)
}

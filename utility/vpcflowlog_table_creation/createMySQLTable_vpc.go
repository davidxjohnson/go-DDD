package main

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var tablename = "aws_flowlogs"

func createTable(dbhandler *sql.DB, table string) *sql.Stmt {
	statement := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id                    INTEGER NOT NULL AUTO_INCREMENT, 
			version               VARCHAR(128) NOT NULL,
			account               VARCHAR(128) NOT NULL, 
			interface_id          VARCHAR(128) NOT NULL, 
			srcaddr               VARCHAR(128) NOT NULL,
			dstaddr               VARCHAR(128) NOT NULL, 
			srcport               VARCHAR(128) NOT NULL,
           	        dstport               VARCHAR(128) NOT NULL,
			protocol              VARCHAR(128) NOT NULL,
			packets               VARCHAR(128) NOT NULL,
			bytes                 VARCHAR(128) NOT NULL,
			start                 VARCHAR(128) NOT NULL,
			end                   VARCHAR(128) NOT NULL,
			action                VARCHAR(128) NOT NULL,
			log_status            VARCHAR(128) NOT NULL,
			inserted_timestamp    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			row_id                VARCHAR(128) NOT NULL,
			to_resource           VARCHAR(128) NOT NULL,
			from_resource         VARCHAR(128) NOT NULL,
			processed             BOOLEAN DEFAULT 0, 
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

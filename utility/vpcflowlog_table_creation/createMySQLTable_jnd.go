package main

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var tablename = "staged_neptune_data"

func createTable(dbhandler *sql.DB, table string) *sql.Stmt {
	statement := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id                    INTEGER NOT NULL AUTO_INCREMENT,
			row_id                VARCHAR(128) NOT NULL, 
			version               VARCHAR(128) NOT NULL,
			label                 VARCHAR(128) NOT NULL,
			account               VARCHAR(128) NOT NULL, 
			networkinterfaceid    VARCHAR(128) NOT NULL, 
			srcaddr               VARCHAR(128) NOT NULL,
			dstaddr               VARCHAR(128) NOT NULL, 
			srcport               VARCHAR(128) NOT NULL,
            dstport               varchar(128) not null,
			protocol              VARCHAR(128) NOT NULL,
			packets               VARCHAR(128) NOT NULL,
			bytes                 VARCHAR(128) NOT NULL,
			starttime             VARCHAR(128) NOT NULL, 
			endtime               VARCHAR(128) NOT NULL, 
			inserted_date         VARCHAR(128) NOT NULL,
			action                VARCHAR(128) NOT NULL,
			logstatus             VARCHAR(128) NOT NULL,
			date_time_stamp       VARCHAR(128) NOT NULL,
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

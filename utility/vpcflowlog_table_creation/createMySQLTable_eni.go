package main

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var tablename = "aws_resources_testing"

func createTable(dbhandler *sql.DB, table string) *sql.Stmt {
	statement := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id                    INTEGER NOT NULL AUTO_INCREMENT,
			row_id                VARCHAR(128) NOT NULL, 
			label                 VARCHAR(128) NOT NULL,
			node_type             VARCHAR(128) NOT NULL, 
			networkinterfaceid    VARCHAR(128) NOT NULL, 
			description           VARCHAR(128) NOT NULL,
			account               VARCHAR(128) NOT NULL, 
			region                VARCHAR(128) NOT NULL,
			availability_zone     VARCHAR(128) NOT NULL, 
			vpc_id                VARCHAR(128) NOT NULL,
                        subnet_id             VARCHAR(128) NOT NULL,
			private_ip_address    VARCHAR(128) NOT NULL,
			private_dns_name      VARCHAR(128) NOT NULL,
			secondary_ip_address  VARCHAR(128) NOT NULL,
			name_tag              VARCHAR(128) NOT NULL, 
			business_unit_tag     VARCHAR(128) NOT NULL,
			business_region_tag   VARCHAR(128) NOT NULL,
			platform_tag          VARCHAR(128) NOT NULL,
			client_tag            VARCHAR(128) NOT NULL,
			resourcecreatedby_tag VARCHAR(128) NOT NULL, 
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

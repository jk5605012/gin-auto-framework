package main

import (
	"flag"
	"gin-auto-framework/create"
)

var (
	dbType string
	ini    string
)

func init() {
	flag.StringVar(&ini, "init", "", "init new porject name")
	flag.StringVar(&dbType, "db", "", "mysql, postgres, sqlite or mssql")
}

func main() {
	flag.Parse()

	if ini != "" && isValidDB() {
		create.CreateFolder(ini)
		create.CreateAllFile(ini, dbType)
	} else {
		flag.Usage()
	}

}

func isValidDB() bool {
	return dbType != "" && (dbType == "mysql" || dbType == "postgres" || dbType == "sqlite" || dbType == "mssql")
}

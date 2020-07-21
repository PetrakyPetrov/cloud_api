package libs

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql" // mysql ..
)

// DBmap ..
var DBmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cloud_storage")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	// err = dbmap.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

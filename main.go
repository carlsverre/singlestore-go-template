package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var sdbHost = flag.String("sdb-host", "127.0.0.1", "The SingleStore DB host to connect to")
var sdbPort = flag.Int("sdb-port", 3306, "The SingleStore DB port to connect to")
var sdbUser = flag.String("sdb-user", "root", "The SingleStore DB user to use")
var sdbPass = flag.String("sdb-password", "", "The SingleStore DB user's password")

func connectSingleStore() (*sql.DB, error) {
	connParams := strings.Join([]string{
		// convert timestame and date to time.Time
		"parseTime=true",
		// don't use the binary protocol
		"interpolateParams=true",
		// set a sane connection timeout rather than the default infinity
		"timeout=10s",

		// strict mode
		"collation_server=utf8_general_ci",
		"sql_select_limit=18446744073709551615",
		"compile_only=false",
		"enable_auto_profile=false",
		"sql_mode='STRICT_ALL_TABLES,ONLY_FULL_GROUP_BY'",

		// use SSL when server advertises support
		"tls=preferred",
	}, "&")

	host := *sdbHost
	if host == "0.0.0.0" {
		host = "127.0.0.1"
	}

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/information_schema?%s",
		*sdbUser,
		*sdbPass,
		host,
		*sdbPort,
		connParams,
	)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	// the maximum amount of time a connection may be reused
	db.SetConnMaxLifetime(time.Hour * 6)

	// the maximum number of open connections to the database
	db.SetMaxOpenConns(10)

	//the maximum number of connections in the idle connection pool
	db.SetMaxIdleConns(10)

	return db, nil
}

func main() {
	flag.Parse()

	handleError := func(err error, msg string, args ...interface{}) {
		if err != nil {
			msg := fmt.Sprintf(msg, args...)
			log.Fatalf("%s: %+v", msg, err)
		}
	}

	log.Println("connecting to SingleStore")
	conn, err := connectSingleStore()
	handleError(err, "error while connecting to SingleStore")

	log.Println("creating example database")
	_, err = conn.Exec("create database if not exists example")
	handleError(err, "error while creating example database")

	log.Println("creating example table")
	_, err = conn.Exec(`
		create table if not exists example.test (
			id int primary key,
			s text
		)
	`)
	handleError(err, "error while creating test table")

	log.Println("inserting sample data")
	_, err = conn.Exec(`
		insert ignore into example.test values
			(1, "hello"),
			(2, "world"),
			(3, "this is an example")
	`)
	handleError(err, "error while inserting sample data")

	result, err := conn.Query("select * from example.test")
	handleError(err, "error while querying test table")

	i := 0
	for result.Next() {
		i++
		var (
			id int
			s  string
		)
		err := result.Scan(&id, &s)
		handleError(err, "error while scanning row %d", i)

		log.Printf("Row %d: (%d, %s)", i, id, s)
	}
}

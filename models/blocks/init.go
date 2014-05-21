package blocks

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
)

var (
	db *sql.DB
)

func init() {
	// Load config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve credentials from config file
	username := to.String(conf.Get("database", "username"))
	password := to.String(conf.Get("database", "password"))
	name := to.String(conf.Get("database", "name"))

	// Open mysql connection
	dsn := fmt.Sprintf("%s:%s@/%s?%s", username, password, name, "parseTime=true")
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Sets the maximum number of connections in the idle connection pool
	db.SetMaxIdleConns(100)

}

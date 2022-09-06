package test

import (
	"flag"
	"fmt"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "sql2008", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "10.69.16.5", "the database server")
	user          = flag.String("user", "sa", "the database user")
	dbname        = flag.String("name", "Lorentz_China", "the default database name")
)

func testflags() {
	flag.Parse()
	others := flag.Args()
	fmt.Println(others)
}

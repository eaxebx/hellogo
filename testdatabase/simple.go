package testdatabase

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "sql2008", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "10.69.16.5", "the database server")
	user          = flag.String("user", "sa", "the database user")
	dbname        = flag.String("name", "Lorentz_China", "the default database name")
)

func Testdatabasemain() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *dbname)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	defer conn.Close()

	QueryRows(conn)
}

func QueryRows(db *sql.DB) {
	rows, err := db.Query("select Itemcode,itemname from oitm where itemcode like '30-0109%'")
	if err != nil {
		log.Fatal("ecc", err.Error())
	}
	for rows.Next() {
		var a, b string
		err := rows.Scan(&a, &b)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(a, b)
	}
}
func prepareQueryRow(db *sql.DB) {
	stmt, err := db.Prepare("select ItemCode,ItemName from OITM where ItemCode='30-001766'")
	if err != nil {
		log.Fatal("Prepare failed:", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var ItemCode string
	var ItemName string
	err = row.Scan(&ItemCode, &ItemName)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	fmt.Printf("ItemCode: %s\n", ItemCode)
	fmt.Printf("ItemName: %s\n", ItemName)
}

func prepareQueryRows(db *sql.DB) {
	stmt, err := db.Prepare("select ItemCode,ItemName from OITM where ItemCode like ?")
	if err != nil {
		log.Fatal("Prepare failed:", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query("30-0109%")
	if err != nil {
		log.Fatal("this", err)
	}
	defer rows.Close()

	for rows.Next() {
		var a string
		var b string
		err = rows.Scan(&a, &b)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(a, b)
	}
}

package main

import (
	"database/sql"
	"log"
	"os"
	"pdf-esignature-server/db"
	"pdf-esignature-server/web"

	_ "github.com/lib/pq"
)

func main() {
	d, err := sql.Open("postgres", dataSource())
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.PDFSigningServer(db.NewDB(d), cors)
	err = app.Serve()
	log.Println("Error", err)
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "postgresql://" + host + ":5432/goxygen" +
		"?user=goxygen&sslmode=disable&password=" + pass
}

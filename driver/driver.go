package driver

import("os"
	"log"
	"github.com/lib/pq"
	"database/sql")

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL"))

	logFatal(err)
	log.Println(pgUrl)

	db, err = sql.Open("postgres",pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}

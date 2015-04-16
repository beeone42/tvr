package db_playlist

import "database/sql"

func OpenDb(driver, dataSourceName) {
		db, err := sql.Open(driver, dataSourceName)
		if err := nill {
			log.Fatal(err)
		}
}

func OpenDb(driver, dataSourceName) {
		db, err := sql.Close(driver, dataSourceName)
		if err := nill {
			log.Fatal(err)
		}
}

func ExecDb(cmd) {
		result , err := db.Exec(string)
		if err := nill {
			log.Fatal(err)
		}
}
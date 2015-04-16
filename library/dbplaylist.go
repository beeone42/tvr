package dbplaylist

import ("database/sql"
		"fmt"
)

func OpenDb(driver, dataSourceName) {
		db, err := sql.Open(driver, dataSourceName)
		if err := nill {
			log.Fatal(err)
		}
}

func CloseDb(driver, dataSourceName) {
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

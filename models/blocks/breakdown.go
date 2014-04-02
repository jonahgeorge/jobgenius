package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type BreakdownChart []Value

type Value struct {
	// Key   sql.NullInt64
	Title sql.NullString
	Value sql.NullInt64
}

func (b BreakdownChart) RetrieveById(db *sql.DB, id string) []Value {

	sql := `SELECT 
				nid, 
				tid
			FROM
				F_DAILY_BREAKDOWN
			WHERE
				F_DAILY_BREAKDOWN.iid = ?`

	var values []Value

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var value Value

		err = rows.Scan(&value.Title, &value.Value)
		if err != nil {
			log.Printf("%s", err)
		}

		values = append(values, value)
	}

	return values
}

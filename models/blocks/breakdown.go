package blocks

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type BreakdownChart []Value

type Value struct {
	// Key   sql.NullInt64
	Title sql.NullString
	Value sql.NullInt64
}

func (b BreakdownChart) RetrieveById(db *sql.DB, id string) []Value {

	sql := `
	SELECT 
		nid, 
		tid
	FROM
		Interviews_Daily_Breakdown	
	WHERE
		Interviews_Daily_Breakdown.iid = ?`

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

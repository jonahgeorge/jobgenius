package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type GroupworkChart struct {
	Solo  sql.NullInt64
	Group sql.NullInt64
}

func (g GroupworkChart) RetrieveById(db *sql.DB, id string) GroupworkChart {

	sql := `SELECT 
				vid AS 'Solo',
				(100 - vid) AS 'Group'
			FROM
				F_SOLO_GROUP
			WHERE
				F_SOLO_GROUP.iid = ?`

	var chart GroupworkChart

	row := db.QueryRow(sql, id)
	err := row.Scan(&chart.Solo, &chart.Group)
	if err != nil {
		log.Printf("%s", err)
	}

	return chart
}

package blocks

import (
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type GroupworkChart struct {
	Solo  *int
	Group *int
}

func (g GroupworkChart) RetrieveById(id string) GroupworkChart {

	sql := `
	SELECT 
		vid AS 'Solo',
		(100 - vid) AS 'Group'
	FROM
		Interviews_Groupwork
	WHERE
		Interviews_Groupwork.iid = ?`

	var chart GroupworkChart

	row := db.QueryRow(sql, id)
	err := row.Scan(&chart.Solo, &chart.Group)
	if err != nil {
		log.Printf("%s", err)
	}

	return chart
}

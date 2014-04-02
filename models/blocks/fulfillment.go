package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type FulfillmentChart struct {
	Development  sql.NullInt64
	Independence sql.NullInt64
	Impact       sql.NullInt64
	Personal     sql.NullInt64
}

func (f FulfillmentChart) RetrieveById(db *sql.DB, id string) FulfillmentChart {

	sql := `SELECT
				development,
				independence,
				impact,
				personal
			FROM
				F_FULFILLMENT
			WHERE
				F_FULFILLMENT.iid = ?`

	var chart FulfillmentChart

	row := db.QueryRow(sql, id)
	err := row.Scan(&chart.Development, &chart.Independence, &chart.Impact, &chart.Personal)
	if err != nil {
		log.Printf("%s", err)
	}

	return chart
}

// Consult Edward on how to average people with multiple industries
// Integral to how this function will correctly return values
func (f FulfillmentChart) RetrieveIndustryAverage(db *sql.DB, id string) FulfillmentChart {

	// sql := `SELECT
	// 			*
	// 		FROM
	// 			F_FULFILLMENT
	// 		WHERE
	// 			F_FULFILLMENT.iid = ?`

	// var chart FulfillmentChart

	// row := db.QueryRow(sql, id)
	// err := row.Scan(&chart.Development, &chart.Independence, &chart.Impact, &chart.Personal)
	// if err != nil {
	// 	log.Printf("%s", err)
	// }

	chart := FulfillmentChart{
		Development:  sql.NullInt64{Valid: true, Int64: 5},
		Independence: sql.NullInt64{Valid: true, Int64: 5},
		Impact:       sql.NullInt64{Valid: true, Int64: 5},
		Personal:     sql.NullInt64{Valid: true, Int64: 5},
	}

	return chart
}

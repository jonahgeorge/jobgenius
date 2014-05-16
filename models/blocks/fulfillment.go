package blocks

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type FulfillmentChart struct {
	Development  *int
	Independence *int
	Impact       *int
	Personal     *int
}

func (f FulfillmentChart) RetrieveById(db *sql.DB, id string) FulfillmentChart {

	sql := `
	SELECT
		development,
		independence,
		impact,
		personal
	FROM
		Interviews_Fulfillment
	WHERE
		Interviews_Fulfillment.iid = ?`

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

	average := 6

	chart := FulfillmentChart{
		Development:  &average,
		Independence: &average,
		Impact:       &average,
		Personal:     &average,
	}

	return chart
}

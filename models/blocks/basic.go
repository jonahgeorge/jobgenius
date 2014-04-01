package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type BasicBlock struct {
	Type            sql.NullString
	Sector          sql.NullString
	Industry        []Field
	Experience      sql.NullString
	Environment     []Field
	Salary          sql.NullString
	HoursPerWeek    sql.NullString
	WeekendsWorked  sql.NullString
	OvernightTravel sql.NullString
}

func (b BasicBlock) RetrieveById(db *sql.DB, id string) BasicBlock {

	sql := `SELECT
    			L_TYPE.value,
    			L_SECTOR.value,
    			L_EXPERIENCE.value,
    			L_SALARY.value,
    			L_HOURS_PER_WEEK.value,
    			L_WEEKENDS_WORKED.value,
    			L_OVERNIGHT_TRAVEL.value

    		FROM
    			C_INTERVIEW

            LEFT JOIN 
            	F_TYPE ON F_TYPE.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_TYPE ON L_TYPE.id = F_TYPE.vid
            
            LEFT JOIN 
            	F_SECTOR ON F_SECTOR.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_SECTOR ON L_SECTOR.id = F_SECTOR.vid
            
            LEFT JOIN 
            	F_EXPERIENCE ON F_EXPERIENCE.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_EXPERIENCE ON L_EXPERIENCE.id = F_EXPERIENCE.vid
            
            LEFT JOIN 
            	F_SALARY ON F_SALARY.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_SALARY ON L_SALARY.id = F_SALARY.vid
            
            LEFT JOIN 
            	F_HOURS_PER_WEEK ON F_HOURS_PER_WEEK.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_HOURS_PER_WEEK ON L_HOURS_PER_WEEK.id = F_HOURS_PER_WEEK.vid
            
            LEFT JOIN 
            	F_WEEKENDS_WORKED ON F_WEEKENDS_WORKED.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_WEEKENDS_WORKED ON L_WEEKENDS_WORKED.id = F_WEEKENDS_WORKED.vid
            
            LEFT JOIN 
            	F_OVERNIGHT_TRAVEL ON F_OVERNIGHT_TRAVEL.iid = C_INTERVIEW.id
            LEFT JOIN 
            	L_OVERNIGHT_TRAVEL ON L_OVERNIGHT_TRAVEL.id = F_OVERNIGHT_TRAVEL.vid

            WHERE
            	C_INTERVIEW.id = ?`

	var block BasicBlock

	row := db.QueryRow(sql, id)
	err := row.Scan(&block.Type, &block.Sector, &block.Experience, &block.Salary, &block.HoursPerWeek, &block.WeekendsWorked, &block.OvernightTravel)
	if err != nil {
		log.Printf("%s", err)
	}

	block.Industry = BasicBlock{}.RetrieveIndustry(db, id)
	block.Environment = BasicBlock{}.RetrieveEnvironment(db, id)

	return block
}

func (b BasicBlock) RetrieveIndustry(db *sql.DB, id string) []Field {

	sql := `SELECT
                L_INDUSTRY.id,
                L_INDUSTRY.value
            FROM
                F_INDUSTRY
            LEFT JOIN
                L_INDUSTRY ON L_INDUSTRY.id = F_INDUSTRY.vid
            WHERE
                F_INDUSTRY.iid = ?`

	var fields []Field

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var field Field

		err = rows.Scan(&field.Key, &field.Value)
		if err != nil {
			log.Printf("%s", err)
		}

		fields = append(fields, field)
	}
	return fields
}

func (b BasicBlock) RetrieveEnvironment(db *sql.DB, id string) []Field {

	sql := `SELECT
                L_ENVIRONMENT.id,
                L_ENVIRONMENT.value
            FROM
                F_ENVIRONMENT
            LEFT JOIN
                L_ENVIRONMENT ON L_ENVIRONMENT.id = F_ENVIRONMENT.vid
            WHERE
                F_ENVIRONMENT.iid = ?`

	var fields []Field

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var field Field

		err = rows.Scan(&field.Key, &field.Value)
		if err != nil {
			log.Printf("%s", err)
		}

		fields = append(fields, field)
	}
	return fields
}

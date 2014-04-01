package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type EducationBlock struct{}

type Degree struct {
	Degree        sql.NullString
	Concentration sql.NullString
	University    sql.NullString
	Year          sql.NullInt64
}

func (e EducationBlock) RetrieveById(db *sql.DB, id string) []Degree {

	sql := `SELECT  
				L_DEGREE.value as degree, 
				F_EDUCATION.concentration, 
				L_UNIVERSITY.value as university, 
				F_EDUCATION.year
            FROM
            	F_EDUCATION
            LEFT JOIN
                L_UNIVERSITY on L_UNIVERSITY.id = F_EDUCATION.university
            LEFT JOIN
            	L_DEGREE on F_EDUCATION.degree = L_DEGREE.id
            WHERE
            	F_EDUCATION.iid = ?`

	var degrees []Degree

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var degree Degree

		err = rows.Scan(&degree.Degree, &degree.Concentration, &degree.University, &degree.Year)
		if err != nil {
			log.Fatal(err)
		}

		degrees = append(degrees, degree)
	}

	return degrees
}

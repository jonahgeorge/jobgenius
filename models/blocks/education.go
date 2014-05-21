package blocks

import (
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type EducationBlock struct{}

type DegreeModel struct {
	Degree        *string
	Concentration *string
	University    *string
	Year          *int
}

func (e EducationBlock) RetrieveById(id string) []DegreeModel {

	sql := `
	SELECT  
		Interviews_Degree_Lookup.value as degree, 
		Interviews_Education.concentration, 
		Interviews_University_Lookup.value as university,
		Interviews_Education.year
	FROM
		Interviews_Education
	LEFT JOIN
		Interviews_University_Lookup on Interviews_University_Lookup.id = Interviews_Education.university
	LEFT JOIN
		Interviews_Degree_Lookup on Interviews_Education.degree = Interviews_Degree_Lookup.id
	WHERE
		Interviews_Education.iid = ?`

	var degrees []DegreeModel

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var degree DegreeModel

		err = rows.Scan(&degree.Degree, &degree.Concentration, &degree.University, &degree.Year)
		if err != nil {
			log.Fatal(err)
		}

		degrees = append(degrees, degree)
	}

	return degrees
}

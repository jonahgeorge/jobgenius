package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type RequirementsBlock struct {
	Certifications []Field
	Skills         []Field
}

type Field struct {
	Key   sql.NullInt64
	Value sql.NullString
}

func (r RequirementsBlock) Retrieve(db *sql.DB, id string) RequirementsBlock {

	rb := RequirementsBlock{
		Certifications: RequirementsBlock{}.RetrieveCertifications(db, id),
		Skills:         RequirementsBlock{}.RetrieveSkills(db, id),
	}

	return rb

}

func (r RequirementsBlock) RetrieveCertifications(db *sql.DB, id string) []Field {

	sql := `SELECT
				L_CERTIFICATION.id,
				L_CERTIFICATION.value
			FROM
				F_CERTIFICATION
			LEFT JOIN
				L_CERTIFICATION on L_CERTIFICATION.id = F_CERTIFICATION.vid
			WHERE
				F_CERTIFICATION.iid = ?`

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

func (r RequirementsBlock) RetrieveSkills(db *sql.DB, id string) []Field {

	sql := `SELECT
				F_SKILL.id,
				F_SKILL.value
			FROM
				F_SKILL
			WHERE
				F_SKILL.iid = ?`

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

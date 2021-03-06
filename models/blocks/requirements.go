package blocks

import (
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type RequirementsBlock struct {
	Certifications []Field
	Skills         []Field
}

type Field struct {
	Key   *int
	Value *string
}

func (r RequirementsBlock) Retrieve(id string) RequirementsBlock {

	rb := RequirementsBlock{
		Certifications: RequirementsBlock{}.RetrieveCertifications(id),
		Skills:         RequirementsBlock{}.RetrieveSkills(id),
	}

	return rb

}

func (r RequirementsBlock) RetrieveCertifications(id string) []Field {

	sql := `
	SELECT
		Interviews_Certification_Lookup.id,
		Interviews_Certification_Lookup.value
	FROM
		Interviews_Certification
	LEFT JOIN
		Interviews_Certification_Lookup on Interviews_Certification_Lookup.id = Interviews_Certification.vid
	WHERE
		Interviews_Certification.iid = ?`

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

func (r RequirementsBlock) RetrieveSkills(id string) []Field {

	sql := `
	SELECT
		Interviews_Skill.id,
		Interviews_Skill.value
	FROM
		Interviews_Skill
	WHERE
		Interviews_Skill.iid = ?`

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

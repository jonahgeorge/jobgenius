package blocks

import (
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type ToolsBlock struct {
	Skills []Field
	Tools  []Field
}

func (t ToolsBlock) Retrieve(id string) ToolsBlock {

	tb := ToolsBlock{
		Skills: ToolsBlock{}.RetrieveSkills(id),
		Tools:  ToolsBlock{}.RetrieveTools(id),
	}

	return tb
}

func (t ToolsBlock) RetrieveTools(id string) []Field {

	sql := `
	SELECT
		Interviews_Tools.id,
		Interviews_Tools.value
	FROM
		Interviews_Tools
	WHERE
		Interviews_Tools.iid = ?`

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

func (t ToolsBlock) RetrieveSkills(id string) []Field {

	sql := `
	SELECT
		Interviews_Soft_Skills_Lookup.id,
		Interviews_Soft_Skills_Lookup.value
	FROM
		Interviews_Soft_Skills	
	LEFT JOIN
		Interviews_Soft_Skills_Lookup ON Interviews_Soft_Skills_Lookup.id = Interviews_Soft_Skills.vid
	WHERE
		Interviews_Soft_Skills.iid = ?`

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

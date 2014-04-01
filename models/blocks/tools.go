package blocks

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type ToolsBlock struct {
	Skills []Field
	Tools  []Field
}

func (t ToolsBlock) Retrieve(db *sql.DB, id string) ToolsBlock {

	tb := ToolsBlock{
		Skills: ToolsBlock{}.RetrieveSkills(db, id),
		Tools:  ToolsBlock{}.RetrieveTools(db, id),
	}

	return tb
}

func (t ToolsBlock) RetrieveTools(db *sql.DB, id string) []Field {

	sql := `SELECT
				F_TOOL.id,
				F_TOOL.value
			FROM
				F_TOOL
			WHERE
			    F_TOOL.iid = ?`

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

func (t ToolsBlock) RetrieveSkills(db *sql.DB, id string) []Field {

	sql := `SELECT
				L_SOFT_SKILL.id,
				L_SOFT_SKILL.value
			FROM
				F_SOFT_SKILL
			LEFT JOIN
				L_SOFT_SKILL ON L_SOFT_SKILL.id = F_SOFT_SKILL.vid
			WHERE
			    F_SOFT_SKILL.iid = ?`

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

package models

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/models/blocks"
	"log"
)

type InterviewModel struct {
	Id       sql.NullInt64
	Name     sql.NullString
	Position sql.NullString
	AuthorId sql.NullInt64
	Author   sql.NullString
	Basic    BasicBlock
}

func (i InterviewModel) Create(db *sql.DB) error {
	return nil
}

func (i InterviewModel) RetrieveAll(db *sql.DB) []InterviewModel {

	var interviews []InterviewModel

	sql := `SELECT
				C_INTERVIEW.id, 
				C_INTERVIEW.name, 
				C_INTERVIEW.position, 
				C_USER.uid,
				C_USER.display_name
          	FROM 
          		C_INTERVIEW
            LEFT JOIN 
            	C_USER ON C_INTERVIEW.uid = C_USER.uid
          	WHERE 
          		C_INTERVIEW.published = 1`

	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var interview InterviewModel

		err = rows.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.AuthorId, &interview.Author)
		if err != nil {
			log.Printf("%s", err)
		}

		interviews = append(interviews, interview)
	}
	return interviews
}

func (i InterviewModel) RetrieveByAuthor(db *sql.DB, id int) []InterviewModel {
	var interviews []InterviewModel

	sql := `SELECT
				C_INTERVIEW.id, 
				C_INTERVIEW.name, 
				C_INTERVIEW.position, 
				C_USER.uid,
				C_USER.display_name
          	FROM 
          		C_INTERVIEW
            LEFT JOIN 
            	C_USER ON C_INTERVIEW.uid = C_USER.uid
          	WHERE 
          		C_INTERVIEW.published = 1
          	AND
          		C_USER.uid = ?`

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var interview InterviewModel
		err = rows.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.AuthorId, &interview.Author)
		if err != nil {
			log.Printf("%s", err)
		}
		interviews = append(interviews, interview)
	}

	return interviews
}

func (i InterviewModel) RetrieveById(db *sql.DB, id string) InterviewModel {

	sql := `SELECT
				C_INTERVIEW.id, 
				C_INTERVIEW.name, 
				C_INTERVIEW.position, 
				C_USER.uid,
				C_USER.display_name
          	FROM 
          		C_INTERVIEW
            LEFT JOIN 
            	C_USER ON C_INTERVIEW.uid = C_USER.uid
          	WHERE 
          		C_INTERVIEW.published = 1
          	AND
          		C_INTERVIEW.id = ?`

	var interview InterviewModel

	row := db.QueryRow(sql, id)
	err := row.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.AuthorId, &interview.Author)
	if err != nil {
		log.Printf("%s", err)
	}

	interview.Basic = BasicBlock{}.RetrieveById(db, id)

	return interview
}

func (i InterviewModel) Update(db *sql.DB) error {
	return nil
}

func (i InterviewModel) Delete(db *sql.DB) error {
	return nil
}

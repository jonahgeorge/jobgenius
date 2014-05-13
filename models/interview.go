package models

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/models/blocks"
)

type InterviewFactory struct {
}

func (i InterviewFactory) RetrieveAll(db *sql.DB) []InterviewModel {

	var interviews []InterviewModel

	sql := `SELECT
				C_INTERVIEW.id, 
				C_INTERVIEW.name, 
				C_INTERVIEW.position, 
				C_USER.uid,
				C_USER.display_name,
				C_USER.email_hash
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

		err = rows.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.AuthorId, &interview.Author, &interview.Picture)
		if err != nil {
			log.Printf("%s", err)
		}

		interviews = append(interviews, interview)
	}
	return interviews
}

func (i InterviewFactory) RetrieveByAuthor(db *sql.DB, id int) []InterviewModel {
	var interviews []InterviewModel

	sql := `SELECT
				C_INTERVIEW.id, 
				C_INTERVIEW.name, 
				C_INTERVIEW.position, 
				C_USER.uid,
				C_USER.display_name,
				C_USER.email_hash
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
		err = rows.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.AuthorId, &interview.Author, &interview.Picture)
		if err != nil {
			log.Printf("%s", err)
		}
		interviews = append(interviews, interview)
	}

	return interviews
}

func (i InterviewFactory) RetrieveById(db *sql.DB, id string) InterviewModel {

	sql := `SELECT
				C_INTERVIEW.id, 
				C_INTERVIEW.name, 
				C_INTERVIEW.position, 
				C_USER.uid,
				C_USER.display_name,
				C_USER.email_hash
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
	err := row.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.AuthorId, &interview.Author, &interview.Picture)
	if err != nil {
		log.Printf("%s", err)
	}

	interview.Basic = BasicBlock{}.RetrieveById(db, id)
	interview.Education = EducationBlock{}.RetrieveById(db, id)
	interview.Requirements = RequirementsBlock{}.Retrieve(db, id)
	interview.Tools = ToolsBlock{}.Retrieve(db, id)

	return interview
}

type InterviewModel struct {
	Id           sql.NullInt64
	Name         sql.NullString
	Position     sql.NullString
	AuthorId     sql.NullInt64
	Author       sql.NullString
	Picture      sql.NullString
	Basic        BasicBlock
	Education    []Degree
	Requirements RequirementsBlock
	Tools        ToolsBlock
}

func (i InterviewModel) Create(db *sql.DB) error {
	return nil
}

func (i InterviewModel) Update(db *sql.DB) error {
	return nil
}

func (i InterviewModel) Delete(db *sql.DB) error {
	return nil
}

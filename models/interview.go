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

	sql := `
	SELECT
		Interviews.id
	,	Interviews.name
	,	Interviews.position
	,	Users.uid
	,	Users.display_name
	FROM 
		Interviews
	LEFT JOIN 
		Users ON Interviews.uid = Users.uid
	WHERE 
		Interviews.published = 1`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var interview InterviewModel

		err = rows.Scan(
			&interview.Id, &interview.Name, &interview.Position,
			&interview.User.Id, &interview.User.DisplayName)

		if err != nil {
			log.Println(err)
		}

		interviews = append(interviews, interview)
	}

	return interviews
}

func (i InterviewFactory) RetrieveByAuthor(db *sql.DB, id int) []InterviewModel {
	var interviews []InterviewModel

	sql := `
	SELECT
		Interviews.id, 
		Interviews.name, 
		Interviews.position, 
		Users.uid,
		Users.display_name
	FROM 
		Interviews
	LEFT JOIN 
		Users ON Interviews.uid = Users.uid
	WHERE 
		Interviews.published = 1
	AND
		Users.uid = ?`

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var interview InterviewModel
		err = rows.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.User.Id, &interview.User.DisplayName)
		if err != nil {
			log.Printf("%s", err)
		}
		interviews = append(interviews, interview)
	}

	return interviews
}

func (i InterviewFactory) RetrieveById(db *sql.DB, id string) InterviewModel {

	sql := `
	SELECT
		Interviews.id, 
		Interviews.name, 
		Interviews.position, 
		Users.uid,
		Users.display_name
	FROM 
		Interviews
	LEFT JOIN 
		Users ON Interviews.uid = Users.uid
	WHERE 
		Interviews.published = 1
	AND
		Interviews.id = ?`

	var interview InterviewModel

	row := db.QueryRow(sql, id)
	err := row.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.User.Id, &interview.User.DisplayName)
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
	Id           int
	Name         *string
	Position     *string
	User         UserModel
	Picture      *string
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

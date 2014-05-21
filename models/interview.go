package models

import (
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/models/blocks"
)

type InterviewFactory struct {
}

func (i InterviewFactory) RetrieveAll() []InterviewModel {

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

func (i InterviewFactory) RetrieveByAuthor(id int) []InterviewModel {
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
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var interview InterviewModel
		err = rows.Scan(&interview.Id, &interview.Name, &interview.Position, &interview.User.Id, &interview.User.DisplayName)
		if err != nil {
			log.Println(err)
		}
		interviews = append(interviews, interview)
	}

	return interviews
}

func (i InterviewFactory) RetrieveById(id string) InterviewModel {

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
	err := row.Scan(
		&interview.Id, &interview.Name, &interview.Position,
		&interview.User.Id, &interview.User.DisplayName)
	if err != nil {
		log.Println(err)
	}

	interview.Basic = BasicBlock{}.RetrieveById(id)
	interview.Education = EducationBlock{}.RetrieveById(id)
	interview.Requirements = RequirementsBlock{}.Retrieve(id)
	interview.Tools = ToolsBlock{}.Retrieve(id)

	return interview
}

func (i InterviewFactory) Filter(title string) []InterviewModel {
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
		Interviews.published = 1
	AND
		Interviews.name LIKE ?
	OR 
		Interviews.position LIKE ?`

	rows, err := db.Query(sql, "%"+title+"%", "%"+title+"%")
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

type InterviewModel struct {
	Id           int
	Name         *string
	Position     *string
	User         UserModel
	Picture      *string
	Basic        BasicBlock
	Education    []DegreeModel
	Requirements RequirementsBlock
	Tools        ToolsBlock
	Comments     []CommentModel
}

func (i InterviewModel) Create() error {
	return nil
}

func (i InterviewModel) Update() error {
	return nil
}

func (i InterviewModel) Delete() error {
	return nil
}

func (i InterviewModel) AddComment() error {
	return nil
}

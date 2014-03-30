package models

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type InterviewModel struct {

	// Info InfoModel

	// Basic BasicModel
	// Education EducationModel
	// Requirements RequirementModel
	// Skills SkillsModel
	// Solo SoloModel
	// Tasks TasksModel

	Id      sql.NullInt64
	Author  sql.NullString
	Date    sql.NullString
	Title   sql.NullString
	Content string
}

func (i InterviewModel) Create(db *sql.DB) error {
	return nil
}

func (i InterviewModel) RetrieveAll(db *sql.DB) ([]InterviewModel, error) {
	var teasers []InterviewModel

	sql := `SELECT U.display_name, A.aid, A.title, A.body, A.timestamp
          FROM C_ARTICLE AS A
            LEFT JOIN C_USER AS U ON A.uid = U.uid
          WHERE A.published = 1`

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i InterviewModel
		err = rows.Scan(&i.Author, &i.Id, &i.Title, &i.Content, &i.Date)
		if err != nil {
			log.Fatal(err)
		}
		teasers = append(teasers, i)
	}
	return teasers, err
}

func (i InterviewModel) RetrieveByAuthor(db *sql.DB, id int) ([]InterviewModel, error) {
	var interviews []InterviewModel

	sql := fmt.Sprintf("SELECT U.display_name, A.aid, A.title, A.body, A.timestamp FROM C_ARTICLE AS A LEFT JOIN C_USER AS U ON A.uid = U.uid WHERE A.uid = %d", id)

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i InterviewModel
		err = rows.Scan(&i.Author, &i.Id, &i.Title, &i.Content, &i.Date)
		if err != nil {
			log.Fatal(err)
		}
		interviews = append(interviews, i)
	}

	return interviews, err
}

func (i InterviewModel) RetrieveOne(db *sql.DB, id string) (InterviewModel, error) {
	return InterviewModel{}, nil
}

func (i InterviewModel) Update(db *sql.DB) error {
	return nil
}

func (i InterviewModel) Delete(db *sql.DB) error {
	return nil
}

package blocks

import (
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type BasicBlock struct {
	Type            Field
	Sector          Field
	Industry        []Field
	Experience      Field
	Environment     []Field
	Salary          Field
	HoursPerWeek    Field
	WeekendsWorked  Field
	OvernightTravel Field
}

func (b BasicBlock) RetrieveById(id string) BasicBlock {

	sql := `
	SELECT
		Interviews_Type_Lookup.id,
		Interviews_Type_Lookup.value,

		Interviews_Sector_Lookup.id,
		Interviews_Sector_Lookup.value,

		Interviews_Experience_Lookup.id,
		Interviews_Experience_Lookup.value,

		Interviews_Salary_Lookup.id,
		Interviews_Salary_Lookup.value,

		Interviews_Hours_Per_Week_Lookup.id,
		Interviews_Hours_Per_Week_Lookup.value,

		Interviews_Weekends_Worked_Lookup.id,
		Interviews_Weekends_Worked_Lookup.value,

		Interviews_Overnight_Travel_Lookup.id,
		Interviews_Overnight_Travel_Lookup.value

	FROM Interviews

	LEFT JOIN Interviews_Type ON Interviews_Type.iid = Interviews.id
	LEFT JOIN Interviews_Type_Lookup ON Interviews_Type_Lookup.id = Interviews_Type.vid

	LEFT JOIN Interviews_Sector ON Interviews_Sector.iid = Interviews.id
	LEFT JOIN Interviews_Sector_Lookup ON Interviews_Sector_Lookup.id = Interviews_Sector.vid

	LEFT JOIN Interviews_Experience ON Interviews_Experience.iid = Interviews.id
	LEFT JOIN Interviews_Experience_Lookup ON Interviews_Experience_Lookup.id = Interviews_Experience.vid

	LEFT JOIN Interviews_Salary ON Interviews_Salary.iid = Interviews.id
	LEFT JOIN Interviews_Salary_Lookup ON Interviews_Salary_Lookup.id = Interviews_Salary.vid

	LEFT JOIN Interviews_Hours_Per_Week ON Interviews_Hours_Per_Week.iid = Interviews.id
	LEFT JOIN Interviews_Hours_Per_Week_Lookup ON Interviews_Hours_Per_Week_Lookup.id = Interviews_Hours_Per_Week.vid

	LEFT JOIN Interviews_Weekends_Worked ON Interviews_Weekends_Worked.iid = Interviews.id
	LEFT JOIN Interviews_Weekends_Worked_Lookup ON Interviews_Weekends_Worked_Lookup.id = Interviews_Weekends_Worked.vid

	LEFT JOIN Interviews_Overnight_Travel ON Interviews_Overnight_Travel.iid = Interviews.id
	LEFT JOIN Interviews_Overnight_Travel_Lookup ON Interviews_Overnight_Travel_Lookup.id = Interviews_Overnight_Travel.vid
	
	WHERE Interviews.id = ?`

	var block BasicBlock

	row := db.QueryRow(sql, id)
	err := row.Scan(
		&block.Type.Key, &block.Type.Value, &block.Sector.Key, &block.Sector.Value,
		&block.Experience.Key, &block.Experience.Value, &block.Salary.Key, &block.Salary.Value,
		&block.HoursPerWeek.Key, &block.HoursPerWeek.Value, &block.WeekendsWorked.Key, &block.WeekendsWorked.Value,
		&block.OvernightTravel.Key, &block.OvernightTravel.Value)
	if err != nil {
		log.Println(err)
	}

	block.Industry = BasicBlock{}.RetrieveIndustry(id)
	block.Environment = BasicBlock{}.RetrieveEnvironment(id)

	return block
}

func (b BasicBlock) RetrieveIndustry(id string) []Field {

	sql := `
	SELECT
		Interviews_Industry_Lookup.id,
		Interviews_Industry_Lookup.value
	FROM
		Interviews_Industry
	LEFT JOIN
		Interviews_Industry_Lookup ON Interviews_Industry_Lookup.id = Interviews_Industry.vid
	WHERE
		Interviews_Industry.iid = ?`

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

func (b BasicBlock) RetrieveEnvironment(id string) []Field {

	sql := `
	SELECT
		Interviews_Environment_Lookup.id,
		Interviews_Environment_Lookup.value
	FROM
		Interviews_Environment
	LEFT JOIN
		Interviews_Environment_Lookup ON Interviews_Environment_Lookup.id = Interviews_Environment.vid
	WHERE
		Interviews_Environment.iid = ?`

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

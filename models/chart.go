package models

import (
	"database/sql"
)

type DailyTasksChart struct {
	Tasks []ValueWithName
}

type MotivationChart struct {
	Development  DoubleValue
	Independence DoubleValue
	Impact       DoubleValue
	Personal     DoubleValue
}

type TeamworkChart struct {
	Value DoubleValue
}

type ValueWithName struct {
	Title sql.NullString
	Value sql.NullInt64
}

type DoubleValue struct {
	Value   sql.NullInt64
	Average sql.NullInt64
}

// Functions go here

package model

type Student struct {
	StudentID   int    `json:"studentid" db:"StudentID"`
	Name        string `json:"name" db:"Name"`
	Description string `json:"description" db:"Description"`
}

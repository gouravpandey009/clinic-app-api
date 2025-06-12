package models

type Patient struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Age       int    `db:"age" json:"age"`
	Gender    string `db:"gender" json:"gender"`
	Diagnosis string `db:"diagnosis" json:"diagnosis"`
	CreatedBy int    `db:"created_by" json:"created_by"`
}

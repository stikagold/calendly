package models

type Pack struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Size int    `db:"size"`
}

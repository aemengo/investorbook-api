package models

type Investor struct {
	ID               int    `db:"id"`
	Name             string `db:"name"`
	ConnectionDegree string
}

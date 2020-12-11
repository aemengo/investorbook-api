package db

import "github.com/aemengo/investorbook-api/models"

type DB struct {}

func New() (*DB, error) {
	return &DB{}, nil
}

//- GET /investors/:investorId/connections
func (d *DB) Connections(id int) ([]models.Investor, error)  {
	return nil, nil
}

//- GET /investors/:investorId/mutual/:otherInvestorId
func (d *DB) MutualConnections(id int, otherId int) ([]models.Investor, error)  {
	return nil, nil
}

//- GET /investors/:investorId/search?q={searchString}
func (d *DB) Search(id int, query string) ([]models.Investor, error) {
	return nil, nil
}
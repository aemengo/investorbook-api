package db

import (
	"github.com/aemengo/investorbook-api/models"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func New(databaseURI string) (*DB, error) {
	db, err := sqlx.Connect("pgx", databaseURI)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

//- GET /investors/:investorId/connections
func (d *DB) Connections(id int) ([]models.Investor, error) {
	investors := []models.Investor{}
	err := d.db.Select(&investors, `
select investor.id, investor.name
from investor join investment on investor.id = investment.investor_id
where not investor.id = $1
and investment.company_id in (select company_id
             from investment
             where investor_id = $1);`, id)
	return investors, err
}

//- GET /investors/:investorId/mutual/:otherInvestorId
func (d *DB) MutualConnections(id int, otherId int) ([]models.Investor, error) {
	connections, err := d.Connections(id)
	if err != nil {
		return nil, err
	}

	otherConnections, err := d.Connections(otherId)
	if err != nil {
		return nil, err
	}

	return mutualConnections(connections, otherConnections), nil
}

//- GET /investors/:investorId/search?q={searchString}
func (d *DB) Search(id int, query string) ([]models.Investor, error) {
	investors := []models.Investor{}
	err := d.db.Select(&investors, "select id, name from investor where lower(name) LIKE ?", "%"+query+"%")
	return investors, err
}

func mutualConnections(connections []models.Investor, otherConnections []models.Investor) []models.Investor {
	ids := map[int]bool{}
	for _, con := range connections {
		ids[con.ID] = true
	}

	result := []models.Investor{}
	for _, con := range otherConnections{
		if _, ok := ids[con.ID]; ok {
			result = append(result, con)
		}
	}
	return result
}

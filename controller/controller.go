package controller

import (
	"github.com/aemengo/investorbook-api/db"
	"github.com/aemengo/investorbook-api/models"
)

type Controller struct {
	database *db.DB
}

func New(database *db.DB) *Controller{
	return &Controller{
		database: database,
	}
}

func (c *Controller) Connections(id int) ([]models.Investor, error)  {
	return c.database.Connections(id)
}

func (c *Controller) MutualConnections(id int, otherId int) ([]models.Investor, error)  {
	return c.database.MutualConnections(id, otherId)
}

func (c *Controller) Search(id int, query string) ([]models.Investor, error) {
	return c.database.Search(id, query)
}
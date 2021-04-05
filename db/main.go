package db

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// Config
type Config struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	client       driver.Client
	database     driver.Database
}

// Initdb
func (c *Config) Initdb() (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{c.Host},
	})
	if err != nil {
		return nil, err
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication(
			c.Username, c.Password),
	})
	if err != nil {
		return nil, err
	}
	c.client = client
	createDbUnlessExists := func() (driver.Database, error) {
		exist, err := client.DatabaseExists(context.TODO(), c.DatabaseName)
		if err != nil {
			return nil, err
		}
		if exist {
			db, err := client.Database(context.TODO(), c.DatabaseName)
			if err != nil {
				return nil, err
			}
			return db, nil
		}
		return client.CreateDatabase(context.TODO(), c.DatabaseName, nil)
	}
	database, err := createDbUnlessExists()
	if err != nil {
		return nil, err
	}
	c.database = database
	return database, nil
}

// Collection
func (c *Config) Collection(name string) (driver.Collection, error) {
	exists, err := c.database.CollectionExists(context.TODO(), name)
	if err != nil {
		return nil, err
	}
	if exists {
		col, err := c.database.Collection(context.TODO(), name)
		if err != nil {
			return nil, err
		}
		return col, nil
	}
	c.database.CreateCollection(context.TODO(), name, nil)
	col, err := c.database.Collection(context.TODO(), name)
	if err != nil {
		return nil, err
	}
	return col, nil
}

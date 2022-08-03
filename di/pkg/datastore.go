package pkg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"syreclabs.com/go/faker"
)

type DataStore struct {
	db *sql.DB
}

func (ds *DataStore) GetCities() ([]*City, error) {
	var cities []*City

	rows, err := ds.db.Query("SELECT id, name FROM cities")
	if err != nil {
		return cities, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)

		err := rows.Scan(&id, &name)
		if err != nil {
			return cities, err
		}
		city := NewCity(id, name)
		cities = append(cities, city)
	}

	return cities, nil
}

func (ds *DataStore) init() error {
	_, err := ds.db.Exec("CREATE TABLE cities (id integer, name varchar(64))")
	if err != nil {
		return err
	}

	sqlStmt := `INSERT INTO cities (id, name) VALUES ($1, $2)`
	for i := 1; i < 4; i++ {
		_, err := ds.db.Exec(sqlStmt, i, faker.Address().City())
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDataStore(db *sql.DB) (*DataStore, error) {
	datastore := DataStore{db: db}

	err := datastore.init()
	if err != nil {
		return nil, err
	}

	return &datastore, nil
}

func OpenConnection(settings *Settings) (*sql.DB, error) {
	return sql.Open("sqlite3", settings.DSN)
}

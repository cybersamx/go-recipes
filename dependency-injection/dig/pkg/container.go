package pkg

import (
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	err := c.Provide(NewSettings)
	if err != nil {
		return nil, err
	}

	err = c.Provide(OpenConnection)
	if err != nil {
		return nil, err
	}

	err = c.Provide(NewDataStore)
	if err != nil {
		return nil, err
	}

	err = c.Provide(NewHTTPServer)
	if err != nil {
		return nil, err
	}

	return c, nil
}

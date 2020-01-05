package pkg

import (
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	// Alternatively if we want to inject something different from
	// NewSettings, we can do the following:
	//
	//data := []byte("port: 7000\ndsn: \":memory:\"")
	//
	//err := c.Provide(func() *Settings {
	//	settings := Settings{}
	//	err := yaml.Unmarshal(data, &settings)
	//	if err != nil {
	//		return &Settings{
	//			Port: 3000,
	//			DSN:  ":memory:",
	//		}
	//	}
	//
	//	return &settings
	//})

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

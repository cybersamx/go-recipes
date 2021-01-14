package services

import (
	"github.com/cybersamx/go-recipes/microservice/simple/models"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type StateService interface {
	GetStates() []*models.State
	GetState(abbreviation string) (*models.State, error)
}

type StateServiceImpl struct{}

var (
	states []*models.State
)

func loadYAML(reader io.Reader) ([]*models.State, error) {
	var parsedStates []*models.State

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &parsedStates); err != nil {
		return nil, err
	}

	return parsedStates, nil
}

func init() {
	file, err := os.Open("states.yaml")
	if err != nil {
		panic(err)
	}
	states, err = loadYAML(file)
	if err != nil {
		panic(err)
	}
}

func (ss StateServiceImpl) GetStates() []*models.State {
	return states
}

func (ss StateServiceImpl) GetState(abbreviation string) (*models.State, error) {
	for _, state := range states {
		if state.Abbreviation == strings.ToUpper(abbreviation) {
			return state, nil
		}
	}

	return nil, newStateNotFoundError(abbreviation)
}

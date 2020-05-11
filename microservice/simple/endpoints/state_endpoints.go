package endpoints

import (
	"context"
	"errors"
	"github.com/cybersamx/go-recipes/microservice/simple/models"
	"github.com/cybersamx/go-recipes/microservice/simple/services"
	"github.com/go-kit/kit/endpoint"
)

var (
	ErrWrongRequestType = errors.New("wrong request type")
)

// --- Requests and Responses ---

type GetStateRequest struct {
	Abbreviation string `json:"abbreviation"`
}

type StateResponse struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	Population   int    `json:"population"`
}

// --- Endpoints ---

func stateToStateRes(state *models.State) *StateResponse {
	return &StateResponse{
		Abbreviation: state.Abbreviation,
		Name: state.Name,
		Population: state.Population,
	}
}

func GetStatesEndpoint(svc services.StateService) endpoint.Endpoint {
	return func (_ context.Context, request interface{}) (interface{}, error) {
		var statesRes []*StateResponse
		for _, state := range svc.GetStates() {
			statesRes = append(statesRes, stateToStateRes(state))
		}

		return statesRes, nil
	}
}

func GetStateEndpoint(svc services.StateService) endpoint.Endpoint {
	return func (_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetStateRequest)
		if !ok {
			return nil, ErrWrongRequestType
		}
		state, err := svc.GetState(req.Abbreviation)
		if err != nil {
			return StateResponse{}, err
		}

		return stateToStateRes(state), nil
	}
}

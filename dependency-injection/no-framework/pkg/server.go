package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPServer struct {
	settings *Settings
	dataStore *DataStore
}

func (hs *HTTPServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hs.citiesHandlerFunc())

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", hs.settings.HTTPPort),
		Handler: mux,
	}

	return (&httpServer).ListenAndServe()
}

func (hs *HTTPServer) citiesHandlerFunc() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, fmt.Sprintf("doesn't support method %s", r.Method), http.StatusMethodNotAllowed)
			return
		}

		cities, err := hs.dataStore.GetCities()
		if err != nil {
			http.Error(w, "can't produce json payload", http.StatusInternalServerError)
			return
		}
		payload, err := json.Marshal(cities)
		if err != nil {
			http.Error(w, "can't produce json payload", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(payload)
		if err != nil {
			http.Error(w, "can't produce json payload", http.StatusInternalServerError)
			return
		}
	}
}

func NewHTTPServer(settings *Settings, datastore *DataStore) *HTTPServer {
	return &HTTPServer{
		settings:  settings,
		dataStore: datastore,
	}
}

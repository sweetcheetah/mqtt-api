package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRootPath(t *testing.T) {
	t.Run("Root Path", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerRoot)
		handler.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestBusOneTempPath(t *testing.T) {
	t.Run("bus-1-temp Path", func(t *testing.T) {
		router := mux.NewRouter()
		srv := NewServer(router)

		req, err := http.NewRequest(http.MethodGet, "/bus-1-temp", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		srv.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

}

func TestBusOneHumidityPath(t *testing.T) {
	t.Run("bus-1-humidity Path", func(t *testing.T) {
		router := mux.NewRouter()
		srv := NewServer(router)

		req, err := http.NewRequest(http.MethodGet, "/bus-1-humidity", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		srv.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

}

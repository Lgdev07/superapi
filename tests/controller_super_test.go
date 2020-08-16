package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lgdev07/superapi/api/models"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateSuperSuccess(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	inputJSON := `{"name": "Batman"}`

	req, err := http.NewRequest(http.MethodPost, "/supers", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CreateSuper)
	handler.ServeHTTP(rr, req)

	responseInterface := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusCreated)
	assert.Equal(t, responseInterface["name"], "Batman")
	assert.Equal(t, responseInterface["power"], float64(63))

}

func TestCreateSuperNotFound(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	inputJSON := `{"name": "Not a Super"}`

	req, err := http.NewRequest(http.MethodPost, "/supers", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CreateSuper)
	handler.ServeHTTP(rr, req)

	responseInterface := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, responseInterface["error"], "No Super found with this name")

}

func TestCreateSuperAlreadyRegistered(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	inputJSON := `{"name": "Batman"}`

	req, err := http.NewRequest(http.MethodPost, "/supers", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CreateSuper)
	handler.ServeHTTP(rr, req)

	responseInterface := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusCreated)
	assert.Equal(t, responseInterface["name"], "Batman")
	assert.Equal(t, responseInterface["power"], float64(63))

	inputJSON = `{"name": "Batman"}`

	req, err = http.NewRequest(http.MethodPost, "/supers", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(app.CreateSuper)
	handler.ServeHTTP(rr, req)

	responseInterface = make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, responseInterface["error"], "There is Already a Super With That Same Name Registered")

}

func TestDeleteSuper(t *testing.T) {
	var params map[string]string

	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	super, _ := seedOneSuper()

	supers, err := models.FindSupers(app.DB, params)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*supers), 1)

	req, err := http.NewRequest(http.MethodDelete, "/supers/"+super.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/supers/{id}", app.DeleteSuper)
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	supers, err = models.FindSupers(app.DB, params)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*supers), 0)
}

func TestShowSuper(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	super, _ := seedOneSuper()

	req, err := http.NewRequest(http.MethodGet, "/supers/"+super.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/supers/{id}", app.ShowSuper)
	router.ServeHTTP(rr, req)

	responseInterface := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, responseInterface["name"], "Super Test")
	assert.Equal(t, responseInterface["power"], float64(50))
}

func TestShowSuperWrongID(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	_, _ = seedOneSuper()

	randomUUID := uuid.NewV4()

	req, err := http.NewRequest(http.MethodGet, "/supers/"+randomUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/supers/{id}", app.ShowSuper)
	router.ServeHTTP(rr, req)

	responseInterface := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, responseInterface["error"], "record not found")
}

func TestListSupers(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	_, _ = seedOneSuper()

	req, err := http.NewRequest(http.MethodGet, "/supers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.ListSupers)
	handler.ServeHTTP(rr, req)

	var responseInterface []map[string]interface{}

	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, responseInterface[0]["name"], "Super Test")
	assert.Equal(t, responseInterface[0]["power"], float64(50))
}

func TestListSupersByName(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	_, _ = seedOneSuper()

	// When the query param is incorrect it should return nothing
	req, err := http.NewRequest(http.MethodGet, "/supers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("name", "Wrong Name")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.ListSupers)
	handler.ServeHTTP(rr, req)

	var responseInterface []map[string]interface{}

	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(responseInterface), 0)

	// When the query param is correct it should return the super
	req, err = http.NewRequest(http.MethodGet, "/supers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q = req.URL.Query()
	q.Add("name", "Super Test")

	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(app.ListSupers)
	handler.ServeHTTP(rr, req)

	var response2Interface []map[string]interface{}

	err = json.Unmarshal([]byte(rr.Body.String()), &response2Interface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(response2Interface), 1)
	assert.Equal(t, response2Interface[0]["name"], "Super Test")
	assert.Equal(t, response2Interface[0]["power"], float64(50))

}

func TestListSupersByAlignment(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	_, _ = seedOneSuper()

	// When the query param is correct it should return the super
	req, err := http.NewRequest(http.MethodGet, "/supers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("alignment", "good")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.ListSupers)
	handler.ServeHTTP(rr, req)

	var responseInterface []map[string]interface{}

	err = json.Unmarshal([]byte(rr.Body.String()), &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(responseInterface), 1)

	// When the query param is incorrect it should return the error
	req, err = http.NewRequest(http.MethodGet, "/supers", nil)
	if err != nil {
		t.Fatal(err)
	}

	q = req.URL.Query()
	q.Add("alignment", "wrongalignment")

	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(app.ListSupers)
	handler.ServeHTTP(rr, req)

	var response2Interface map[string]interface{}

	err = json.Unmarshal([]byte(rr.Body.String()), &response2Interface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, response2Interface["error"], "alignment must be good or bad")

}

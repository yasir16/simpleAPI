package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user").Methods("GET")
	return router
}
func TestGetUser(t *testing.T) {
	// server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// 	// Test request parameters
	// 	equals(t, req.URL.String(), "/some/path")
	// 	// Send response to be tested
	// 	rw.Write([]byte(`OK`))
	// }))

	// user := &structs.User{}
	// jsonPerson, _ := json.Marshal(user)
	// request, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonPerson))
	// response := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/product/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Product not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

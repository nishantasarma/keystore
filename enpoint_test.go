package main
import (

	"testing"
	"net/http/httptest"
	"net/http"
	endpoint "example/keystore/endpoints"
	"bytes"
)


func TestGetKeys(t *testing.T) {

	req, err := http.NewRequest("GET", "/get/abc-1", nil)

	if err != nil {

		t.Fatal(err)

	}
	
	rr := httptest.NewRecorder()
	teststore := endpoint.NewstoreHandlers()
	handler := http.HandlerFunc(teststore.Getkeys)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


	// Check the response body 
	expected := `"this"`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}



func TestSearchKeysPrefix(t *testing.T) {

	req, err := http.NewRequest("GET", "/search", nil)
	
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("prefix", "abc")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	teststore := endpoint.NewstoreHandlers()
	handler := http.HandlerFunc(teststore.Searchkeys)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	
	expected := `["abc-1","abc-2"]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	
}


func TestSearchKeysSufix(t *testing.T) {

	req, err := http.NewRequest("GET", "/search", nil)
	
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("suffix", "-1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	teststore := endpoint.NewstoreHandlers()
	handler := http.HandlerFunc(teststore.Searchkeys)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	
	expected := `["abc-1","xyz-1"]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	
}


func TestSetKeys(t *testing.T){

	var jsonStr = []byte(`{"id":"test"}`)

	req, err := http.NewRequest("POST", "/SetKeys", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	teststore := endpoint.NewstoreHandlers()
	handler := http.HandlerFunc(teststore.Setkeys)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"abc-1":"this","abc-2":"is","id":"test","xyz-1":"a","xyz-2":"test"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
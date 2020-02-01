package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {

	req, err := http.NewRequest("GET", "/companies", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAll)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Chamada retornou um código incorreto de status: retorno %v esperado %v",
			status, http.StatusOK)
	}

	// Checagem de corpo de retorno
	cantFindData := `{"message": "Can't find any Companies on Database"}`
	if rr.Body.String() == cantFindData {
		t.Errorf("Chamada não consegiu encontrar dados no server : %v",
			rr.Body.String())
	}
}

func TestPost(t *testing.T) {

	req, err := http.NewRequest("POST", "/companies/load", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Chamada retornou um código incorreto de status: retorno %v esperado %v",
			status, http.StatusOK)
	}
}

func TestGet(t *testing.T) {

	req, err := http.NewRequest("GET", "/companies/find/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Chamada retornou um código incorreto de status: retorno %v esperado %v",
			status, http.StatusOK)
	}

	// Checagem de corpo de retorno zip em branco e name "espaço"
	expected := `[{"_id":25,"name":"EPICBOARDSHOP BRANCH","zip":"","website":""}]`
	if rr.Body.String() != expected {
		t.Errorf("Chamada retornou corpo de dados diferente do esperado: retorno %v esperado %v",
			rr.Body.String(), expected)
	}
}

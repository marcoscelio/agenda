package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"../models"
)

type MockContactDao struct {
}

func (m *MockContactDao) GetAll() ([]models.Contact, error) {
	var contacts = []models.Contact{
		models.Contact{
			ID: "123", Name: "Marcos", Email: "marcos@email.com", Active: true,
		},
		models.Contact{
			ID: "456", Name: "Cecilia", Email: "cecilia@email.com", Active: true,
		},
		models.Contact{
			ID: "789", Name: "Theo", Email: "theo@email.com", Active: true,
		},
	}
	return contacts, nil
}

func (m *MockContactDao) GetByID(id string) (models.Contact, error) {
	var contact = models.Contact{Name: "Marcos", Email: "marcos@email.com", Active: true}
	return contact, nil
}

func (m *MockContactDao) FindByName(name string) ([]models.Contact, error) {
	var contacts = []models.Contact{
		models.Contact{
			ID: "123", Name: "Marcos", Email: "marcos@email.com", Active: true,
		},
		models.Contact{
			ID: "456", Name: "Cecilia", Email: "cecilia@email.com", Active: true,
		},
		models.Contact{
			ID: "789", Name: "Theo", Email: "theo@email.com", Active: true,
		},
	}
	return contacts, nil
}

func (m *MockContactDao) Create(contact models.Contact) error {
	return nil
}

func (m *MockContactDao) Delete(id string) error {
	return nil
}

func (m *MockContactDao) Update(id string, contact models.Contact) error {
	return nil
}

func TestContactRouterGetByID(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}

	var contactRouter = ContactService{Dao: &MockContactDao{}}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contactRouter.GetByID)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"","name":"Marcos","email":"marcos@email.com","active":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestContactRouterGetAll(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}

	var contactRouter = ContactService{Dao: &MockContactDao{}}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contactRouter.GetAll)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var contacts []interface{}
	result := rr.Body.String()

	err = json.Unmarshal([]byte(result), &contacts)

	if len(contacts) != 3 {
		t.Errorf("handler returned wrong list elements: got %v want %v",
			len(contacts), 3)
	}

}

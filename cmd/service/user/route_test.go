package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rohitpandeydev/go_learn_project1/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUseStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalide", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "pandey",
			Email:     "",
			Password:  "addjf",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUseStore struct{}

func (m *mockUseStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUseStore) GetUserbyId(ID int) (*types.User, error) {
	return nil, nil
}

func (m *mockUseStore) CreateUser(types.User) error {
	return nil
}

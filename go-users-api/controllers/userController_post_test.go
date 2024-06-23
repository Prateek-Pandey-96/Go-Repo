package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-users-api/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user", CreateUser(db))

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	t.Run("User created successfully", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id;`)).
			WithArgs("johndoe", "johndoe@example.com").
			WillReturnRows(rows)
		user := &models.UserRequest{
			Username: "johndoe",
			Email:    "johndoe@example.com",
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling JSON", err)
		}

		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("User creation failed due to bad request", func(t *testing.T) {
		type WrongUser struct {
			Name string
		}
		// this will fail because models.UserRequest is required
		user := &WrongUser{
			Name: "johndoe",
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling JSON", err)
		}

		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

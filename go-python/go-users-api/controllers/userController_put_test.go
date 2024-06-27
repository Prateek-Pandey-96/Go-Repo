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

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/user/:id", UpdateUser(db))

	t.Run("User updated successfully", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectQuery(regexp.QuoteMeta(`UPDATE users SET username = $1 where id = $2 RETURNING id;`)).
			WithArgs("updatedjohndoe", "1").
			WillReturnRows(rows)
		user := &models.UserRequest{
			Username: "updatedjohndoe",
			Email:    "johndoe@example.com",
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling JSON", err)
		}

		req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(jsonData))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("User update failed if bad request", func(t *testing.T) {
		type WrongUser struct {
			Name string
		}
		user := &WrongUser{
			Name: "johndoe",
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling JSON", err)
		}

		req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(jsonData))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

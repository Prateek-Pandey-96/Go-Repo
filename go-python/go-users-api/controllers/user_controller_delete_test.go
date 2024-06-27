package controllers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/user/:id", DeleteUser(db))

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	t.Run("User exists and is deleted", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM users where id = $1 RETURNING id;`)).
			WithArgs("1").
			WillReturnRows(rows)
		req, _ := http.NewRequest("DELETE", "/user/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("User does not exist", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM users where id = $1 RETURNING id;`)).
			WithArgs("2").
			WillReturnError(sql.ErrNoRows)

		req, _ := http.NewRequest("DELETE", "/user/2", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

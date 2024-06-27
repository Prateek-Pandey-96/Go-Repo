package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	created_at := time.Now().Truncate(time.Second)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:id", GetUser(db))

	rows := sqlmock.NewRows([]string{"id", "username", "email", "created_at"}).
		AddRow(1, "johndoe", "johndoe@example.com", created_at)
	t.Run("User exists", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * from users where id = $1;`)).
			WithArgs("1").
			WillReturnRows(rows)
		req, _ := http.NewRequest("GET", "/user/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedJSON := fmt.Sprintf(`{"user":{"Id":1,"Username":"johndoe","Email":"johndoe@example.com","Created_at":"%s"}}`, created_at.Format(time.RFC3339))
		assert.JSONEq(t, expectedJSON, w.Body.String())

	})

	t.Run("UserId doesn't exist", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * from users where id = $1;`)).
			WithArgs("2").
			WillReturnError(sql.ErrNoRows)

		req, _ := http.NewRequest("GET", "/user/2", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

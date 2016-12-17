package services_test

import (
	"testing"
	"time"

	"github.com/replicatedcom/gin-example/models"
	. "github.com/replicatedcom/gin-example/services"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserCreate(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	user := &models.User{
		Email: "ethan@replicated.com",
	}
	userPassword := "testing123"

	mock.ExpectBegin()
	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(user.Email, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	columns := []string{"email", "password", "created_at", "updated_at", "password_updated_at"}
	now := time.Now()
	row1 := sqlmock.NewRows(columns).
		AddRow(user.Email, userPassword, now, now, now)
	mock.
		ExpectQuery(`SELECT .+ FROM users`).
		WithArgs(user.Email).
		WillReturnRows(row1)
	mock.ExpectCommit()

	userService := User(db)
	err = userService.Create(user, userPassword)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, user.Email, "ethan@replicated.com")
	assert.NotEmpty(t, user.PasswordHash) // TODO
	assert.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), user.UpdatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), user.PasswordUpdatedAt, time.Second)
}

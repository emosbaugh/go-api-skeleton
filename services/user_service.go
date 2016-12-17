package services

import (
	"github.com/replicatedcom/gin-example/db/queries"
	"github.com/replicatedcom/gin-example/models"

	"github.com/jmoiron/sqlx"
	"github.com/nleof/goyesql"
	passlib "gopkg.in/hlandau/passlib.v1"
)

var (
	userQueries = goyesql.MustParseBytes(queries.MustAsset("users.sql"))
)

type IUserService interface {
	Get(email string) (*models.User, error)
	List() ([]*models.User, error)
	Create(user *models.User, password string) error
	UpdatePassword(user *models.User, password string) error
}

type UserService struct {
	DB *sqlx.DB
}

func User(db *sqlx.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) Get(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Get(&user, userQueries["get"], email)
	return &user, err
}

func (s *UserService) List() ([]*models.User, error) {
	var users []*models.User
	err := s.DB.Select(&users, userQueries["list"])
	return users, err
}

func (s *UserService) Create(user *models.User, password string) error {
	hash, err := passlib.Hash(password)
	if err != nil {
		return err
	}
	user.PasswordHash = hash
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.NamedExec(userQueries["create"], user)
	if err != nil {
		return err
	}
	err = tx.Get(user, userQueries["get"], user.Email)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *UserService) UpdatePassword(user *models.User, password string) error {
	hash, err := passlib.Hash(password)
	if err != nil {
		return err
	}
	user.PasswordHash = hash
	_, err = s.DB.NamedExec(userQueries["update-password"], user)
	return err
}

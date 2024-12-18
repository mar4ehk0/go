package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Add(dto *CreateDto) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id"
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		wrappedErr := fmt.Errorf("can't do prepare query user {%s, %s, %s} error: %w", dto.Name, dto.Email, dto.Password, err)
		return 0, wrappedErr
	}

	var id int
	err = stmt.Get(&id, dto)
	if err != nil {
		msgErr := fmt.Sprintf("can't do insert user {%s, %s, %s}", dto.Name, dto.Email, dto.Password)
		err = db.ProcessError(err, msgErr)
		return 0, err
	}

	return id, nil
}

func (r *Repo) GetByID(id int) (User, error) {
	var user User

	err := r.db.QueryRowx("SELECT id, name, email, password FROM users WHERE id=$1", id).StructScan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return user, db.ErrDBNotFound
	}
	if err != nil {
		wrappedErr := fmt.Errorf("can't do select user by id {%d} error: %w", id, err)
		return user, wrappedErr
	}

	return user, nil
}

func (r *Repo) Update(id int, dto *UpdateDto) error {
	msgErr := fmt.Sprintf("can't do prepare update user {%s, %s}", dto.Name, dto.Email)

	result, err := r.db.NamedExec("UPDATE users SET name=:name, email=:email WHERE id=:id", struct {
		ID    int
		Name  string
		Email string
	}{ID: id, Name: dto.Name, Email: dto.Email})
	if err != nil {
		err = db.ProcessError(err, msgErr)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = db.ProcessError(err, msgErr)
		return err
	}
	if rowsAffected == 0 {
		return db.ErrDBNotFound
	}

	return nil
}

func (r *Repo) GetByINWithTx(tx *sqlx.Tx, id int) (User, error) {
	var user User
	err := tx.QueryRowx("SELECT id, name, email, password FROM users WHERE id=$1;", id).StructScan(&user)
	if err != nil {
		return User{}, err
	}
	if user.ID != id {
		return User{}, db.ErrDBNotFound
	}

	return user, nil
}

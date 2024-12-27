package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw16_docker/pkg/db"
)

type RepoUser struct {
	db *db.Connect
}

func NewRepoUser(connect *db.Connect) *RepoUser {
	return &RepoUser{db: connect}
}

func (r *RepoUser) Add(dto *EntryCreateDto) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id"
	stmt, err := r.db.Connect.PrepareNamed(query)
	if err != nil {
		return 0, fmt.Errorf("prepare query: %w", err)
	}

	var id int
	err = stmt.Get(&id, dto)
	if err != nil {
		if r.db.IsErrDuplicate(err) {
			return 0, fmt.Errorf("already exist user{%v}: %w", dto, errors.Join(db.ErrDBDuplicateKey, err))
		}
		return 0, fmt.Errorf("sql insert user{%v}: %w", dto, err)
	}

	return id, nil
}

func (r *RepoUser) GetByID(id int) (User, error) {
	var user User

	err := r.db.Connect.QueryRowx("SELECT id, name, email, password FROM users WHERE id=$1", id).StructScan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return user, fmt.Errorf("not found user by id {%d}: %w", id, errors.Join(db.ErrDBNotFound, err))
	}
	if err != nil {
		wrappedErr := fmt.Errorf("sql select user by id {%d}: %w", id, err)
		return user, wrappedErr
	}

	return user, nil
}

func (r *RepoUser) Update(id int, dto *EntryUpdateDto) error {
	result, err := r.db.Connect.NamedExec("UPDATE users SET name=:name, email=:email WHERE id=:id", struct {
		ID    int
		Name  string
		Email string
	}{ID: id, Name: dto.Name, Email: dto.Email})
	if err != nil {
		if r.db.IsErrDuplicate(err) {
			return fmt.Errorf("already exist user {%d, %v}: %w", id, dto, errors.Join(db.ErrDBDuplicateKey, err))
		}
		return fmt.Errorf("update user{%d, %v}: %w", id, dto, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected user {%d}: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("update user {%d}: %w", id, db.ErrDBNotFound)
	}

	return nil
}

func (r *RepoUser) GetByIDWithTx(tx *sqlx.Tx, id int) (User, error) {
	var user User
	err := tx.QueryRowx("SELECT id, name, email, password FROM users WHERE id=$1;", id).StructScan(&user)
	if err != nil {
		return User{}, fmt.Errorf("failed to select user by id {%d}: %w", id, err)
	}
	if user.ID != id {
		return User{}, db.ErrDBNotFound
	}

	return user, nil
}

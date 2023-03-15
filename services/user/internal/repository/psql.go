package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"grpcrest/services/user/domain/entity"
	domain "grpcrest/services/user/domain/repository"
	"grpcrest/services/user/internal/config"

	"github.com/google/uuid"
)

const (
	errNoRowsAffected = "no rows affected"
)

type UserRepository struct {
	sql *sql.DB
	cfg *config.AppConfig
}

// Create implements repository.UserRepository
func (u *UserRepository) Create(d *entity.User) error {
	res, err := u.sql.Exec(fmt.Sprintf("INSERT INTO %s.users(id,name,email,address) VALUES ($1,$2,$3,$4)", u.cfg.DBSchema), d.ID, d.Name, d.Email, d.Address)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New(errNoRowsAffected)
	}
	return nil
}

// Delete implements repository.UserRepository
func (u *UserRepository) Delete(d *entity.User) error {
	//convert to uuid
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return err
	}
	res, err := u.sql.Exec(fmt.Sprintf("DELETE FROM %s.users WHERE text(id) = $1", u.cfg.DBSchema), id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New(errNoRowsAffected)
	}
	return nil
}

// GetAll implements repository.UserRepository
func (u *UserRepository) GetAll(page *int, limit *int) ([]entity.User, error) {
	var data []entity.User
	offset := (*page - 1) * *limit
	rows, err := u.sql.Query(fmt.Sprintf("SELECT text(id),name,email,address FROM %s.users OFFSET $1 LIMIT $2", u.cfg.DBSchema), offset, *limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tmp entity.User
		err = rows.Scan(&tmp.ID, &tmp.Name, &tmp.Email, &tmp.Address)
		if err != nil {
			return nil, err
		}
		data = append(data, tmp)
	}
	return data, nil
}

// IsExistEmail implements repository.UserRepository
func (u *UserRepository) IsExistEmail(email *string) bool {
	var id string
	err := u.sql.QueryRow(fmt.Sprintf("SELECT id from %s.users WHERE email = $1", u.cfg.DBSchema), *email).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

type fieldUpdate struct {
	field string
	value string
}

// Update implements repository.UserRepository
func (u *UserRepository) Update(user *entity.User) error {
	var upField []fieldUpdate
	if user.Email != "" {
		upField = append(upField, fieldUpdate{field: "email", value: user.Email})
	}
	if user.Name != "" {
		upField = append(upField, fieldUpdate{field: "name", value: user.Name})
	}
	if user.Address != "" {
		upField = append(upField, fieldUpdate{field: "address", value: user.Address})
	}
	err := u.dynamicUpdate(user.ID, upField...)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) dynamicUpdate(userID string, f ...fieldUpdate) error {
	//begin transaction
	tx, err := u.sql.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	//loop
	for _, v := range f {
		_, err = tx.Exec(fmt.Sprintf("UPDATE %s.users SET %s = $1 WHERE id = $3", u.cfg.DBSchema, v.field), v.value, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//commit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func NewUserRepository(sql *sql.DB, cfg *config.AppConfig) domain.UserRepository {
	return &UserRepository{
		sql: sql,
		cfg: cfg,
	}
}

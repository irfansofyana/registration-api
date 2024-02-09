package repository

import (
	"context"
	"database/sql"
	"errors"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (output *GetUserByPhoneNumberOutput, err error) {
	output = &GetUserByPhoneNumberOutput{}
	err = r.Db.QueryRowContext(ctx, "SELECT id FROM users WHERE phone_number = $1", input.PhoneNumber).Scan(&output.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			output = nil
			err = nil
			return
		}
	}
	return
}

func (r *Repository) SaveUser(ctx context.Context, input SaveUserInput) (id string, err error) {
	id = ""
	err = r.Db.QueryRowContext(ctx,
		"INSERT INTO users (full_name, phone_number, password) VALUES ($1, $2, $3) RETURNING id",
		input.FullName, input.PhoneNumber, input.Password).Scan(&id)
	return
}

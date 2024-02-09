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
	err = r.Db.QueryRowContext(ctx,
		"SELECT id, full_name, password FROM users WHERE phone_number = $1",
		input.PhoneNumber).Scan(
		&output.Id,
		&output.FullName,
		&output.Password,
	)
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

func (r *Repository) UpdateUserLoginCount(ctx context.Context, input UpdateUserCountInput) (err error) {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"UPDATE users SET successful_login_count = successful_login_count + 1 WHERE id = $1",
		input.Id,
	)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetProfile(ctx context.Context, input GetProfileInput) (output *GetProfileOutput, err error) {
	output = &GetProfileOutput{}
	err = r.Db.QueryRowContext(ctx,
		"SELECT full_name, phone_number FROM users WHERE id = $1",
		input.Id).Scan(
		&output.FullName,
		&output.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			output = nil
			err = nil
			return
		}
		return
	}
	return
}

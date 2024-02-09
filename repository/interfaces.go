// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (output *GetUserByPhoneNumberOutput, err error)
	SaveUser(ctx context.Context, input SaveUserInput) (id string, err error)
	UpdateUserLoginCount(ctx context.Context, input UpdateUserCountInput) (err error)
	GetProfile(ctx context.Context, input GetProfileInput) (output *GetProfileOutput, err error)
}

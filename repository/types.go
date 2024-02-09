// This file contains types that are used in the repository layer.
package repository

type GetUserByPhoneNumberInput struct {
	PhoneNumber string
}

type GetUserByPhoneNumberOutput struct {
	Id       string
	FullName string
	Password string
}

type SaveUserInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

type UpdateUserCountInput struct {
	Id string
}

type GetProfileInput struct {
	Id string
}

type GetProfileOutput struct {
	Id          string
	FullName    string
	PhoneNumber string
}

type UpdateUserProfileInput struct {
	Id          string
	FullName    string
	PhoneNumber string
}

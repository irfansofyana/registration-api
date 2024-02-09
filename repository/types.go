// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type GetUserByPhoneNumberInput struct {
	PhoneNumber string
}

type GetUserByPhoneNumberOutput struct {
	Id string
}

type SaveUserInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

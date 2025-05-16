package dto

import "time"

type AdminWithCountDTO struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Role           string `json:"role"`
	DependentCount int    `json:"dependent_count"`
}

type CreateAdminDTO struct {
	FirstName   string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string    `json:"last_name" validate:"required,min=2,max=50"`
	Email       string    `json:"email" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" validate:"required,e164"`
	Password    string    `json:"password" validate:"required,min=8"`
	Role        string    `json:"role" validate:"omitempty,oneof=staff admin superadmin"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type UpdateAdminDTO struct {
	FirstName   *string    `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName    *string    `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Email       *string    `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber *string    `json:"phone_number,omitempty" validate:"omitempty,e164"`
	Role        *string    `json:"role,omitempty" validate:"omitempty,oneof=staff admin superadmin"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
}

type AdminBasicDTO struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

package dto

import "time"

type DependentWithCountDTO struct {
	ID             uint    `json:"id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Email          *string `json:"email"`
	PhoneNumber    *string `json:"phone_number"`
	MembershipTier string  `json:"membership_tier"`
	GuardianCount  int     `json:"guardian_count"`
}

type CreateDependentDTO struct {
	FirstName      string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName       string    `json:"last_name" validate:"required,min=2,max=50"`
	Email          string    `json:"email" validate:"required,email"`
	PhoneNumber    string    `json:"phone_number" validate:"required,e164"`
	Password       string    `json:"password" validate:"required,min=8"`
	MembershipTier string    `json:"membership_tier" validate:"omitempty,oneof=free basic prime elite"`
	DateOfBirth    time.Time `json:"date_of_birth" validate:"required"`
	AdminID        *uint     `json:"admin_id" validate:"omitempty"`
}

type UpdateDependentDTO struct {
	FirstName      *string    `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName       *string    `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Email          *string    `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber    *string    `json:"phone_number,omitempty" validate:"omitempty,e164"`
	MembershipTier *string    `json:"membership_tier,omitempty" validate:"omitempty,oneof=free basic prime elite"`
	DateOfBirth    *time.Time `json:"date_of_birth,omitempty"`
	AdminID        *uint      `json:"admin_id,omitempty"`
}

type DependentBasicDTO struct {
	ID          uint    `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
}

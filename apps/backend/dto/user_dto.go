package dto

import "time"

type UserWithCountDTO struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	MembershipTier string `json:"membership_tier"`
	DependentCount int    `json:"dependent_count"`
}

type CreateUserDTO struct {
	FirstName      string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName       string    `json:"last_name" validate:"required,min=2,max=50"`
	Email          string    `json:"email" validate:"required,email"`
	PhoneNumber    string    `json:"phone_number" validate:"required,e164"`
	Password       string    `json:"password" validate:"required,min=8"`
	MembershipTier string    `json:"membership_tier" validate:"omitempty,oneof=free basic prime elite"`
	DateOfBirth    time.Time `json:"date_of_birth" validate:"required"`
}

type UpdateUserDTO struct {
	FirstName      *string    `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName       *string    `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Email          *string    `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber    *string    `json:"phone_number,omitempty" validate:"omitempty,e164"`
	Password       *string    `json:"password,omitempty" validate:"omitempty,min=8"`
	MembershipTier *string    `json:"membership_tier,omitempty" validate:"omitempty,oneof=free basic prime elite"`
	DateOfBirth    *time.Time `json:"date_of_birth,omitempty" validate:"omitempty"`
}

type UpdatePasswordDTO struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type EmergencyContactDTO struct {
	ID          uint       `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Relation    string     `json:"relation"`
}

type DependentDTO struct {
	ID                uint                  `json:"id"`
	FirstName         string                `json:"first_name"`
	LastName          string                `json:"last_name"`
	Email             *string               `json:"email"`
	PhoneNumber       *string               `json:"phone_number"`
	DateOfBirth       *time.Time            `json:"date_of_birth"`
	MembershipTier    string                `json:"membership_tier"`
	EmergencyContacts []EmergencyContactDTO `json:"emergency_contacts"`
}

type UserWithDependentsDTO struct {
	ID             uint           `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Email          string         `json:"email"`
	PhoneNumber    string         `json:"phone_number"`
	MembershipTier string         `json:"membership_tier"`
	DateOfBirth    *time.Time     `json:"date_of_birth"`
	Dependents     []DependentDTO `json:"dependents"`
}

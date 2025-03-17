package models

import (
	"database/sql"
	"time"
)

type UserStatus string

const (
	ActiveUserStatus    UserStatus = "active"
	SuspendedUserStatus UserStatus = "suspended"
	DeletedUserStatus   UserStatus = "deleted"
)

type UserRole string

const (
	AdminUserRole   UserRole = "admin"
	PatientUserRole UserRole = "patient"
	DoctorUserRole  UserRole = "doctor"
)

type User struct {
	ID          string         `gorm:"default:uuid_generate_v4();primary_key" json:"id"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Name        string         `json:"name" validate:"required,name"`
	Email       string         `gorm:"unique" json:"email" validate:"required,email"`
	PhoneNo     sql.NullString `json:"phone_no"`
	Status      UserStatus     `gorm:"default:active" json:"status"`
	Role        UserRole       `json:"role" validate:"required,role"`
	Otp         sql.NullString `json:"otp"`
	PatientData Patient
}

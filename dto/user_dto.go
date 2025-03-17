package dto

type UserCreateAdminDTO struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	PhoneNo string `json:"phone_no"`
}

type UserCreatePatientDTO struct {
	UserData    UserCreateAdminDTO `json:"user_data" validate:"required"`
	PatientData PatientCreateDto   `json:"patient_data" validate:"required"`
}

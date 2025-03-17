package services

import (
	"CarepluseBackend/database"
	"CarepluseBackend/dto"
	"CarepluseBackend/models"
	"database/sql"
	"errors"
	"time"
)

func CreateAdminUser(userData *dto.UserCreateAdminDTO) (*models.User, error) {
	// confirm if user with same email exists
	userExists := &models.User{}

	if err := database.DB.Db.Where("email = ?", userData.Email).First(userExists).Error; err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, errors.New("user already exists")
	}

	phoneNo := userData.PhoneNo
	user := &models.User{
		Name:    userData.Name,
		Email:   userData.Email,
		PhoneNo: sql.NullString{String: phoneNo, Valid: phoneNo != ""},
	}

	if err := database.DB.Db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil

}

func GetUser(id string) (*models.User, error) {
	user := &models.User{}

	if err := database.DB.Db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := database.DB.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreatePatientAccount(userData *dto.UserCreateAdminDTO, patientData *dto.PatientCreateDto) (*models.User, error) {
	userExists := &models.User{}

	if err := database.DB.Db.Where("email = ?", userData.Email).First(userExists).Error; err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, errors.New("user already exists")
	}

	phoneNo := userData.PhoneNo
	user := &models.User{
		Name:    userData.Name,
		Email:   userData.Email,
		PhoneNo: sql.NullString{String: phoneNo, Valid: phoneNo != ""},
	}

	if err := database.DB.Db.Create(user).Error; err != nil {
		return nil, err
	}

	dob, _ := time.Parse("2006-01-01", patientData.Dob)

	patient := &models.Patient{
		Dob:                             dob,
		Gender:                          patientData.Gender,
		Address:                         patientData.Address,
		Occupation:                      &patientData.Occupation,
		EmergencyContactName:            &patientData.EmergencyContactName,
		EmergencyContactPhoneNo:         &patientData.EmergencyContactPhoneNo,
		InsurancePolicyNumber:           &patientData.InsurancePolicyNumber,
		InsuranceProvider:               &patientData.InsuranceProvider,
		Allergies:                       &patientData.Allergies,
		CurrentMedications:              &patientData.CurrentMedications,
		FamilyHistory:                   &patientData.FamilyHistory,
		PastMedicalHistory:              &patientData.PastMedicalHistory,
		IdentificationType:              models.IdentificationType(patientData.IdentificationType),
		IdentificationNumber:            patientData.IdentificationNumber,
		AcceptedReceiveTreatment:        patientData.AcceptedReceiveTreatment,
		AcceptedDisclosureOfInformation: patientData.AcceptedDisclosureOfInformation,
		AcceptedPrivacyPolicy:           patientData.AcceptedPrivacyPolicy,
		UserID:                          &user.ID,
	}

	if err := database.DB.Db.Create(patient).Error; err != nil {
		return nil, err
	}

	return user, nil

}

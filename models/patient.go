package models

import "time"

type IdentificationType string

const (
	PatientIdentificationTypeNational IdentificationType = "national"
	PatientIdentificationTypeAlien    IdentificationType = "alien"
	PatientIdentificationTypePassport IdentificationType = "passport"
	PatientIdentificationTypeOther    IdentificationType = "other"
)

type Patient struct {
	Dob                             time.Time          `gorm:"type:date"`
	Gender                          string             `json:"gender"`
	Address                         string             `json:"address"`
	Occupation                      *string            `json:"occupation"`
	EmergencyContactName            *string            `json:"emergency_contact_name"`
	EmergencyContactPhoneNo         *string            `json:"emergency_contact_phone_no"`
	InsuranceProvider               *string            `json:"insurance_provider"`
	InsurancePolicyNumber           *string            `json:"insurance_policy_number"`
	Allergies                       *string            `json:"allergies"`
	CurrentMedications              *string            `json:"current_medications"`
	FamilyHistory                   *string            `json:"family_history"`
	PastMedicalHistory              *string            `json:"past_medical_history"`
	IdentificationType              IdentificationType `json:"identification_type"`
	IdentificationNumber            string             `json:"identification_number"`
	IdentificationUrl               string             `gorm:"type:text" json:"identification_url"`
	AcceptedReceiveTreatment        bool               `json:"accepted_receive_treatment"`
	AcceptedDisclosureOfInformation bool               `gorm:"default:false" json:"accepted_disclosure_of_information"`
	AcceptedPrivacyPolicy           bool               `gorm:"default:false" json:"accepted_privacy_policy"`
	UserID                          *string
}

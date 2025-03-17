package dto

type PatientCreateDto struct {
	Dob                             string `json:"dob" validate:"required,date"`
	Gender                          string `json:"gender" validate:"required"`
	Address                         string `json:"address" validate:"required"`
	Occupation                      string `json:"occupation"`
	EmergencyContactName            string `json:"emergency_contact_name"`
	EmergencyContactPhoneNo         string `json:"emergency_contact_phone_no"`
	InsuranceProvider               string `json:"insurance_provider"`
	InsurancePolicyNumber           string `json:"insurance_policy_number"`
	Allergies                       string `json:"allergies"`
	CurrentMedications              string `json:"current_medications"`
	FamilyHistory                   string `json:"family_history"`
	PastMedicalHistory              string `json:"past_medical_history"`
	IdentificationType              string `json:"identification_type" validate:"required,oneof=national alien passport other"`
	IdentificationNumber            string `json:"identification_number" validate:"required"`
	IdentificationUrl               string `json:"identification_url" validate:"required"`
	AcceptedReceiveTreatment        bool   `json:"accepted_receive_treatment"`
	AcceptedDisclosureOfInformation bool   `json:"accepted_disclosure_of_information"`
	AcceptedPrivacyPolicy           bool   `json:"accepted_privacy_policy"`
}

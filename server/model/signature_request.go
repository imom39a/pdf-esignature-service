package model

type SignatureRequest struct {
	ID                  int64  `json:"id"`
	OriginalDocumentID  string `json:"originalDocumentId"`
	OriginalDocumentURL string `json:"originalDocumentUrl"`
	GeneratedID         string `json:"generatedDocumentId"`
	ExpiresOn           string `json:"expiresOn"`
	ApproverEmail       string `json:"approverEmail"`
	RequestedDate       string `json:"requestedDate"`
}

type CompletedSignatureRequest struct {
	ID                  int64  `json:"id"`
	OriginalDocumentID  string `json:"originalDocumentId"`
	OriginalDocumentURL string `json:"originalDocumentUrl"`
	SignedDocumentID    string `json:"signedDocumentId"`
	SignedDocumentURL   string `json:"signedDocumentUrl"`
	SignedOn            string `json:"signedOn"`
	SignedByFirstName   string `json:"signedByFirstName"`
	SignedByLastName    string `json:"signedByLastName"`
	SignedByEmail       string `json:"signedByEmail"`
	SignedFromIPAddress string `json:"signedFromIpAddress"`
	SignedDocumentHash  string `json:"signedDocumentHash"`
	TermsAndConditionID string `json:"termsAndConditionId"`
	SignImageBase64     string `json:"signImageBase64"`
}

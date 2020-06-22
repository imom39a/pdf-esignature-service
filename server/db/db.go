package db

import (
	"database/sql"
	"pdf-esignature-server/model"
	"pdf-esignature-server/utils"
)

type DB interface {
	GetSigningRequests() ([]*model.SignatureRequest, error)
	CreateSigningRequests(model.SignatureRequest) error
	CompleteSingingRequest(model.CompletedSignatureRequest) error
}

type PostgresDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return PostgresDB{db: db}
}

func (appPostgresDB PostgresDB) GetSigningRequests() ([]*model.SignatureRequest, error) {
	rows, err := appPostgresDB.db.Query(`
	select 
		id, original_document_id, original_document_url, generated_id, expires_on, approver_email, requested_date 
	from signature_requests
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var signatureRequests []*model.SignatureRequest
	for rows.Next() {
		signatureRequest := new(model.SignatureRequest)
		err = rows.Scan(
			&signatureRequest.ID,
			&signatureRequest.OriginalDocumentID,
			&signatureRequest.OriginalDocumentURL,
			&signatureRequest.GeneratedID,
			&signatureRequest.ExpiresOn,
			&signatureRequest.ApproverEmail,
			&signatureRequest.RequestedDate,
		)
		if err != nil {
			return nil, err
		}
		signatureRequests = append(signatureRequests, signatureRequest)
	}
	return signatureRequests, nil
}

func (appPostgresDB PostgresDB) CreateSigningRequests(signatureRequest model.SignatureRequest) error {
	_, err := appPostgresDB.db.Exec(`
	insert into signature_requests (original_document_id, original_document_url, generated_id, approver_email) 
	values ($1, $2, $3, $4)
	`,
		signatureRequest.OriginalDocumentID,
		signatureRequest.OriginalDocumentURL,
		utils.CreateHash(signatureRequest.OriginalDocumentID),
		signatureRequest.ApproverEmail,
	)
	if err != nil {
		return err
	}

	return nil
}

func (appPostgresDB PostgresDB) CompleteSingingRequest(completedSignatureRequest model.CompletedSignatureRequest) error {
	_, err := appPostgresDB.db.Exec(`
	insert into completed_signature_requests (original_document_id, original_document_url, signed_document_id, signed_document_url, signed_by_first_name, signed_by_last_name, signed_by_email, signed_from_ip_address, signed_document_hash, terms_and_condition_id) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`,
		completedSignatureRequest.OriginalDocumentID,
		completedSignatureRequest.OriginalDocumentURL,
		completedSignatureRequest.SignedDocumentID,
		completedSignatureRequest.SignedDocumentURL,
		completedSignatureRequest.SignedByFirstName,
		completedSignatureRequest.SignedByLastName,
		completedSignatureRequest.SignedByEmail,
		completedSignatureRequest.SignedFromIPAddress,
		completedSignatureRequest.SignedDocumentHash,
		completedSignatureRequest.TermsAndConditionID,
	)
	if err != nil {
		return err
	}
	return nil
}

export interface SignatureRequest {
	id: BigInteger
	originalDocumentId: string
	originalDocumentUrl: string
	generatedDocumentId: string
	expiresOn: Date
	approverEmail: string
	requestedDate: Date
}

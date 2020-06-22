drop table if exists signature_requests;

CREATE TABLE IF NOT EXISTS signature_requests (
  id BIGSERIAL,
  original_document_id TEXT,
  original_document_url TEXT,
  generated_id TEXT,
  expires_on TIMESTAMP default NOW() + interval '2 day',
  approver_email TEXT,
  requested_date TIMESTAMP default NOW(),
  CONSTRAINT signature_requests_pkey PRIMARY KEY (id)
);

drop table if exists complted_signature_requests;

CREATE TABLE IF NOT EXISTS completed_signature_requests (
  id BIGSERIAL,
  original_document_id TEXT,
  original_document_url TEXT,
  signed_document_id TEXT,
  signed_document_url TEXT,
  signed_on TIMESTAMP default NOW(),
  signed_by_first_name TEXT,
  signed_by_last_name TEXT,
  signed_by_email TEXT,
  signed_from_ip_address TEXT,
  generated_id TEXT,
  signed_document_hash TEXT,
  terms_and_condition_id TEXT,
  CONSTRAINT compelted_signature_requests_pkey PRIMARY KEY (id)
);

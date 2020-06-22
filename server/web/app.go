package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pdf-esignature-server/blobstore"
	"pdf-esignature-server/db"
	"pdf-esignature-server/mail"
	"pdf-esignature-server/model"
	"pdf-esignature-server/pdfutils"
	"strings"
	"time"
)

type PDFSigningApp struct {
	appPostgresDB db.DB
	handlers      map[string]http.HandlerFunc
}

func PDFSigningServer(appPostgresDB db.DB, cors bool) PDFSigningApp {
	app := PDFSigningApp{
		appPostgresDB: appPostgresDB,
		handlers:      make(map[string]http.HandlerFunc),
	}
	signingRequestsHandler := app.SigningManager
	completeSigningRequestsHandler := app.ComeplteSingingRequestHandler
	fileUploadHandler := app.FileUploadManager
	fileDownloadHandler := app.FileDownloadManager
	testHandler := app.TestHandler

	if !cors {
		signingRequestsHandler = disableCors(signingRequestsHandler)
	}
	app.handlers["/api/signing-requests"] = signingRequestsHandler
	app.handlers["/api/signing-requests/complete"] = completeSigningRequestsHandler
	app.handlers["/api/upload"] = fileUploadHandler
	app.handlers["/api/download/"] = fileDownloadHandler
	app.handlers["/api/test"] = testHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *PDFSigningApp) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (pdfSigningApp *PDFSigningApp) SigningManager(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Println("Received a GET request")
		responseWriter.Header().Set("Content-Type", "application/json")
		signingRequests, err := pdfSigningApp.appPostgresDB.GetSigningRequests()
		if err != nil {
			sendErr(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}
		err = json.NewEncoder(responseWriter).Encode(signingRequests)
		if err != nil {
			sendErr(responseWriter, http.StatusInternalServerError, err.Error())
		}
	case "POST":
		fmt.Println("Received a POST request")
		var signatureRequest model.SignatureRequest
		err := json.NewDecoder(request.Body).Decode(&signatureRequest)
		if err != nil {
			sendErr(responseWriter, http.StatusBadRequest, err.Error())
		}
		fmt.Printf("%+v\n", signatureRequest)
		err = pdfSigningApp.appPostgresDB.CreateSigningRequests(signatureRequest)
		if err != nil {
			sendErr(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		now := time.Now()
		expiresOn := now.AddDate(0, 0, 2)

		//  once entry is accepted, create and send email
		templateData := struct {
			OriginalDocumentID  string
			OriginalDocumentURL string
			ExpiresOn           string
		}{
			OriginalDocumentID:  signatureRequest.OriginalDocumentID,
			OriginalDocumentURL: signatureRequest.OriginalDocumentURL,
			ExpiresOn:           expiresOn.String(),
		}
		r := mail.NewRequest([]string{signatureRequest.ApproverEmail}, "Request to signing PDF document "+signatureRequest.OriginalDocumentID, "")
		err = r.ParseTemplate(templateData)
		if err != nil {
			fmt.Println("Failed to send email")
		}
		ok, _ := r.SendEmail()
		fmt.Println(ok)

		err = json.NewEncoder(responseWriter).Encode(http.StatusAccepted)
		if err != nil {
			sendErr(responseWriter, http.StatusInternalServerError, err.Error())
		}
	}
}

func (pdfSigningApp *PDFSigningApp) ComeplteSingingRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET menthod not supporte for this API")
		sendErr(w, http.StatusNotFound, "GET menthod not supporte for this API")
	case "POST":
		fmt.Println("Received a POST request")

		var completedSignatureRequest model.CompletedSignatureRequest
		err := json.NewDecoder(r.Body).Decode(&completedSignatureRequest)
		if err != nil {
			sendErr(w, http.StatusBadRequest, err.Error())
		}
		fmt.Printf("%+v\n", completedSignatureRequest)

		// sign document using document id and signature from user
		signedFilePath := pdfutils.SignFile(completedSignatureRequest.OriginalDocumentID, completedSignatureRequest.SignImageBase64)

		// update signed document details
		completedSignatureRequest.SignedDocumentID = "Signed-" + completedSignatureRequest.OriginalDocumentID
		completedSignatureRequest.SignedDocumentURL = "Signed-" + completedSignatureRequest.OriginalDocumentID

		// upload signed document to S3
		blobstore.AddFileToS3_(signedFilePath, "Signed-"+completedSignatureRequest.SignedDocumentID)
		// save uplaoded document id and document url to

		err = pdfSigningApp.appPostgresDB.CompleteSingingRequest(completedSignatureRequest)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

}

func (pdfSigningApp *PDFSigningApp) FileUploadManager(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file-1")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	err = blobstore.AddFileToS3(file, handler)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, "Failed Uploaded File into S3")
	} else {
		fmt.Fprintf(w, "Successfully Uploaded File into S3\n")
	}

}

func (pdfSigningApp *PDFSigningApp) FileDownloadManager(w http.ResponseWriter, r *http.Request) {
	documentID := strings.TrimPrefix(r.URL.Path, "/api/download/")
	if documentID != "" {
		tempFile, err := blobstore.GetFileFromS3(documentID)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, "Failed Download File from S3")
		} else {
			http.ServeFile(w, r, tempFile.Name())
		}

	} else {
		sendErr(w, http.StatusBadRequest, "No file ID passed")
	}
}

func (pdfSigningApp *PDFSigningApp) TestHandler(w http.ResponseWriter, r *http.Request) {
	//pdfutils.GetFile()
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}

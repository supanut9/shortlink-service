package httpService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/supanut9/shortlink-service/internal/repository"
)

type FileService interface {
	UploadFile(bucketName, folderPath, filename string, file *bytes.Reader) (string, error)
}

type fileService struct {
	fileServiceBaseURL string
}

func NewFileService(fileServiceBaseURL string) FileService {
	return &fileService{
		fileServiceBaseURL: fileServiceBaseURL,
	}
}

type FileUploadResponse struct {
	Filename string `json:"filename"`
	Message  string `json:"message"`
	URL      string `json:"url"`
}

func (f *fileService) UploadFile(bucketName, folderPath, filename string, file *bytes.Reader) (string, error) {
	log.Printf("🟢 Starting file upload: filename=%s bucket=%s folder=%s", filename, bucketName, folderPath)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		log.Printf("❌ Error creating form file: %v", err)
		return "", fmt.Errorf("create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		log.Printf("❌ Error copying file content: %v", err)
		return "", fmt.Errorf("copy file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		log.Printf("❌ Error closing multipart writer: %v", err)
		return "", fmt.Errorf("close writer: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/files?bucketName=%s&folderPath=%s", f.fileServiceBaseURL, bucketName, folderPath)
	log.Printf("📤 Sending POST request to: %s", url)

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		log.Printf("❌ Error creating HTTP request: %v", err)
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("❌ HTTP request failed: %v", err)
		return "", fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Upload failed. Status: %d Body: %s", resp.StatusCode, string(body))

		// If the upstream service returns 507, return our specific error
		if resp.StatusCode == http.StatusInsufficientStorage {
			return "", repository.ErrInsufficientStorage
		}

		// For all other errors, wrap the general file upload error
		return "", fmt.Errorf("%w: %s", repository.ErrFileUploadFailed, string(body))
	}

	log.Printf("✅ Upload successful. Parsing response...")

	var result FileUploadResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("❌ JSON decode failed: %v. Raw body: %s", err, string(body))
		return "", fmt.Errorf("decode response failed: %w", err)
	}

	log.Printf("✅ File uploaded successfully: %s", result.URL)
	return result.URL, nil
}

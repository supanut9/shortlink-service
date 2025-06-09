package repository

import "errors"

// ErrNotFound is returned when a record is not found
var ErrNotFound = errors.New("not found")

// NEW: Add custom errors for specific failure scenarios.
// These can be used across the service and handler layers.

// ErrUniqueSlugGenerationFailed is returned when the service cannot create a unique slug.
var ErrUniqueSlugGenerationFailed = errors.New("failed to generate unique slug after multiple attempts")

// ErrFileUploadFailed indicates a general failure from the file service.
var ErrFileUploadFailed = errors.New("file upload failed")

// ErrInsufficientStorage indicates a specific 507 error from the file service.
var ErrInsufficientStorage = errors.New("insufficient storage for file upload")

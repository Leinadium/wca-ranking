package port

import (
	"io"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type FileService interface {
	// FromReader creates a *domain.File out of a io.Reader
	FromReader(reader io.Reader) (*domain.File, error)

	// ExtractZip extracts a file from a zip, searching inside the zipped folder using the searchFor function
	ExtractZip(zipFile *domain.File, searchFor func(string) bool) (*domain.File, error)

	// Delete deletes the files from the system
	Delete(file ...domain.File) error
}

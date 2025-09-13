package port

import (
	"leinadium.dev/wca-ranking/internal/core/domain"
)

type FileService interface {
	// FromReader creates a domain.File out of a domain.RawFile
	FromRaw(reader domain.RawFile) (domain.File, error)

	// ExtractZip extracts a file from a zip, searching inside the zipped folder using the searchFor function
	ExtractZip(zipFile domain.File, searchFor func(string) bool) (domain.File, error)

	// Delete deletes the files from the system
	Delete(file ...domain.FileDeleter) error
}

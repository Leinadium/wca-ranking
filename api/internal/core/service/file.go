package service

import (
	"archive/zip"
	"errors"
	"io"
	"os"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

// type FileService interface {
// 	// FromReader creates a *domain.File out of a io.Reader
// 	FromReader(reader io.Reader) (*domain.File, error)

// 	// ExtractZip extracts a file from a zip, searching inside the zipped folder using the searchFor function
// 	ExtractZip(zipFile *domain.File, searchFor func(string) bool) (*domain.File, error)

// 	// Delete deletes the files from the system
// 	Delete(file ...domain.File) error
// }

const (
	FilePattern = "wca"
)

type File struct {
	OSFile *os.File
}

func (f *File) Name() string {
	return f.OSFile.Name()
}

func (f *File) Delete() error {
	return os.Remove(f.Name())
}

func createFile() (*File, error) {
	f, err := os.CreateTemp("", FilePattern)
	if err != nil {
		return nil, err
	}
	return &File{OSFile: f}, nil
}

type FileService struct {
}

func (f *FileService) FromRaw(reader domain.RawFile) (domain.File, error) {
	file, err := createFile()
	if err != nil {
		return nil, err
	}
	n, err := io.Copy(file.OSFile, reader)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		// TODO: const error
		return nil, errors.New("empty reader")
	}
	return file, nil
}

func (f *FileService) ExtractZip(zipFile domain.File, searchFor func(string) bool) (domain.File, error) {
	// create zip reader
	r, err := zip.OpenReader(zipFile.Name())
	if err != nil {
		return nil, err
	}

	// defer removing file
	defer r.Close()
	if fd, ok := zipFile.(domain.FileDeleter); ok {
		defer f.Delete(fd)
	}

	// creating final uncompressed file
	finalFile, err := createFile()
	if err != nil {
		return nil, err
	}

	// iterating over zip folder
	for _, file := range r.File {
		reader, err := file.Open()
		if err != nil {
			continue
		}
		defer reader.Close()

		// skipping if found other files
		if !searchFor(file.Name) {
			continue
		}

		// if found, then copy to destination
		_, err = io.Copy(finalFile.OSFile, reader)
		if err != nil {
			return nil, err
		}
		defer finalFile.OSFile.Close()
		return finalFile, err
	}
	// TODO: const error
	return nil, errors.New("file not found inside zip")

}

func (f *FileService) Delete(file ...domain.FileDeleter) error {
	for _, fl := range file {
		if err := fl.Delete(); err != nil {
			return err
		}
	}
	return nil
}

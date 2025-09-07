package domain

import "io"

type RawFile io.ReadCloser

type File interface {
	Name() string
}

type FileDeleter interface {
	File
	Delete() error
}

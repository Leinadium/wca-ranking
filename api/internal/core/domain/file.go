package domain

type File interface {
	Name() string
}

type FileDeleter interface {
	File
	Delete() error
}

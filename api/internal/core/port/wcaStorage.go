package port

import (
	"time"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type WCAStorage interface {
	// ImportFile imports a sql file to the storage
	ImportFile(file *domain.File) error

	// Update updates app and datalake tables from the sql dump
	Update() error

	// Refresh updates the rankings without recreating most of the tables
	Refresh() error

	// CurrentDate gets the current dump date stored in the database
	CurrentDate() (time.Time, error)

	// SetCurrentDate updates the current dump date
	SetCurrentDate(time.Time) error
}

package port

import (
	"context"
	"time"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type SyncRepository interface {
	CurrentDate(context.Context) (*time.Time, error)
	ImportFile(string, domain.File) error
	Update(context.Context) error
	Refresh(context.Context) error
	SetCurrentDate(context.Context, time.Time) error
}

type SyncService interface {
	// ImportFile imports a sql file to the storage
	ImportFile(bin string, file domain.File) error

	// Update updates app and datalake tables from the sql dump
	Update(context.Context) error

	// Refresh updates the rankings without recreating most of the tables
	Refresh(context.Context) error

	// CurrentDate gets the current dump date stored in the database
	CurrentDate(context.Context) (*time.Time, error)

	// SetCurrentDate updates the current dump date
	SetCurrentDate(context.Context, time.Time) error
}

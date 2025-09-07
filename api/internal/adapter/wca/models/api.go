package models

type WCALatestData struct {
	ExportDate string `json:"export_date"`
	SqlUrl     string `json:"sql_url"`
	TsvUrl     string `json:"tsv_url"`
}

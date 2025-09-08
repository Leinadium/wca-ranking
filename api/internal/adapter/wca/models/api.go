package models

type WCALatestData struct {
	ExportDate string `json:"export_date"`
	SqlUrl     string `json:"sql_url"`
	TsvUrl     string `json:"tsv_url"`
}

type WCAUserInfo struct {
	Me struct {
		Name    string `json:"name"`
		WCAID   string `json:"wca_id"`
		Country string `json:"country_iso2"`
	} `json:"me"`
}

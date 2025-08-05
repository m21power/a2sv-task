package domain

import "time"

type CMSP struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	LicensedDate string `json:"licensedDate"`
	Description  string `json:"description"`
	Plan         string `json:"plan"`
	Source       string `json:"source"`
}

type FilterModel struct {
	Type           string     `json:"type"`
	LicensedBefore *time.Time `json:"licensedBefore"`
	LicensedAfter  *time.Time `json:"licensedAfter"`
}

type CMSPRepository interface {
	GetAllCMSPs(filter FilterModel) ([]CMSP, error)
	GetCMSPByID(id int) (*CMSP, error)
}

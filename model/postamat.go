package model

import (
	"gorm.io/gorm"
)

type Postamat struct {
	gorm.Model

	Number                     int16  `json:"number"`
	Owner                      string `json: "owner"`
	OfficeCode                 string `json: "officeCode"`
	Region                     string `json: "region"`
	Address                    string `json: "address"`
	FirmwareStatus             string `json:"firmwareStatus"`
	ActivationBlocker          string `json:"activationBlocker"`
	FirmwareVersion            string `json:"firmwareVersion"`
	TerminalId                 int    `json:"terminalId"`
	PaymentOn                  bool   `json:"paymentOn"`
	OnMaintenance              bool   `json:"onMaintenance"`
	MaintenanceAgreementNumber string `json:"maintenanceAgreementNumber"`
	Visible                    bool   `json:"visible"`
	PostamatCellsNumber        string `json:"postamatCellsNumber"`
}

package models

type ElectrometerResponse struct {
	CodigoSuministro string `json:"CodigoSuministro"`
	CodigoMedidor    string `json:"CodigoMedidor"`
	NombreMarca      string `json:"NombreMarca"`
	NombreModelo     string `json:"NombreModelo"`
	DigitosDecimal   string `json:"DigitosDecimal"`
	LecturaActual    string `json:"LecturaActual"`
}

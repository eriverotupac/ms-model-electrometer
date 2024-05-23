package services

import (
	"ms-model-electrometer/internal/models"
)

type IService interface {
	GetInfo(periodo string, sucursal string, zona string) ([]models.ElectrometerResponse, error)
}

package request

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Create struct {
	Name         string `json:"name" validate:"required"`
	SerialNumber string `json:"serial_number" validate:"required"`
	Location     string `json:"location,omitempty"`
	Type         string `json:"type,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
}

func (r *Create) ToEntity() *entity.Machine {
	return &entity.Machine{
		Name:         r.Name,
		SerialNumber: r.SerialNumber,
		Location:     r.Location,
		Type:         r.Type,
		Vendor:       r.Vendor,
		Status:       "offline",
	}
}

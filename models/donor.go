// models/donor.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Donor struct {
	ID           uuid.UUID `json:"id"`
	CampaignID   uuid.UUID `json:"campaign_id"`
	Address      string    `json:"address"`
	Amount       float64   `json:"amount"`
	DonatedAt    time.Time `json:"donated_at"`
}
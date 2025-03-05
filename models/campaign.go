// models/campaign.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID              uuid.UUID `json:"id"`
	Owner           string    `json:"owner"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Target          float64   `json:"target"`
	Deadline        time.Time `json:"deadline"`
	AmountCollected float64   `json:"amount_collected"`
	Image           string    `json:"image"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateCampaignRequest struct {
	Owner       string    `json:"owner" validate:"required,ethereum_address"`
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	Target      float64   `json:"target" validate:"required,gt=0"`
	Deadline    time.Time `json:"deadline" validate:"required,future"`
	Image       string    `json:"image" validate:"required,url"`
}
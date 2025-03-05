// repositories/campaign_repository.go
package repositories

import (
	"database/sql"
	"errors"
	"time"

	"crowdfunding/models"

	"github.com/google/uuid"
)

type CampaignRepository struct {
	db *sql.DB
}

func NewCampaignRepository(db *sql.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (r *CampaignRepository) Create(campaign *models.Campaign) error {
	query := `
		INSERT INTO campaigns 
		(id, owner, title, description, target, deadline, amount_collected, image, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	now := time.Now()
	_, err := r.db.Exec(
		query,
		campaign.ID,
		campaign.Owner,
		campaign.Title,
		campaign.Description,
		campaign.Target,
		campaign.Deadline,
		campaign.AmountCollected,
		campaign.Image,
		now,
		now,
	)

	return err
}

func (r *CampaignRepository) GetByID(id uuid.UUID) (*models.Campaign, error) {
	query := `
		SELECT id, owner, title, description, target, deadline, 
		amount_collected, image, created_at, updated_at 
		FROM campaigns 
		WHERE id = $1
	`

	campaign := &models.Campaign{}
	err := r.db.QueryRow(query, id).Scan(
		&campaign.ID,
		&campaign.Owner,
		&campaign.Title,
		&campaign.Description,
		&campaign.Target,
		&campaign.Deadline,
		&campaign.AmountCollected,
		&campaign.Image,
		&campaign.CreatedAt,
		&campaign.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("campaign not found")
		}
		return nil, err
	}

	return campaign, nil
}

func (r *CampaignRepository) GetAll() ([]models.Campaign, error) {
	query := `
		SELECT id, owner, title, description, target, deadline, 
		amount_collected, image, created_at, updated_at 
		FROM campaigns
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []models.Campaign
	for rows.Next() {
		var campaign models.Campaign
		err := rows.Scan(
			&campaign.ID,
			&campaign.Owner,
			&campaign.Title,
			&campaign.Description,
			&campaign.Target,
			&campaign.Deadline,
			&campaign.AmountCollected,
			&campaign.Image,
			&campaign.CreatedAt,
			&campaign.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}
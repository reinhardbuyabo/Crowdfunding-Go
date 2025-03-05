// services/campaign_service.go
package services

import (
	"errors"
	"time"

	"crowdfunding/models"
	"crowdfunding/repositories"
	"crowdfunding/utils"

	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
)

type CampaignService struct {
	repo      *repositories.CampaignRepository
	validator *validator.Validate
}

func NewCampaignService(
	repo *repositories.CampaignRepository, 
	validator *validator.Validate
) *CampaignService {
	return &CampaignService{
		repo:      repo,
		validator: validator,
	}
}

func (s *CampaignService) CreateCampaign(req *models.CreateCampaignRequest) (*models.Campaign, error) {
	// Validate input
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	// Validate deadline is in the future
	if req.Deadline.Before(time.Now()) {
		return nil, errors.New("deadline must be in the future")
	}

	// Create campaign model
	campaign := &models.Campaign{
		ID:              uuid.New(),
		Owner:           req.Owner,
		Title:           req.Title,
		Description:     req.Description,
		Target:          req.Target,
		Deadline:        req.Deadline,
		AmountCollected: 0,
		Image:           req.Image,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Save to repository
	if err := s.repo.Create(campaign); err != nil {
		return nil, err
	}

	return campaign, nil
}

func (s *CampaignService) GetCampaignByID(id uuid.UUID) (*models.Campaign, error) {
	return s.repo.GetByID(id)
}

func (s *CampaignService) GetAllCampaigns() ([]models.Campaign, error) {
	return s.repo.GetAll()
}
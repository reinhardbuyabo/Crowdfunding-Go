// controllers/campaign_controller.go
package controllers

import (
	"net/http"

	"crowdfunding/services"
	"crowdfunding/models"
	"crowdfunding/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CampaignController struct {
	service *services.CampaignService
}

func NewCampaignController(service *services.CampaignService) *CampaignController {
	return &CampaignController{service: service}
}

func (c *CampaignController) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCampaignRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	campaign, err := c.service.CreateCampaign(&req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, campaign)
}

func (c *CampaignController) GetCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, errors.New("invalid campaign ID"))
		return
	}

	campaign, err := c.service.GetCampaignByID(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, campaign)
}

func (c *CampaignController) GetAllCampaigns(w http.ResponseWriter, r *http.Request) {
	campaigns, err := c.service.GetAllCampaigns()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, campaigns)
}
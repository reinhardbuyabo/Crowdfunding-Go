// routes/routes.go
package routes

import (
	"database/sql"

	"crowdfunding/controllers"
	"crowdfunding/repositories"
	"crowdfunding/services"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	// Initialize validator
	validate := validator.New()

	// Initialize repositories
	campaignRepo := repositories.NewCampaignRepository(db)

	// Initialize services
	campaignService := services.NewCampaignService(campaignRepo, validate)

	// Initialize controllers
	campaignController := controllers.NewCampaignController(campaignService)

	// Campaign routes
	router.HandleFunc("/campaigns", campaignController.CreateCampaign).Methods("POST")
	router.HandleFunc("/campaigns", campaignController.GetAllCampaigns).Methods("GET")
	router.HandleFunc("/campaigns/{id}", campaignController.GetCampaign).Methods("GET")
}
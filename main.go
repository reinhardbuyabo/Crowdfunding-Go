package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Model
type Campaign struct {
	ID              int       `json:"id"`
	Owner           string    `json:"owner"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Target          float64   `json:"target"`
	Deadline        time.Time `json:"deadline"`
	AmountCollected float64   `json:"amountCollected"`
	Image           string    `json:"image"`
}

// Repository Interface
type CampaignRepository interface {
	CreateCampaign(campaign *Campaign) error
	GetCampaigns() ([]Campaign, error)
}

// PostgreSQL Repository Implementation
type PostgresCampaignRepository struct {
	db *sql.DB
}

// Connection String Example
// postgresql://username:password@localhost:5432/dbname?sslmode=disable
func NewPostgresCampaignRepository(connectionString string) (*PostgresCampaignRepository, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresCampaignRepository{db: db}, nil
}

// Create Campaign Method
func (r *PostgresCampaignRepository) CreateCampaign(campaign *Campaign) error {
	query := `
		INSERT INTO campaigns 
		(owner, title, description, target, deadline, amount_collected, image) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id
	`
	err := r.db.QueryRow(
		query,
		campaign.Owner,
		campaign.Title,
		campaign.Description,
		campaign.Target,
		campaign.Deadline,
		campaign.AmountCollected,
		campaign.Image,
	).Scan(&campaign.ID)

	return err
}

// Get Campaigns Method
func (r *PostgresCampaignRepository) GetCampaigns() ([]Campaign, error) {
	query := "SELECT * FROM campaigns"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var c Campaign
		err := rows.Scan(
			&c.ID,
			&c.Owner,
			&c.Title,
			&c.Description,
			&c.Target,
			&c.Deadline,
			&c.AmountCollected,
			&c.Image,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, c)
	}

	return campaigns, nil
}

// Service Layer
type CampaignService struct {
	repo CampaignRepository
}

func NewCampaignService(repo CampaignRepository) *CampaignService {
	return &CampaignService{repo: repo}
}

func (s *CampaignService) CreateCampaign(campaign *Campaign) error {
	// Additional business logic can be added here
	return s.repo.CreateCampaign(campaign)
}

func (s *CampaignService) GetCampaigns() ([]Campaign, error) {
	return s.repo.GetCampaigns()
}

// Controller
type CampaignController struct {
	service *CampaignService
}

func NewCampaignController(service *CampaignService) *CampaignController {
	return &CampaignController{service: service}
}

func (c *CampaignController) CreateCampaign(ctx *gin.Context) {
	var campaign Campaign
	if err := ctx.BindJSON(&campaign); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCampaign(&campaign); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, campaign)
}

func (c *CampaignController) GetCampaigns(ctx *gin.Context) {
	campaigns, err := c.service.GetCampaigns()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, campaigns)
}

// env variables
// DatabaseConfig struct to hold configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadDatabaseConfig reads configuration from environment variables
func LoadDatabaseConfig() DatabaseConfig {
	// Load .env file if it exists (for local development)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "crowdfunding"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// getEnv retrieves an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// BuildConnectionString constructs the PostgreSQL connection string
func BuildConnectionString(config DatabaseConfig) string {
	// Validate required fields
	if config.User == "" || config.Password == "" {
		log.Fatal("Database user and password must be provided")
	}

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
}

// Main Setup
func main() {
	// Load database configuration
	dbConfig := LoadDatabaseConfig()

	// Build connection string
	connectionString := BuildConnectionString(dbConfig)

	// Sample Connection String
	// connectionString := "postgres://username:password@localhost:5432/crowdfunding?sslmode=disable"

	// Create Repository
	repo, err := NewPostgresCampaignRepository(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create Service
	service := NewCampaignService(repo)

	// Create Controller
	controller := NewCampaignController(service)

	// Setup Router
	router := gin.Default()

	// Routes
	router.POST("/api/campaigns", controller.CreateCampaign)
	router.GET("/api/campaigns", controller.GetCampaigns)

	// Start Server
	router.Run(":8080")
}

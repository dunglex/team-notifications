package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Appplication struct {
	Config            *AppConfig
	WebApp            *fiber.App
	SelfHealthChecker *time.Ticker
}

var app Appplication

func init() {
	// Load the config
	app.Config = &AppConfig{}
	var err = app.Config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		os.Exit(1)
	}

	// Print the config to the console
	fmt.Println("WebhookURL:" + app.Config.WebhookURL)
	fmt.Println("HealthCheckURL:" + app.Config.HealthCheckURL)
	fmt.Println("HealthCheckInterval:", app.Config.HealthCheckInterval)
	fmt.Println("Port:" + app.Config.Port)

	// Initialize the Fiber app
	app.WebApp = fiber.New()

	// Init the self health checker
	app.SelfHealthChecker = time.NewTicker(time.Duration(app.Config.HealthCheckInterval) * time.Second)
}

func main() {

	app.WebApp.Post("/pr/created", prCreatedHandler) // Endpoint for handling new pull request notifications
	app.WebApp.Get("/healthz", healthCheckHandler)   // Endpoint for health checks

	go app.StartSelfHealthCheck() // Start the periodic health check

	fmt.Printf("Server is running on port %s\n", app.Config.Port)
	app.WebApp.Listen(fmt.Sprintf(":%s", app.Config.Port))
}

func prCreatedHandler(c *fiber.Ctx) error {
	var payload PrCreatedRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	// Dump the payload to the console
	requestBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("Received new pull request notification: %v", string(requestBytes))

	// Convert the payload to a PullRequest object
	pr := payload.ToPullRequest()

	// Create the Adaptive Card object
	card := pr.CreateAdaptiveCard()

	// Send the Adaptive Card to the Power Automate webhook URL
	card.sendAdaptiveCard(app.Config.WebhookURL, true)

	return c.Status(fiber.StatusCreated).SendString("New pull request has notification has been sent to power automate")
}

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func selfHealthCheck() {
	resp, err := http.Get(app.Config.HealthCheckURL)
	if err != nil {
		log.Printf("Health check failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Health check failed with status code: %d", resp.StatusCode)
	} else {
		log.Println("Health check succeeded")
	}
}

func (app *Appplication) StartSelfHealthCheck() {
	defer app.SelfHealthChecker.Stop()

	for range app.SelfHealthChecker.C {
		selfHealthCheck()
	}
}

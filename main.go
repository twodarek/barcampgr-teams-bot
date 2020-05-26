package main

import (
	"fmt"
	"github.com/twodarek/barcampgr-teams-bot/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"github.com/twodarek/barcampgr-teams-bot/server"
)

func main() {
	log.Println("Attempting to start barcampgr-teams-bot")
	router := mux.NewRouter()

	httpClient := &http.Client{}

	conf := barcampgr.Config{
		APIToken: os.Getenv("CISCO_TEAMS_API_TOKEN"),
		MySqlUser: os.Getenv("MYSQL_USER"),
		MySqlPass: os.Getenv("MYSQL_PASS"),
		MySqlServer: os.Getenv("MYSQL_SERVER"),
		MySqlPort: os.Getenv("MYSQL_PORT"),
		MySqlDatabase: os.Getenv("MYSQL_DATABASE"),
		AdminPassword: os.Getenv("BARCAMPGR_ADMIN_PASSWORD"),
		WebexRoomID: os.Getenv("WEBEX_ROOM_ID"),
		WebexCallbackURL: os.Getenv("WEBEX_CALLBACK_URL"),
	}
	log.Println("Attempting to start webex teams client")
	teamsClient := webexteams.NewClient()
	initTeamsClient(teamsClient, conf)
	log.Println("Webex teams client started, connecting to database")

	sdb := database.NewDatabase(conf.MySqlUser, conf.MySqlPass, conf.MySqlServer, conf.MySqlPort, conf.MySqlDatabase)
	log.Println("Database connected")

	ac := barcampgr.NewAppController(
		teamsClient,
		httpClient,
		sdb,
		conf,
	)

	s := server.New(ac, conf, router)

	// Multiple codepaths use the DefaultServeMux so we start listening at the top
	go http.ListenAndServe("0.0.0.0:8080", s)

	log.Println("Barcampgr-teams-bot started")

	select {}
}

func initTeamsClient(client *webexteams.Client, config barcampgr.Config) error {
	client.SetAuthToken(config.APIToken)
	webHookURL := config.WebexCallbackURL

	// Clean up old webhooks
	webhooksQueryParams := &webexteams.ListWebhooksQueryParams{
		Max: 10,
	}

	webhooks, _, err := client.Webhooks.ListWebhooks(webhooksQueryParams)
	if err != nil {
		log.Printf("Unable to get old webhooks, continuing anyway")
	}
	for _, webhook := range webhooks.Items {
		_, err := client.Webhooks.DeleteWebhook(webhook.ID)
		if err != nil {
			log.Printf("Unable to clean up old webhook")
		}
	}

	// Create new webhook
	webhookRequest := &webexteams.WebhookCreateRequest{
		Name:      "BarCampGR Webhook - Test",
		TargetURL: webHookURL,
		Resource:  "messages",
		Event:     "created",

	}

	testWebhook, _, err := client.Webhooks.CreateWebhook(webhookRequest)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to create webhook: %s", err))
	}

	log.Printf("Created webhook. ID: %s, Name: %s, target URL: %s, created: %s", testWebhook.ID, testWebhook.Name, testWebhook.TargetURL, testWebhook.Created)
	return nil
}
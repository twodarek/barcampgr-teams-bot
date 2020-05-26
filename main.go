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
	}
	teamsClient := webexteams.NewClient()
	initTeamsClient(teamsClient, conf)

	sdb := database.NewDatabase(conf)

	ac := barcampgr.NewAppController(
		teamsClient,
		httpClient,
		sdb,
		conf,
	)

	s := server.New(ac, conf, router)

	// Multiple codepaths use the DefaultServeMux so we start listening at the top
	go http.ListenAndServe("0.0.0.0:8080", s)

	select {}
}

func initTeamsClient(client *webexteams.Client, config barcampgr.Config) error {
	client.SetAuthToken(config.APIToken)
	myRoomID := ""   // Change to your testing room
	webHookURL := "" // Change this to your test URL

	// POST webhooks

	webhookRequest := &webexteams.WebhookCreateRequest{
		Name:      "BarCampGR Webhook - Test",
		TargetURL: webHookURL,
		Resource:  "messages",
		Event:     "created",
		Filter:    "roomId=" + myRoomID,
	}

	testWebhook, _, err := client.Webhooks.CreateWebhook(webhookRequest)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to create webhook: %s", err))\
	}

	fmt.Printf("Created webhook. ID: %s, Name: %s, target URL: %s, created: %s", testWebhook.ID, testWebhook.Name, testWebhook.TargetURL, testWebhook.Created)
	return nil
}
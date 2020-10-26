package main

import (
	"fmt"
	"github.com/twodarek/barcampgr-teams-bot/barcampgr/slack"
	"github.com/twodarek/barcampgr-teams-bot/barcampgr/teams"
	"github.com/twodarek/barcampgr-teams-bot/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	webexteams "github.com/twodarek/go-cisco-webex-teams/sdk"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"github.com/twodarek/barcampgr-teams-bot/server"
)

func main() {
	log.Println("Attempting to start barcampgr-teams-bot")
	router := mux.NewRouter()

	httpClient := &http.Client{}

	conf := barcampgr.Config{
		TeamsAPIToken:    os.Getenv("CISCO_TEAMS_API_TOKEN"),
		BarCampGRWebexId: os.Getenv("BARCAMPGR_WEBEX_ID"),
		BaseCallbackURL:  os.Getenv("BARCAMPGR_BASE_CALLBACK_URL"),
		MySqlUser:        os.Getenv("MYSQL_USER"),
		MySqlPass:        os.Getenv("MYSQL_PASS"),
		MySqlServer:      os.Getenv("MYSQL_SERVER"),
		MySqlPort:        os.Getenv("MYSQL_PORT"),
		MySqlDatabase:    os.Getenv("MYSQL_DATABASE"),
		AdminPassword:    os.Getenv("BARCAMPGR_ADMIN_PASSWORD"),
		InvitePassword:   os.Getenv("BARCAMPGR_INVITE_PASSWORD"),
		WebexTeamID:      os.Getenv("BARCAMPGR_TEAM_ID"),
		WebexOrgID: os.Getenv("WEBEX_ORG_ID"),
		WebexRoomID: os.Getenv("WEBEX_ROOM_ID"),
		WebexCallbackURL: os.Getenv("WEBEX_CALLBACK_URL"),
		WebexMembershipCallbackURL: os.Getenv("WEBEX_MEMBERSHIP_CALLBACK_URL"),
	}
	conf.SetWebexAllRooms(os.Getenv("WEBEX_ALL_ROOMS"))
	log.Println("Attempting to start webex teams client")
	teamsClient := webexteams.NewClient()
	initTeamsClient(teamsClient, conf)
	log.Println("Webex teams client started, connecting to database")

	sdb := database.NewDatabase(conf.MySqlUser, conf.MySqlPass, conf.MySqlServer, conf.MySqlPort, conf.MySqlDatabase)
	log.Println("Database connected")

	ac := barcampgr.NewAppController(
		httpClient,
		sdb,
		conf,
	)

	sac := slack.NewAppController(
		ac,
		teamsClient,
		httpClient,
		sdb,
		conf,
	)

	tac := teams.NewAppController(
		ac,
		teamsClient,
		httpClient,
		sdb,
		conf,
	)

	s := server.New(ac, sac, tac, conf, router)

	// Multiple codepaths use the DefaultServeMux so we start listening at the top
	go http.ListenAndServe("0.0.0.0:8080", s)

	log.Println("Barcampgr-teams-bot started")

	select {}
}

func initSlackClient() error {
	return nil
}

func initTeamsClient(client *webexteams.Client, config barcampgr.Config) error {
	client.SetAuthToken(config.TeamsAPIToken)

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
			log.Printf("Unable to clean up old webhook %s on endpoint %s", webhook.ID, webhook.TargetURL)
		}
	}

	// Create new @bot message webhook
	webhookRequest := &webexteams.WebhookCreateRequest{
		Name:      "BarCampGR Webhook",
		TargetURL: fmt.Sprintf("%s%s", config.BaseCallbackURL, config.WebexCallbackURL),
		Resource:  "messages",
		Event:     "created",

	}

	testWebhook, _, err := client.Webhooks.CreateWebhook(webhookRequest)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to create webhook: %s", err))
	}

	log.Printf("Created chatop webhook. ID: %s, Name: %s, target URL: %s, created: %s", testWebhook.ID, testWebhook.Name, testWebhook.TargetURL, testWebhook.Created)

	// Create new memberships webhook
	membershipWebhookRequest := &webexteams.WebhookCreateRequest{
		Name:      "BarCampGR Memberships Webhook",
		TargetURL: fmt.Sprintf("%s%s", config.BaseCallbackURL, config.WebexMembershipCallbackURL),
		Resource:  "memberships",
		Event:     "created",
		Filter:    fmt.Sprintf("roomId=%s", config.WebexRoomID),
	}

	testMembershipWebhook, _, err := client.Webhooks.CreateWebhook(membershipWebhookRequest)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to create webhook: %s", err))
	}
	log.Printf("Created membership webhook. ID: %s, Name: %s, target URL: %s, created: %s", testMembershipWebhook.ID, testMembershipWebhook.Name, testMembershipWebhook.TargetURL, testMembershipWebhook.Created)

	return nil
}
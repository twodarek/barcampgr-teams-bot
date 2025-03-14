package barcampgr

import (
	"strings"
)

type Config struct {
	SlackAPIToken              string
	TeamsAPIToken              string
	DiscordAPIToken            string
	BarCampGRWebexId           string
	BaseCallbackURL            string
	MySqlUser                  string
	MySqlPass                  string
	MySqlServer                string
	MySqlPort                  string
	MySqlDatabase              string
	AdminPassword              string
	InvitePassword             string
	DiscordAppId               string
	DiscordPublicKey           string
	SlackCallbackURL           string
	SlackUsername              string
	SlackVerificationToken     string
	WebexTeamID                string
	WebexRoomID                string
	WebexOrgID                 string
	WebexCallbackURL           string
	WebexMembershipCallbackURL string
	WebexAllRooms              []string
}

func New(
	slackApiToken string,
	teamsApiToken string,
	discordApiToken string,
	barCampGRWebexId string,
	baseCallbackURL string,
	mySqlUser string,
	mySqlPass string,
	mySqlServer string,
	mySqlPort string,
	mySqlDatabase string,
	adminPassword string,
	invitePassword string,
	discordAppId string,
	discordPublicKey string,
	slackCallbackURL string,
	slackUsername string,
	slackVerificationToken string,
	webexTeamID string,
	webexRoomID string,
	webexOrgID string,
	webexCallbackURL string,
	webexMembershipCallbackURL string,
	webexAllRooms []string,
) *Config {
	c := &Config{
		SlackAPIToken:              slackApiToken,
		TeamsAPIToken:              teamsApiToken,
		DiscordAPIToken:            discordApiToken,
		BarCampGRWebexId:           barCampGRWebexId,
		BaseCallbackURL:            baseCallbackURL,
		MySqlUser:                  mySqlUser,
		MySqlPass:                  mySqlPass,
		MySqlServer:                mySqlServer,
		MySqlPort:                  mySqlPort,
		MySqlDatabase:              mySqlDatabase,
		AdminPassword:              adminPassword,
		InvitePassword:             invitePassword,
		DiscordAppId:               discordAppId,
		DiscordPublicKey:           discordPublicKey,
		SlackCallbackURL:           slackCallbackURL,
		SlackUsername:              slackUsername,
		SlackVerificationToken:     slackVerificationToken,
		WebexTeamID:                webexTeamID,
		WebexRoomID:                webexRoomID,
		WebexOrgID:                 webexOrgID,
		WebexCallbackURL:           webexCallbackURL,
		WebexMembershipCallbackURL: webexMembershipCallbackURL,
		WebexAllRooms:              webexAllRooms,
	}
	return c
}

func (c *Config) SetWebexAllRooms(input string) {
	c.WebexAllRooms = strings.Split(input, ":")
}

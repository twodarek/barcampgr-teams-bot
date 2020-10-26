package barcampgr

import (
	"strings"
)

type Config struct {
	SlackAPIToken	 string
	TeamsAPIToken    string
	BarCampGRWebexId string
	BaseCallbackURL  string
	MySqlUser        string
	MySqlPass        string
	MySqlServer      string
	MySqlPort        string
	MySqlDatabase    string
	AdminPassword    string
	InvitePassword   string
	SlackCallbackURL string
	SlackUsername    string
	SlackVerificationToken string
	WebexTeamID      string
	WebexRoomID string
	WebexOrgID string
	WebexCallbackURL string
	WebexMembershipCallbackURL string
	WebexAllRooms []string
}

func New(
	slackApiToken string,
	teamsApiToken string,
	barCampGRWebexId string,
	baseCallbackURL string,
	mySqlUser string,
	mySqlPass string,
	mySqlServer string,
	mySqlPort string,
	mySqlDatabase string,
	adminPassword string,
	invitePassword string,
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
		SlackAPIToken:    slackApiToken,
		TeamsAPIToken:    teamsApiToken,
		BarCampGRWebexId: barCampGRWebexId,
		BaseCallbackURL:  baseCallbackURL,
		MySqlUser:        mySqlUser,
		MySqlPass:        mySqlPass,
		MySqlServer:      mySqlServer,
		MySqlPort:        mySqlPort,
		MySqlDatabase:    mySqlDatabase,
		AdminPassword:    adminPassword,
		InvitePassword:   invitePassword,
		SlackCallbackURL: slackCallbackURL,
		SlackUsername:    slackUsername,
		SlackVerificationToken: slackVerificationToken,
		WebexTeamID:      webexTeamID,
		WebexRoomID: webexRoomID,
		WebexOrgID: webexOrgID,
		WebexCallbackURL: webexCallbackURL,
		WebexMembershipCallbackURL: webexMembershipCallbackURL,
		WebexAllRooms: webexAllRooms,
	}
	return c
}

func (c *Config) SetWebexAllRooms(input string) {
	c.WebexAllRooms = strings.Split(input, ":")
}
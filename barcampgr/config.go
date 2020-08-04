package barcampgr

import "strings"

type Config struct {
	APIToken string
	BarCampGRWebexId string
	MySqlUser string
	MySqlPass string
	MySqlServer string
	MySqlPort string
	MySqlDatabase string
	AdminPassword string
	InvitePassword string
	WebexTeamID string
	WebexRoomID string
	WebexOrgID string
	WebexCallbackURL string
	WebexMembershipCallbackURL string
	WebexAllRooms []string
}

func New(
	apiToken string,
	barCampGRWebexId string,
	mySqlUser string,
	mySqlPass string,
	mySqlServer string,
	mySqlPort string,
	mySqlDatabase string,
	adminPassword string,
	invitePassword string,
	webexTeamID string,
	webexRoomID string,
	webexOrgID string,
	webexCallbackURL string,
	webexMembershipCallbackURL string,
	webexAllRooms []string,
) *Config {
	c := &Config{
		APIToken: apiToken,
		BarCampGRWebexId: barCampGRWebexId,
		MySqlUser: mySqlUser,
		MySqlPass: mySqlPass,
		MySqlServer: mySqlServer,
		MySqlPort: mySqlPort,
		MySqlDatabase: mySqlDatabase,
		AdminPassword: adminPassword,
		InvitePassword: invitePassword,
		WebexTeamID: webexTeamID,
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
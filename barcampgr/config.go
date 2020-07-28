package barcampgr

type Config struct {
	APIToken string
	BarCampGRWebexId string
	MySqlUser string
	MySqlPass string
	MySqlServer string
	MySqlPort string
	MySqlDatabase string
	AdminPassword string
	WebexRoomID string
	WebexCallbackURL string
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
	webexRoomID string,
	webexCallbackURL string,
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
		WebexRoomID: webexRoomID,
		WebexCallbackURL: webexCallbackURL,
	}
	return c
}
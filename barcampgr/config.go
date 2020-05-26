package barcampgr

type Config struct {
	APIToken string
	MySqlUser string
	MySqlPass string
	MySqlServer string
	MySqlPort string
	MySqlDatabase string
	AdminPassword string
}

func New(
	apiToken string,
	mySqlUser string,
	mySqlPass string,
	mySqlServer string,
	mySqlPort string,
	mySqlDatabase string,
	adminPassword string,
) *Config {
	c := &Config{
		APIToken: apiToken,
		MySqlUser: mySqlUser,
		MySqlPass: mySqlPass,
		MySqlServer: mySqlServer,
		MySqlPort: mySqlPort,
		MySqlDatabase: mySqlDatabase,
		AdminPassword: adminPassword,
	}
	return c
}
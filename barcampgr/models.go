package barcampgr

type Schedule struct {
	RefreshedAt string `json:"refreshedAt"`
	LastUpdate  string `json:"lastUpdate"`
	Times       []ScheduleTime `json:"times"`
	Rows        []ScheduleRow `json:"rows"`
}

type ScheduleTime struct {
	ID int `json:"id"`
	Start string `json:"start"`
	End string `json:"end"`
	Day string `json:"day"`
	Displayable bool `json:"displayable"`
}

type ScheduleRow struct {
	Room string `json:"room"`
	Sessions []ScheduleSession `json:"sessions"`
}

type ScheduleSession struct {
	Time int `json:"time,string"`
	Room int `json:"room,string"`
	Title string `json:"title"`
	Speaker string `json:"speaker"`
	UniqueString string `json:"uniqueString"`
	Version int `json:"-"`
	AltText string `json:"altText"`
}

type ScheduleRoom struct {
	Name string `json:"name"`
	ID int `json:"id"`
}

type InvitePerson struct {
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

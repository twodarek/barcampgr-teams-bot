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
}

type ScheduleRow struct {
	Room string `json:"room"`
	Sessions []ScheduleSession `json:"sessions"`
}

type ScheduleSession struct {
	Time int `json:"time"`
	Room int `json:"room"`
	Title string `json:"title"`
	Speaker string `json:"speaker"`
}

type ScheduleRoom struct {
	Name string `json:"name"`
	ID int `json:"id"`
}

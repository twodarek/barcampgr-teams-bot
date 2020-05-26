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
	time int `json:"time"`
	room int `json:"room"`
	title string `json:"title"`
	speaker string `json:"speaker"`
}
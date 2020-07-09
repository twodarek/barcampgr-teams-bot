package barcampgr

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"

	"github.com/twodarek/barcampgr-teams-bot/database"
)

type Controller struct {
	teamsClient	*webexteams.Client
	httpClient   *http.Client
	sdb          *database.ScheduleDatabase
	config    Config
}

func NewAppController(
	teamsClient	*webexteams.Client,
	httpClient *http.Client,
	sdb        *database.ScheduleDatabase,
	config Config,

) *Controller {
	return &Controller{
		teamsClient:  teamsClient,
		httpClient:   httpClient,
		sdb:          sdb,
		config:     config,
	}
}

func (ac *Controller) HandleChatop(requestData webexteams.WebhookRequest) (string, error) {
	message, _, err := ac.teamsClient.Messages.GetMessage(requestData.Data.ID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to get message id %s", requestData.Data.ID))
	}
	log.Printf("Received message `%s` as message id %s in room %s from person %s", message.Text, requestData.Data.ID, message.RoomID, message.PersonID)
	// chatbot_1  | 2020/05/26 03:34:56 Received message `BarcampGRBot this is a test` as message id Y2lzY29zcGFyazovL3VzL01FU1NBR0UvZGU3ZmMyYzAtOWYwMS0xMWVhLWE0YWUtZDk2MjUyNjYwNzI2 in room Y2lzY29zcGFyazovL3VzL1JPT00vMDVlMjg2NzAtOWUyZC0xMWVhLTkwZGItOWRlOGYwYmY1NzZk from person Y2lzY29zcGFyazovL3VzL1BFT1BMRS9jZjNlMGU4Zi0wZmY3LTRjYzgtODM5MS05NTIxNzQzYjVkMzI
	person, _, err :=ac.teamsClient.People.GetPerson(message.PersonID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to get person id %s", message.PersonID))
	}
	if person.ID == "Y2lzY29zcGFyazovL3VzL1BFT1BMRS9lYTZhNWVlOC02Y2VjLTQxNzUtYjk5Mi03NGZhMzcwMmU2ZDc" {
		log.Printf("Rejecting message from myself, returning cleanly")
		return "", nil
	}
	log.Printf("Got message from person.  Display: %s, Nick: %s, Name: %s %s", person.DisplayName, person.NickName, person.FirstName, person.LastName)
	room, _, err := ac.teamsClient.Rooms.GetRoom(message.RoomID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to get room id %s", message.RoomID))
	}
	log.Printf("Get message from room. Title: %s, Type: %s", room.Title, room.RoomType)

	//TODO(twodarek): Figure out what the message sent was and deal with it
	replyText, err := ac.handleCommand(message.Text, person.DisplayName)
	if err != nil || replyText == "" {
		replyText = fmt.Sprintf("Hello %s!  I have received your request of '%s', but I don't know how to do that right now.", person.DisplayName, message.Text)
	}
	replyMessage := &webexteams.MessageCreateRequest{
		RoomID:        message.RoomID,
		Markdown:      replyText,
	}
	_, resp, err := ac.teamsClient.Messages.CreateMessage(replyMessage)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to reply to message %s", message.ID))
	}
	log.Printf("Replied with %s, got http code %d, body %s", replyText, resp.StatusCode(), resp.Body())


	return "", nil
}

func (ac *Controller) handleCommand (message, displayName string) (string, error) {
	message = strings.TrimPrefix(message, "BarcampGRBot")
	message = strings.TrimPrefix(message, " ")
	commandArray := strings.Split(message, " ")
	switch strings.ToLower(commandArray[0]) {
		case "schedule":
			log.Printf("Hi %s, I'm attempting to schedule block with message %s, commandArray %s", displayName, message, commandArray)
			return "I'm attempting to schedule you for a talk", nil
		case "test":
			log.Printf("Test message %s, commandArray %s", message, commandArray)
			return fmt.Sprintf("Hi Test!!!!, I received your message of %s from %s", message, displayName), nil
		case "help":
			return fmt.Sprintf("I accept the following commands:\n - `Schedule me in ROOM at START_TIME for TITLE` to schedule a talk\n - `test MESSAGE_TO_ECHO` to test this bot\n - `help` to get this message"), nil
		default:
			return "", errors.New("command unknown")
	}
}

func (ac *Controller) GetScheduleJson() (Schedule, error) {
	var times []database.DBScheduleTime
	var sessions []database.DBScheduleSession
	var rooms []database.DBScheduleRoom

	ac.sdb.Orm.Find(&times)
	outTimes := ac.convertTimes(times)

	ac.sdb.Orm.Find(&sessions)
	ac.sdb.Orm.Find(&rooms)

	outRows := ac.buildRows(sessions, rooms)

	schedule := Schedule{
		RefreshedAt: "",
		LastUpdate:  "",
		Times:       outTimes,
		Rows:        outRows,
	}

	return schedule, nil
}

func (ac *Controller) GetTimesJson() ([]ScheduleTime, error) {
	var times []database.DBScheduleTime
	ac.sdb.Orm.Find(&times)
	return ac.convertTimes(times), nil
}

func (ac *Controller) GetRoomsJson() ([]ScheduleRoom, error) {
	var rooms []database.DBScheduleRoom
	ac.sdb.Orm.Find(&rooms)
	return ac.convertRooms(rooms), nil
}

func (ac *Controller) convertTimes(times []database.DBScheduleTime) []ScheduleTime {
	var resultant []ScheduleTime
	for _, t := range times {
		resultant = append(resultant, ScheduleTime{
			ID:    int(t.ID),
			Start: t.Start,
			End:   t.End,
		})
	}
	return resultant
}

func (ac *Controller) convertRooms(rooms []database.DBScheduleRoom) []ScheduleRoom {
	var resultant []ScheduleRoom
	for _, r := range rooms {
		resultant = append(resultant, ScheduleRoom{
			Name: r.Name,
			ID:   int(r.ID),
		})
	}
	return resultant
}

func (ac *Controller) buildRows(sessions []database.DBScheduleSession, rooms []database.DBScheduleRoom) []ScheduleRow {
	var resultant []ScheduleRow

	for _,r := range rooms {
		resultant = append(resultant, ScheduleRow{
			Room:    r.Name,
			Sessions: ac.getSessionsInRoom(r.Name, sessions),
		})
	}
	return resultant
}

func (ac *Controller) MigrateDB() error {
	return ac.sdb.MigrateDB()
}

func (ac *Controller) getSessionsInRoom(name string, sessions []database.DBScheduleSession) []ScheduleSession {
	var resultant []ScheduleSession
	for _,s := range sessions {
		if s.Room.Name == name {
			resultant = append(resultant, ScheduleSession{
				Time:    int(s.Time.ID),
				Room:    int(s.Room.ID),
				Title:   s.Title,
				Speaker: s.Speaker,
			})
		}
	}
	return resultant
}

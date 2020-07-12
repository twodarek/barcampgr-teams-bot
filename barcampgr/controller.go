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
		replyText = fmt.Sprintf("Hello %s!  I have received your request of '%s', but I'm unable to do that right now.  Message: %s, Error: %s", person.DisplayName, message.Text, replyText, err)
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
			log.Printf("I'm attempting to schedule block with message %s, commandArray %s, for %s", message, commandArray, displayName)
			message, err := ac.parseAndScheduleTalk(displayName, commandArray[1:])
			return message, err
		case "test":
			log.Printf("Test message %s, commandArray %s", message, commandArray)
			return fmt.Sprintf("Hi Test!!!!, I received your message of %s from %s", message, displayName), nil
		case "help":
			return fmt.Sprintf("I accept the following commands:\n - `Schedule me at START_TIME in ROOM for TITLE` to schedule a talk\n - `test MESSAGE_TO_ECHO` to test this bot\n - `help` to get this message"), nil
		default:
			return "", errors.New(fmt.Sprintf("Unknown command %s", ac.commandArrayToString(commandArray)))
	}
}

func (ac *Controller) parseAndScheduleTalk(displayName string, commandArray []string) (string, error) {
	var name string
	var time string
	var room string
	var title string
	currentArrayPos := 0
	message := ""
	// me at 10:00am in The Hotdog Stand for Speaking to Bots, a Minecraft Story

	if strings.ToLower(commandArray[0]) == "me" {
		name = displayName
		currentArrayPos = 1
		// skip 'at' command word
	} else {
		for i, s := range commandArray {
			name = name + " " + s
			if ac.isCommandWord(commandArray[i + 1]) {
				currentArrayPos = i + 1
				break
			}
		}
	}
	name = strings.TrimPrefix(name, " ")

	// at 10:00am in The Hotdog Stand for Speaking to Bots, a Minecraft Story
	// skip 'at' command word
	currentArrayPos++
	time = commandArray[currentArrayPos]

	// in The Hotdog Stand for Speaking to Bots, a Minecraft Story
	// skip 'in' command word
	currentArrayPos++
	for i, s := range commandArray[currentArrayPos:] {
		room = room + " " + s
		if ac.isCommandWord(commandArray[i + currentArrayPos]) {
			currentArrayPos = i + currentArrayPos
			break
		}
	}
	room = strings.TrimPrefix(room, " ")

	// for Speaking to Bots, a Minecraft Story
	// skip 'for' command word
	currentArrayPos++
	for _, s := range commandArray[currentArrayPos:] {
		title = title + " " + s
	}
	title = strings.TrimPrefix(title, " ")

	var timeObj database.DBScheduleTime
	result := ac.sdb.Orm.Where("start = ?", time).Find(&timeObj)
	if result.Error != nil {
		log.Printf("Received error %s when trying to query for time starting at %s", result.Error, time)
		return fmt.Sprintf("Unable to find a time starting at %s", time), result.Error
	}

	var roomObj database.DBScheduleRoom
	result = ac.sdb.Orm.Where("name = ?", time).Find(&roomObj)
	if result.Error != nil {
		log.Printf("Received error %s when trying to query for room %s", result.Error, room)
		return fmt.Sprintf("Unable to find a room named %s", room), result.Error
	}

	session := database.DBScheduleSession{
		Time:    timeObj,
		Room:    roomObj,
		Updater: displayName,
		Title:   title,
		Speaker: name,
	}

	result = ac.sdb.Orm.Create(&session)
	if result.Error != nil {
		log.Printf("Received error %s when trying to create talk %s", result.Error, ac.commandArrayToString(commandArray))
		return message, result.Error
	} else {
		log.Printf("Created talk %s with %d rows affected", ac.commandArrayToString(commandArray), result.RowsAffected)
		message = fmt.Sprintf("I've scheduled your session %s", session.ToString())
	}

	return message, nil
}

func (ac *Controller) commandArrayToString(array []string) string {
	var resultant string
	for _,s := range array {
		resultant = resultant + " " + s
	}
	return strings.TrimPrefix(resultant, " ")
}

func (ac *Controller) isCommandWord(check string) bool {
	switch check {
	case "at", "in", "for":
		return true
	default:
		return false
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

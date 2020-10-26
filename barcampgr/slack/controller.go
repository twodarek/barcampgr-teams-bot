package slack

import (
	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"math/rand"
	"net/http"
	"time"

	webexteams "github.com/twodarek/go-cisco-webex-teams/sdk"

	"github.com/twodarek/barcampgr-teams-bot/database"
)

type Controller struct {
	bc  *barcampgr.Controller
	slackClient *webexteams.Client
	httpClient  *http.Client
	sdb         *database.ScheduleDatabase
	config      barcampgr.Config
	sRand       *rand.Rand
}

func NewAppController(
	barcampgrController *barcampgr.Controller,
	slackClient	*webexteams.Client,
	httpClient  *http.Client,
	sdb         *database.ScheduleDatabase,
	config barcampgr.Config,
) *Controller {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Controller{
		bc:    barcampgrController,
		slackClient:  slackClient,
		httpClient:   httpClient,
		sdb:          sdb,
		config:       config,
		sRand:        seededRand,
	}
}
//
//const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//const help_message = "I accept the following commands:\n - `help` to get this message\n - `get schedule`, `get grid`, or `get talks` to get a link to the schedule grid\n - `get links` to get all of the unique links for your talks\n - `dm` to open a direct message connection with me\n - `Schedule me at START_TIME in ROOM for TITLE` to schedule a talk\n - `Schedule web` to schedule a talk via web form\n\nMake sure to `@barcampgrbot` at the start or I won't get the message!"
//
//func (ac *Controller) HandleChatop(requestData webexteams.WebhookRequest) (string, error) {
//	// Filter to make sure it's only from BarCampGR
//	if requestData.OrgID != ac.config.WebexOrgID {
//		return "", errors.New(fmt.Sprintf("Unable to handle messages from non-BarCampGR orgs %s", requestData.Data.ID))
//	}
//
//	message, _, err := ac.teamsClient.Messages.GetMessage(requestData.Data.ID)
//	if err != nil {
//		return "", errors.New(fmt.Sprintf("Unable to get message id %s", requestData.Data.ID))
//	}
//	log.Printf("Received message `%s` as message id %s in room %s from person %s", message.Text, requestData.Data.ID, message.RoomID, message.PersonID)
//
//	person, _, err :=ac.teamsClient.People.GetPerson(message.PersonID)
//	if err != nil {
//		return "", errors.New(fmt.Sprintf("Unable to get person id %s", message.PersonID))
//	}
//	if person.ID == ac.config.BarCampGRWebexId {
//		log.Printf("Rejecting message from myself, returning cleanly")
//		return "", nil
//	}
//	log.Printf("Got message from person.  Display: %s, Nick: %s, Name: %s %s", person.DisplayName, person.NickName, person.FirstName, person.LastName)
//
//	room, _, err := ac.teamsClient.Rooms.GetRoom(message.RoomID)
//	if err != nil {
//		return "", errors.New(fmt.Sprintf("Unable to get room id %s", message.RoomID))
//	}
//	log.Printf("Get message from room. Title: %s, Type: %s", room.Title, room.RoomType)
//
//	replyText, dmText, err := ac.handleCommand(message.Text, person)
//	if err != nil {
//		replyText = fmt.Sprintf("Hello %s!  I have received your request of '%s', but I'm unable to do that right now.  Message: %s, Error: %s", person.DisplayName, message.Text, replyText, err)
//	}
//
//	if replyText != "" {
//		replyMessage := &webexteams.MessageCreateRequest{
//			RoomID:   message.RoomID,
//			Markdown: replyText,
//		}
//		_, resp, err := ac.teamsClient.Messages.CreateMessage(replyMessage)
//		if err != nil {
//			return "", errors.New(fmt.Sprintf("Unable to reply to message %s", message.ID))
//		}
//		log.Printf("Replied with %s, got http code %d, body %s", replyText, resp.StatusCode(), resp.Body())
//	}
//
//	if dmText != "" {
//		err = ac.sendDM(person.ID, dmText)
//		if err != nil {
//			return "", errors.New(fmt.Sprintf("Unable to send dm reply to message %s", message.ID))
//		}
//	}
//
//	return "", nil
//}
//
//func (ac *Controller) sendDM(personID, message string) error {
//	replyDmMessage := &webexteams.MessageCreateRequest{
//		ToPersonID:		personID,
//		Markdown:      message,
//	}
//	_, resp, err := ac.teamsClient.Messages.CreateMessage(replyDmMessage)
//	log.Printf("Attempted to DM with %s, got http code %d, body %s", message, resp.StatusCode(), resp.Body())
//	return err
//}
//
//func (ac *Controller) InviteNewPeople(requestData webexteams.WebhookRequest) (string, error) {
//	personID := requestData.Data.PersonID
//	if personID == "" {
//		log.Printf("Webhook from Cisco did not include PersonID")
//		return "", errors.New("No person ID provided")
//	}
//	log.Printf("Inviting Person %s to all rooms", personID)
//	for _, roomID := range ac.config.WebexAllRooms {
//		membershipRequest := &webexteams.MembershipCreateRequest{
//			RoomID:      roomID,
//			PersonID:    personID,
//			IsModerator: false,
//		}
//
//		membershipResult, _, err := ac.teamsClient.Memberships.CreateMembership(membershipRequest)
//		if err != nil {
//			log.Printf("Unable to join Person %s to Room %s because %s", personID, roomID, err)
//		} else {
//			log.Printf("Sent request to join Person %s to Room %s, Membership %s, Created at %s", membershipResult.PersonID, membershipResult.RoomID, membershipResult.ID, membershipResult.Created)
//		}
//		time.Sleep(100 * time.Millisecond)
//
//	}
//	return "", nil
//}
//
//func (ac *Controller) InviteNewEmail(requestData barcampgr.InvitePerson) (string, error) {
//	emails := []string{requestData.Email}
//
//	personRequest := &webexteams.PersonRequest{
//		Emails:      emails,
//		FirstName:   requestData.FirstName,
//		LastName:    requestData.LastName,
//		OrgID:	     ac.config.WebexOrgID,
//	}
//
//	personResult, _, err := ac.teamsClient.People.CreatePerson(personRequest)
//	log.Printf("Person Result: %#v", personResult)
//	log.Printf("Error: %s", err)
//	if err != nil {
//		log.Printf("Unable to create Person %s because %s", requestData.Email, err)
//	} else {
//		log.Printf("Sent request to create Person %s, ID %s, Created at %s", personResult.Emails, personResult.ID, personResult.Created)
//	}
//
//	log.Printf("Inviting person %s to the team", personResult.ID)
//
//	teamMembership := &webexteams.TeamMembershipCreateRequest{
//		TeamID:      ac.config.WebexTeamID,
//		PersonID:    personResult.ID,
//		IsModerator: false,
//	}
//
//	teamMembershipResult, _, err := ac.teamsClient.TeamMemberships.CreateTeamMembership(teamMembership)
//	log.Printf("Team Result: %#v", teamMembershipResult)
//	log.Printf("Error: %s", err)
//	if err != nil {
//		log.Printf("Unable to invite Person %s to the Team because %s", personResult.ID, err)
//	} else {
//		log.Printf("Sent request to join Email %s to the Team, Membership %s, Created at %s", teamMembershipResult.PersonEmail, teamMembershipResult.ID, teamMembershipResult.Created)
//	}
//
//	return "", nil
//}
//
//func (ac *Controller) handleCommand (message string, person *webexteams.Person) (string, string, error) {
//	message = strings.TrimPrefix(message, "@BarcampGRBot")
//	message = strings.TrimPrefix(message, "@barcampgrbot")
//	message = strings.TrimPrefix(message, "BarcampGRBot")
//	message = strings.TrimPrefix(message, "barcampgrbot")
//	message = strings.TrimPrefix(message, " ")
//	commandArray := strings.Split(message, " ")
//	displayName := person.DisplayName
//	switch strings.ToLower(commandArray[0]) {
//	case "schedule":
//		switch strings.ToLower(commandArray[1]) {
//		case "web":
//			return ac.scheduleWeb(person)
//		default:
//			log.Printf("I'm attempting to schedule block with message %s, commandArray %s, for %s", message, commandArray, person.DisplayName)
//			message, dmMessage, err := ac.parseAndScheduleTalk(person, commandArray[1:])
//			return message, dmMessage, err
//		}
//	case "get":
//		if len(commandArray) < 2 {
//			return "", "", errors.New("the command `get` must have arguments, such as `get schedule`")
//		}
//		switch strings.ToLower(commandArray[1]) {
//		case "schedule", "talks", "grid":
//			log.Printf("Talk grid message %s, commandArray %s", message, commandArray)
//			return fmt.Sprintf("The talk grid can be found at https://talks.barcampgr.org/"), "", nil
//		case "links":
//			message, dmMessage, err := ac.getAllMyLinks(person)
//			return message, dmMessage, err
//		default:
//			return "", "", errors.New(fmt.Sprintf("Unknown command %s", ac.commandArrayToString(commandArray)))
//		}
//	case "test":
//		log.Printf("Test message %s, commandArray %s", message, commandArray)
//		return fmt.Sprintf("Hi Test!!!!, I received your message of %s from %s", message, displayName), "", nil
//	case "talks", "talk", "grid":
//		log.Printf("Talk grid message %s, commandArray %s", message, commandArray)
//		return fmt.Sprintf("The talk grid can be found at https://talks.barcampgr.org/"), "", nil
//	case "ping":
//		log.Printf("Ping from %s", displayName)
//		return "Pong", "", nil
//	case "dmping", "dm":
//		log.Printf("DMping from %s", displayName)
//		return "Pong", "Pong", nil
//	case "help", "hi", "hi!", "hello", "hello!", "/help":
//		return fmt.Sprintf("Hi!  I'm BarCampGR's automation bot!  %s", help_message), "", nil
//	case "admin":
//		if len(commandArray) < 2 {
//			return "Admins hide in `Organizer Chat`.", "", nil
//		}
//		return ac.handleAdminAction(commandArray[1:], person)
//	default:
//		return fmt.Sprintf("Sorry, I don't know how to handle '%s'.  %s", ac.commandArrayToString(commandArray), help_message), "", nil
//	}
//}
//
//func (ac *Controller) scheduleWeb(person *webexteams.Person) (string, string, error) {
//	stubTalk := database.DBScheduleSession{
//		UpdaterName:          person.DisplayName,
//		UpdaterID:            person.ID,
//		UniqueString:         ac.generateUniqueString(),
//		Version:              0,
//		OutDated:             false,
//	}
//	result := ac.sdb.Orm.Create(&stubTalk)
//	if result.Error != nil {
//		return fmt.Sprintf("I'm sorry, an error occurred when attempting to generate your link.  Error: %s", result.Error), "", nil
//	}
//	return "I'll DM you a link to continue with your talk information.", fmt.Sprintf("Please go to this link to complete your talk information: https://talks.barcampgr.org/actions/?unique_str=%s", stubTalk.UniqueString), nil
//}
//
//func (ac *Controller) getAllMyLinks(person *webexteams.Person) (string, string, error) {
//	var sessions []database.DBScheduleSession;
//	results := ac.sdb.Orm.Where("updater_id = ? AND out_dated = 0", person.ID).Find(&sessions)
//	if results.Error != nil {
//		return "", "", results.Error
//	}
//	if len(sessions) < 1 {
//		return "", "I'm sorry, I couldn't find any sessions created by you.", nil
//	}
//	stringOut := "Here are your edit session links:"
//	for _, session := range sessions {
//		stringOut = fmt.Sprintf("%s\n - %s", stringOut, session.GetEditInfo())
//	}
//	return "", stringOut, nil
//}
//
//func (ac *Controller) parseAndScheduleTalk(person *webexteams.Person, commandArray []string) (string, string, error) {
//	var name string
//	var time string
//	var room string
//	var title string
//	currentArrayPos := 0
//	message := ""
//	// me at 10:00am in The Hotdog Stand for Speaking to Bots, a Minecraft Story
//
//	if len(commandArray) < 7 {
//		return "You must provide all arguments for `schedule <person|me> at <time> in <room> for <title>` \n Example: `schedule me at 7:00pm in Wellness for An Awesome Talk` or `schedule Jane at 7:30pm in Wellness for Another Awesome Talk`", "", nil
//	}
//
//	if strings.ToLower(commandArray[0]) == "me" {
//		name = person.DisplayName
//		currentArrayPos = 1
//		// skip 'at' command word
//	} else {
//		for i, s := range commandArray {
//			name = name + " " + s
//			if ac.isCommandWord(commandArray[i + 1]) {
//				currentArrayPos = i + 1
//				break
//			}
//		}
//	}
//	name = strings.TrimPrefix(name, " ")
//
//	// at 10:00am in The Hotdog Stand for Speaking to Bots, a Minecraft Story
//	// skip 'at' command word
//	currentArrayPos++
//	time = ac.standardizeTime(commandArray[currentArrayPos])
//	currentArrayPos++
//
//	// in The Hotdog Stand for Speaking to Bots, a Minecraft Story
//	// skip 'in' command word
//	currentArrayPos++
//	for i, s := range commandArray[currentArrayPos:] {
//		room = room + " " + s
//		if ac.isCommandWord(commandArray[i + currentArrayPos + 1]) && strings.ToLower(s) != "life" {
//			currentArrayPos = i + currentArrayPos + 1
//			break
//		}
//	}
//	room = strings.TrimPrefix(room, " ")
//
//	// for Speaking to Bots, a Minecraft Story
//	// skip command word "for"
//	currentArrayPos++
//	for _, s := range commandArray[currentArrayPos:] {
//		title = title + " " + s
//	}
//	title = strings.TrimPrefix(title, " ")
//
//	var timeObj database.DBScheduleTime
//	result := ac.sdb.Orm.Where("lower(start) = ? AND displayable = 1", strings.ToLower(time)).Find(&timeObj)
//	if result.Error != nil {
//		log.Printf("Received error %s when trying to query for time starting at %s", result.Error, time)
//		return fmt.Sprintf("Unable to find a scheduleable time starting at %s", time), "", result.Error
//	}
//
//	var roomObj database.DBScheduleRoom
//	result = ac.sdb.Orm.Where("lower(name) = ?", strings.ToLower(room)).Find(&roomObj)
//	if result.Error != nil {
//		log.Printf("Received error %s when trying to query for room %s", result.Error, room)
//		return fmt.Sprintf("Unable to find a room named %s", room), "", result.Error
//	}
//
//	session := database.DBScheduleSession{
//		Time:    &timeObj,
//		Room:    &roomObj,
//		UpdaterName: person.DisplayName,
//		UpdaterID: person.ID,
//		Title:   title,
//		Speaker: name,
//		TimeID:  int(timeObj.ID),
//		RoomID:  int(roomObj.ID),
//		Version: 0,
//		UniqueString: ac.generateUniqueString(),
//	}
//
//	var sessionObj database.DBScheduleSession
//	if err := ac.sdb.Orm.Where("room_id = ? AND time_id = ? AND out_dated = 0", session.RoomID, session.TimeID).First(&sessionObj).Error; err != nil {
//		if gorm.IsRecordNotFoundError(err){
//			result = ac.sdb.Orm.Create(&session)
//			if result.Error != nil {
//				log.Printf("Received error %s when trying to create talk %s", result.Error, ac.commandArrayToString(commandArray))
//				return message, "", result.Error
//			} else {
//				log.Printf("Created talk %s with %d rows affected", ac.commandArrayToString(commandArray), result.RowsAffected)
//				message = fmt.Sprintf("I've scheduled your session %s.  A link has been DM'd to you to manage your session entry on the grid.", session.ToString())
//				dmMessage := fmt.Sprintf("Here I just scheduled this session for you: %s", session.ToDmString())
//				return message, dmMessage, nil
//			}
//		} else {
//			log.Printf("Received error %s when trying to create talk %s", result.Error, ac.commandArrayToString(commandArray))
//			return message, "", result.Error
//		}
//	}
//
//	log.Printf("Session already exists at %s in room %s.", session.Time.Start, session.Room.Name)
//	return fmt.Sprintf("I'm sorry, a session already exists at %s in room %s.", session.Time.Start, session.Room.Name), "", nil
//}
//
//func (ac *Controller) generateUniqueString() string {
//	for i := 0; i < 10; i++ {
//		resultant := make([]byte, 64)
//		for i := range resultant {
//			resultant[i] = charset[ac.sRand.Intn(len(charset))]
//		}
//		resultStr := string(resultant)
//		if ac.sessionStrNotUsed(resultStr) {
//			return resultStr
//		}
//	}
//	return ""
//}
//
//func (ac *Controller) sessionStrNotUsed (sessionStr string) bool {
//	session := database.DBScheduleSession{}
//	result := ac.sdb.Orm.Where("unique_string = ?", sessionStr).Find(&session)
//	if result.Error != nil {
//		if gorm.IsRecordNotFoundError(result.Error){
//			return true
//		}
//	}
//	return false
//}
//
//func (ac *Controller) commandArrayToString(array []string) string {
//	var resultant string
//	for _,s := range array {
//		resultant = resultant + " " + s
//	}
//	return strings.TrimPrefix(resultant, " ")
//}
//
//func (ac *Controller) isCommandWord(check string) bool {
//	switch check {
//	case "at", "in", "for":
//		return true
//	default:
//		return false
//	}
//}
//
//func (ac *Controller) GetScheduleJson() (barcampgr.Schedule, error) {
//	var times []database.DBScheduleTime
//	var rooms []database.DBScheduleRoom
//	var sessions []database.DBScheduleSession
//
//	ac.sdb.Orm.Find(&times)
//	outTimes := ac.convertTimes(times)
//
//	ac.sdb.Orm.Find(&rooms)
//	ac.sdb.Orm.Where("out_dated = 0").Find(&sessions)
//	sessions = ac.fillTimes(sessions, times)
//	outRows := ac.buildRows(sessions, rooms)
//
//	schedule := barcampgr.Schedule{
//		RefreshedAt: "",
//		LastUpdate:  "",
//		Times:       outTimes,
//		Rows:        outRows,
//	}
//
//	return schedule, nil
//}
//
//func (ac *Controller) GetSessionByStr(sessionStr string) (barcampgr.ScheduleSession, error) {
//	var session database.DBScheduleSession
//	ac.sdb.Orm.Where("unique_string = ? AND out_dated = 0", sessionStr).Find(&session)
//	return ac.convertSession(session), nil
//}
//
//func (ac *Controller) GetTimesJson() ([]barcampgr.ScheduleTime, error) {
//	var times []database.DBScheduleTime
//	results := ac.sdb.Orm.Where("displayable = 1").Order("start").Find(&times)
//	if results.Error != nil {
//		return []barcampgr.ScheduleTime{}, results.Error
//	}
//	return ac.convertTimes(times), nil
//}
//
//func (ac *Controller) GetRoomsJson() ([]barcampgr.ScheduleRoom, error) {
//	var rooms []database.DBScheduleRoom
//	results := ac.sdb.Orm.Order("name").Find(&rooms)
//	if results.Error != nil {
//		return []barcampgr.ScheduleRoom{}, results.Error
//	}
//	return ac.convertRooms(rooms), nil
//}
//
//func (ac *Controller) convertTimes(times []database.DBScheduleTime) []barcampgr.ScheduleTime {
//	var resultant []barcampgr.ScheduleTime
//	for _, t := range times {
//		if t.Displayable {
//			resultant = append(resultant, barcampgr.ScheduleTime{
//				ID:          int(t.ID),
//				Start:       t.Start,
//				End:         t.End,
//				Day:         t.Day,
//				Displayable: t.Displayable,
//			})
//		}
//	}
//	return resultant
//}
//
//func (ac *Controller) convertRooms(rooms []database.DBScheduleRoom) []barcampgr.ScheduleRoom {
//	var resultant []barcampgr.ScheduleRoom
//	for _, r := range rooms {
//		resultant = append(resultant, barcampgr.ScheduleRoom{
//			Name: r.Name,
//			ID:   int(r.ID),
//		})
//	}
//	return resultant
//}
//
//func (ac *Controller) convertSession(session database.DBScheduleSession) barcampgr.ScheduleSession {
//	return barcampgr.ScheduleSession{
//		Time:         session.TimeID,
//		Room:         session.RoomID,
//		Title:        session.Title,
//		Speaker:      session.Speaker,
//		UniqueString: session.UniqueString,
//		Version:      session.Version,
//	}
//}
//
//func (ac *Controller) buildRows(sessions []database.DBScheduleSession, rooms []database.DBScheduleRoom) []barcampgr.ScheduleRow {
//	var resultant []barcampgr.ScheduleRow
//
//	log.Printf("Sessions: %d Rooms: %d", len(sessions), len(rooms))
//	for _,r := range rooms {
//		resultant = append(resultant, barcampgr.ScheduleRow{
//			Room:    r.Name,
//			Sessions: ac.getSessionsInRoom(sessions, r),
//		})
//	}
//	return resultant
//}
//
//func (ac *Controller) RollSchedule(scheduleBlock string) error  {
//	switch scheduleBlock {
//	case "fri-pm":
//		var times []barcampgr.ScheduleTime
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "6:30pm",
//			End:   "6:55pm",
//			Day:   "Friday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "7:00pm",
//			End:   "7:25pm",
//			Day:   "Friday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "7:30pm",
//			End:   "7:55pm",
//			Day:   "Friday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "8:00pm",
//			End:   "8:25pm",
//			Day:   "Friday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "8:30pm",
//			End:   "8:55pm",
//			Day:   "Friday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "9:00pm",
//			End:   "9:25pm",
//			Day:   "Friday",
//		})
//		result := ac.createTimeBlockAndDisableOthers(times)
//		if len(result) < 0 {
//			return result[0]
//		}
//	case "sat-am":
//		var times []barcampgr.ScheduleTime
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "10:00am",
//			End:   "10:25am",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "10:30am",
//			End:   "10:55am",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "11:00am",
//			End:   "11:25am",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "11:30am",
//			End:   "11:55am",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "12:00pm",
//			End:   "12:55pm",
//			Day:   "Saturday",
//		})
//		ac.createTimeBlockAndDisableOthers(times)
//	case "sat-pm":
//		var times []barcampgr.ScheduleTime
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "1:00pm",
//			End:   "1:25pm",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "1:30pm",
//			End:   "1:55pm",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "2:00pm",
//			End:   "2:25pm",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "2:30pm",
//			End:   "2:55pm",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "3:00pm",
//			End:   "3:25pm",
//			Day:   "Saturday",
//		})
//		times = append(times, barcampgr.ScheduleTime{
//			Start: "3:30pm",
//			End:   "3:55pm",
//			Day:   "Saturday",
//		})
//		ac.createTimeBlockAndDisableOthers(times)
//	case "rooms":
//		return ac.confirmAndGenerateRooms()
//	default:
//		return errors.New("not allowed")
//	}
//	return nil
//}
//
//func (ac *Controller) confirmAndGenerateRooms() error {
//	rooms := [8]string{"General", "Life in 2020", "Creative Corner", "Programming", "Room 120", "Room 140", "Room 170", "Wellness"}
//	for _, room := range rooms {
//		var roomObj database.DBScheduleRoom
//		result := ac.sdb.Orm.Where("name = ?", room).Find(roomObj)
//		if result.Error != nil {
//			log.Printf("Error in finding existing room %s, Type: %T, Message: %s", room, result.Error, result.Error)
//			//TODO(twodarek): Change this to use https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
//			result = ac.sdb.Orm.Create(&database.DBScheduleRoom{
//				Name: room,
//			})
//			if result.Error != nil {
//				log.Printf("Error in creating room %s, Type: %T, Message: %s", room, result.Error, result.Error)
//			}
//		}
//	}
//	return nil
//}
//
//func (ac *Controller) createTimeBlockAndDisableOthers(times []barcampgr.ScheduleTime) []error {
//	timeObj := database.DBScheduleTime{}
//	var resultErrors []error
//	result := ac.sdb.Orm.Model(&timeObj).Where("displayable = ?", true).Update("displayable", false)
//	if result.Error != nil {
//		resultErrors = append(resultErrors, result.Error)
//		return resultErrors
//	}
//	for _, time := range times {
//		timeObj := database.DBScheduleTime{}
//		if err := ac.sdb.Orm.Where("start = ? AND day = ?", time.Start, time.Day).First(&timeObj).Error; err != nil {
//			if gorm.IsRecordNotFoundError(err){
//				log.Printf("Error in finding existing time for start %s on %s, error: %s", time.Start, time.Day, err)
//				timeObj = database.DBScheduleTime{
//					Day:         time.Day,
//					Start:       time.Start,
//					End:         time.End,
//					Displayable: true,
//				}
//				result = ac.sdb.Orm.Create(&timeObj)
//				if result.Error != nil {
//					resultErrors = append(resultErrors, result.Error)
//				}
//			}
//		}else{
//			result = ac.sdb.Orm.Model(&timeObj).Where("id = ?", timeObj.ID).Update("displayable", true)
//			if result.Error != nil {
//				resultErrors = append(resultErrors, result.Error)
//			}
//		}
//	}
//	return resultErrors
//}
//
//func (ac *Controller) MigrateDB() error {
//	return ac.sdb.MigrateDB()
//}
//
//func (ac *Controller) getSessionsInRoom(sessions []database.DBScheduleSession, room database.DBScheduleRoom) []barcampgr.ScheduleSession {
//	var resultant []barcampgr.ScheduleSession
//	for _,s := range sessions {
//		if s.Time == nil {
//			continue
//		}
//		if s.Time.Displayable && s.RoomID == int(room.ID) {
//			var altText string
//			if s.Title == "Blocked" {
//				altText = "Unscheduleable time"
//			} else {
//				altText = fmt.Sprintf("%s by %s in %s at %s", s.Title, s.Speaker, room.Name, s.Time.Start)
//			}
//			resultant = append(resultant, barcampgr.ScheduleSession{
//				Time:    int(s.Time.ID),
//				Room:    int(room.ID),
//				Title:   s.Title,
//				Speaker: s.Speaker,
//				UniqueString: s.UniqueString,
//				AltText: altText,
//			})
//		}
//	}
//	return resultant
//}
//
//func (ac *Controller) fillTime(session *database.DBScheduleSession) {
//	time := database.DBScheduleTime{}
//	ac.sdb.Orm.Where("id = ?", session.TimeID).Find(&time)
//	session.Time = &time
//}
//
//func (ac *Controller) fillRoom(session *database.DBScheduleSession) {
//	room := database.DBScheduleRoom{}
//	ac.sdb.Orm.Where("id = ?", session.RoomID).Find(&room)
//	session.Room = &room
//}
//
//func (ac *Controller) fillTimes(sessions []database.DBScheduleSession, times []database.DBScheduleTime) []database.DBScheduleSession {
//	for i, s := range sessions {
//		for j, t := range times {
//			if s.TimeID == int(t.ID) {
//				sessions[i].Time = &times[j]
//			}
//		}
//	}
//	return sessions
//}
//
//const outputTimeLayout = "3:04om"
//// My sincerest apologies for what you are about to read, time is messy and horrible.
//func (ac *Controller) standardizeTime(input string) string {
//	input = strings.ToLower(input)
//
//	splitTime, err := time.Parse(outputTimeLayout, "9:01")
//
//	timeOutput, err := time.Parse(outputTimeLayout, input)
//	if err == nil {
//		return timeOutput.String()
//	}
//	testable := strings.Replace(strings.Replace(input, "am", "", 1), "pm", "", 1)
//	if len(testable) == 3 && !strings.Contains(input, ":") {
//		input = fmt.Sprintf("%s:%s", input[:1], input[1:])
//		log.Printf("Look at this abomination: before: %s, after: %s", testable, input)
//	}
//	if len(testable) == 4 && !strings.Contains(input, ":"){
//		input = fmt.Sprintf("%s:%s", input[:2], input[2:])
//		log.Printf("Look at this abomination: before: %s, after: %s", testable, input)
//	}
//	timeOutput, err = time.Parse("3pm", input)
//	if err == nil {
//		return strings.Replace(timeOutput.Format(outputTimeLayout), "om", input[len(input)-2:], -1)
//	}
//	timeOutput, err = time.Parse("34pm", input)
//	if err == nil {
//		return timeOutput.Format(outputTimeLayout)
//	}
//	log.Printf("error log: %s", err)
//	timeOutput, err = time.Parse("3", input)
//	if err == nil {
//		if timeOutput.Before(splitTime) {
//			return strings.Replace(timeOutput.Format(outputTimeLayout), "om", "pm", -1)
//		} else if timeOutput.After(splitTime) {
//			return strings.Replace(timeOutput.Format(outputTimeLayout), "om", "am", -1)
//		} else {
//			log.Printf("You dun goofed")
//			return input
//		}
//	}
//	timeOutput, err = time.Parse("3:4", input)
//	if err == nil {
//		if timeOutput.Before(splitTime) {
//			return strings.Replace(fmt.Sprintf("%spm", timeOutput.Format(outputTimeLayout)), "om", "", -1)
//		} else if timeOutput.After(splitTime) {
//			return strings.Replace(fmt.Sprintf("%sam", timeOutput.Format(outputTimeLayout)), "om", "", -1)
//		} else {
//			log.Printf("You dun goofed")
//			return input
//		}
//	}
//	log.Printf("Time parsing failed, good luck!  Input: %s", input)
//	return input
//}
//
//func (ac *Controller) UpdateSession(sessionStr string, sessionInbound barcampgr.ScheduleSession) error {
//	sessionObj := database.DBScheduleSession{}
//	ac.sdb.Orm.Where("unique_string = ? AND out_dated = 0", sessionStr).Find(&sessionObj)
//	updated := false
//	if sessionObj.Speaker != sessionInbound.Speaker {
//		updated = true
//	}
//	if sessionObj.Title != sessionInbound.Title {
//		updated = true
//	}
//	if sessionObj.RoomID != sessionInbound.Room {
//		updated = true
//	}
//	if sessionObj.TimeID != sessionInbound.Time {
//		updated = true
//	}
//	if updated {
//		newSession := database.DBScheduleSession{
//			RoomID:       sessionInbound.Room,
//			TimeID:       sessionInbound.Time,
//			UpdaterName:  sessionObj.UpdaterName,
//			UpdaterID:    sessionObj.UpdaterID,
//			Title:        sessionInbound.Title,
//			Speaker:      sessionInbound.Speaker,
//			UniqueString: ac.generateUniqueString(),
//			PreviousUniqueString: sessionObj.UniqueString,
//			Version:      sessionObj.Version + 1,
//			OutDated:     false,
//		}
//
//		sessionTest := database.DBScheduleSession{}
//		err := ac.sdb.Orm.Where("room_id = ? AND time_id = ? AND out_dated = 0", sessionInbound.Room, sessionInbound.Time).First(&sessionTest).Error
//		if err != nil {
//			if gorm.IsRecordNotFoundError(err){
//				sessionObj.OutDated = true
//				result := ac.sdb.Orm.Save(&sessionObj)
//				if result.Error != nil {
//					return result.Error
//				}
//				result = ac.sdb.Orm.Create(&newSession)
//				if result.Error != nil {
//					log.Printf("Received error %s when trying to update talk %s from %s", result.Error, sessionInbound.UniqueString, sessionObj.UniqueString)
//					return result.Error
//				} else {
//					log.Printf("Updated talk %s from %s with %d rows affected", newSession.UniqueString, sessionObj.UniqueString, result.RowsAffected)
//					ac.fillTime(&newSession)
//					ac.fillRoom(&newSession)
//					ac.sendDM(newSession.UpdaterID, fmt.Sprintf(fmt.Sprintf("I've scheduled your session %s", newSession.ToDmString())))
//					return nil
//				}
//			} else {
//				log.Printf("Received error %s when trying to update talk %s from %s", err, newSession.UniqueString, sessionObj.UniqueString)
//				return err
//			}
//		}
//		if sessionTest.UniqueString != sessionInbound.UniqueString {
//			log.Printf("error in updating session %s, Error: %s", sessionInbound.UniqueString, err)
//			return errors.New("Sorry, a session already is scheduled for that time and room.  Please select an available slot.")
//		}
//		sessionObj.OutDated = true
//		result := ac.sdb.Orm.Save(&sessionObj)
//		if result.Error != nil {
//			return result.Error
//		}
//		result = ac.sdb.Orm.Create(&newSession)
//		if result.Error != nil {
//			log.Printf("Received error %s when trying to update talk %s from %s", result.Error, sessionInbound.UniqueString, sessionObj.UniqueString)
//			return result.Error
//		} else {
//			log.Printf("Updated talk %s from %s with %d rows affected", newSession.UniqueString, sessionObj.UniqueString, result.RowsAffected)
//			ac.fillTime(&newSession)
//			ac.fillRoom(&newSession)
//			ac.sendDM(newSession.UpdaterID, fmt.Sprintf(fmt.Sprintf("I've scheduled your session %s", newSession.ToDmString())))
//			return nil
//		}
//	}
//	return errors.New("Previous session not found to update.")
//}
//
//func (ac *Controller) DeleteSession(sessionStr string) error {
//	return ac.sdb.Orm.Where("unique_string = ?", sessionStr).Delete(database.DBScheduleSession{}).Error
//}
//
//func (ac *Controller) handleAdminAction(commandArray []string, person *webexteams.Person) (string, string, error) {
//	params := &webexteams.ListMembershipsQueryParams{
//		RoomID:      ac.config.WebexRoomID,
//		PersonID:    person.ID,
//	}
//	memberships, _, err := ac.teamsClient.Memberships.ListMemberships(params)
//	if err != nil {
//		return "Unable to authenticate you as an admin user.", "", err
//	}
//	isAdmin := false
//	for _, membership := range memberships.Items {
//		if membership.IsModerator {
//			isAdmin = true
//		}
//	}
//	if !isAdmin {
//		return "Unable to authenticate you as an admin user.", "", err
//	}
//	switch commandArray[0] {
//	case "roll":
//		err := ac.RollSchedule(commandArray[1])
//		if err != nil {
//			return "Unable to roll the schedule.", "", err
//		}
//		return fmt.Sprintf("The schedule has been successfully rolled to %s", commandArray[1]), "", nil
//	default:
//		return "I'm sorry, I don't know how to run that admin command.", "", errors.New("Command not recognized.")
//	}
//}
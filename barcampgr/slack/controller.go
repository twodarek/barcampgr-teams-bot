package slack

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/twodarek/barcampgr-teams-bot/database"
)

type Controller struct {
	bc  *barcampgr.Controller
	slackClient *slack.Client
	httpClient  *http.Client
	sdb         *database.ScheduleDatabase
	config      barcampgr.Config
	sRand       *rand.Rand
}

func NewAppController(
	barcampgrController *barcampgr.Controller,
	slackClient	*slack.Client,
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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const help_message = "I accept the following commands:\n - `help` to get this message\n - `get schedule`, `get grid`, or `get talks` to get a link to the schedule grid\n - `get links` to get all of the unique links for your talks\n - `dm` to open a direct message connection with me\n - `Schedule me at START_TIME in ROOM for TITLE` to schedule a talk\n - `Schedule web` to schedule a talk via web form\n\nMake sure to `@barcampgrbot` at the start or I won't get the message!"

func (ac *Controller) HandleChatop(eventsAPIEvent slackevents.EventsAPIEvent) (string, error) {
	if eventsAPIEvent.Type != slackevents.CallbackEvent {
		return "", nil
	}

	//	// Filter to make sure it's only from BarCampGR
	//	if requestData.OrgID != ac.config.WebexOrgID {
	//		return "", errors.New(fmt.Sprintf("Unable to handle messages from non-BarCampGR orgs %s", requestData.Data.ID))
	//	}
	//

	//
	//	person, _, err :=ac.teamsClient.People.GetPerson(message.PersonID)
	//	if err != nil {
	//		return "", errors.New(fmt.Sprintf("Unable to get person id %s", message.PersonID))
	//	}

	//log.Printf("Got message from person.  Display: %s, Nick: %s, Name: %s %s", person.DisplayName, person.NickName, person.FirstName, person.LastName)

	innerEvent := eventsAPIEvent.InnerEvent
	switch ev := innerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		message := eventsAPIEvent.InnerEvent.Data.(*slackevents.AppMentionEvent)

		log.Printf("Received message `%s` in room %s from person %s, username: %s", message.Text, message.Channel, message.User, message.User)

		if message.User == ac.config.SlackUsername {
			log.Printf("Rejecting message from myself, returning cleanly")
			return "", nil
		}

		user, err :=ac.slackClient.GetUserProfile(message.User, true)
		if err != nil {
			log.Printf("Unable to get profile of user %s.  Error: %s", message.User, err)
			user = &slack.UserProfile{
				DisplayName: message.User,
			}
		}

		reply, dmReply, err := ac.bc.HandleCommand(message.Text, user.DisplayName, message.User)
		if err != nil {
			ac.slackClient.PostMessage(ev.Channel, slack.MsgOptionText("I'm sorry, something went wrong.", false))
		}

		log.Printf("Here's what I would have replied with if I writed up to do so.  Public message: %s, Direct message: %s", reply, dmReply)

		respChannel, timestamp, err := ac.slackClient.PostMessage(ev.Channel, slack.MsgOptionText(reply, false))
		log.Printf("Attempted to post to channel %s at timestamp %s, error: %s", respChannel, timestamp, err)
		log.Printf("Got this text: %s", message.Text)
	case *slackevents.MessageEvent:
		message := eventsAPIEvent.InnerEvent.Data.(*slackevents.MessageEvent)
		log.Printf("Received message `%s` as message id %s in room %s from person %s, username: %s", message.Text, message.ClientMsgID, message.Channel, message.User, message.Username)
	}
	return "", nil
}
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

//func (ac *Controller) scheduleWeb(person *webexteams.Person) (string, string, error) {
//	stubTalk := database.DBScheduleSession{
//		UpdaterName:          person.DisplayName,
//		UpdaterID:            person.ID,
//		UniqueString:         ac.bc.GenerateUniqueString(),
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
//			if ac.bc.IsCommandWord(commandArray[i + 1]) {
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
//	time = ac.bc.StandardizeTime(commandArray[currentArrayPos])
//	currentArrayPos++
//
//	// in The Hotdog Stand for Speaking to Bots, a Minecraft Story
//	// skip 'in' command word
//	currentArrayPos++
//	for i, s := range commandArray[currentArrayPos:] {
//		room = room + " " + s
//		if ac.bc.IsCommandWord(commandArray[i + currentArrayPos + 1]) && strings.ToLower(s) != "life" {
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
//		UniqueString: ac.bc.GenerateUniqueString(),
//	}
//
//	var sessionObj database.DBScheduleSession
//	if err := ac.sdb.Orm.Where("room_id = ? AND time_id = ? AND out_dated = 0", session.RoomID, session.TimeID).First(&sessionObj).Error; err != nil {
//		if gorm.IsRecordNotFoundError(err){
//			result = ac.sdb.Orm.Create(&session)
//			if result.Error != nil {
//				log.Printf("Received error %s when trying to create talk %s", result.Error, ac.bc.CommandArrayToString(commandArray))
//				return message, "", result.Error
//			} else {
//				log.Printf("Created talk %s with %d rows affected", ac.bc.CommandArrayToString(commandArray), result.RowsAffected)
//				message = fmt.Sprintf("I've scheduled your session %s.  A link has been DM'd to you to manage your session entry on the grid.", session.ToString())
//				dmMessage := fmt.Sprintf("Here I just scheduled this session for you: %s", session.ToDmString())
//				return message, dmMessage, nil
//			}
//		} else {
//			log.Printf("Received error %s when trying to create talk %s", result.Error, ac.bc.CommandArrayToString(commandArray))
//			return message, "", result.Error
//		}
//	}
//
//	log.Printf("Session already exists at %s in room %s.", session.Time.Start, session.Room.Name)
//	return fmt.Sprintf("I'm sorry, a session already exists at %s in room %s.", session.Time.Start, session.Room.Name), "", nil
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
//		err := ac.bc.RollSchedule(commandArray[1])
//		if err != nil {
//			return "Unable to roll the schedule.", "", err
//		}
//		return fmt.Sprintf("The schedule has been successfully rolled to %s", commandArray[1]), "", nil
//	default:
//		return "I'm sorry, I don't know how to run that admin command.", "", errors.New("Command not recognized.")
//	}
//}

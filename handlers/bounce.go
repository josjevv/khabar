package handlers

import (
	"log"
	"net/http"

	"gopkg.in/bulletind/khabar.v1/core"
	"gopkg.in/bulletind/khabar.v1/db"
	"gopkg.in/bulletind/khabar.v1/dbapi/available_topics"
	"gopkg.in/bulletind/khabar.v1/dbapi/saved_item"
	"gopkg.in/bulletind/khabar.v1/dbapi/topics"
	"gopkg.in/bulletind/khabar.v1/utils"
	"gopkg.in/simversity/gottp.v2"
	gottpUtils "gopkg.in/simversity/gottp.v2/utils"
)

type Bounce struct {
	gottp.BaseHandler
}

const BounceNotification = "Bounce"

type snsNotice struct {
	Type      string `json:"Type"`
	MessageId string `json:"MessageId" required:"true"`
	TopicArn  string `json:"TopicArn"`
	Subject   string `json:"Subject"`
	Message   string `json:"Message" required:"true"`
	Timestamp string `json:"Timestamp"`
	Signature string `json:"Signature" required:"true"`
}

type bounceMessage struct {
	Type   string `json:"notificationType" required:"true"`
	Bounce struct {
		Recipients []struct {
			Email string `json:"emailAddress" required:"true"`
		} `json:"bouncedRecipients" required:"true"`
	} `json:"bounce" required:"true"`
}

func (self *Bounce) Post(request *gottp.Request) {

	args := new(snsNotice)

	request.ConvertArguments(&args)

	if !utils.ValidateAndRaiseError(request, args) {
		log.Println("Invalid Request", request.GetArguments())
		return
	}

	msg := bounceMessage{}
	gottpUtils.Decoder([]byte(args.Message), &msg)

	errs := gottpUtils.Validate(&msg)
	if len(*errs) > 0 {
		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			ConcatenateErrors(errs),
		})

		return
	}

	if msg.Type != BounceNotification {
		log.Println("Invalid Bounce Request", request.GetArguments())

		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			"Invalid Bounce Request",
		})

		return
	}

	for _, entry := range msg.Bounce.Recipients {
		DisableBounceEmail(entry.Email, request)
	}

	request.Write(utils.R{
		StatusCode: http.StatusOK,
		Data:       nil,
	})

}

func DisableBounceEmail(email string, request *gottp.Request) error {
	userId, orgs := saved_item.GetSentOrganizations(db.SavedEmailCollection, email)
	all_topics := available_topics.GetAllTopics()
	topics.DisableUserChannel(orgs, all_topics, userId, core.EMAIL)

	return nil
}
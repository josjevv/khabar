package core

import (
	"log"

	"github.com/changer/khabar/dbapi/pending"
)

func webHandler(item *pending.PendingItem, text string, settings map[string]interface{}) {
	log.Println("Sending Web Notification...")
	return
}

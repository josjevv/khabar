package processed

import (
	"github.com/bulletind/khabar/db"
	"github.com/bulletind/khabar/utils"
)

func IsProcessed(user, org string) bool {
	return db.Conn.Count(db.ProcessedCollection, utils.M{
		"user": user,
		"org":  org,
	}) != 0
}

func MarkAsProcessed(user, org string) (error, bool) {
	if IsProcessed(user, org) {
		return nil, false
	}

	processed := &db.Processed{User: user, Organization: org}
	processed.PrepareSave()

	err := db.Conn.Upsert(db.ProcessedCollection, utils.M{
		"user": user,
		"org":  org,
	}, utils.M{"$set": processed})

	if err != nil {
		return err, false
	} else {
		return nil, true
	}
}

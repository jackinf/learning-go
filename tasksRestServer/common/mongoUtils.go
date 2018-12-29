package common

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		session = createDbSession()
	}
	return session
}

func createDbSession() *mgo.Session {
	sess, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:	[]string{AppConfig.MogoDbHost},
		Timeout: 	60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s", err)
	}
	return sess
}

func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key: 			[]string{"email"},
		Unique: 		true,
		Background: 	true,
		Sparse: 		true,
	}
	taskIndex := mgo.Index{
		Key: 			[]string{"createdBy"},
		Unique: 		false,
		Background: 	true,
		Sparse: 		true,
	}
	noteIndex := mgo.Index{
		Key: 			[]string{"taskId"},
		Unique: 		false,
		Background: 	true,
		Sparse: 		true,
	}
	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB(AppConfig.Database).C("users")
	tasksCol := session.DB(AppConfig.Database).C("tasks")
	notesCol := session.DB(AppConfig.Database).C("notes")
	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s", err)
	}
	err = tasksCol.EnsureIndex(taskIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s", err)
	}
	err = notesCol.EnsureIndex(noteIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s", err)
	}
}
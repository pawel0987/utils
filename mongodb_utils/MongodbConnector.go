// Author: Paweł Konopko
// License: MIT

package mongodb_utils

import (
	"strconv"
	"gopkg.in/mgo.v2"
	"time"
	"github.com/pawel0987/log"
)

type MongodbConnector struct {
	config             MongodbConfig
	fatalErrorCallback func(error)
	session            *mgo.Session
}

func NewMongodbConnector(config *MongodbConfig, fatalErrorCallback func(error)) *MongodbConnector {
	if fatalErrorCallback == nil {
		fatalErrorCallback = func(err error) {
			log.Session("DefaultFatalErrorCallback").Fatal("DB crashed!", log.Data{"err": err})
		}
	}
	return &MongodbConnector{
		config: *config,
		fatalErrorCallback: fatalErrorCallback,
		session: nil,
	}
}

func (db *MongodbConnector) GetCollection(name string) *mgo.Collection {
	return db.GetSession().DB("").C(name)
}

func (db *MongodbConnector) GetSession() *mgo.Session {
	var err error

	for i:=uint(0); i<db.config.MaxRetries; i++ {
		if db.session == nil {
			err = db.connect()
			if err == nil {
				return db.session
			}
		} else {
			err = db.session.Ping()
			if err != nil {
				db.session.Close()
				err = db.connect()
				if err == nil {
					return db.session
				}
			} else {
				return db.session
			}
		}

		log.Session("mongodbConnector.GetSession").Error("Error getting mongodb session", log.Data{"attempt": i+1, "err": err.Error()})
		time.Sleep(time.Duration(db.config.MillisecondsBetweenRetries) * time.Millisecond)
	}

	db.HandleDatabaseError(err)
	return nil
}

func (db *MongodbConnector) connect() error {
	// create connection uri:
	var uri string
	if len(db.config.User) == 0 {
		uri = "mongodb://" + db.config.Host + ":" + strconv.Itoa(int(db.config.Port)) + "/" + db.config.DatabaseName
	} else {
		uri = "mongodb://" + db.config.User + ":" + db.config.Password + "@" + db.config.Host + ":" + strconv.Itoa(int(db.config.Port)) + "/" + db.config.DatabaseName
	}

	// create mongodb connection:
	var err error
	db.session, err = mgo.Dial(uri)
	if err != nil {
		db.session = nil
		return err
	}

	// not sure if it's required but enable this just for case
	db.session.SetSafe(&mgo.Safe{})
	return nil
}

func (db *MongodbConnector) HandleDatabaseError(err error) {
	if err == nil {
		return
	}
	log.Session("mongodbConnector.HandleDatabaseError").Fatal("error: " + err.Error())
	db.fatalErrorCallback(err)
}

func (db *MongodbConnector) RemoveAllData() {
	db.HandleDatabaseError(db.GetSession().DB("").DropDatabase())
}

func (db *MongodbConnector) Dispose() {
	if db.session == nil {
		log.Session("MongodbConnector.Dispose").Warning("DB session not exists")
		return
	}

	db.session.Close()
	db.session = nil
}

func (db *MongodbConnector) Ping() error {
	return db.GetSession().Ping()
}
package dao

import (
	"log"

	"../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ContactDao struct {
	Server   string
	Database string
}

var db *mgo.Database
var collection *mgo.Collection

const (
	COLLECTION = "contacts"
)

type IContactDao interface {
	GetAll() ([]models.Contact, error)
	GetByID(id string) (models.Contact, error)
	FindByName(name string) ([]models.Contact, error)
	Create(contact models.Contact) error
	Update(id string, contact models.Contact) error
	Delete(id string) error
}

func (m *ContactDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	collection = db.C(COLLECTION)
}

func (m *ContactDao) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := collection.Find(bson.M{}).All(&contacts)
	return contacts, err
}

func (m *ContactDao) GetByID(id string) (models.Contact, error) {
	var contact models.Contact
	err := collection.FindId(bson.ObjectIdHex(id)).One(&contact)
	return contact, err
}

func (m *ContactDao) FindByName(name string) ([]models.Contact, error) {
	var contacts []models.Contact
	err := collection.Find(bson.M{"name": bson.RegEx{
		Pattern: name,
		Options: "i",
	},
	},
	).All(&contacts)
	return contacts, err
}

func (m *ContactDao) Create(contact models.Contact) error {
	err := collection.Insert(&contact)
	return err
}

func (m *ContactDao) Delete(id string) error {
	err := collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ContactDao) Update(id string, contact models.Contact) error {
	err := collection.UpdateId(bson.ObjectIdHex(id), &contact)
	return err
}

package db

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type Music struct {
	Id   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name, omitempty"`
}

const MusicCollection = "music"

var ErrDuplicatedMusic = errors.New("Duplicated music")

type MusicRepository struct {
	session *mgo.Session
}

func (r *MusicRepository) Create(p *Music) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	err := collection.Insert(p)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedMusic
		}
	}
	return err
}

func (r *MusicRepository) Update(p *Music) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	return collection.Update(bson.M{"_id": p.Id}, p)
}

func (r *MusicRepository) Remove(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	return collection.Remove(bson.M{"_id": id})
}

func (r *MusicRepository) FindAllActive() ([]*Music, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	query := bson.M{"inative": false}

	documents := make([]*Music, 0)

	err := collection.Find(query).All(&documents)
	return documents, err
}

func (r *MusicRepository) FindById(id string) (*Music, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	query := bson.M{"_id": id}

	person := &Music{}

	err := collection.Find(query).One(person)
	return person, err
}

func NewMusicRepository(session *mgo.Session) *MusicRepository {
	return &MusicRepository{session}
}

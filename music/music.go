package music

import (

	"gopkg.in/mgo.v2"
	"log"

	"github.com/edersonbrilhante/go-course-project/db"
)

func NewMusic(m *db.Music) error{
	session, err := mgo.Dial("localhost:27017/go-course")
	repo := db.NewMusicRepository(session)
	err = repo.Create(m)

	if err == db.ErrDuplicatedMusic {
		log.Println("Duplicated music: ", err)
		return db.ErrDuplicatedMusic
	} else if err != nil {
		log.Println("Failed to create a music: ", err)
		return err
	}
	return err
}

func DelMusic(id string) error {
	session, err := mgo.Dial("localhost:27017/go-course")
	repo := db.NewMusicRepository(session)
	err = repo.Remove(id)

	if err != nil {
		log.Println("Failed to delete a music: ", err)
		return err
	}
	return err
}

func GetMusic(id string) (*db.Music, error) {
	session, err := mgo.Dial("localhost:27017/go-course")
	repo := db.NewMusicRepository(session)
	m, err := repo.FindById(id)
	return m, err
}

func UpdateMusic(m *db.Music) error{
	session, err := mgo.Dial("localhost:27017/go-course")
	repo := db.NewMusicRepository(session)
	err = repo.Update(m)
	return err
}

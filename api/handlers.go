package api

import (
	"log"
	"encoding/json"
	"net/http"

	"github.com/edersonbrilhante/go-course-project/db"
	"github.com/edersonbrilhante/go-course-project/music"

	"github.com/dimfeld/httptreemux"
)

type UpdateMusicHandler struct{}

func (h *UpdateMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	obj := &db.Music{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	music.UpdateMusic(obj)
}

type GetMusicHandler struct{}

func (h *GetMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	m, err2 := music.GetMusic(params["id"])
	if err2 != nil {
		log.Println(err2)
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type DeleteMusicHandler struct{}

func (h *DeleteMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	err2 := music.DelMusic(params["id"])
	if err2 != nil {
		log.Println(err2)
	}
}

type PostMusicHandler struct{}

func (h *PostMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	obj := &db.Music{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	music.NewMusic(obj)
}

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Song struct {
	ID         string `json:"id"`
	ArtistName string `json:"artistname"`
	SongName   string `json:"songname"`
	Genre      string `json:"genre"`
}

var songs []Song

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range songs {
		if item.ID == params["id"] {
			songs = append(songs[:i], songs[i+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range songs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
}

func addSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var son Song
	_ = json.NewDecoder(r.Body).Decode(&son)
	son.ID = strconv.Itoa(rand.Intn(1000))
	songs = append(songs, son)
	json.NewEncoder(w).Encode(son)
}

func UpdSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range songs {
		if item.ID == params["id"] {
			songs = append(songs[:i], songs[i+1:]...)
			var updatedSong Song
			_ = json.NewDecoder(r.Body).Decode(&updatedSong)
			updatedSong.ID = params["id"]
			songs = append(songs, updatedSong)
			json.NewEncoder(w).Encode(updatedSong)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})

}
func main() {
	r := mux.NewRouter()
	songs = append(songs, Song{ID: "1", ArtistName: "Radiohead", SongName: "BulletProofIWishIWas", Genre: "Suicide"})
	songs = append(songs, Song{ID: "2", ArtistName: "Radiohead", SongName: "FakePlasticTrees", Genre: "Suicide"})

	r.HandleFunc("/songs", getSongs).Methods("GET")
	r.HandleFunc("/songs/{id}", getSong).Methods("GET")
	r.HandleFunc("/songs", addSong).Methods("POST")
	r.HandleFunc("/songs/{id}", deleteSong).Methods("DELETE")
	r.HandleFunc("/songs/{id}", UpdSong).Methods("PUT")

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

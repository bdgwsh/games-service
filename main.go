package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type server struct {
	db *sql.DB
}

 type Game struct {
	Id string  `json:"id"`
	Title   string `json:"title"`
	Developer  string `json:"developer"`
	ReleaseDate string `json:"release_date"`
	Description string `json:"description"`
	Metacritic int   `json:"metacritic"`

}




func (s *server) getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := s.db.Query("SELECT * FROM games")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	games := make([]*Game, 0)
	for rows.Next() {
		game := new(Game)
		err := rows.Scan(&game.Id, &game.Title, &game.Developer, &game.ReleaseDate, &game.Description, &game.Metacritic)

		games = append(games, game)

		if err != nil {
			panic(err)
		}
	}
	json.NewEncoder(w).Encode(games)
}

/* func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range books {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
} */

func createGame(w http.ResponseWriter, r *http.Request) {

}

func updateGame(w http.ResponseWriter, r *http.Request) {

}

func deleteGame(w http.ResponseWriter, r *http.Request) {

}

func main() {

	var err error
	db, err := sql.Open("postgres", "user=super_postgres_user password=pass dbname=gamesdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	s := server{db: db}

	r := mux.NewRouter()

	// books = append(books, Book{ID:"1", Isbn:"448743", Title:"Go programming", Author: &Author{ Firstname:"Golang", Lastname: "Master"}})
	// books = append(books, Book{ID:"2", Isbn:"443713", Title:"Python programming", Author: &Author{ Firstname:"Python", Lastname: "Master"}})



	r.HandleFunc("/games", s.getGames).Methods("GET")
	// r.HandleFunc("/games/{id}", getGame).Methods("GET")
	r.HandleFunc("/games", createGame).Methods("POST")
	r.HandleFunc("/games/{id}", updateGame).Methods("PUT")
	r.HandleFunc("/games/{id}", deleteGame).Methods("DELETE")
	http.ListenAndServe(":8000", r)
}

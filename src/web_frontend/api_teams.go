package web_frontend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tobyjwebb/teamchess/src/teams"
)

func setJSON(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}

func (s *Server) setupTeamsRoutes() *chi.Mux {
	teams := chi.NewRouter()
	teams.Get("/", s.listTeams)
	teams.Post("/", s.CreateTeamHandler)
	teams.Post("/{team_id}/join", s.joinTeam)
	teams.Post("/{team_id}/leave", s.leaveTeam)
	return teams
}

func (s *Server) listTeams(rw http.ResponseWriter, r *http.Request) {
	// XXX implement listTeams action
	setJSON(rw)
	fmt.Fprintf(rw, `
[
    {
        "name": "team1",
        "id": "id1",
        "status": {
			"battleID":"aaaaaabbbbbbcc-1122-33-44444444",
            "status": "Battling XXXX team",
            "timestamp": "2022-22-33T11:22:33Z"
        },
        "rank": 9,
        "members": 22
    },
    {
        "name": "team2",
        "id": "id2",
        "status": {
			"battleID":"aaaaaabbbbbbcc-1122-33-44444444",
            "status": "Battling XXXX team",
            "timestamp": "2022-22-33T11:22:33Z"
        },
        "rank": 99,
        "members": 22
    },
    {
        "name": "team2.5",
        "id": "id2andahalf",
        "status": {
            "status": "idle",
            "timestamp": "2022-22-33T11:22:33Z"
        },
        "rank": 9922,
        "members": 5
    },
    {
        "name": "team3",
        "id": "id3",
        "status": {
            "status": "idle",
            "timestamp": "2022-22-33T11:22:33Z"
        },
        "rank": 922,
        "members": 5
    }
]
			`)
}

func (s *Server) CreateTeamHandler(rw http.ResponseWriter, r *http.Request) {
	owner := r.FormValue("owner")
	team := &teams.Team{
		Name:    r.FormValue("name"),
		Owner:   owner,
		Members: []string{owner},
	}

	if err := s.TeamService.CreateTeam(team); err != nil {
		log.Printf("Error creating team: %v", err)
		panic(err)
	}

	setJSON(rw)
	rw.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(team); err != nil {
		panic(err)
	}
}

func (s *Server) joinTeam(rw http.ResponseWriter, r *http.Request) {
	// XXX implement joinTeam action
	setJSON(rw)
	fmt.Fprintf(rw, `{"warning":"not implemented"}`)
}

func (s *Server) leaveTeam(rw http.ResponseWriter, r *http.Request) {
	// XXX implement leaveTeam action
	setJSON(rw)
	fmt.Fprintf(rw, `{"warning":"not implemented"}`)
}

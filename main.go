package main

import (
	"go-db/pkg/players"
	"go-db/pkg/team"
	"go-db/storage"
	"log"
)

func main() {

	storage.NewMySQL()

	myPlayers := storage.NewPlayersMySQL(storage.Pool())
	myPlayersService := players.NewService(myPlayers)

	myTeam := storage.NewTeamMySQL(storage.Pool())
	myTeamService := team.NewService(myTeam)

	if err := myPlayersService.Migrate(); err != nil {
		log.Fatalf("players.Migrate() %v", err)
	}

	if err := myTeamService.Migrate(); err != nil {
		log.Fatalf("team.Migrate() %v", err)
	}
}

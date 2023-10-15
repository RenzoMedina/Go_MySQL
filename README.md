# Conection MySQL
## Migration
We to use function of migrate, we have implement the next script
``` go
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

```
## Creation and Insert data
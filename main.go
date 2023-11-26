package main

import (
	"SongJob_Server/repo"
	"SongJob_Server/web"
	"log"
)

func main() {
	client, err := repo.GetMongoClient()
	if err != nil {
		log.Fatal("Can't get mongo client")
		return
	}
	defer repo.CloseMongoClient(client)

	router := web.NewRouter()
	if err = router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

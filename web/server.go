package web

import (
	"SongJob_Server/info"
	"SongJob_Server/repo"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// NewRouter `Gin Engine`을 초기화하고 반환한다
func NewRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // or use "*" to allow all origins
	//config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	//config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	r.Use(cors.New(config))

	r.GET("/", homeHandler)
	r.GET("/dummy", dummyHandler)
	r.GET("/jumpit", jumpItHandler)
	return r
}

// homeHandler GET("/") 엔드포인트
func homeHandler(c *gin.Context) {
	c.JSON(int(StatusOK), gin.H{
		"message": "Welcome to the Chat Server",
	})
}

// dummyHandler GET("/dummy") 엔드포인트
func dummyHandler(c *gin.Context) {
	jobPosts := repo.GetDummyJobPosts()
	c.JSON(int(StatusOK), jobPosts)
}

// jumpItHandler GET("/jumpit") 엔드포인트
func jumpItHandler(c *gin.Context) {
	client, err := repo.GetMongoClient()
	if err != nil {
		c.JSON(int(StatusInternalServerError), gin.H{
			"message": "Can't access to database",
		})
		return
	}
	//defer repo.CloseMongoClient(client)

	find, err := repo.Find(client, "job", "job_info", bson.M{})
	if err != nil {
		c.JSON(int(StatusInternalServerError), gin.H{
			"message": "Can't find data in database",
		})
		return
	}
	defer repo.CloseMongoCursor(find)

	var jobPosts []info.JobPost
	for find.Next(context.TODO()) {
		var result info.JobPost
		err = find.Decode(&result)
		if err != nil {
			log.Printf("Failed to decode data to job post: %+v", err)
			continue
		}
		jobPosts = append(jobPosts, result)
	}
	c.JSON(int(StatusOK), jobPosts)
}

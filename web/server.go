package web

import (
	"SongJob_Server/constants"
	"SongJob_Server/info"
	"SongJob_Server/repo"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
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
	r.GET("/job-posts", jobPosts)
	r.GET("/favorite-job-posts", favoriteJobPosts)
	r.PUT("/add-to-favorites", addToFavorites)
	r.PUT("/remove-from-favorites", removeFromFavorites)
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

// jobPosts GET("/job-posts") 엔드포인트
func jobPosts(c *gin.Context) {
	client, err := repo.GetMongoClient()
	if err != nil {
		c.JSON(int(StatusInternalServerError), gin.H{
			"message": "Can't access to database",
		})
		return
	}

	opts := options.Find().SetLimit(3)
	find, err := repo.Find(
		client,
		string(constants.JobDataBaseName),
		string(constants.JobInfoCollectionName),
		bson.M{string(constants.HasBeenViewed): false},
		opts)
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

// favoriteJobPosts GET("/favorite-job-posts") 엔드포인트
func favoriteJobPosts(c *gin.Context) {
	client, err := repo.GetMongoClient()
	if err != nil {
		c.JSON(int(StatusInternalServerError), gin.H{
			"message": "Can't access to database",
		})
		return
	}

	find, err := repo.Find(
		client,
		string(constants.JobDataBaseName),
		string(constants.JobInfoCollectionName),
		bson.M{
			string(constants.Favorite): true,
		},
		nil)
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

// addToFavorites PUT("/add-to-favorites") 엔트포인트
func addToFavorites(c *gin.Context) {
	var request UpdateFavoriteRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	client, err := repo.GetMongoClient()
	if err != nil {
		c.JSON(int(StatusInternalServerError), gin.H{
			"message": "Can't access to database",
		})
		return
	}

	updateResult, err := repo.Update(
		client,
		string(constants.JobDataBaseName),
		string(constants.JobInfoCollectionName),
		bson.M{string(constants.Link): request.Link},
		bson.M{"$set": bson.M{
			string(constants.Favorite):      true,
			string(constants.HasBeenViewed): true,
		}})
	if err != nil {
		log.Printf("err: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating the document",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Favorite updated successfully",
		"result":  updateResult,
	})
}

// removeFromFavorites PUT("/remove-from-favorites") 엔드포인트
func removeFromFavorites(c *gin.Context) {
	var request UpdateFavoriteRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	client, err := repo.GetMongoClient()
	if err != nil {
		c.JSON(int(StatusInternalServerError), gin.H{
			"message": "Can't access to database",
		})
		return
	}

	updateResult, err := repo.Update(
		client,
		string(constants.JobDataBaseName),
		string(constants.JobInfoCollectionName),
		bson.M{string(constants.Link): request.Link},
		bson.M{"$set": bson.M{
			string(constants.Favorite):      false,
			string(constants.HasBeenViewed): true,
		}})
	if err != nil {
		log.Printf("err: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating the document",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Favorite updated successfully",
		"result":  updateResult,
	})
}

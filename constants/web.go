package constants

type DatabaseName string

const (
	JobDataBaseName DatabaseName = "job"
)

type CollectionName string

const (
	JobInfoCollectionName CollectionName = "job_info"
)

type JobInfoField string

const (
	Link          JobInfoField = "link"
	HasBeenViewed JobInfoField = "hasBeenViewed"
	Favorite      JobInfoField = "favorite"
)

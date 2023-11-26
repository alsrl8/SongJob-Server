package repo

import "fmt"

type CannotConnectToMongoCloudError struct {
}

func (e *CannotConnectToMongoCloudError) Error() string {
	return fmt.Sprintf("can't connect to mongo cloud")
}

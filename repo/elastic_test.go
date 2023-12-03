package repo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewElasticSearchClient(t *testing.T) {
	client := NewElasticSearchClient()
	assert.NotNil(t, client)
}

func TestElasticSearchClient_IndexDocuments(t *testing.T) {
	indexName := "myindex"

	client := NewElasticSearchClient()
	assert.NotNil(t, client)

	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch",
	}

	client.IndexDocuments(indexName, document)
}

func TestElasticSearchClient_GetDocument(t *testing.T) {
	indexName := "myindex"
	docId := "oC8LL4wBouBJq3Xl7yLW"

	client := NewElasticSearchClient()
	assert.NotNil(t, client)

	client.GetDocument(indexName, docId)
}

func TestElasticSearchClient_SearchAllDocumentsDocument(t *testing.T) {
	indexName := "myindex"

	client := NewElasticSearchClient()
	assert.NotNil(t, client)

	documents, err := client.SearchAllDocuments(indexName)
	assert.Nil(t, err)

	for _, doc := range documents {
		for k, v := range doc {
			fmt.Printf("%s: %v, ", k, v)
		}
		fmt.Printf("\n")
	}
}

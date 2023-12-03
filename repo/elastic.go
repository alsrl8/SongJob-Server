package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type ElasticSearchClient struct {
	instance *elasticsearch.Client
}

type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewElasticSearchClient() *ElasticSearchClient {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * 1,
		},
	}
	instance, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err = instance.Ping(instance.Ping.WithContext(ctx))
	if err != nil {
		log.Printf("Error pinging ElasticSearch: %+v", err)
		return nil
	}

	return &ElasticSearchClient{
		instance: instance,
	}
}

func (es *ElasticSearchClient) CreateIndex(indexName string) {
	_, err := es.instance.Indices.Create(indexName)
	if err != nil {
		log.Printf("Failed creating index(%s): %+v", indexName, err)
		return
	}
}

func (es *ElasticSearchClient) GetDocument(indexName string, docId string) {
	ret, err := es.instance.Get(indexName, docId)
	if err != nil {
		log.Printf("Failed getting document(%s) from index(%s)", docId, indexName)
		return
	}
	log.Printf("ret: %+v", ret)
}

func (es *ElasticSearchClient) IndexDocuments(indexName string, document interface{}) {
	data, err := json.Marshal(document)
	if err != nil {
		log.Printf("Failed marshing document(%+v) to array of bytes: %+v", document, err)
		return
	}
	ret, err := es.instance.Index(indexName, bytes.NewReader(data))
	if err != nil {
		log.Printf("Failed indexing document(%+v) from index(%s)", document, indexName)
		return
	}
	log.Printf("ret: %+v", ret)
}

func (es *ElasticSearchClient) SearchAllDocuments(indexName string) ([]map[string]interface{}, error) {
	query := `{ "query": { "match_all": {} } }`
	return es.SearchDocuments(indexName, query)
}

func (es *ElasticSearchClient) SearchDocuments(indexName string, query string) ([]map[string]interface{}, error) {
	res, err := es.instance.Search(
		es.instance.Search.WithIndex(indexName),
		es.instance.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Printf("Failed seraching document from index(%s)", indexName)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Failed closing response body after searching index(%s): %+v", indexName, err)
		}
	}(res.Body)

	var r SearchResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Failed decoding JSON response from Elasticsearch index(%s): %+v", indexName, err)
		return nil, err
	}

	var documents []map[string]interface{}
	for _, hit := range r.Hits.Hits {
		documents = append(documents, hit.Source)
	}

	return documents, nil
}

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	elastic "gopkg.in/olivere/elastic.v6"
)

type WikiEntry struct {
	User    string                `json:"user"`
	Body    string                `json:"body"`
	Title   string                `json:"title"`
	Created time.Time             `json:"created,omitempty"`
	Suggest *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"wikiEntry":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"body":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
        "title":{
          "type":"text",
          "store":true,
          "fielddata": true
        },
				"created":{
					"type":"date"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

/*********GET ES CLIENT************/
func getClient() *elastic.Client {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return client
}

/**********VERIFY OR CREATE INDEX***********/
func verifyIndex(client *elastic.Client, index string) {

	ctx := context.Background()

	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		fmt.Printf("failed to determine if index exists...")
	}

	if !exists {

		newIndex, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("new index created . . . ")
		}

		if !newIndex.Acknowledged {
			fmt.Printf("Could not create new ES index 'msg'")
		}
	}
}

/**********ADD ENTRY TO ES********************/
func addWikiEntry(client *elastic.Client, entry WikiEntry, id string) {

	ctx := context.Background()

	put, err := client.Index().
		Index("wiki").
		Type("wikiEntry").
		Id(id).
		BodyJson(entry).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed wiki entry %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

/************GET ENTRY BY ID FROM INDEX*****************/
func getWikiEntry(client *elastic.Client, id string) (WikiEntry, error) {

	ctx := context.Background()
	var result WikiEntry

	get, err := client.Get().
		Index("wiki").
		Type("wikiEntry").
		Id(id).
		Do(ctx)

	if err != nil {
		fmt.Printf("no entry fond for id - create entry...")
		return result, errors.New("no entry found...")
	} else {
		if get.Found {
			bytes, err := json.Marshal(get.Source)
			if err != nil {
				fmt.Printf("uable to marshal entry")
			}
			json.Unmarshal(bytes, &result)
		}
		return result, nil
	}
}

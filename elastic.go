package main

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/olivere/elastic.v6"
)

type Msg struct {
	User    string                `json:"user"`
	Message string                `json:"message"`
	Created time.Time             `json:"created,omitempty"`
	Title   string                `json:"title"`
	Suggest *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"msg":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
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

func makeOrCreateIndex(client *elastic.Client, index string) {

	ctx := context.Background()

	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("index exists...")
	}

	if !exists {
		newIndex, err := client.CreateIndex("msgz").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}

		if !newIndex.Acknowledged {
			fmt.Printf("Could not create new ES index 'msg'")
		}
	}
}

func addMsg(client *elastic.Client, msg Msg, id string) {
	//msg1 := Msg{User: "fliropp", Message: "Hey Ho, let's go!!", Title: "ES entry #1"}
	ctx := context.Background()

	put1, err := client.Index().
		Index("msgz").
		Type("msg").
		Id(id).
		BodyJson(msg).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed msg %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func getMsg(client *elastic.Client, id string) {

	ctx := context.Background()

	get1, err := client.Get().
		Index("msgz").
		Type("msg").
		Id(id).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

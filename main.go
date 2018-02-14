package main

import (
	"context"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v6"
)

func main() {
	/*http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)*/

	/* init elastic client */
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

	makeOrCreateIndex(client, "msgz")
	getMsg(client, "1")
	msg := Msg{User: "foobar", Message: "No no, let's stay at home!!", Title: "ES entry #2"}
	addMsg(client, msg, "2")
	getMsg(client, "2")

}

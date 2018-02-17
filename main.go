package main

import "net/http"

func main() {

	/* init elastic client */
	/*client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)*/

	/* init controllers */
	http.HandleFunc("/view/", viewHandler)
	//http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)

	/*verifyIndex(client, "wiki")
	entry := WikiEntry{User: "fliropp", Body: "wiki entry no3#", Title: "wiki#3"}
	addWikiEntry(client, entry, "3")
	getWikiEntry(client, "3")*/

}

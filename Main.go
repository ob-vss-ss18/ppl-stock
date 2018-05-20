package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/graphql"
	"encoding/json"
	"io/ioutil"
	"github.com/ob-vss-ss18/ppl-stock/pplStock"
)

func main() {

	http.HandleFunc("/graphql", func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		// httpRequest.URL.Query().Get("query"):
		// returns "message" (without the qoutes) for the example below
		bodyBytes, error := ioutil.ReadAll(httpRequest.Body)
		if error == nil {
			bodyString := string(bodyBytes[:len(bodyBytes)])
			result := executeQuery(bodyString)
			json.NewEncoder(httpResponseWriter).Encode(result)
		} else {
			fmt.Println(error)
		}
		httpRequest.Body.Close()
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl --data 'query{greet(name:\"nico\") message}' http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)

}

/*
	A small function to which will call graphql.Do(), which will handle resolving the
	request and returns the results.
 */
func executeQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: pplStock.PPLStockSchema,
		// this variable is set to "message" for the example below
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

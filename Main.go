package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/ob-vss-ss18/ppl-stock/pplStock"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/ppl-stock", func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
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
	fmt.Println("Curl test: curl --data 'query{ski(id:10){brand id useCase}}' http://localhost:8080/ppl-stock")
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

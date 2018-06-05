package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/ob-vss-ss18/ppl-stock/pplStock"
	"net/http"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/graphql", func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		fmt.Print("Req: "); fmt.Println(httpRequest);
		fmt.Print("Url: "); fmt.Println(httpRequest.URL);
		fmt.Print("Qry: "); fmt.Println(httpRequest.URL.Query());
		fmt.Print("Qdt: "); fmt.Println(httpRequest.URL.Query()["query"]);
		dataArr, err := httpRequest.URL.Query()["query"];
		if  (!err || len(dataArr) < 1)  {
			fmt.Println("Get Data from URL-Query failed, trying to get from body instead...");

			// returns "message" (without the qoutes) for the example below
			bodyBytes, error := ioutil.ReadAll(httpRequest.Body)
			if error == nil {
				bodyString := string(bodyBytes[:len(bodyBytes)])
				result := executeQuery(bodyString)
				json.NewEncoder(httpResponseWriter).Encode(result)
			} else {
				fmt.Println(error)
			}
		}  else  {
			data := dataArr[0];
			result := executeQuery(data);
			json.NewEncoder(httpResponseWriter).Encode(result);
		}

		httpRequest.Body.Close()
	});

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Curl test: curl --data 'query{Ski(id:10){brand id useCase}}' http://localhost:8080/graphql")
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
		fmt.Printf("wrong result, unexpected errors: %v\n\n", result.Errors)
	}
	return result
}

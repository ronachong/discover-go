package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* following the suggestions of http://www.freenerd.de/accessing-the-github-api-with-golang/ */

type movieData struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	Response   string
}

func main() {

	// request data stream from api? seems to match protocol for http.Get on https://golang.org/doc/articles/wiki/
	res, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		log.Fatal(err)
	}

	// read body of returned data (?)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	log.Printf("The status code is 200 / OK.")

	// parse json!
	createStruct()      // calling function.. not sure I'm supposed to do it this way
	data := movieData{} // set data to structure

	err = json.Unmarshal(body, &data) //I think Franck said that this puts the data into the structure??
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The movie: %s was released in %s - the IMBD rating is %s with %s votes", data.Title, data.Year, data.imdbRating, data.imdbVotes)
	return
}

/*
func main() {
  resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
  if err != nil {
    log.Fatal(err) //handles error?
  }
  defer resp.Body.Close()
  if err := json.NewDecoder(resp.Body).Decode(&s); err != nil { // if NewDecoder has error decoding json
    fmt.Errorf("error parsing body %v", err) //return error
    return
  }
  //else
  fmt.Printf("The status code is: %v - OK\n", StatusOK)
  fmt.Printf("The movie: %s was released in %d - the IMBD rating is %d with %d votes", resp.Title, resp.Year, resp.imdbRating)
}

*/

/* getContent(url) accepts a URL, creates a HTTP client(?), sends a request, and returns the response or any errors.
func getContent(url string) ([]byte , error) {

  // http.NewRequest compiles a request?
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err // this doesn't make 100% sense to me...
  }

  // HTTP client sends the request (?)
  client := &http.Client{}
  resp, err := client.Do(req) // what does this line do?
  if err != nil {

    return nil, err

  }

  // Defer the closing of the body
  defer resp.Body.Close()

  // Read the content into ..? in the exp, a byte array
  body, err := io.util.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return body, nil // return line returns body as bytes?
}

func parse(response string) (*)

*/

/*
func createStruct() data {
	type data struct {
		Title      string
		Year       string
		Rated      string
		Released   string
		Runtime    string
		Genre      string
		Director   string
		Writer     string
		Actors     string
		Plot       string
		Language   string
		Country    string
		Awards     string
		Poster     string
		Metascore  string
		imdbRating string
		imdbVotes  string
		imdbID     string
		Type       string
		Response   string
	}
}

*/

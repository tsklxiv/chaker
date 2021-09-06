/*
	This is Hecker, a Hacker News 'client' written in Go.
	(Currently it is more like a scraper than a client)
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"github.com/PuerkitoBio/goquery"
	"github.com/muesli/termenv"
)

// The submission's data struct
type Submission struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// List of links and titles (For feeding the TUI part)
var submissions []Submission = []Submission{}

// Helper functions
func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func check_status_code(msg string, res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("%s | %d %s", msg, res.StatusCode, res.Status)
	}
}

func Scrape() []Submission {	
	// Scrape the news
	res, err := http.Get("https://news.ycombinator.com/news")
	check_err(err)
	defer res.Body.Close()
	check_status_code("Status code error", res)

	// Create a document for scraping
	doc, err := goquery.NewDocumentFromReader(res.Body)
	check_err(err)

	// Scrape the submissions
	doc.Find("tr .athing").Each(func(i int, s *goquery.Selection) {
		// Take the ID of the submission (very important for scraping the submissions data)
		id, _ := s.Attr("id")

		// Scrape the submissions data using the ID from Hacker New's Firebase DB
		json_res, err := http.Get(spf("https://hacker-news.firebaseio.com/v0/item/%s.json?print=pretty", id))
		check_err(err)
		defer json_res.Body.Close()
		check_status_code(spf("Cannot scrape data of ID %s", id), json_res)

		doc, err := goquery.NewDocumentFromReader(json_res.Body)
		check_err(err)
		json_data := doc.Children().Text()

		var submission Submission

		// Unmarshal the JSON data to Submission
		err = json.Unmarshal([]byte(json_data), &submission)
		check_err(err)
		
		// In case that the submission doesn't have an URL (like with a local post and not a link)
		// We will just give the post link instead
		if submission.URL == "" {
			submission.URL = spf("https://news.ycombinator.com/item?id=%s", id)
		}

		submissions = append(submissions, submission) 
	})

	return submissions
}

func main() {
	termenv.ClearScreen()
	termenv.AltScreen()

	fmt.Println("Please be patient...")

	s := Scrape()

	tui(s)
}

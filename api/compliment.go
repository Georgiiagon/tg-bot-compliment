package api

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetRandomCompliment() string {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(80) + 1

	res, err := http.Get("http://kompli.me/komplimenty/page/" + strconv.Itoa(randInt))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var compliments []string
	// Find the review items
	doc.Find(".post-card .post-card__title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		compliment := s.Find("a").Text()
		compliments = append(compliments, compliment)
	})

	return compliments[rand.Intn(len(compliments))]
}

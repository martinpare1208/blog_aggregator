package command

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/martinpare1208/gator/internal/database"
)

// const link = "https://www.wagslane.dev/index.xml"



type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	// create get reqeust
	req, err := http.NewRequestWithContext(ctx,"GET", feedURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// decode xml
	defer resp.Body.Close()

	var rss RSSFeed
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}


	return &rss, nil
}


func Agg(s *State, cmd Command) (error) {

	if len (cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <duration of feed fetch>", cmd.Name)
	}

	timeBetweenRequests := cmd.Args[0]
	
	duration, err := time.ParseDuration(timeBetweenRequests)
	if err != nil {
		return err
	}


	ticker := time.NewTicker(duration)
	for ; ; <- ticker.C {
		scrapeFeeds(s)
	}



}


func scrapeFeeds(s *State) (error) {
	context := context.Background()
	feed, err := s.DBConnection.GetNextFeedToFetch(context)
	if err != nil {
		return err
	}

	// If successful, update the feed that was fetched
	_, err = s.DBConnection.MarkFeedFetched(context, database.MarkFeedFetchedParams{LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}, UpdatedAt: time.Now(), ID: feed.ID})
	if err != nil {
		return err
	}


	fetchedFeed, err := FetchFeed(context, feed.Url)
	if err != nil {
		return err
	}


	// Print to console
	fmt.Print("Grabbing feed...\n")
	title := html.UnescapeString(fetchedFeed.Channel.Item[0].Title)

	fmt.Printf("%s\n", title)

	return nil



}




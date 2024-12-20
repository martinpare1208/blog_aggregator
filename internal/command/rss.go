package command

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

const link = "https://www.wagslane.dev/index.xml"



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

	context := context.Background()
	rss, err := FetchFeed(context, link)
	if err != nil {
		log.Fatal(err)
		return err
	}

	title := html.UnescapeString(rss.Channel.Item[0].Title)
	desc := html.UnescapeString(rss.Channel.Item[0].Description)

	fmt.Printf("%s\n", title)
	fmt.Printf("%s\n", desc)

	return nil
}




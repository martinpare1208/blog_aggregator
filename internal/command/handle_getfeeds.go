package command

import (
	"context"
	"fmt"
	"log"
)

func HandlerGetFeeds(s* State, cmd Command) (error) {

	context := context.Background()
	data, err := s.DBConnection.GetFeeds(context)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, feed := range data {
		userData, err := s.DBConnection.GetUserById(context, feed.UserID)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("Name: %s | URL: %s | Created by: %s\n", feed.Name, feed.Url, userData.Name)
	}

	return nil
}
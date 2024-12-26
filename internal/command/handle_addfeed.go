package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/martinpare1208/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	
	// save to database
	context := context.Background()
	name := cmd.Args[0]
	url := cmd.Args[1]

	feedData, err := s.DBConnection.CreateFeed(context, database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: user.ID,})
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = s.DBConnection.CreateFeedFollow(context, database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feedData.ID})
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Feed: %s added!\n", name)


	return nil
}
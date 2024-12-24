package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/martinpare1208/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	// process input
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name of feed> %s <url of feed>", cmd.Args[0], cmd.Args[1])
	}

	// save to database
	context := context.Background()
	name := cmd.Args[0]
	url := cmd.Args[1]
	userInfo, err := s.DBConnection.GetUser(context, s.CfgPtr.CurrentUser)
	if err != nil {
		log.Fatal(err)
		return err
	}

	feedData, err := s.DBConnection.CreateFeed(context, database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: userInfo.ID,})
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = s.DBConnection.CreateFeedFollow(context, database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: userInfo.ID, FeedID: feedData.ID})
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Feed: %s added!\n", name)


	return nil
}
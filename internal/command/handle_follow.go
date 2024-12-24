package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/martinpare1208/gator/internal/database"
)

func HandlerFollowFeed(s *State, cmd Command) error {
	// process input
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url of feed>", cmd.Args[0])
	}

	// save to database
	context := context.Background()
	url := cmd.Args[0]
	userInfo, err := s.DBConnection.GetUser(context, s.CfgPtr.CurrentUser)
	if err != nil {
		log.Fatal(err)
		return err
	}

	feedInfo, err := s.DBConnection.GetFeedByUrl(context, url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = s.DBConnection.CreateFeedFollow(context, database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: userInfo.ID, FeedID: feedInfo.ID})
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Following feed")


	return nil
}
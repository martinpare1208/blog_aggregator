package command

import (
	"context"
	"fmt"
	"log"
)

func HandlerFollowing(s *State, cmd Command) error {

	// access to db
	context := context.Background()
	userInfo, err := s.DBConnection.GetUser(context, s.CfgPtr.CurrentUser)
	if err != nil {
		log.Fatal(err)
		return err
	}


	feeds, err := s.DBConnection.GetFeedFollowsForUser(context, userInfo.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}
	
	fmt.Println("Feed printing complete.")

	return nil
}
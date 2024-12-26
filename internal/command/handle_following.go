package command

import (
	"context"
	"fmt"
	"log"

	"github.com/martinpare1208/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {


	fmt.Printf("currentuser: %s\n", user.Name)
	context := context.Background()
	feeds, err := s.DBConnection.GetFeedFollowsForUser(context, user.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}


	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
		fmt.Printf("%s\n", feed.UserName)
	}
	
	fmt.Println("Feed printing complete.")

	return nil
}
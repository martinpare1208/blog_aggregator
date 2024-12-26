package command

import (
	"context"
	"fmt"

	"github.com/martinpare1208/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url of link>", cmd.Args[0])
	}
	url := cmd.Args[0]

	// unfollow and delete from database
	context := context.Background()



	_, err := s.DBConnection.UnfollowFeed(context, database.UnfollowFeedParams{UserID: user.ID, Url: url})
	if err != nil {
		return err
	}

	return nil
}
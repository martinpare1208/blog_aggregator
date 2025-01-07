package command

import (
	"context"
	"fmt"
	"html"
	"strconv"

	"github.com/martinpare1208/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <limit rows>", cmd.Name)
	}

	limit := cmd.Args[0]
	i, err := strconv.Atoi(limit)
	if err != nil {
		return fmt.Errorf("not an int")
	}

	i32 := int32(i)

	posts, err := s.DBConnection.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: i32})

	if err != nil {
		return(err)
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("%s\n", html.UnescapeString(post.Description))
	}



	return nil
}
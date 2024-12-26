package command

import (
	"context"
	"github.com/martinpare1208/gator/internal/database"
)

func MiddlewareLoggedIn(
    handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {


		return func(s *State, cmd Command) (error) {

	

			context := context.Background()
			userInfo, err := s.DBConnection.GetUser(context, s.CfgPtr.CurrentUser)
			if err != nil {
				return err
			}

			return handler(s, cmd, userInfo)

		}
	} 
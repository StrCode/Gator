package main

import (
	"context"
	"fmt"
	"time"

	"github.com/StrCode/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	url := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("issues retrieving feed information: %v", err)
	}

	if feed.Name == "" {
		return fmt.Errorf("feed does not exist on that url")
	}

	ffRow, err := s.db.CreateFeedFellow(context.Background(), database.CreateFeedFellowParams{
		uuid.New(),
		user.ID,
		feed.ID,
		time.Now().UTC(),
		time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	fmt.Printf("* Name:          %s\n", ffRow.FeedName)
	fmt.Printf("* UserName:           %s\n", ffRow.UserName)

	return nil
}

func handerListUserFollows(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.Feedname)
	}

	return nil
}

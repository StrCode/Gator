package main

import (
	"context"
	"fmt"
	"time"

	"github.com/StrCode/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	url := cmd.Args[0]

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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		user.ID,
		feedUrl,
	})
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	fmt.Println("Feed has been unfollowed successfully!")

	return nil
}

func handerListUserFollows(s *state, cmd command, user database.User) error {
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

package main

import (
	"context"
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/tonyserranodev/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	var err error

	if len(cmd.args) > 0 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for i := range posts {
		formatPost(&posts[i])
		printPost(posts[i])
	}

	return nil
}

func formatPost(post *database.Post) {
	post.Title = html.UnescapeString(post.Title)
	post.Url = html.UnescapeString(post.Url)
	post.Description = html.UnescapeString(post.Description)
}

func printPost(post database.Post) {
	description := post.Description
	if len(description) > 200 {
		description = strings.TrimSpace(description[:200]) + "..."
	}

	fmt.Println("====================================")
	fmt.Printf("* Title:       %s\n", post.Title)
	fmt.Printf("* URL:         %s\n", post.Url)
	fmt.Printf("* Description: %s\n", description)
	fmt.Printf("* Published:   %s\n", post.PublishedAt.Format("Mon, 02 Jan 2006 15:04:05 MST"))
	fmt.Println("====================================")
	fmt.Println()
}

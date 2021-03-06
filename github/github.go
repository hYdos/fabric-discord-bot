package github

import (
	"context"
	"github.com/FabricMC/fabric-discord-bot/utils"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

var (
	client       *github.Client
	Organization string
)

func Connect() error {
	token, err := utils.GetEnv("GITHUB_TOKEN")
	if err != nil {
		return err
	}

	org, err := utils.GetEnv("GITHUB_ORG")
	if err != nil {
		return err
	}
	Organization = org

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
	return nil
}

// Blocks a user if they are not already blocked
//
// Returns true if they were previously blocked, false if newly blocked
func BlockUser(org string, user string) (bool, error) {
	blocked, _, err := client.Organizations.IsBlocked(context.Background(), org, user)
	if err != nil {
		return false, err
	}

	if blocked {
		return true, nil
	}

	_, err = client.Organizations.BlockUser(context.Background(), org, user)
	if err != nil {
		return false, err
	}

	return false, nil
}

// Unblocks a user if they are not already blocked
//
// Returns false if they were not blocked before, true when unblocked
func UnblockUser(org string, user string) (bool, error) {
	blocked, _, err := client.Organizations.IsBlocked(context.Background(), org, user)
	if err != nil {
		return false, err
	}

	if !blocked {
		return false, nil
	}

	_, err = client.Organizations.UnblockUser(context.Background(), org, user)
	if err != nil {
		return false, err
	}

	return true, nil
}

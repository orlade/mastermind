package main

import (
	"context"
	"github.com/orlade/mastermind/internal/handlers"
	"log"

	"github.com/anz-bank/sysl-go/core"

	mastermind "github.com/orlade/mastermind/gen/pkg/servers/Mastermind"
)

func main() {
	type AppConfig struct {
		// Define app-level config fields here.
	}
	log.Fatal(mastermind.Serve(context.Background(),
		func(ctx context.Context, config AppConfig) (*mastermind.ServiceInterface, *core.Hooks, error) {
			// Perform one-time setup based on config here.
			return &mastermind.ServiceInterface{
				GetReposIssuesList: handlers.GetReposIssuesList,
			}, nil, nil
		},
	))
}

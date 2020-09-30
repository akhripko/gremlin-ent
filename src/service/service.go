package service

import (
	"context"
	"log"

	"github.com/akhripko/gremlin-ent/ent"
)

type Service struct {
	ctx    context.Context
	client *ent.Client
}

func New(ctx context.Context, addr string) (*Service, error) {
	client, err := ent.Open("gremlin", addr)
	if err != nil {
		log.Fatal(err)
	}
	return &Service{
		client: client,
		ctx:    ctx,
	}, nil
}

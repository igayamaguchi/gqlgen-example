package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/igayamaguchi/gqlgen-example/enum/graph/generated"
	"github.com/igayamaguchi/gqlgen-example/enum/graph/model"
)

var data = []Article{
	{"1", "title1", "status"},
	{"2", "title2", "PUBLISHED"},
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	for _, d := range data {
		if d.Id == id {
			s := model.Status(d.Status)
			if !s.IsValid() {
				log.Printf("不正なステータス:%s", s)
				return nil, errors.New("internal server error")
			}
			return &model.Article{
				ID:     d.Id,
				Title:  d.Title,
				Status: s,
			}, nil
		}
	}

	return nil, errors.New("not found")
}

func (r *queryResolver) Articles(ctx context.Context, status model.Status) ([]*model.Article, error) {
	result := []*model.Article{}

	for _, d := range data {
		s := model.Status(d.Status)
		if !s.IsValid() {
			log.Printf("不正なステータス:%s", s)
			continue
		}
		if s == status {
			result = append(result, &model.Article{
				ID:     d.Id,
				Title:  d.Title,
				Status: s,
			})
		}
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type Article struct {
	Id     string
	Title  string
	Status string
}

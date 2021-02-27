package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/valdirmendesdev/fc2-graphql/graph/generated"
	"github.com/valdirmendesdev/fc2-graphql/graph/model"
)

func genID() string {
	return fmt.Sprintf("T%d", rand.Int())
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := model.Category{
		ID:          genID(),
		Name:        input.Name,
		Description: &input.Description,
	}

	r.Categories = append(r.Resolver.Categories, &category)
	return &category, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	var category *model.Category

	for _, c := range r.Categories {
		if c.ID == input.CategoryID {
			category = c
		}
	}

	course := model.Course{
		ID:       genID(),
		Description: &input.Description,
		Name:     input.Name,
		Category: category,
	}

	r.Courses = append(r.Courses, &course)

	return &course, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	var course *model.Course

	for _, c := range r.Courses {
		if c.ID == input.CourseID {
			course = c
		}
	}

	chapter := model.Chapter{
		ID:       genID(),
		Name:     input.Name,
		Course: course,
	}

	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

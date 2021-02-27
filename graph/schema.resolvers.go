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

func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	var courses []*model.Course

	for _, course := range r.Resolver.Courses {
		if course.Category.ID == obj.ID {
			courses = append(courses, course)
		}
	}
	return courses, nil
}

func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	var chapters []*model.Chapter

	for _, chapter := range r.Resolver.Chapters {
		if chapter.Course.ID == obj.ID {
			chapters = append(chapters, chapter)
		}
	}
	return chapters, nil
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
		ID:          genID(),
		Description: &input.Description,
		Name:        input.Name,
		Category:    category,
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
		ID:     genID(),
		Name:   input.Name,
		Course: course,
	}

	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func genID() string {
	return fmt.Sprintf("T%d", rand.Int())
}

package application

import (
	"context"

	"../domain"
	"../repository"
)

// UserInteractor provides use-case
type UserInteractor struct {
	Repository repository.UserRepository
}

// GetUser returns user
func (intractor UserInteractor) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return intractor.Repository.Get(ctx, id)
}

// GetUsers returns user list
func (intractor UserInteractor) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return intractor.Repository.GetAll(ctx)
}

// AddUser saves new user
func (intractor UserInteractor) AddUser(ctx context.Context, name string, pass string, dob string) error {
	u, err := domain.NewUser(name, pass, dob)
	if err != nil {
		return err
	}
	return intractor.Repository.Save(ctx, u)
}

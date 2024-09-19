package task

import (
	"context"

	"github.com/hadroncorp/geck/data/persistence"
	"github.com/hadroncorp/geck/systemerror"
	"github.com/hadroncorp/geck/validation"
)

type Service struct {
	Repository Repository
	Validator  validation.Validator
}

func NewService(repository Repository, validator validation.Validator) Service {
	return Service{
		Repository: repository,
		Validator:  validator,
	}
}

func (s Service) Create(ctx context.Context, cmd CreateCommand) error {
	if err := s.Validator.Validate(ctx, cmd); err != nil {
		return err
	}
	entity := Task{
		Auditable: persistence.NewAuditable(ctx),
		ID:        cmd.TaskID,
		Name:      cmd.Name,
		Status:    cmd.Status,
	}
	return s.Repository.Save(ctx, entity)
}

func (s Service) Get(ctx context.Context, key string) (Task, error) {
	entity, err := s.Repository.FindByKey(ctx, key)
	if err != nil {
		return Task{}, err
	} else if entity == nil {
		return Task{}, systemerror.NewResourceNotFound[Task](key)
	}
	return *entity, nil
}

func (s Service) Delete(ctx context.Context, key string) error {
	entity, err := s.Get(ctx, key)
	if err != nil {
		return err
	}
	return s.Repository.Remove(ctx, entity)
}

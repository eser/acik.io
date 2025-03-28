package profiles

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFailedToGetRecord   = errors.New("failed to get record")
	ErrFailedToListRecords = errors.New("failed to list records")
	// ErrFailedToCreateRecord = errors.New("failed to create record").
)

type Repository interface {
	GetProfileById(ctx context.Context, id string) (*Profile, error)
	GetProfileBySlug(ctx context.Context, slug string) (*Profile, error)
	ListProfiles(ctx context.Context) ([]*Profile, error)
	// CreateProfile(ctx context.Context, arg CreateProfileParams) (*Profile, error)
	// UpdateProfile(ctx context.Context, arg UpdateProfileParams) (int64, error)
	// DeleteProfile(ctx context.Context, id string) (int64, error)
}

type Service struct {
	repo Repository

	idGenerator RecordIDGenerator
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo, idGenerator: DefaultIDGenerator}
}

func (s *Service) GetById(ctx context.Context, id string) (*Profile, error) {
	record, err := s.repo.GetProfileById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w(id: %s): %w", ErrFailedToGetRecord, id, err)
	}

	return record, nil
}

func (s *Service) GetBySlug(ctx context.Context, slug string) (*Profile, error) {
	record, err := s.repo.GetProfileBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("%w(slug: %s): %w", ErrFailedToGetRecord, slug, err)
	}

	return record, nil
}

func (s *Service) List(ctx context.Context) ([]*Profile, error) {
	records, err := s.repo.ListProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToListRecords, err)
	}

	return records, nil
}

// func (s *Service) Create(ctx context.Context, input *Profile) (*Profile, error) {
// 	record, err := s.repo.CreateProfile(ctx, input)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %w", ErrFailedToCreateRecord, err)
// 	}

// 	return record, nil
// }

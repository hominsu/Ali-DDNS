package biz

import (
	"context"
	"go.uber.org/zap"
)

type DomainUser struct {
	Username   string
	Password   string
	UUID       string
	DomainName string
}

type DomainUserRepo interface {
	AddUser(ctx context.Context, du *DomainUser) (bool, error)
	DelUser(ctx context.Context, du *DomainUser) (int64, error)
	IsUserExists(ctx context.Context, du *DomainUser) (bool, error)
	GetUserPassword(ctx context.Context, du *DomainUser) (string, error)

	AddDevice(ctx context.Context, du *DomainUser) (bool, error)
	GetDevice(ctx context.Context, du *DomainUser) ([]string, error)
	DelDevice(ctx context.Context, du *DomainUser) (int64, error)
	GetAllDevice(ctx context.Context, du *DomainUser) ([]string, error)

	AddDomainName(ctx context.Context, du *DomainUser) (bool, error)
	GetDomainName(ctx context.Context, du *DomainUser) ([]string, error)
	DelDomainName(ctx context.Context, du *DomainUser) (int64, error)
	GetAllDomainName(ctx context.Context, du *DomainUser) ([]string, error)
}

type DomainUserUsecase struct {
	repo   DomainUserRepo
	logger *zap.SugaredLogger
}

func NewDomainUserUsecase(repo DomainUserRepo, logger *zap.Logger) *DomainUserUsecase {
	return &DomainUserUsecase{
		repo:   repo,
		logger: logger.Sugar(),
	}
}

func (uc *DomainUserUsecase) AddUser(ctx context.Context, du *DomainUser) (bool, error) {
	return uc.repo.AddUser(ctx, du)
}

func (uc *DomainUserUsecase) DelUser(ctx context.Context, du *DomainUser) (int64, error) {
	return uc.repo.DelUser(ctx, du)
}

func (uc *DomainUserUsecase) IsUserExists(ctx context.Context, du *DomainUser) (bool, error) {
	return uc.repo.IsUserExists(ctx, du)
}

func (uc *DomainUserUsecase) GetUserPassword(ctx context.Context, du *DomainUser) (string, error) {
	return uc.repo.GetUserPassword(ctx, du)
}

func (uc *DomainUserUsecase) AddDevice(ctx context.Context, du *DomainUser) (bool, error) {
	return uc.repo.AddDevice(ctx, du)
}

func (uc *DomainUserUsecase) GetDevice(ctx context.Context, du *DomainUser) ([]string, error) {
	return uc.repo.GetDevice(ctx, du)
}

func (uc *DomainUserUsecase) DelDevice(ctx context.Context, du *DomainUser) (int64, error) {
	return uc.repo.DelDevice(ctx, du)
}

func (uc *DomainUserUsecase) GetAllDevice(ctx context.Context, du *DomainUser) ([]string, error) {
	return uc.repo.GetAllDevice(ctx, du)
}

func (uc *DomainUserUsecase) AddDomainName(ctx context.Context, du *DomainUser) (bool, error) {
	return uc.repo.AddDomainName(ctx, du)
}

func (uc *DomainUserUsecase) GetDomainName(ctx context.Context, du *DomainUser) ([]string, error) {
	return uc.repo.GetDomainName(ctx, du)
}

func (uc *DomainUserUsecase) DelDomainName(ctx context.Context, du *DomainUser) (int64, error) {
	return uc.repo.DelDomainName(ctx, du)
}

func (uc *DomainUserUsecase) GetAllDomainName(ctx context.Context, du *DomainUser) ([]string, error) {
	return uc.repo.GetAllDomainName(ctx, du)
}

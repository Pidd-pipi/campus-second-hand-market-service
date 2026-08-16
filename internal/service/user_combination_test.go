package service

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/lp/campus-market/internal/constants"
	"github.com/lp/campus-market/internal/dto"
	"github.com/lp/campus-market/internal/model"
	"github.com/lp/campus-market/internal/util"
)

type nilUserRepo struct{}

func (m *nilUserRepo) Create(ctx context.Context, u *model.User) error { return nil }
func (m *nilUserRepo) FindByPhone(ctx context.Context, phone string) (*model.User, error) { return nil, nil }
func (m *nilUserRepo) FindByID(ctx context.Context, id uint) (*model.User, error) { return nil, nil }
func (m *nilUserRepo) UpdateProfile(ctx context.Context, id uint, nickname, avatar, campus string) error { return nil }
func (m *nilUserRepo) AddCredit(ctx context.Context, id uint, delta int) error { return nil }
func (m *nilUserRepo) Count(ctx context.Context) (int64, error) { return 0, nil }

func newNilUserService() *UserService {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	return NewUserService(&nilUserRepo{}, "test-secret", 72, logger)
}

func TestGetProfileNilUserReturnsError(t *testing.T) {
	svc := newNilUserService()
	user, err := svc.GetProfile(context.Background(), 1)
	if err == nil {
		t.Fatalf("expected error for nil user, got user=%+v", user)
	}
}

func TestLoginNilUserUnauthorized(t *testing.T) {
	svc := newNilUserService()
	_, err := svc.Login(context.Background(), &dto.LoginRequest{Phone: "13800009999", Password: "123456"})
	if err == nil {
		t.Fatal("expected unauthorized error")
	}
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.Code != constants.CodeUnauthorized {
		t.Fatalf("expected CodeUnauthorized, got %v", err)
	}
}

func TestRegisterDuplicateConflict(t *testing.T) {
	svc := newTestUserService(newFakeUserRepo())
	if _, err := svc.Register(context.Background(), &dto.RegisterRequest{Phone: "13800001111", Password: "123456", Nickname: "测试", Campus: "东校区"}); err != nil {
		t.Fatalf("first register failed: %v", err)
	}
	_, err := svc.Register(context.Background(), &dto.RegisterRequest{Phone: "13800001111", Password: "123456", Nickname: "测试", Campus: "东校区"})
	if err == nil {
		t.Fatal("duplicate register should fail")
	}
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.Code != constants.CodeConflict {
		t.Fatalf("expected CodeConflict, got %v", err)
	}
}

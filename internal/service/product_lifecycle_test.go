package service

import (
	"context"
	"log/slog"
	"testing"

	"github.com/lp/campus-market/internal/constants"
	"github.com/lp/campus-market/internal/dto"
	"github.com/lp/campus-market/internal/util"
)

func TestProductLifecycleValidation(t *testing.T) {
	svc := NewProductService(newFakeProductRepo(), slog.Default())
	_, err := svc.Create(context.Background(), 1, &dto.CreateProductRequest{Title: "非法分类", Price: 10, Category: "sports", Condition: "全新", Campus: "东校区", TradeLocation: "东门"})
	if err == nil {
		t.Fatal("invalid category should be rejected")
	}
	if util.CategoryText(constants.ProductCategoryBooks) != "书籍" {
		t.Fatalf("ProductCategoryText(books) = %s, want 书籍", util.CategoryText(constants.ProductCategoryBooks))
	}
	if util.ProductStatusText(constants.ProductStatusSold) != "已售出" {
		t.Fatalf("ProductStatusText(sold) = %s, want 已售出", util.ProductStatusText(constants.ProductStatusSold))
	}
}

func TestProductLifecycleOwnershipAndSold(t *testing.T) {
	repo := newFakeProductRepo()
	svc := NewProductService(repo, slog.Default())
	created, _ := svc.Create(context.Background(), 1, &dto.CreateProductRequest{Title: "我的书", Price: 10, Category: constants.ProductCategoryBooks, Condition: "全新", Campus: "东校区", TradeLocation: "东门"})
	if _, err := svc.Remove(context.Background(), 99, created.ID); err == nil {
		t.Fatal("non-owner remove should be forbidden")
	}
	if err := svc.MarkSold(context.Background(), created.ID); err != nil {
		t.Fatalf("mark sold failed: %v", err)
	}
	got, _ := svc.Get(context.Background(), created.ID)
	if got.Status != constants.ProductStatusSold {
		t.Fatalf("status = %s, want sold", got.Status)
	}
}

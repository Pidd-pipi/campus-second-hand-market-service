package service

import (
	"context"
	"log/slog"
	"testing"

	"github.com/lp/campus-market/internal/constants"
	"github.com/lp/campus-market/internal/dto"
	"github.com/lp/campus-market/internal/util"
)

func TestCreditAndProductStatusBoundary(t *testing.T) {
	if got := util.CreditDelta(constants.ReviewRatingGood); got != 5 {
		t.Fatalf("CreditDelta(good) = %d, want 5", got)
	}
	if got := util.ClampCredit(-10); got != 0 {
		t.Fatalf("ClampCredit(-10) = %d, want 0", got)
	}
	if got := util.ClampCredit(350); got != 300 {
		t.Fatalf("ClampCredit(350) = %d, want 300", got)
	}
	if got := util.ProductStatusText(constants.ProductStatusSold); got != "已售出" {
		t.Fatalf("ProductStatusText(sold) = %s, want 已售出", got)
	}
}

func TestMarkSoldStatus(t *testing.T) {
	svc := NewProductService(newFakeProductRepo(), slog.Default())
	p, _ := svc.Create(context.Background(), 1, &dto.CreateProductRequest{Title: "书", Price: 10, Category: constants.ProductCategoryBooks, Condition: "全新", Campus: "东校区", TradeLocation: "东门"})
	if err := svc.MarkSold(context.Background(), p.ID); err != nil {
		t.Fatalf("mark sold failed: %v", err)
	}
	got, _ := svc.Get(context.Background(), p.ID)
	if got.Status != constants.ProductStatusSold {
		t.Fatalf("status = %s, want sold", got.Status)
	}
}

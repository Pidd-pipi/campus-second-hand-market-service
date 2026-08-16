package util

import (
	"testing"

	"github.com/lp/campus-market/internal/constants"
)

func TestTradeStatusBoundary(t *testing.T) {
	if !constants.IsTradeStatus(constants.TradeStatusConfirmed) {
		t.Fatal("IsTradeStatus(confirmed) should be true")
	}
	if got := TradeStatusText(constants.TradeStatusPending); got != "待确认" {
		t.Fatalf("TradeStatusText(pending) = %s, want 待确认", got)
	}
	if got := constants.ReviewRatingText(constants.ReviewRatingGood); got != "好评" {
		t.Fatalf("ReviewRatingText(good) = %s, want 好评", got)
	}
}

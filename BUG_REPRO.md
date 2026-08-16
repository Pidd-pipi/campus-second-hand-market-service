# BUG 复现说明（campus-market__004）

## Bug 是什么
信用分计算与商品状态文案/状态变更联合失效：好评扣分、信用分上下限颠倒、已售出商品标记为已下架且文案显示在售。

## 如何触发
```bash
go test ./internal/service -run 'TestCreditAndProductStatusBoundary|TestMarkSoldStatus'
```

## 错误信息
```
--- FAIL: TestCreditAndProductStatusBoundary
    credit_product_diagnosis_test.go:15: CreditDelta(good) = -5, want 5
--- FAIL: TestMarkSoldStatus
    credit_product_diagnosis_test.go:36: status = removed, want sold
```

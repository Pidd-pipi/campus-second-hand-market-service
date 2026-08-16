# BUG 复现说明（campus-market__001）

## Bug 是什么
商品生命周期多处联合失效：非法分类可发布，非卖家可下架商品，标记已售出被改成已下架，分类/状态文案错误。

## 如何触发
```bash
go test ./internal/service -run 'TestProductLifecycleValidation|TestProductLifecycleOwnershipAndSold' -count=1
```

## 错误信息
```
--- FAIL: TestProductLifecycleValidation
    product_lifecycle_test.go:17: invalid category should be rejected
--- FAIL: TestProductLifecycleOwnershipAndSold
    product_lifecycle_test.go:32: non-owner remove should be forbidden
```

# BUG 复现说明（campus-market__005）

## Bug 是什么
交易状态机与状态/评价文案多处错位：已确认状态被判定为非法，待确认文案显示为已完成，好评文案显示为差评。

## 如何触发
```bash
go test ./internal/util -run 'TestTradeStatusBoundary'
```

## 错误信息
```
--- FAIL: TestTradeStatusBoundary
    trade_status_diagnosis_test.go:11: IsTradeStatus(confirmed) should be true
```

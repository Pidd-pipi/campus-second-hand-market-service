# BUG 复现说明（campus-market__002）

## Bug 是什么
用户查询/登录/注册的 nil 与错误码联合失效：未找到用户返回 (nil,nil)，登录 nil 用户会 panic，重复注册错误码错误。

## 如何触发
```bash
go test ./internal/service -run 'TestGetProfileNilUserReturnsError|TestLoginNilUserUnauthorized|TestRegisterDuplicateConflict' -count=1
```

## 错误信息
```
--- FAIL: TestGetProfileNilUserReturnsError
    user_combination_test.go:34: expected error for nil user, got user=<nil>
--- FAIL: TestLoginNilUserUnauthorized
    user_combination_test.go:45: expected CodeUnauthorized
--- FAIL: TestRegisterDuplicateConflict
    user_combination_test.go:61: expected CodeConflict
```

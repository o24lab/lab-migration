package migration

import "pkg/errorx"

var (
	ErrUserAlreadyExisted = errorx.Define(1, "user already existed")
	ErrInvalidIdType      = errorx.Define(2, "invalid id type")
	ErrKycNotExisted      = errorx.Define(3, "kyc not existed")
)

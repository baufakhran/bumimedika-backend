package errorz

import (
	"errors"
	"fmt"
)

// Errorz adalah custom error type dengan error code dan reason.
type Errorz struct {
	Raw    error
	Code   ErrorCode
	Reason string
}

// Error mengembalikan representasi string dari Errorz.
func (e *Errorz) Error() string {
	if e.Raw != nil {
		return fmt.Sprintf("[%d] %s: %s", e.Code, e.Reason, e.Raw.Error())
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Reason)
}

// Unwrap mendukung penggunaan errors.Unwrap untuk nested error.
func (e *Errorz) Unwrap() error {
	return e.Raw
}

// Is mendukung errors.Is untuk membandingkan berdasarkan ErrorCode.
func (e *Errorz) Is(target error) bool {
	t, ok := target.(*Errorz)
	if !ok {
		return false
	}
	return e.Code == t.Code
}
func (e *Errorz) IsCode(target ErrorCode) bool {
	return e.Code == target
}

// New membuat Errorz baru dari ErrorCode, reason, dan error.
func New(code ErrorCode, reason string, err error) *Errorz {
	return &Errorz{
		Code:   code,
		Reason: reason,
		Raw:    err,
	}
}

// NewSimple membuat Errorz tanpa error asli (misalnya hanya untuk kode dan reason).
func NewByCode(code ErrorCode) *Errorz {
	return &Errorz{
		Code:   code,
		Reason: code.ToString(),
	}
}

// NewSimple membuat Errorz tanpa error asli (misalnya hanya untuk kode dan reason).
func NewSimpleWithRaw(code ErrorCode, err error) *Errorz {
	return &Errorz{
		Raw:    err,
		Code:   code,
		Reason: code.ToString(),
	}
}

// AsErrorz mencoba casting error menjadi *Errorz.
func AsErrorz(err error) (*Errorz, bool) {
	var e *Errorz
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}

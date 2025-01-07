// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/storage/storage.proto

package storage

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on UploadChunksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadChunksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadChunksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadChunksRequestMultiError, or nil if none found.
func (m *UploadChunksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadChunksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChunkId

	// no validation rules for Data

	if len(errors) > 0 {
		return UploadChunksRequestMultiError(errors)
	}

	return nil
}

// UploadChunksRequestMultiError is an error wrapping multiple validation
// errors returned by UploadChunksRequest.ValidateAll() if the designated
// constraints aren't met.
type UploadChunksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadChunksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadChunksRequestMultiError) AllErrors() []error { return m }

// UploadChunksRequestValidationError is the validation error returned by
// UploadChunksRequest.Validate if the designated constraints aren't met.
type UploadChunksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadChunksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadChunksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadChunksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadChunksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadChunksRequestValidationError) ErrorName() string {
	return "UploadChunksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadChunksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadChunksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadChunksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadChunksRequestValidationError{}

// Validate checks the field values on UploadChunksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadChunksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadChunksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadChunksResponseMultiError, or nil if none found.
func (m *UploadChunksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadChunksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if len(errors) > 0 {
		return UploadChunksResponseMultiError(errors)
	}

	return nil
}

// UploadChunksResponseMultiError is an error wrapping multiple validation
// errors returned by UploadChunksResponse.ValidateAll() if the designated
// constraints aren't met.
type UploadChunksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadChunksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadChunksResponseMultiError) AllErrors() []error { return m }

// UploadChunksResponseValidationError is the validation error returned by
// UploadChunksResponse.Validate if the designated constraints aren't met.
type UploadChunksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadChunksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadChunksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadChunksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadChunksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadChunksResponseValidationError) ErrorName() string {
	return "UploadChunksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UploadChunksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadChunksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadChunksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadChunksResponseValidationError{}

// Validate checks the field values on GetChunksRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetChunksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetChunksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetChunksRequestMultiError, or nil if none found.
func (m *GetChunksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetChunksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetChunksRequestMultiError(errors)
	}

	return nil
}

// GetChunksRequestMultiError is an error wrapping multiple validation errors
// returned by GetChunksRequest.ValidateAll() if the designated constraints
// aren't met.
type GetChunksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetChunksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetChunksRequestMultiError) AllErrors() []error { return m }

// GetChunksRequestValidationError is the validation error returned by
// GetChunksRequest.Validate if the designated constraints aren't met.
type GetChunksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetChunksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetChunksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetChunksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetChunksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetChunksRequestValidationError) ErrorName() string { return "GetChunksRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetChunksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetChunksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetChunksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetChunksRequestValidationError{}

// Validate checks the field values on GetChunksResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetChunksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetChunksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetChunksResponseMultiError, or nil if none found.
func (m *GetChunksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetChunksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChunkId

	// no validation rules for Data

	if len(errors) > 0 {
		return GetChunksResponseMultiError(errors)
	}

	return nil
}

// GetChunksResponseMultiError is an error wrapping multiple validation errors
// returned by GetChunksResponse.ValidateAll() if the designated constraints
// aren't met.
type GetChunksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetChunksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetChunksResponseMultiError) AllErrors() []error { return m }

// GetChunksResponseValidationError is the validation error returned by
// GetChunksResponse.Validate if the designated constraints aren't met.
type GetChunksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetChunksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetChunksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetChunksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetChunksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetChunksResponseValidationError) ErrorName() string {
	return "GetChunksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetChunksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetChunksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetChunksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetChunksResponseValidationError{}
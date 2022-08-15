// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/role.proto

package v1

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

// Validate checks the field values on CreateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoleRequestMultiError, or nil if none found.
func (m *CreateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return CreateRoleRequestMultiError(errors)
	}

	return nil
}

// CreateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoleRequestMultiError) AllErrors() []error { return m }

// CreateRoleRequestValidationError is the validation error returned by
// CreateRoleRequest.Validate if the designated constraints aren't met.
type CreateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleRequestValidationError) ErrorName() string {
	return "CreateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleRequestValidationError{}

// Validate checks the field values on CreateRoleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateRoleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoleReplyMultiError, or nil if none found.
func (m *CreateRoleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return CreateRoleReplyMultiError(errors)
	}

	return nil
}

// CreateRoleReplyMultiError is an error wrapping multiple validation errors
// returned by CreateRoleReply.ValidateAll() if the designated constraints
// aren't met.
type CreateRoleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoleReplyMultiError) AllErrors() []error { return m }

// CreateRoleReplyValidationError is the validation error returned by
// CreateRoleReply.Validate if the designated constraints aren't met.
type CreateRoleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleReplyValidationError) ErrorName() string { return "CreateRoleReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateRoleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleReplyValidationError{}

// Validate checks the field values on UpdateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleRequestMultiError, or nil if none found.
func (m *UpdateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateRoleRequestMultiError(errors)
	}

	return nil
}

// UpdateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleRequestMultiError) AllErrors() []error { return m }

// UpdateRoleRequestValidationError is the validation error returned by
// UpdateRoleRequest.Validate if the designated constraints aren't met.
type UpdateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleRequestValidationError) ErrorName() string {
	return "UpdateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleRequestValidationError{}

// Validate checks the field values on UpdateRoleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleReplyMultiError, or nil if none found.
func (m *UpdateRoleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateRoleReplyMultiError(errors)
	}

	return nil
}

// UpdateRoleReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateRoleReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateRoleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleReplyMultiError) AllErrors() []error { return m }

// UpdateRoleReplyValidationError is the validation error returned by
// UpdateRoleReply.Validate if the designated constraints aren't met.
type UpdateRoleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleReplyValidationError) ErrorName() string { return "UpdateRoleReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRoleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleReplyValidationError{}

// Validate checks the field values on DeleteRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoleRequestMultiError, or nil if none found.
func (m *DeleteRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRoleRequestMultiError(errors)
	}

	return nil
}

// DeleteRoleRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoleRequestMultiError) AllErrors() []error { return m }

// DeleteRoleRequestValidationError is the validation error returned by
// DeleteRoleRequest.Validate if the designated constraints aren't met.
type DeleteRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleRequestValidationError) ErrorName() string {
	return "DeleteRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleRequestValidationError{}

// Validate checks the field values on DeleteRoleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoleReplyMultiError, or nil if none found.
func (m *DeleteRoleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRoleReplyMultiError(errors)
	}

	return nil
}

// DeleteRoleReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteRoleReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteRoleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoleReplyMultiError) AllErrors() []error { return m }

// DeleteRoleReplyValidationError is the validation error returned by
// DeleteRoleReply.Validate if the designated constraints aren't met.
type DeleteRoleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleReplyValidationError) ErrorName() string { return "DeleteRoleReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRoleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleReplyValidationError{}

// Validate checks the field values on GetRoleRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRoleRequestMultiError,
// or nil if none found.
func (m *GetRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRoleRequestMultiError(errors)
	}

	return nil
}

// GetRoleRequestMultiError is an error wrapping multiple validation errors
// returned by GetRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleRequestMultiError) AllErrors() []error { return m }

// GetRoleRequestValidationError is the validation error returned by
// GetRoleRequest.Validate if the designated constraints aren't met.
type GetRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleRequestValidationError) ErrorName() string { return "GetRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleRequestValidationError{}

// Validate checks the field values on GetRoleReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRoleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRoleReplyMultiError, or
// nil if none found.
func (m *GetRoleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRoleReplyMultiError(errors)
	}

	return nil
}

// GetRoleReplyMultiError is an error wrapping multiple validation errors
// returned by GetRoleReply.ValidateAll() if the designated constraints aren't met.
type GetRoleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleReplyMultiError) AllErrors() []error { return m }

// GetRoleReplyValidationError is the validation error returned by
// GetRoleReply.Validate if the designated constraints aren't met.
type GetRoleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleReplyValidationError) ErrorName() string { return "GetRoleReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleReplyValidationError{}

// Validate checks the field values on ListRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoleRequestMultiError, or nil if none found.
func (m *ListRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRoleRequestMultiError(errors)
	}

	return nil
}

// ListRoleRequestMultiError is an error wrapping multiple validation errors
// returned by ListRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type ListRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoleRequestMultiError) AllErrors() []error { return m }

// ListRoleRequestValidationError is the validation error returned by
// ListRoleRequest.Validate if the designated constraints aren't met.
type ListRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleRequestValidationError) ErrorName() string { return "ListRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleRequestValidationError{}

// Validate checks the field values on ListRoleReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRoleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRoleReplyMultiError, or
// nil if none found.
func (m *ListRoleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRoleReplyMultiError(errors)
	}

	return nil
}

// ListRoleReplyMultiError is an error wrapping multiple validation errors
// returned by ListRoleReply.ValidateAll() if the designated constraints
// aren't met.
type ListRoleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoleReplyMultiError) AllErrors() []error { return m }

// ListRoleReplyValidationError is the validation error returned by
// ListRoleReply.Validate if the designated constraints aren't met.
type ListRoleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleReplyValidationError) ErrorName() string { return "ListRoleReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListRoleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleReplyValidationError{}
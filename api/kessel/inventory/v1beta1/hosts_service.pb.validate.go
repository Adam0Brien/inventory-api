// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/hosts_service.proto

package v1beta1

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

// Validate checks the field values on CreateRHELHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRHELHostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRHELHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRHELHostRequestMultiError, or nil if none found.
func (m *CreateRHELHostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRHELHostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRHELHostRequestValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRHELHostRequestValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRHELHostRequestValidationError{
				field:  "Host",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRHELHostRequestMultiError(errors)
	}

	return nil
}

// CreateRHELHostRequestMultiError is an error wrapping multiple validation
// errors returned by CreateRHELHostRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateRHELHostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRHELHostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRHELHostRequestMultiError) AllErrors() []error { return m }

// CreateRHELHostRequestValidationError is the validation error returned by
// CreateRHELHostRequest.Validate if the designated constraints aren't met.
type CreateRHELHostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRHELHostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRHELHostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRHELHostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRHELHostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRHELHostRequestValidationError) ErrorName() string {
	return "CreateRHELHostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRHELHostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRHELHostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRHELHostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRHELHostRequestValidationError{}

// Validate checks the field values on CreateRHELHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRHELHostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRHELHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRHELHostResponseMultiError, or nil if none found.
func (m *CreateRHELHostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRHELHostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRHELHostResponseValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRHELHostResponseValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRHELHostResponseValidationError{
				field:  "Host",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRHELHostResponseMultiError(errors)
	}

	return nil
}

// CreateRHELHostResponseMultiError is an error wrapping multiple validation
// errors returned by CreateRHELHostResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateRHELHostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRHELHostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRHELHostResponseMultiError) AllErrors() []error { return m }

// CreateRHELHostResponseValidationError is the validation error returned by
// CreateRHELHostResponse.Validate if the designated constraints aren't met.
type CreateRHELHostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRHELHostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRHELHostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRHELHostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRHELHostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRHELHostResponseValidationError) ErrorName() string {
	return "CreateRHELHostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRHELHostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRHELHostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRHELHostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRHELHostResponseValidationError{}

// Validate checks the field values on UpdateRHELHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRHELHostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRHELHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRHELHostRequestMultiError, or nil if none found.
func (m *UpdateRHELHostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRHELHostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if all {
		switch v := interface{}(m.GetHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRHELHostRequestValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRHELHostRequestValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRHELHostRequestValidationError{
				field:  "Host",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateRHELHostRequestMultiError(errors)
	}

	return nil
}

// UpdateRHELHostRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateRHELHostRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateRHELHostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRHELHostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRHELHostRequestMultiError) AllErrors() []error { return m }

// UpdateRHELHostRequestValidationError is the validation error returned by
// UpdateRHELHostRequest.Validate if the designated constraints aren't met.
type UpdateRHELHostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRHELHostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRHELHostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRHELHostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRHELHostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRHELHostRequestValidationError) ErrorName() string {
	return "UpdateRHELHostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRHELHostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRHELHostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRHELHostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRHELHostRequestValidationError{}

// Validate checks the field values on UpdateRHELHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRHELHostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRHELHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRHELHostResponseMultiError, or nil if none found.
func (m *UpdateRHELHostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRHELHostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateRHELHostResponseMultiError(errors)
	}

	return nil
}

// UpdateRHELHostResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateRHELHostResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateRHELHostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRHELHostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRHELHostResponseMultiError) AllErrors() []error { return m }

// UpdateRHELHostResponseValidationError is the validation error returned by
// UpdateRHELHostResponse.Validate if the designated constraints aren't met.
type UpdateRHELHostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRHELHostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRHELHostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRHELHostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRHELHostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRHELHostResponseValidationError) ErrorName() string {
	return "UpdateRHELHostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRHELHostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRHELHostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRHELHostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRHELHostResponseValidationError{}

// Validate checks the field values on DeleteRHELHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRHELHostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRHELHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRHELHostRequestMultiError, or nil if none found.
func (m *DeleteRHELHostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRHELHostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if len(errors) > 0 {
		return DeleteRHELHostRequestMultiError(errors)
	}

	return nil
}

// DeleteRHELHostRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteRHELHostRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteRHELHostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRHELHostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRHELHostRequestMultiError) AllErrors() []error { return m }

// DeleteRHELHostRequestValidationError is the validation error returned by
// DeleteRHELHostRequest.Validate if the designated constraints aren't met.
type DeleteRHELHostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRHELHostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRHELHostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRHELHostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRHELHostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRHELHostRequestValidationError) ErrorName() string {
	return "DeleteRHELHostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRHELHostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRHELHostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRHELHostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRHELHostRequestValidationError{}

// Validate checks the field values on DeleteRHELHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRHELHostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRHELHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRHELHostResponseMultiError, or nil if none found.
func (m *DeleteRHELHostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRHELHostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRHELHostResponseMultiError(errors)
	}

	return nil
}

// DeleteRHELHostResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteRHELHostResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteRHELHostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRHELHostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRHELHostResponseMultiError) AllErrors() []error { return m }

// DeleteRHELHostResponseValidationError is the validation error returned by
// DeleteRHELHostResponse.Validate if the designated constraints aren't met.
type DeleteRHELHostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRHELHostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRHELHostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRHELHostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRHELHostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRHELHostResponseValidationError) ErrorName() string {
	return "DeleteRHELHostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRHELHostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRHELHostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRHELHostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRHELHostResponseValidationError{}
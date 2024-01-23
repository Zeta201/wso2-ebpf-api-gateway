// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/v3/ratelimit_strategy.proto

package typev3

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

// Validate checks the field values on RateLimitStrategy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RateLimitStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitStrategyMultiError, or nil if none found.
func (m *RateLimitStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Strategy.(type) {

	case *RateLimitStrategy_BlanketRule_:

		if _, ok := RateLimitStrategy_BlanketRule_name[int32(m.GetBlanketRule())]; !ok {
			err := RateLimitStrategyValidationError{
				field:  "BlanketRule",
				reason: "value must be one of the defined enum values",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *RateLimitStrategy_RequestsPerTimeUnit_:

		if all {
			switch v := interface{}(m.GetRequestsPerTimeUnit()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitStrategyValidationError{
						field:  "RequestsPerTimeUnit",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitStrategyValidationError{
						field:  "RequestsPerTimeUnit",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRequestsPerTimeUnit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitStrategyValidationError{
					field:  "RequestsPerTimeUnit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RateLimitStrategy_TokenBucket:

		if all {
			switch v := interface{}(m.GetTokenBucket()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitStrategyValidationError{
						field:  "TokenBucket",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitStrategyValidationError{
						field:  "TokenBucket",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTokenBucket()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitStrategyValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := RateLimitStrategyValidationError{
			field:  "Strategy",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return RateLimitStrategyMultiError(errors)
	}
	return nil
}

// RateLimitStrategyMultiError is an error wrapping multiple validation errors
// returned by RateLimitStrategy.ValidateAll() if the designated constraints
// aren't met.
type RateLimitStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitStrategyMultiError) AllErrors() []error { return m }

// RateLimitStrategyValidationError is the validation error returned by
// RateLimitStrategy.Validate if the designated constraints aren't met.
type RateLimitStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitStrategyValidationError) ErrorName() string {
	return "RateLimitStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitStrategyValidationError{}

// Validate checks the field values on RateLimitStrategy_RequestsPerTimeUnit
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *RateLimitStrategy_RequestsPerTimeUnit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitStrategy_RequestsPerTimeUnit
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RateLimitStrategy_RequestsPerTimeUnitMultiError, or nil if none found.
func (m *RateLimitStrategy_RequestsPerTimeUnit) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitStrategy_RequestsPerTimeUnit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestsPerTimeUnit

	if _, ok := RateLimitUnit_name[int32(m.GetTimeUnit())]; !ok {
		err := RateLimitStrategy_RequestsPerTimeUnitValidationError{
			field:  "TimeUnit",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RateLimitStrategy_RequestsPerTimeUnitMultiError(errors)
	}
	return nil
}

// RateLimitStrategy_RequestsPerTimeUnitMultiError is an error wrapping
// multiple validation errors returned by
// RateLimitStrategy_RequestsPerTimeUnit.ValidateAll() if the designated
// constraints aren't met.
type RateLimitStrategy_RequestsPerTimeUnitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitStrategy_RequestsPerTimeUnitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitStrategy_RequestsPerTimeUnitMultiError) AllErrors() []error { return m }

// RateLimitStrategy_RequestsPerTimeUnitValidationError is the validation error
// returned by RateLimitStrategy_RequestsPerTimeUnit.Validate if the
// designated constraints aren't met.
type RateLimitStrategy_RequestsPerTimeUnitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitStrategy_RequestsPerTimeUnitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitStrategy_RequestsPerTimeUnitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitStrategy_RequestsPerTimeUnitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitStrategy_RequestsPerTimeUnitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitStrategy_RequestsPerTimeUnitValidationError) ErrorName() string {
	return "RateLimitStrategy_RequestsPerTimeUnitValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitStrategy_RequestsPerTimeUnitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitStrategy_RequestsPerTimeUnit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitStrategy_RequestsPerTimeUnitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitStrategy_RequestsPerTimeUnitValidationError{}

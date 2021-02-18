// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_http_ratelimit_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _rate_limit_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RateLimit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetDomain()) < 1 {
		return RateLimitValidationError{
			field:  "Domain",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetStage() > 10 {
		return RateLimitValidationError{
			field:  "Stage",
			reason: "value must be less than or equal to 10",
		}
	}

	if _, ok := _RateLimit_RequestType_InLookup[m.GetRequestType()]; !ok {
		return RateLimitValidationError{
			field:  "RequestType",
			reason: "value must be in list [internal external both ]",
		}
	}

	if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailureModeDeny

	// no validation rules for RateLimitedAsResourceExhausted

	if m.GetRateLimitService() == nil {
		return RateLimitValidationError{
			field:  "RateLimitService",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRateLimitService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "RateLimitService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := RateLimit_XRateLimitHeadersRFCVersion_name[int32(m.GetEnableXRatelimitHeaders())]; !ok {
		return RateLimitValidationError{
			field:  "EnableXRatelimitHeaders",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// RateLimitValidationError is the validation error returned by
// RateLimit.Validate if the designated constraints aren't met.
type RateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitValidationError) ErrorName() string { return "RateLimitValidationError" }

// Error satisfies the builtin error interface
func (e RateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitValidationError{}

var _RateLimit_RequestType_InLookup = map[string]struct{}{
	"internal": {},
	"external": {},
	"both":     {},
	"":         {},
}

// Validate checks the field values on RateLimitPerRoute with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RateLimitPerRoute) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := RateLimitPerRoute_VhRateLimitsOptions_name[int32(m.GetVhRateLimits())]; !ok {
		return RateLimitPerRouteValidationError{
			field:  "VhRateLimits",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// RateLimitPerRouteValidationError is the validation error returned by
// RateLimitPerRoute.Validate if the designated constraints aren't met.
type RateLimitPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitPerRouteValidationError) ErrorName() string {
	return "RateLimitPerRouteValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitPerRouteValidationError{}

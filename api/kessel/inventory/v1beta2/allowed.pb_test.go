package v1beta2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestAllowedEnum(t *testing.T) {
	tests := []struct {
		value    Allowed
		expected string
		number   protoreflect.EnumNumber
	}{
		{Allowed_ALLOWED_UNSPECIFIED, "ALLOWED_UNSPECIFIED", 0},
		{Allowed_ALLOWED_TRUE, "ALLOWED_TRUE", 1},
		{Allowed_ALLOWED_FALSE, "ALLOWED_FALSE", 2},
	}

	for _, tc := range tests {
		t.Run(tc.expected, func(t *testing.T) {
			// Test Enum() returns pointer to the correct value
			ptr := tc.value.Enum()
			assert.NotNil(t, ptr)
			assert.Equal(t, tc.value, *ptr)

			// Test String() returns correct enum name
			assert.Equal(t, tc.expected, tc.value.String())

			// Test Number() returns correct enum number
			assert.Equal(t, tc.number, tc.value.Number())
		})
	}
}

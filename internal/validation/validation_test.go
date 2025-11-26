package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected error
	}{
		{
			name:     "Normal1",
			value:    "Vasily",
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    "Petya",
			expected: nil,
		},
		{
			name:     "ErrEmpty",
			value:    "",
			expected: ErrValidationNameEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NameValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAgeValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected error
	}{
		{
			name:     "Normal1",
			value:    17,
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    25,
			expected: nil,
		},
		{
			name:     "Normal3",
			value:    80,
			expected: nil,
		},
		{
			name:     "ErrTooSmall",
			value:    10,
			expected: ErrValidationAgeTooSmall,
		},
		{
			name:     "ErrTooBig",
			value:    100,
			expected: ErrValidationAgeTooBig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AgeValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSexValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected error
	}{
		{
			name:     "Normal1",
			value:    "man",
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    "woman",
			expected: nil,
		},
		{
			name:     "ErrInvalidValue1",
			value:    "Womans",
			expected: ErrValidationSexInvalid,
		},
		{
			name:     "ErrInvalidValue2",
			value:    "MAAN",
			expected: ErrValidationSexInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SexValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCourseValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected error
	}{
		{
			name:     "Normal1",
			value:    1,
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    3,
			expected: nil,
		},
		{
			name:     "Normal3",
			value:    6,
			expected: nil,
		},
		{
			name:     "ErrTooSmall",
			value:    0,
			expected: ErrValidationCourseTooSmall,
		},
		{
			name:     "ErrTooBig",
			value:    10,
			expected: ErrValidationCourseTooBig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CourseValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

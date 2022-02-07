// Package envisage is a lightweight package that makes easier and safer to deal with environment variables.
package envisage

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IsThere returns true if the variable is present in the environment
func IsThere(key string) bool {
	_, ok := os.LookupEnv(key)

	return ok
}

// String returns the env var value as string
func String(key string, defaultValue string) string {
	if s, ok := os.LookupEnv(key); ok {
		return s
	}

	return defaultValue
}

// Int returns the env var value as int
func Int(key string, defaultValue int) int {
	if s, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(s); err == nil {
			return i
		}
	}

	return defaultValue
}

// I64 returns the env var value as int64
func I64(key string, defaultValue int64) int64 {
	if s, ok := os.LookupEnv(key); ok {
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return i
		}
	}

	return defaultValue
}

// Int64 returns the env var value as int64
// It's an idiomatic convenience alias for I64
func Int64(key string, defaultValue int64) int64 {
	return I64(key, defaultValue)
}

// Bool returns the env var value as boolean
func Bool(key string, defaultValue bool) bool {
	if s, ok := os.LookupEnv(key); ok {
		if b, err := strconv.ParseBool(s); err == nil {
			return b
		}
	}

	return defaultValue
}

// F64 returns the env var value as float64
func F64(key string, commaDecimalSeparator bool, defaultValue float64) float64 {
	if s, ok := os.LookupEnv(key); ok {
		if commaDecimalSeparator {
			s = strings.Replace(s, ",", ".", 1)
		}

		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f
		}
	}

	return defaultValue
}

// Float64 returns the env var value as float64
// It's an idiomatic convenience alias for F64
func Float64(key string, commaDecimalSeparator bool, defaultValue float64) float64 {
	return F64(key, commaDecimalSeparator, defaultValue)
}

// StringS returns the env var value as []string
func StringS(key, separator string, defaultValue []string) []string {
	if s, ok := os.LookupEnv(key); ok {
		return strings.Split(s, separator)
	}

	return defaultValue
}

// IntS returns the env var value as []int
func IntS(key, listItemSeparator string, defaultValue []int) ([]int, error) {
	if s, ok := os.LookupEnv(key); ok {
		ss := strings.Split(s, listItemSeparator)

		if len(ss) > 0 {
			var a []int

			for _, x := range ss {
				if i, err := strconv.Atoi(x); err != nil {
					return defaultValue, err
				} else {
					a = append(a, i)
				}
			}

			return a, nil
		}
	}

	return defaultValue, nil
}

// IntSlice returns the env var value as []int
// It's an idiomatic convenience alias for IntS
func IntSlice(key, listItemSeparator string, defaultValue []int) ([]int, error) {
	return IntS(key, listItemSeparator, defaultValue)
}

// F64S returns the env var value as []float64
func F64S(key, listItemSeparator string, commaDecimalSeparator bool, defaultValue []float64) ([]float64, error) {
	if s, ok := os.LookupEnv(key); ok {
		ss := strings.Split(s, listItemSeparator)

		if len(ss) > 0 {
			var a []float64

			for _, x := range ss {
				if commaDecimalSeparator {
					x = strings.Replace(x, ",", ".", 1)
				}

				if f, err := strconv.ParseFloat(x, 64); err != nil {
					return a, err
				} else {
					a = append(a, f)
				}
			}

			return a, nil
		}
	}

	return defaultValue, nil
}

// Float64Slice returns the env var value as []float64
// It's an idiomatic convenience alias for F64S
func Float64Slice(key, listItemSeparator string, commaDecimalSeparator bool, defaultValue []float64) ([]float64, error) {
	return F64S(key, listItemSeparator, commaDecimalSeparator, defaultValue)
}

// SetString sets the value of the environment variable named by the key.
func SetString(key, value string) error {
	return os.Setenv(key, value)
}

// SetInt sets the value of the environment variable named by the key.
func SetInt(key string, value int) error {
	return os.Setenv(key, strconv.Itoa(value))
}

// SetI64 sets the value of the environment variable named by the key.
func SetI64(key string, value int64) error {
	return os.Setenv(key, strconv.FormatInt(value, 10))
}

// SetInt64 sets the value of the environment variable named by the key.
// It's an idiomatic convenience alias for SetI64
func SetInt64(key string, value int64) error {
	return SetI64(key, value)
}

// SetF64 sets the value of the environment variable named by the key.
func SetF64(key string, value float64) error {
	// For implementation details please refer to https://stackoverflow.com/questions/19101419/formatfloat-convert-float-number-to-string/19101700#19101700
	return os.Setenv(key, strconv.FormatFloat(value, 'f', -1, 64))
}

// SetFloat64 sets the value of the environment variable named by the key.
// It's an idiomatic convenience alias for SetF64
func SetFloat64(key string, value float64) error {
	return SetF64(key, value)
}

// SetBool sets the value of the environment variable named by the key.
func SetBool(key string, value bool) error {
	return os.Setenv(key, fmt.Sprintf("%t", value))
}

// Check Test environment variables according given directives.
// defaultValue returned if the value is not present/set in the environment.
// twoWay updates the environment with the defaultValue, in case of the environment variable is not present/set.
// canBeEmpty forces an error if the variable has empty value.
func Check(key, defaultValue string, twoWay, canBeEmpty bool) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		if defaultValue != "" {
			s = defaultValue
		}

		if twoWay {
			if err := SetString(key, s); err != nil {
				return err
			}
		}
	}

	if s == "" && !canBeEmpty {
		return fmt.Errorf("environment variable %s can't be empty", key)
	}

	return nil
}

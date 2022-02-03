// Package envisage is a lightweight package that makes easier and safer to deal with environment variables.
package envisage

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSetString(t *testing.T) {
	type testCase struct {
		title,
		key,
		value string
	}

	tests := []testCase{
		{
			title: "simple test",
			key:   "SIMPLE_TEST",
			value: "SIMPLE_VALUE",
		},
		{
			title: "string with spaces",
			key:   "WITH_SPACE",
			value: "string with spaces",
		},
		{
			title: "empty string",
			key:   "EMPTY_STRING",
			value: "",
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if err := SetString(x.key, x.value); err != nil {
				t.Error(err)
			}

			if got := os.Getenv(x.key); got != x.value {
				t.Errorf("failed. expecting %s, got %s", x.value, got)
			}
		})
	}
}

func TestSetInt(t *testing.T) {
	type testCase struct {
		title,
		key string
		value int
	}

	tests := []testCase{
		{
			title: "simple test",
			key:   "SIMPLE_TEST",
			value: 0,
		},
		{
			title: "big value",
			key:   "BIG_VALUE",
			value: 545545464564156,
		},
		{
			title: "negative value",
			key:   "NEGATIVE_VALUE",
			value: -987654,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if err := SetInt(x.key, x.value); err != nil {
				t.Error(err)
			}

			got := os.Getenv(x.key)

			if i, _ := strconv.Atoi(got); i != x.value {
				t.Errorf("failed. expecting %d, got %d", x.value, i)
			}
		})
	}
}

func TestSetI64(t *testing.T) {
	type testCase struct {
		title,
		key string
		value int64
	}

	tests := []testCase{
		{
			title: "simple test",
			key:   "SIMPLE_TEST",
			value: 0,
		},
		{
			title: "big value",
			key:   "BIG_VALUE",
			value: 545545464564145454,
		},
		{
			title: "negative value",
			key:   "NEGATIVE_VALUE",
			value: -54554546456545454,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if err := SetI64(x.key, x.value); err != nil {
				t.Error(err)
			}

			got := os.Getenv(x.key)

			if i, _ := strconv.ParseInt(got, 10, 64); i != x.value {
				t.Errorf("failed. expecting %d, got %d", x.value, i)
			}
		})
	}
}

func TestSetF64(t *testing.T) {
	type testCase struct {
		title,
		key string
		value float64
	}

	tests := []testCase{
		{
			title: "zeroed value",
			key:   "SIMPLE_TEST",
			value: .0,
		},
		{
			title: "simple test",
			key:   "SIMPLE_TEST",
			value: 98765.4321,
		},
		{
			title: "big value",
			key:   "BIG_VALUE",
			value: 545545464564156545454.9797879879879879879,
		},
		{
			title: "negative value",
			key:   "NEGATIVE_VALUE",
			value: -545545464564156545454.6546546565465464454616161654,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if err := SetF64(x.key, x.value); err != nil {
				t.Error(err)
			}

			got := os.Getenv(x.key)

			if i, _ := strconv.ParseFloat(got, 64); i != x.value {
				t.Errorf("failed. expecting %f, got %f", x.value, i)
			}
		})
	}
}

func TestSetBool(t *testing.T) {
	type testCase struct {
		title,
		key string
		value bool
	}

	tests := []testCase{
		{
			title: "false value",
			key:   "FALSE_VALUE",
			value: false,
		},
		{
			title: "true value",
			key:   "TRUE_VALUE",
			value: true,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if err := SetBool(x.key, x.value); err != nil {
				t.Error(err)
			}

			got := os.Getenv(x.key)

			if i, _ := strconv.ParseBool(got); i != x.value {
				t.Errorf("failed. expecting %t, got %t", x.value, i)
			}
		})
	}
}

func TestIsThere(t *testing.T) {
	type testCase struct {
		title,
		key,
		value string
	}

	tests := []testCase{
		{
			title: "envisage simple test",
			key:   "ENVISAGE_SIMPLE_TEST",
			value: "AAAAAAAAAAAAAAA",
		},
		{
			title: "envisage test with empty string",
			key:   "ENVISAGE_TEST_WITH_EMPTY_STRING",
			value: "",
		},
		{
			title: "envisage lowercase test",
			key:   "envisage_lower_case_test",
			value: "fjdlhfkjdhfjkahkjdsfkljasdhsaklkd",
		},
		{
			title: "envisage mixed case test",
			key:   "envisage_MixedCaseTest",
			value: "blablablah",
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if err := SetString(x.key, x.value); err != nil {
				t.Error(err)
			}

			if !IsThere(x.key) {
				t.Errorf("failed. %s variable isnt' there", x.title)
			}
		})
	}
}

func TestString(t *testing.T) {
	type testCase struct {
		title,
		key,
		value,
		defaultValue string
	}

	tests := []testCase{
		{
			title: "envisage simple test",
			key:   "ENVISAGE_SIMPLE_TEST",
			value: "AAAAAAAAAAAAAAA",
		},
		{
			title: "envisage test with empty string",
			key:   "ENVISAGE_TEST_WITH_EMPTY_STRING",
			value: "",
		},
		{
			title: "envisage testing string with spaces",
			key:   "ENVISAGE_TESTING_STRING_WITH_SPACES",
			value: "word 1, word 2 and word 3",
		},
		{
			title: "envisage mixed case test",
			key:   "envisage_MixedCaseTest",
			value: "blablablah",
		},
		{
			title:        "envisage testing default value",
			key:          "ENVISAGE_TESTING_DEFAULT_VALUE",
			value:        "blablablah",
			defaultValue: "blablablah",
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if x.defaultValue == "" {
				if err := SetString(x.key, x.value); err != nil {
					t.Error(err)
				}
			}

			if got := String(x.key, x.defaultValue); got != x.value {
				t.Errorf("failed %s. expecting %s, got %s", x.title, x.value, got)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type testCase struct {
		title,
		key string
		value,
		defaultValue int
	}

	tests := []testCase{
		{
			title: "envisage simple test",
			key:   "TEST_INT_SIMPLE_TEST",
			value: 987654321,
		},
		{
			title: "zeroed value",
			key:   "TEST_INT_ZEROED_VALUE",
			value: 0,
		},
		{
			title: "negative value",
			key:   "TEST_INT_NEGATIVE_VALUE",
			value: -987654321,
		},
		{
			title:        "default value",
			key:          "TEST_INT_DEFAULT_VALUE",
			value:        123456789,
			defaultValue: 123456789,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if x.defaultValue == 0 {
				if err := SetInt(x.key, x.value); err != nil {
					t.Error(err)
				}
			}

			if got := Int(x.key, x.defaultValue); got != x.value {
				t.Errorf("failed %s. expecting %d, got %d", x.title, x.value, got)
			}
		})
	}
}

func TestI64(t *testing.T) {
	type testCase struct {
		title,
		key string
		value,
		defaultValue int64
	}

	tests := []testCase{
		{
			title: "envisage simple test",
			key:   "TEST_I64_SIMPLE_TEST",
			value: 9876543210123,
		},
		{
			title: "zeroed value",
			key:   "TEST_I64_ZEROED_VALUE",
			value: 0,
		},
		{
			title: "negative value",
			key:   "TEST_I64_NEGATIVE_VALUE",
			value: -9876543210123,
		},
		{
			title:        "default value",
			key:          "TEST_I64_DEFAULT_VALUE",
			value:        123456789,
			defaultValue: 123456789,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if x.defaultValue == 0 {
				if err := SetI64(x.key, x.value); err != nil {
					t.Error(err)
				}
			}

			if got := I64(x.key, x.defaultValue); got != x.value {
				t.Errorf("failed %s. expecting %d, got %d", x.title, x.value, got)
			}
		})
	}
}

func TestF64(t *testing.T) {
	type testCase struct {
		title,
		key string
		value,
		defaultValue float64
	}

	tests := []testCase{
		{
			title: "envisage simple test",
			key:   "TEST_I64_SIMPLE_TEST",
			value: 9876543210123.987654321,
		},
		{
			title: "zeroed value",
			key:   "TEST_I64_ZEROED_VALUE",
			value: .0,
		},
		{
			title: "negative value",
			key:   "TEST_I64_NEGATIVE_VALUE",
			value: -9876543210123.987654321,
		},
		{
			title:        "default value",
			key:          "TEST_I64_DEFAULT_VALUE",
			value:        12345.67890,
			defaultValue: 12345.67890,
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if x.defaultValue == 0 {
				if err := SetF64(x.key, x.value); err != nil {
					t.Error(err)
				}
			}

			if got := F64(x.key, false, x.defaultValue); got != x.value {
				t.Errorf("failed %s. expecting %f, got %f", x.title, x.value, got)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type testCase struct {
		title,
		key string
		value        bool
		defaultValue *bool
	}

	tests := []testCase{
		{
			title: "testing true",
			key:   "T_BOOL_TRUE",
			value: true,
		},
		{
			title: "testing false",
			key:   "T_BOOL_FALSE",
			value: false,
		},
		{
			title:        "testing default value",
			key:          "T_BOOL_DEFAULT_VALUE",
			value:        true,
			defaultValue: func() *bool { x := true; return &x }(),
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if x.defaultValue == nil {
				if err := SetBool(x.key, x.value); err != nil {
					t.Error(err)
				}
			}

			defaultValue := func() bool {
				if x.defaultValue != nil {
					return *x.defaultValue
				}
				return false
			}()

			if got := Bool(x.key, defaultValue); got != x.value {
				t.Errorf("failed %s. expecting %t, got %t", x.title, x.value, got)
			}
		})
	}
}

func TestStringS(t *testing.T) {
	type testCase struct {
		title,
		key string
		value        []string
		defaultValue *[]string
	}

	tests := []testCase{
		{
			title: "simple test",
			key:   "SIMPLE_TEST",
			value: []string{"a", "abcde", "01235", "", "-"},
		},
		{
			title: "empty slice",
			key:   "T_STRINGS_EMPTY_SLICE",
			value: []string{},
		},
		{
			title:        "default value",
			key:          "T_STRINGS_DEFAULT_VALUE",
			value:        []string{"a", "abcde", "01235", "", "-"},
			defaultValue: &[]string{"a", "abcde", "01235", "", "-"},
		},
	}

	for _, x := range tests {
		t.Run(x.title, func(t *testing.T) {
			if x.defaultValue == nil {
				v := ""

				if len(x.value) > 0 {
					v = strings.Join(x.value, ",")
				}

				if err := SetString(x.key, v); err != nil {
					t.Error(err)
				}
			}

			defaultValue := func() []string {
				if x.defaultValue != nil {
					return *x.defaultValue
				}
				return nil
			}()

			if got := StringS(x.key, ",", defaultValue); !reflect.DeepEqual(x.value, got) {
				t.Errorf("failed. expecting %#v, got %#v", x.value, got)
			}
		})
	}
}

/*
func TestCheck(t*testing.T){
	type testCase struct {
		title,
		key,
		expected,
		twoWay string
		mandatory,
		canBeEmpty bool
	}

	tests := []testCase{
		{
			title:      "",
			key:        "",
			expected:   "",
			twoWay:     "",
			mandatory:  false,
			canBeEmpty: false,
		},
		{
			title:      "",
			key:        "",
			expected:   "",
			twoWay:     "",
			mandatory:  false,
			canBeEmpty: false,
		},
		{
			title:      "",
			key:        "",
			expected:   "",
			twoWay:     "",
			mandatory:  false,
			canBeEmpty: false,
		},
		{
			title:      "",
			key:        "",
			expected:   "",
			twoWay:     "",
			mandatory:  false,
			canBeEmpty: false,
		},
		{
			title:      "",
			key:        "",
			expected:   "",
			twoWay:     "",
			mandatory:  false,
			canBeEmpty: false,
		},
		{
			title:      "",
			key:        "",
			expected:   "",
			twoWay:     "",
			mandatory:  false,
			canBeEmpty: false,
		},
	}

	key, defaultValue
} string, twoWay, mandatory, canBeEmpty bool) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		if mandatory {
			return fmt.Errorf("missing mandatory environment variable %s", key)
		}

		if defaultValue != "" {
			s = defaultValue
		}

		if twoWay {
			if err := os.Setenv(key, s); err != nil {
				return err
			}
		}
	}

	if s == "" && !canBeEmpty {
		return fmt.Errorf("environment variable %s can't be empty", key)
	}

	return nil
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

*/

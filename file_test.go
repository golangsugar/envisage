package envisage

import (
	"os"
	"testing"
)

const fileName = "./envisage.test.env"

type fileTestCase struct {
	row     string
	k       string
	v       string
	isThere bool
}

func createLoadFromFileTestCases() []fileTestCase {
	return []fileTestCase{
		{
			row:     "=INVALID",
			k:       "",
			v:       "INVALID",
			isThere: false,
		}, {
			row:     "#THISIS=COMMENT",
			k:       "#THISIS",
			v:       "COMMENT",
			isThere: false,
		}, {
			row:     " STARTWITH=SPACE",
			k:       "STARTWITH",
			v:       "space",
			isThere: false,
		}, {
			row:     "_INVALIDXX1=DELETE",
			k:       "_INVALIDXX1",
			v:       "DELETE",
			isThere: false,
		}, {
			row:     "10INVALIDXX11=432423423",
			k:       "10INVALIDXX11",
			v:       "432423423",
			isThere: false,
		}, {
			row:     "INVALIDXX2 SYNTAX",
			k:       "INVALIDXX2",
			v:       "SYNTAX",
			isThere: false,
		}, {
			row:     "lowercase=valid",
			k:       "lowercase",
			v:       "valid",
			isThere: true,
		}, {
			row:     "VALIDUPPERCASE=VALID",
			k:       "VALIDUPPERCASE",
			v:       "VALID",
			isThere: true,
		},
	}
}

func createFile(tests []fileTestCase) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()

	for _, x := range tests {
		if _, werr := f.WriteString(x.row); werr != nil {
			return werr
		}
	}

	return nil
}

func TestLoadFromFile(t *testing.T) {
	tests := createLoadFromFileTestCases()

	if err := createFile(tests); err != nil {
		t.Fatal(err)
	}

	const (
		updateEnvironment = true
		impossibleValue   = "..."
	)

	envmap, err := LoadFromFile(fileName, updateEnvironment, false, true)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		for mapKey, mapValue := range envmap {
			if test.k != mapKey {
				continue
			}

			t.Run("testing key "+mapKey, func(t *testing.T) {
				keyIsThere := IsThere(mapKey)

				if test.isThere != keyIsThere {
					t.Errorf("IsThere failed. expected %t, but got %t", test.isThere, keyIsThere)
				}

				if !test.isThere {
					return
				}

				if test.v != mapValue {
					t.Errorf("test failed on map value assertion. expected %s, got %s", test.v, mapValue)
				}

				if updateEnvironment {
					got := String(mapKey, impossibleValue)

					if test.v != got {
						t.Errorf("test failed on key %s environment value assertion. expected %s, got %s", mapKey, test.v, got)
					}
				}
			})
		}
	}
}

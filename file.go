package envisage

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const (
	configRegex  = `^([a-zA-Z][a-zA-Z0-9_]+)=([[\S ]*]?)$`
	rxGroupKey   = 1
	rxGroupValue = 2
)

var rxConfig = regexp.MustCompile(configRegex)

func envMap(configFile string, errorIfFileDoesntExist bool) (map[string]string, error) {
	f, err := os.Open(configFile)
	if err != nil {
		if err == os.ErrNotExist && !errorIfFileDoesntExist {
			return nil, nil
		}

		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	m := make(map[string]string)

	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())

		if s == "" || strings.HasPrefix(s, "#") { // line is commented or empty
			continue
		}

		if matches := rxConfig.FindStringSubmatch(s); len(matches) > 0 {
			k := matches[rxGroupKey]
			v := ""

			if len(matches) > 2 {
				v = matches[rxGroupValue]
			}

			m[k] = v
		}
	}

	return m, nil
}

// LoadFromFile loads environment variables values from a given text file in to a map[string]string.
// configFile is the file name, with the complete path if necessary.
// if updateEnvironment is true, all valid variable values found will be set on environment as well.
// if skipIfAlreadyDefined is true, the found variable will be added to the map anyway, but only updated in environment if not defined.
// if errorIfFileDoesntExist, the function returns with an error in case of the given file doesn't exist.
// valid lines must comply with regex ^([A-Z][A-Z0-9_]+)([=]{1})([[\S ]*]?)$.
// Examples of valid lines:
// ABC=prd
// XYZ=
// ABC="42378462%&&3 178964@"
// mnoPQR=42378462%&&3 ###
//
// Examples of *invalid* lines:
// Commented/ignored: #XYZ=4334343434 ( starts with # ).
// invalid/ignored: opt=ler0ler0 ( has to be all caps/uppercase ).
// Invalid/Ignored: _LETTERS=4334343434 ( has to start with a letter ).
// Invalid/Ignored: X=4334343434 ( should contain 2 or more chars ).
// Environment variables reference for curious: https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap08.html.
func LoadFromFile(configFile string, updateEnvironment, skipIfAlreadyDefined, errorIfFileDoesntExist bool) (map[string]string, error) {
	m, err := envMap(configFile, errorIfFileDoesntExist)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, nil
	}

	if !updateEnvironment {
		return m, nil
	}

	for k, v := range m {
		if _, ok := os.LookupEnv(k); ok && skipIfAlreadyDefined {
			continue
		}

		if serr := os.Setenv(k, v); serr != nil {
			return nil, serr
		}
	}

	return m, nil
}

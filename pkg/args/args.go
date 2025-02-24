package args

import (
	"os"
	"regexp"
	"strings"
)

// Get All Arguments as a map.
func AsMap() map[string]string {
	return arguments()
}

func arguments() map[string]string {
	arguments := make(map[string]string)

	args := os.Args
	r, _ := regexp.Compile("--(.*?)")
	for _, arg := range args {
		if r.MatchString(arg) {
			split := strings.Split(arg, "=")
			arg_name := strings.Replace(split[0], "--", "", 1)
			arguments[arg_name] = strings.Trim(split[1], " ")
		}
	}
	return arguments
}

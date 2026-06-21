package linux

import (
	"os"
	"strings"
)

func GetEnvironment(pid string) ([]string, error) {
	data, err := os.ReadFile("/proc/" + pid + "/environ")  // same thing , getting the environments
	if err != nil {
		return nil, err
	}

	raw := strings.Split(string(data), "\x00")

	envs := []string{}
	for _, env := range raw {
		if strings.TrimSpace(env) != "" {
			envs = append(envs, env)
		}
	}

	return envs, nil
}
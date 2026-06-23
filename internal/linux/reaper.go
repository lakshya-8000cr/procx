package linux    // will read the procx then sare pid then status of pis if pid contains z then add to zombe list

import "strings"

func FindZombies() ([]ProcessInfo, error) {
	pids, err := GetPIDs()
	if err != nil {
		return nil, err
	}

	zombies := []ProcessInfo{}

	for _, pid := range pids {
		info, err := GetProcessInfo(pid)
		if err != nil {
			continue
		}

		if strings.Contains(info.State, "Z") {
			zombies = append(zombies, *info)
		}
	}

	return zombies, nil
}
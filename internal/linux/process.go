package linux

import (
	"strconv"
    "os"
	"strings"
)


type ProcessInfo struct {
	PID     string
	Name    string
	State   string
	Memory  string
	Threads string
}


func IsPID(name string) bool {

	_, err := strconv.Atoi(name)

	return err == nil
}


func GetPIDs() ([]string, error) {

	entries, err := os.ReadDir("/proc")  // in this part we are reding all the numeric folder means processes in map

	if err != nil {

		return nil, err
	}

	var pids []string

	for _, entry := range entries {  // looping over the map to return all the pid we will find

		if !entry.IsDir() {

			continue
		}

		if IsPID(entry.Name()) {

			pids = append(    // one more func ispid 
				pids,
				entry.Name(),
			)
		}
	}

	return pids, nil
}


func GetProcessInfo(pid string) (*ProcessInfo, error) {   // now furthere we going deep to extract the info 
	data, err := os.ReadFile("/proc/" + pid + "/status")  // like status after getting the pid
	if err != nil {
		return nil, err
	}

	info := &ProcessInfo{ 
		PID: pid,  
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "Name:") {
			info.Name = strings.TrimSpace(strings.TrimPrefix(line, "Name:"))  // some fixing 
		}

		if strings.HasPrefix(line, "State:") {
			info.State = strings.TrimSpace(strings.TrimPrefix(line, "State:"))
		}

		if strings.HasPrefix(line, "VmRSS:") {
			info.Memory = strings.TrimSpace(strings.TrimPrefix(line, "VmRSS:"))
		}

		if strings.HasPrefix(line, "Threads:") {
			info.Threads = strings.TrimSpace(strings.TrimPrefix(line, "Threads:"))
		}
	}

	return info, nil
}
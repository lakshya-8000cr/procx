package linux

import (
	"os/exec"
	"strings"
)

type PortInfo struct {
	Port    string
	Command string
	PID     string
	User    string
	Raw     string
}


func InspectPort(port string) (*PortInfo, error) {  // we will get the port info from the same pattern but in diff way , we will run the command now in linux 
	out, err := exec.Command(     // will run lsof -i :53 then text output then parse the second line in struct , after that pretty print in cmd/port.go
		"lsof",
		"-i",
		":"+port,
	).CombinedOutput()

	output := strings.TrimSpace(string(out))

	if output == "" {
		return nil, nil
	}

	if err != nil && !strings.Contains(output, "COMMAND") {
		return nil, err
	}

	lines := strings.Split(output, "\n")

	if len(lines) < 2 {
		return nil, nil
	}

	fields := strings.Fields(lines[1])

	if len(fields) < 3 {
		return nil, nil
	}

	return &PortInfo{
		Port:    port,
		Command: fields[0],
		PID:     fields[1],
		User:    fields[2],
		Raw:     lines[1],
	}, nil
}
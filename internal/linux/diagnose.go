package linux

import (
	"os"

	"strings"
)

type Diagnosis struct {

	Process ProcessInfo

	FDs int

	Threads int

	MaxFD string

	Warnings []string
}

func GetFDCount(
	pid string,
) (int, error) {

	entries, err := os.ReadDir(
		"/proc/" + pid + "/fd",   // simple linux command is executing /proc/pid/fd
	) 

	if err != nil {

		return 0, err
	}

	return len(entries), nil
}


func GetThreadCount(
	pid string,
) (int, error) {

	entries, err := os.ReadDir(
		"/proc/" + pid + "/task",
	)

	if err != nil {

		return 0, err
	}

	return len(entries), nil
}


// this function is calling all the subfunction whihc is needed for the information
func Diagnose(
	pid string,
) (*Diagnosis, error) {

	process, err := GetProcessInfo(  
		pid,
	)

	if err != nil {

		return nil, err
	}

	fds, _ := GetFDCount(
		pid,
	)

	threads, _ := GetThreadCount(
		pid,
	)

	d := &Diagnosis{

		Process: *process,

		FDs: fds,

		Threads: threads,
	}

	BuildWarnings(d)

	return d, nil
}


func BuildWarnings(
	d *Diagnosis,
) {

	if strings.Contains(
		d.Process.State,
		"Z",
	) {

		d.Warnings = append(
			d.Warnings,

			"Zombie process detected",
		)
	}

	if d.FDs > 100 {

		d.Warnings = append(
			d.Warnings,

			"High file descriptor usage",
		)
	}

	if d.Threads > 50 {

		d.Warnings = append(
			d.Warnings,

			"High thread count",
		)
	}
}
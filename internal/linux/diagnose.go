package linux

import (
	"os/exec"

	"os"

	"strings"
)

type Diagnosis struct {

	Process ProcessInfo

	FDs int

	Threads int

	MaxFD string

	Warnings []string

	Executable  string

    WorkingDir  string

    CommandLine string

    Uptime string
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

	exe, _ := GetExecutable(pid)
cwd, _ := GetWorkingDir(pid)
cmd, _ := GetCommandLine(pid)
uptime := GetUptime(pid)

	d := &Diagnosis{

		Process: *process,

		FDs: fds,

		Threads: threads,

		Executable:  exe,
WorkingDir:  cwd,
CommandLine: cmd,
		Uptime: uptime,
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




 
func GetExecutable(pid string) (string, error) {
	return os.Readlink("/proc/" + pid + "/exe")  // executable path like /usr/bin/node
}

func GetWorkingDir(pid string) (string, error) {
	return os.Readlink("/proc/" + pid + "/cwd")  // current working dir /home/lakshya/forge/backend
}

func GetCommandLine(pid string) (string, error) { 
	data, err := os.ReadFile("/proc/" + pid + "/cmdline")  // node server.js
	if err != nil {
		return "", err
	}

	cmd := strings.ReplaceAll(string(data), "\x00", " ")
	return strings.TrimSpace(cmd), nil
}


//getuptime
func GetUptime(
	pid string,
) string {

	out, err := exec.Command(

		"ps",

		"-p",

		pid,

		"-o",

		"etime=",

	).Output()

	if err != nil {

		return "unknown"
	}

	return strings.TrimSpace(

		string(out),
	)
}
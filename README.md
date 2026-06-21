# Procx

Procx is a Linux process inspection and diagnosis CLI tool built with Go.

The purpose of this project was not to replace existing Linux tools, but to understand how Linux works internally and how applications interact with the operating system programmatically.

While learning Linux, I noticed that we often memorize commands without understanding where the information actually comes from.

Instead of treating Linux as a collection of commands, this project explores Linux as a system that exposes its information through virtual files and system interfaces.

> Linux exposes itself through virtual files. Procx simply reads and explains them.

---

# Installation

```bash
curl -fsSL https://raw.githubusercontent.com/lakshya-8000cr/procx/main/install.sh | bash
```

Verify in new terminal:

```bash
procx list
```

---

# Why I Built This

When an application crashes, we usually jump between multiple commands.

```bash
ps aux | grep myapp

cat /proc/<pid>/status

ls /proc/<pid>/fd

lsof -i :8080

journalctl

kill -15 <pid>
```
---

Each command provides only a small piece of information.

# Features

## List Running Processes

```bash
procx list
```

<img width="768" height="930" alt="image" src="https://github.com/user-attachments/assets/e9d4d492-6413-44ea-93e3-20d98b2d8486" />

Displays:

- PID
- Process name
- Process state
- Memory usage
- Thread count

---

## Diagnose a Process

```bash
procx diagnose <pid>
```

<img width="585" height="512" alt="image" src="https://github.com/user-attachments/assets/62a6c375-4575-43dc-8efb-54594bbb0ea8" />


Provides a summarized diagnosis for a process.

Displays:

- Process information
- Memory usage
- Thread count
- File descriptor count
- Executable path
- Working directory
- Startup command
- Uptime
- Potential warnings

---

## Inspect a Port

```bash
procx port <port>
```

Shows which process is currently using a specific port.

Displays:

- Port number
- Process name
- PID
- User

---

## Environment Variables

```bash
procx env <pid>
```

Displays important environment variables associated with a process ,
use sudo in case of permission denied.


<img width="1097" height="523" alt="image" src="https://github.com/user-attachments/assets/258d08ed-89d1-4df0-b9dc-910ab2dd6ddd" />



To display all variables:

```bash
procx env <pid> --full
```
---

# Architecture

```text
Linux Kernel

↓

Virtual Filesystems

↓

/proc

↓

Go

↓

Diagnosis Engine

↓

Pretty Terminal Output
```

---

# Project Structure

```text
procx/

cmd/

internal/
  linux/

main.go

README.md
```

---

# Tech Stack

- Go
- Cobra
- Linux `/proc`
- `os.ReadFile`
- `os.ReadDir`
- `os.Readlink`
- `os/exec`

---

# Key Takeaway

This project was built primarily as a learning exercise.

The goal was not to create another Linux utility, but to understand how Linux exposes information internally and how developer tools consume that data programmatically.

Instead of memorizing Linux commands, I wanted to understand where the data actually comes from.

---

# Recommended Environment

- Ubuntu
- WSL
- Linux VM

---

# License

MIT

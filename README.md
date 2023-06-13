# OEM 23-24 Electrical Coda CLI

CLI built with Go to interface with Olin Electric Motorsports ELeads 23-24 task tracking table.

## Quick Start
1. Get an API key - go to your [Coda account settings](https://coda.io/account) and scroll down to __Generate API Token__. Copy it and save it in a file in your home directory called `.coda_api_key`

2. Copy the binary, `coda`, to a location in your terminal path, such as `/usr/local/bin/` or `~/.local/bin`

## User guide
Commands:
- `coda list` lists tasks that are currently on the table
    - If you just run `coda list`, it will show your tasks
    - If you run `coda list --all`, it will show all tasks in the table
    - if you run `coda list "<full name of person>"`, it will show tasks for that person  `
- `coda add "<assignee>"` adds a task to the table assigned to the provided person
    - `-p` or `--priority`: Task priority (P0, P1, P2, P3). Default is P2
    - `-s` or `--priority`: Task Status (Not started, In progress, Finished, Blocked). Default is Not started
    - `-t` or `--type`: Task type (Hardware, Firmware, Administrative, or Other). Default is Other
- `coda help` displays a short help message

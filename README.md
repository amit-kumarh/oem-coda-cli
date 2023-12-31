# OEM 23-24 Electrical Coda CLI

CLI built with Go to interface with Olin Electric Motorsports ELeads 23-24 task tracking table.

## Quick Start
1. Get an API key - go to your [Coda account settings](https://coda.io/account) and scroll down to __Generate API Token__. Copy it and save it in a file in your home directory called `.coda_api_key`

2. Copy the binary, `coda`, to a location in your terminal path, such as `/usr/local/bin/` or `~/.local/bin`

## User guide
Commands:
- `coda list` lists open tasks that are on the table
    - If you just run `coda list`, it will show your tasks
    - If you run `coda list --all`, it will show all tasks in the table
    - if you run `coda list "<full name of person>"`, it will show tasks for that person  `
- `coda add "<assignee>"` adds a task to the table assigned to the provided person
    - `-p` or `--priority`: Task priority (P0, P1, P2, P3). Default is P2
    - `-s` or `--priority`: Task Status (Not started, In progress, Finished, Blocked). Default is Not started
    - `-t` or `--type`: Task type (Hardware, Firmware, Administrative, or Other). Default is Other
    - The terminal will then prompt you for the task name
- `coda help` displays a short help message

## Todo
Biggest thing is that's missing is the ability to edit a row to update a task's status, that code is a little more in-depth since now the user has to select a row out of the task table. However, adding that is not a huge priority for me right now (as of 6/23), as I don't see this replacing actually going onto Coda and managing tasks, but rather just a quick way to do simple tasks like check what's been assigned to me, or quickly assign a task before I forget about it.

## Contributing
If you want to contribute, feel free to send in a PR! I'll review it when I can.

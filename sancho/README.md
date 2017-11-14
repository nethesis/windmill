# Sancho

### Configuration:

- `/opt/windmill/sancho/conf.json`

```json
{
  "ssh_port": "981",
  "api_endpoint": "http://localhost:8080/api/"
}
```

### Check status
- `sancho version` (check if is installed and configured)

### Usage
- `sancho <command>`
    - `help [command]` print help of any command
    - `session <command>` execute specific command on sessions
        - `close <session-id>` close session by ID
        - `connect <session-id>` connect to server using session ID
        - `list [session-id]` list all or particular session ID
        - flags:
            - `-j` or `--json` print output in JSON format
            - `-q` or `--quiet` print only session ID
    - `version` print current version
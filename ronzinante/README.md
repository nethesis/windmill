# Ronzinante

### Configuration:

- `/opt/windmill/ronzinante/conf.json`

```
{
    "db_user": "root",
    "db_password": "YourMariaDBPassWordHere"
}
```

### Check status
- `systemctl status ronzinante`

### API enpoints
- `/api/sessions`
    - `GET /` get all connected sessions info (`server ID`, `vpn IP`, `timestamp`)
    - `GET /<session-id>` get specific session info
    - `POST /` create new session (used by `/opt/windmill/helpers/windmill-auth`)
    - `PUT /<session-id>` update session info data (used by `/opt/windmill/helpers/windmill-accounting`)
    - `DELETE /<session-id>` delete a session

- `/api/histories`
    - `GET /` get all histories (sessions information data)
    - `GET /<server-id>` get histories of specific server
    - `PUT /<server-id>` update session history (used by `/opt/windmill/helpers/windmill-disconnect`)
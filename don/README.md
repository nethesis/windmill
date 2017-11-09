# Don

Daemons:

- don-openvpn
- don-sshd

### don-openvpn

OpenVPN connection, authenticated automatically started and stopped by don-sshd.
The authentication is done using `<user>` and a random generated session id.

Config files:

- /usr/share/don/credentials
- /usr/share/don/don.ovpn

Check don-openvpn status:
```
systemctl status don-openvpn
```

### don-sshd

Ad-hoc SSHD instance, accept connections only from tunDON interfaces:

Check don-sshd status:
```
systemctl status don-sshd
```

### Start and stop Don

Start Don:

```
./don start
```

Stop Don
```
./don stop
```

Check general Don status:
```
./don status
```

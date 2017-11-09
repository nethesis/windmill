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

### don-sshd

Ad-hoc SSHD instance, accept connections only from tunDON interfaces.


### Start and stop Don

Start Don:

```
systemctl start don-sshd
```

Stop Don
```
systemctl stop don-sshd
```
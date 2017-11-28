# Don

Daemons:

- don-openvpn
- don-sshd

### don-openvpn

OpenVPN connection, automatically started and stopped by `don-sshd`.
The authentication is done using `<user>` and a random generated session id.

Config files:

- /usr/share/don/credentials
- /usr/share/don/don.ovpn

Check `don-openvpn` status:
```
systemctl status don-openvpn
```

### don-sshd

Ad-hoc SSHD instance, it listens on port `981`.

Check `don-sshd` status:
```
systemctl status don-sshd
```

### Configuration

1. Customize your OpenVPN configuration: `/usr/share/don/don.ovpn`

   - Set `YOUR_WINDMILL_PUBLIC_HOSTNAME` and `YOUR WINDMILL CERTIFICATE CN` variables accordingly
     to your WindMill installation.
     If needed, configure also an HTTP proxy for tunnelling OpenVPN connections.
   - Copy your WindMill OpenVPN CA certificate into `/usr/share/don/ca.pem` file, 
     you can find it inside your WindMill server under `/opt/windmill/easy-rsa/keys/ca.crt`.

2. Copy your WindMill support public key into `/usr/share/don/authorized_keys` file,
   you can find it inside your WindMill server under `/etc/keyholder.d/support.pub`.

3. Customize your firewall rules to allow SSH connections from `tunDON` network device.

4. Customize Don start and stop hooks to create OpenVPN credentials or to configure the system.
   See below section.

#### Hooks

Don can invoke custom executable scripts from following directories:

- `/usr/share/don/pre-start-hook.d`: these scripts are executed before the OpenVPN tunnel has been established.
   Add here a script which saves OpenVPN credentials inside `/run/don/credentials` file.

- `/usr/share/don/start-hook.d`: these scripts are executed after SSHD instance has been started.
   Add here any scripts needed to prepare the system, like creating extra ad-hoc users for remote support.

- `/usr/share/don/post-hook.d`: these scripts are executed after SSHD instance has been stopped.
  Add ere any scripts needed to revert modifications done by the start-hook scripts.

### Start and stop Don

Start Don:

```
don start
```

Stop Don
```
don stop
```

Check general Don status:
```
don status
```

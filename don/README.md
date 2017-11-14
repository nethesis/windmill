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

3. Create a custom `/usr/share/don/create-credentials` script, which saves OpenVPN credentials
   inside `/run/don/credentials` file.
   The credential file must contain two lines: on the first one put the server identification,
   on the second one add a unique session identifier. You can use `uuidgen` command for it.

4. Customize your firewall rules to allow SSH connections from `tunDON` network device.

5. Create a custom `/usr/share/don/post-hook` script if you want to execute some operations
   when Don is started.

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

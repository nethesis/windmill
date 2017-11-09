# WindMill


How many time did you have the need to connect to a customer server but the customer didn't want to open the SSH server?

WindMill is a platform for remote support which simplify accessing remote machines behind NAT or restrictive firewalls.

Characters:

- WindMill: a bastion host which bridges connections from customers and operators
- Don: the client which connects to WindMill
- Sancho: a CLI for the operators to ease remote access
- Ronzinante: the server which does the dirty job on WindMill

## Don

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

# WindMill

<p align="center">
   <img src ="https://github.com/gsanchietti/windmill/raw/master/logo/logo.png" />
</p>

How many time did you have the need to connect to a customer server but the customer didn't want to open the SSH server?
It's almost like tilting at windmills!

WindMill is a platform for remote support which simplify accessing remote machines behind NAT or restrictive firewalls.

Characters:

- **WindMill**: a bastion host which bridges connections from customers and operators
    - [docs](https://github.com/nethesis/windmill/tree/master/windmill)
- **Don**: the client which connects to WindMill
    - [docs](https://github.com/nethesis/windmill/tree/master/don)
- **Sancho**: a CLI for the operators to ease remote access
    - [docs](https://github.com/nethesis/windmill/tree/master/sancho)
- **Ronzinante**: the server which does the dirty job on WindMill
    - [docs](https://github.com/nethesis/windmill/tree/master/ronzinante)

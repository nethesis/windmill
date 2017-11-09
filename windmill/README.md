# WindMill

## Requirements
- packer: [installation](https://www.packer.io/docs/install/index.html)
- ansible: [installation]()

## Configuration
Before launch packer to provision the machine, edit the following files:
- `roles/firewall/defaults/main.yml`
    ```
    ---
    tcp_accessbyip_v4:
        "public_ip/24":
        - "22"
    ```
    change `public_ip` with your company/home/office public IPs, to grant access to bastion only from these ips.

- `roles/windmill/default/main.yml`
    ```
    ---
    caddy_public_name:
        "example.com"
    ```
    change `example.com` with the future bastion host's domain

## Build
Launch packer to provision the machine. There are two different building options:
- DigitalOcean
    ```
    DIGITALOCEAN_API_TOKEN=<your_do_api_token> packer build -only=do-centos packer.json
    ```

- Generic Machine
    ```
    SSH_HOST=<hostname_of_machine> \
    SSH_USERNAME=<username_to_access_machine> \
    SSH_PASSWORD=<password_of_user> \
    packer build -only=null-machine packer.json
    ```

:warning: Provision scripts currently support only CentOS 7 OS

## Check configuration
After provisioning, check if services is correctly configured
- `systemctl status windmill` (for the openvpn server)
- `systemctl status caddy-windmill` (for the caddy web server)
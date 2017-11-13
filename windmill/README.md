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
        "0.0.0.0/0":
            - "22"
    ```
    change `0.0.0.0/0` and `22` with your company/home/office public IPs and SSH port to restric access only from certain ips.

- `roles/windmill/defaults/main.yml`
    ```
    ---
    caddy_public_name:
        "example.com"
    mariadb_root_password:
        "YourMariaDBPassWordHere"
    keyholder_passphrase:
        "YourKeyHolderPassPhrase"
    ```
    change `example.com` with the future bastion host's domain

    change `YourMariaDBPassWordHere` with your MariaDB root password

    change `YourKeyHolderPassPhrase` with your secure passphrase to encrypt public ssh key of support ssh-agent

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
- `systemctl status windmill@windmill` (for the openvpn server default conf)
- `systemctl status windmill@windmill-https` (for the openvpn server with https conf)
- `systemctl status caddy-windmill` (for the caddy web server)
- `systemctl status ronzinante` (for the ronzinante REST API server)
- `systemctl status keyholder-agent` (for the keyholder ssh agent)
- `systemctl status keyholder-proxy` (for the keyholder ssh proxy)

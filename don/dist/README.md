# Building for NethServer

How to build RPMs using NethServer `make-rpms` command.

## NethServer 6

```
dist=el6 mockcfg=nethserver-6-x86_64 /usr/bin/make-rpms don.spec
```

## NethServer 7

```
dist=el7 mockcfg=nethserver-7-x86_64 /usr/bin/make-rpms don.spec
```

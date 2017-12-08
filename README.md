# Iproute2 for MesaLock Linux

An iproute2 utility written in Go. We provide basic ip related commands such as
showing link infomation, setting link up and down, etc. Here is the usage:

```
Usage:
  ip link show
  ip link set <ifname> (up | down)
  ip address add <ifaddr> dev <ifname>
  ip route (show | list)
  ip route add <dst> dev <ifname> [via <gateway>] [src <src>]
```

`giproute2` is based on [netlink](http://github.com/vishvananda/netlink).

## Maintainer

  - Mingshen Sun `<mssun@mesalock-linux.org>` [@mssun](https://github.com/mssun)

## License

Giproute2 is provided under the BSD license.

package main

import (
    "fmt"
    "net"
    "github.com/docopt/docopt-go"
    "github.com/vishvananda/netlink"
)

func main() {
    usage := `giproute2.

Usage:
  ip link show
  ip route <show|list>
  ip address add <ifaddr> dev <ifname>
  ip route add default via <default_gateway> dev <ifname>`

    args, _ := docopt.Parse(usage, nil, true, "giproute2 0.0.1", false)
    if args["link"] == true {
        linkShow()
    } else if args["address"] == true {
        if args["add"] == true {
            addrAdd(args["<ifaddr>"].(string), args["<ifname>"].(string))
        }
    } else if args["route"] == true {
        if args["add"] == true {
            routeAdd(args["<default_gateway>"].(string), args["<ifname>"].(string))
        } else {
            routeShow()
        }
    }
}

func linkShow() {
    links, _ := netlink.LinkList()
    for _, link := range links {
        fmt.Println(link.Attrs())
    }
}

func addrAdd(ifaddr string, ifname string) {
    link, _ := netlink.LinkByName(ifname)
    addr, _ := netlink.ParseAddr(ifaddr)
    err := netlink.AddrAdd(link, addr)
    if err != nil {
        fmt.Println(err)
    }
}

func routeShow() {
    links, _ := netlink.LinkList()
    for _, link := range links {
        routes, _ := netlink.RouteList(link, netlink.FAMILY_V4)
        for _, route := range routes {
            fmt.Println(route)
        }
    }
}

func routeAdd(defaultGateway string, ifname string) {
    link, _ := netlink.LinkByName(ifname)
    defaultGatewayIP := net.ParseIP(defaultGateway)
    route := netlink.Route{LinkIndex: link.Attrs().Index, Gw: defaultGatewayIP}
    err := netlink.RouteAdd(&route)
    if err != nil {
        fmt.Println(err)
    }
}

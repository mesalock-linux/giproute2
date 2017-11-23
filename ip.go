// Copyright (c) 2017, MesaLock Linux Authors.
// All rights reserved.
//
// This work is licensed under the terms of the BSD 3-Clause License.
// For a copy, see the LICENSE file.

package main

import (
    "fmt"
    "net"
    "github.com/docopt/docopt-go"
    "github.com/vishvananda/netlink"
)

func main() {
    usage := `giproute2 is a iproute2 rewritten in Go.

Usage:
  ip link show
  ip link set <ifname> (up | down)
  ip address add <ifaddr> dev <ifname>
  ip route (show | list)
  ip route add <dst> dev <ifname> [via <gateway>] [src <src>]`

    args, _ := docopt.Parse(usage, nil, true, "giproute2 0.0.1", false)
    if args["link"] == true {
        if args["show"] == true {
            linkShow()
        } else if args["set"] == true {
            if args["up"] == true {
                linkSetUp(args["<ifname>"].(string))
            } else if args["down"] == true {
                linkSetDown(args["<ifname>"].(string))
            }
        }
    } else if args["address"] == true {
        if args["add"] == true {
            addrAdd(args["<ifaddr>"].(string), args["<ifname>"].(string))
        }
    } else if args["route"] == true {
        if args["add"] == true {
            dst := ""
            if args["<dst>"] != nil {
                dst = args["<dst>"].(string)
            }

            gateway := ""
            if args["<gateway>"] != nil {
                gateway = args["<gateway>"].(string)
            }
            ifname := ""
            if args["<ifname>"] != nil {
                ifname = args["<ifname>"].(string)
            }

            src := ""
            if args["<src>"] != nil {
                src = args["<src>"].(string)
            }
            routeAdd(dst, gateway, ifname, src)
        } else {
            routeShow()
        }
    }
}

func linkShow() {
    links, err := netlink.LinkList()
    if err != nil {
        fmt.Println(err)
    }
    for _, link := range links {
        fmt.Println(link.Attrs())
    }
}

func addrAdd(ifaddr string, ifname string) {
    link, err := netlink.LinkByName(ifname)
    if err != nil {
        fmt.Println(err)
    }
    addr, err := netlink.ParseAddr(ifaddr)
    if err != nil {
        fmt.Println(err)
    }
    err = netlink.AddrAdd(link, addr)
    if err != nil {
        fmt.Println(err)
    }
}

func routeShow() {
    links, err := netlink.LinkList()
    if err != nil {
        fmt.Println(err)
    }
    for _, link := range links {
        routes, _ := netlink.RouteList(link, netlink.FAMILY_V4)
        for _, route := range routes {
            fmt.Println(route)
        }
    }
}

func routeAdd(dst string, gateway string, ifname string, src string) {
    link, err := netlink.LinkByName(ifname)
    if err != nil {
        fmt.Println(err)
    }
    route := netlink.Route{LinkIndex: link.Attrs().Index}
    if dst != "default" {
        dstAddr, _ := netlink.ParseAddr(dst)
        route.Dst = dstAddr.IPNet
    }
    if src != "" {
        srcIP := net.ParseIP(src)
        route.Src = srcIP
    }
    if gateway != "" {
        gatewayIP := net.ParseIP(gateway)
        route.Gw = gatewayIP
    }
    err = netlink.RouteAdd(&route)
    if err != nil {
        fmt.Println(err)
    }
}

func linkSetUp(ifname string) {
    link, err := netlink.LinkByName(ifname)
    netlink.LinkSetUp(link)
    if err != nil {
        fmt.Println(err)
    }
}

func linkSetDown(ifname string) {
    link, err := netlink.LinkByName(ifname)
    netlink.LinkSetDown(link)
    if err != nil {
        fmt.Println(err)
    }
}

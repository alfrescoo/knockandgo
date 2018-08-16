package main

import (
    "flag"
    "log"

    "github.com/giuliocomi/knockandgo/utility"
    "github.com/giuliocomi/knockandgo/network"
)

var (
    k_port = flag.Int("k", 80, "port to open via knock technique") 
    server_port = flag.Int("s", 8080, "server port") 
    server_ip = flag.String("a", "localhost", "server address (IP)")
    max_forwarding_ports = flag.Int("f", 5, "number of maximum tcp wrappers to instantiate") 
    modality = flag.String("m", "c", "modality of operation: server (s) or client (c)")
    certpath = flag.String("c", "./certs/", "the path to the PEM certificate (public in case of client, private in case of server") 
    timeout = flag.Int("t", 86400, "timeout in seconds") 
    knockable_ports[] int
)
func main() {

    flag.Parse()
        // load and validate the whitelisted knockable ports
    args, err := utility.SliceAtoi(flag.Args())
    if err != nil {
        log.Println(err)
    }

    for _, arg := range args {
        knockable_ports = append(knockable_ports, arg)
    }
    log.Println(knockable_ports)

    switch string( * modality) {
        case string("s"):
            s := network.NewUdpServer( * server_port, knockable_ports, * max_forwarding_ports, * certpath, * timeout)
            s.Run()
        case string("c"):
            c := network.NewUdpClient( * server_ip, * server_port, * k_port, * certpath, * timeout)
            c.Run()
        default:
            panic("unkown modality")
    }
}

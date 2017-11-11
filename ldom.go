package main

import (
    "github.com/miekg/dns"
    "os"
    "net"
    "fmt"
    "log"
    "flag"
    "strings"
)
func ClientConfigFromArgs(servers []string, port string, timeout int) (*dns.ClientConfig, error) {
        c := new(dns.ClientConfig)
        c.Servers = servers // make([]string, 0)
        c.Search = make([]string, 0)
        c.Port = port
        c.Ndots = 1
        c.Timeout = timeout
        c.Attempts = 2
	return c, nil
}

func main() {
    var recordType string;
    flag.StringVar(&recordType, "rtype", "A", "What type of DNS record to lookup")
    var resolvers string;
    flag.StringVar(&resolvers, "resolv", "8.8.8.8,8.8.4.4", "A comma-separated list of resolvers to use, or 'system' to use /etc/resolv.conf")

    flag.Parse()
    //fmt.Println(dns.TypeToString[dns.StringToType["AAAA"]])
    var config *dns.ClientConfig
    if (resolvers == "system") {
        config, _ = dns.ClientConfigFromFile("/etc/resolv.conf")
    } else {
	config, _ = ClientConfigFromArgs(strings.Split(resolvers, ","), "53", 5)
    }
    c := new(dns.Client)

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(flag.Args()[0]), dns.StringToType[recordType])//dns.TypeMX)
    m.RecursionDesired = true

    r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
            log.Fatalf(" *** invalid answer name %s after MX query for %s\n", os.Args[1], os.Args[1])
    }
    // Stuff must be in the answer section
    //if (len(r.Answer)
    fmt.Printf("len=%d \n", len(r.Answer))
    for _, a := range r.Answer {
            fmt.Printf("%v\n", a)
    }
}

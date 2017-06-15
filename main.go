package main

import (
	"bufio"
	"fmt"
	"github.com/miekg/dns"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments: Need a domain or a file with domain list")
		fmt.Println("Usage : " + os.Args[0] + " <domain|domain_file> [DNS server]")
		os.Exit(2)
	}

	// DNS server (default)
	dnsserver := "8.8.8.8:53"
	if len(os.Args)	== 3 {
		dnsserver = os.Args[2]
	}

	// domain or file with domains
	domainArg := os.Args[1]
	var domainList []string

	if _, err := os.Stat(domainArg); os.IsNotExist(err) {
		// single domain
		domainList = append(domainList, domainArg)
	} else {
		// list of domains from file
		f, _ := os.Open(domainArg)
		dscanner := bufio.NewScanner(f)
		for dscanner.Scan() {
			dline := dscanner.Text()
			domainList = append(domainList, dline)
		}
	}

	fmt.Println("== Number of Domains: ", len(domainList))
	dclient := new(dns.Client)

	fmt.Println("== Requests sent to DNS server: ", dnsserver)
	fmt.Printf("== Output format: %s : %s : %s : %s\n\n",
			"domain index", "domain name", "CAA value (if any)", "Full CAA record")

	for index, domain := range domainList {

		// Create DNS Message
		fmt.Printf("%d : %s : ", index, domain )
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn(domain), dns.TypeCAA)

		// Send Query
		in, _, err := dclient.Exchange(message, dnsserver)

		if err != nil {
			fmt.Println(err)
		} else {

			if len(in.Answer) != 0 {
				if t, ok := in.Answer[0].(*dns.CAA); ok {
					fmt.Printf("%s :  %s\n", t.Value, t)
				} else {
					fmt.Printf("%s\n", "Not a CAA response (CNAME alias?)")
				}
			}else{
				fmt.Printf("%s\n", "No")
			}

		}
	}
}

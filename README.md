# caa-twit
Gets you DNS Certification Authority Authorization ([CAA]("https://tools.ietf.org/html/rfc6844")) records for a domain, or a list of domains, resolved against a DNS server or your choice.

Returns CAA value (certificate signing autority), full CAA record if needed, or a "No' if CAA is not implemented.
CNAMEs are not resolved for CAA, you are alerted of the fact that you have submitted and alias. 

Timeouts are observed.

### Usage:

```bash
Usage : ./caa-twit.mac <domain|domain_file> [DNS server]
```
`domain_file` is one domain per line

### Examples:

```bash
$ ./caa-twit.mac ./domains
$ ./caa-twit.mac ./domains 4.2.2.2:53
$ ./caa-twit.mac google.com
$ ./caa-twit.mac amazon.com 8.8.8.8:53
```

### Installation
- Golang build with `go build main.go`
- Or use the prebuilt cross compiled bins for Mac,Linux,Windows


X-compile example here:
```bash
$ ./build.sh
```


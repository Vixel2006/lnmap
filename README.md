# lnmap

a shitty nmap that is faster when it works.

## why

i decided yesterday to learn pentesting. started reading about nmap. got overwhelmed. decided to build my own shitty version to understand how tcp scanning actually works.

so here it is. it connects to ports. it waits. if the port opens, cool. if not, who cares.

can't pass firewalls. syn scan? never heard of it. service version detection? lmao. it sends `GET / HTTP/1.0` and hopes for the best.

but when the firewall isn't in the way it's fast. go brr. goroutines go brrrr.

## install

```sh
git clone https://github.com/Vixel2006/lnmap.git
cd lnmap
go build -o lnmap .
```

## usage

```sh
./lnmap -target 1.2.3.4
./lnmap -target 1.2.3.4 -ports 22,80,443
./lnmap -target 1.2.3.4 -ports 1-65535 -timeout 500ms
./lnmap 1.2.3.4
```

| flag | default | what it does |
|------|---------|-------------|
| `-target` | — | ip or hostname to bully |
| `-ports` | `22,80,443` | ports or ranges: `22,80,1-1000` |
| `-timeout` | `200ms` | how long before we give up |

## how it works

1. open tcp connection (`net.DialTimeout`)
2. wait 200ms for a banner (passive probe)
3. if silent, send `GET / HTTP/1.0` and wait 2s (active probe)
4. print results
5. cry

no raw sockets. no icmp. no packet capture. just tcp. like god intended.

## features (not many)

- concurrent scanning (unlimited goroutines, go brr)
- port ranges (`22,80` or `1-1000`)
- banner grabbing (if the service talks first)
- http probe (if the service doesn't talk first)
- colorful output (the `[-]` and `[+]` are very pretty)

## roadmap

stuff i might add if i get bored (let's be real i probably won't come back to this):

- [ ] **timeout per scan** — separate dial timeout from probe timeout
- [ ] **output formats** — json, xml (copying nmap bad)
- [ ] **port status list** — only show open ports with `-open`
- [ ] **scan pause** — rate limiting / delay between probes
- [ ] **resolve hostnames** — it already does this via `net.DialTimeout` kinda
- [ ] **multiple targets** — scan a subnet or list of hosts
- [ ] **service fingerprinting** — actually try to identify services better
- [ ] **tcp connect scan progress** — show a progress bar for large scans
- [ ] **icmp ping sweep** — find live hosts before scanning them
- [ ] **concurrent host scanning** — scan multiple hosts at once
- [ ] **output to file** — save results so you don't have to scroll
- [ ] **udp scan** — send empty packets into the void
- [ ] **fin / null / xmas scans** — confuse firewalls that don't exist in our userbase
- [ ] **os fingerprinting** — guess the os from ttl and window size
- [ ] **scripting engine** — lua, obviously
- [ ] **become nmap** — ultimate goal

## license

do whatever idc

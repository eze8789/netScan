# netScan

netScan is a concurrent port scanner written in Go

### How to use it

#### Build
```make build```

#### Scan

Single Port:

```netscan -ports 443```

Multiples Ports:

```netscan -ports 1-1024```

#### Help
```Usage of ./netscan:
  -ip string
        Specify IP Address to scan. (default "127.0.0.1")
  -ports string
        Specify port range to perform scan. Allowed formats: 80 or 1-1024 (default "1-1024")
  -w int
        Amount of concurrent process scanning ports (default 8)
```
## zzl
Zzl is a reconnaissance tool that collects subdomains from SSL certificates in IP ranges

### Install zzl
#### Using go
```bash
go install github.com/DEMON1A/zzl/cmd/zzl@latest
```
#### Build from source
##### Windows
```bat
git clone https://github.com/DEMON1A/zzl
cd zzl\
go build cmd\zzl\main.go
main.exe -h
```
##### Linux
```bash
git clone https://github.com/DEMON1A/zzl
cd zzl/
go build cmd/zzl/main.go
./main -h
```
OR
```bash
go install github.com/DEMON1A/zzl/cmd/zzl@latest
```

### Usage
#### IP ranges
zzl automatically generates IP addresses between ranges, you just need to specify the start and the end point for the IP generation function
```bat
go run cmd\zzl\main.go -start-ip 141.95.90.0 -end-ip 141.95.90.255
```

You can only use `-start-ip` too, zzl is made to generate the end up dynamiclly if it isn't provided, for example if the start ip is `192.168.1.0` zzl would set the end ip dynamically to `192.168.1.255`

```bat
go run cmd\zzl\main.go -start-ip 141.95.90.0
```

#### Domain
You can still use zzl to grab SANs from a single domain 
```bat
go run cmd\zzl\main.go -domain x.com
```

Here is a one liner for bash to automate this process with many domains
```bash
for i in `cat domains.txt`; do zzl -domain $i; done
```

#### Validation
zzl do validate every single IP address found inside the IP range for both protocols **HTTP** and **HTTPs** so it never misses a result, you can set zzl to use a single protocol for validation 

```bat
go run cmd\zzl\main.go -domain x.com -only-https
```

```bat
go run cmd\zzl\main.go -domain x.com -only-http
```

#### Timeout
You can choose the timeout in seconds, By default it's set to **1**
```bat
go run cmd\zzl\main.go -domain x.com -timeout 3
```

### Credits
This tool is inspired by [Zwink](https://x.com/_zwink)'s University "Python and Bug bounty" episode, zzl is an enchanced version of his script `sslDomain_v3.py`

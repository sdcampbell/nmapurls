# nmapurls
Nmapurls parses Nmap xml reports from either piped input or command line arg and outputs a list of http(s) URL's to be used in an automation pipeline.

Edit: The output now prints hostnames when available. Otherwise it prints an IP address.

## Usage:

### Basic usage:
```
cat nmap.xml | ./nmapurls
http://192.168.1.1:80
https://192.168.1.1:443
http://192.168.1.4:80
https://192.168.1.4:443
http://192.168.1.4:8080
http://192.168.1.6:80
https://192.168.1.6:443
http://192.168.1.7:80
https://192.168.1.7:443
http://192.168.1.13:80
https://192.168.1.13:443
```

```
./nmapurls -f nmap.xml
http://192.168.1.1:80
https://192.168.1.1:443
http://192.168.1.4:80
https://192.168.1.4:443
http://192.168.1.4:8080
http://192.168.1.6:80
https://192.168.1.6:443
http://192.168.1.7:80
https://192.168.1.7:443
http://192.168.1.13:80
https://192.168.1.13:443
```

### Extract only HTTPS services:
```
./nmapurls -f nmap.xml --https-only
https://192.168.1.1:443
https://192.168.1.4:443
https://192.168.1.6:443
https://192.168.1.7:443
https://192.168.1.13:443
```

```
./nmapurls -f nmap.xml -s
https://192.168.1.1:443
https://192.168.1.4:443
https://192.168.1.6:443
https://192.168.1.7:443
https://192.168.1.13:443
```

## Options:
- `-f, --file`: Nmap XML report file path
- `-s, --https-only`: Extract only HTTPS services

## Build:

You must already have go installed, then run `go install github.com/sdcampbell/nmapurls@latest`.

package main

import (
    "encoding/xml"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

type NmapRun struct {
    XMLName xml.Name `xml:"nmaprun"`
    Hosts   []Host   `xml:"host"`
}

type Host struct {
    Addresses []Address `xml:"address"`
    Hostnames []Hostname `xml:"hostnames>hostname"` // Added hostnames
    Ports     Ports     `xml:"ports"`
}

type Address struct {
    Addr     string `xml:"addr,attr"`
    AddrType string `xml:"addrtype,attr"`
}

type Hostname struct {
    Name string `xml:"name,attr"`
}

type Ports struct {
    Port []Port `xml:"port"`
}

type Port struct {
    Protocol string  `xml:"protocol,attr"`
    PortID   int     `xml:"portid,attr"`
    State    State   `xml:"state"`
    Service  Service `xml:"service"`
}

type State struct {
    State string `xml:"state,attr"`
}

type Service struct {
    Name string `xml:"name,attr"`
}

func parseNmapXML(reader io.Reader) (*NmapRun, error) {
    bytes, err := ioutil.ReadAll(reader)
    if err != nil {
        return nil, err
    }

    var nmapRun NmapRun
    err = xml.Unmarshal(bytes, &nmapRun)
    if err != nil {
        return nil, err
    }

    return &nmapRun, nil
}

func main() {
    fileFlag := flag.String("file", "", "Nmap XML report file path")
    flag.StringVar(fileFlag, "f", "", "Nmap XML report file path (shorthand)")
    flag.Parse()

    var reader io.Reader

    if *fileFlag != "" {
        xmlFile, err := os.Open(*fileFlag)
        if err != nil {
            fmt.Println("Error opening file:", err)
            return
        }
        defer xmlFile.Close()
        reader = xmlFile
    } else {
        reader = os.Stdin
    }

    nmapRun, err := parseNmapXML(reader)
    if err != nil {
        fmt.Println("Error parsing XML:", err)
        return
    }

    for _, host := range nmapRun.Hosts {
        var address string

        if len(host.Hostnames) > 0 {
            address = host.Hostnames[0].Name // Use the first hostname if available
        } else {
            for _, addr := range host.Addresses {
                if addr.AddrType == "ipv4" {
                    address = addr.Addr
                    break
                }
            }
        }

        for _, port := range host.Ports.Port {
            if port.State.State == "open" && (port.Service.Name == "http" || port.Service.Name == "https") {
                protocol := port.Service.Name
                url := fmt.Sprintf("%s://%s:%d", protocol, address, port.PortID)
                fmt.Println(url)
            }
        }
    }
}

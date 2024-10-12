package scanner

import (
    "fmt"
    "net"
    "time"
)

func ScanNetwork() {
    fmt.Println("Scanning network for IoT devices...")
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("Error getting network interfaces:", err)
        return
    }

    for _, iface := range interfaces {
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }
        for _, addr := range addrs {
            ipnet, ok := addr.(*net.IPNet)
            if !ok || ipnet.IP.IsLoopback() {
                continue
            }
            if ipnet.IP.To4() != nil {
                scanRange(ipnet)
            }
        }
    }
}

func scanRange(ipnet *net.IPNet) {
    ip := ipnet.IP.Mask(ipnet.Mask)
    for {
        go tryConnection(ip.String())
        inc(ip)
        if !ipnet.Contains(ip) {
            break
        }
    }
    time.Sleep(2 * time.Second) // Wait for goroutines to finish
}

func tryConnection(ip string) {
    conn, err := net.DialTimeout("tcp", ip+":80", time.Millisecond*500)
    if err == nil {
        conn.Close()
        fmt.Printf("Device found: %s\n", ip)
    }
}

func inc(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}

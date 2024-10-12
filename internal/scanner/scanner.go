package scanner

import (
    "fmt"
    "net"
    "sync"
    "time"
    "github.com/loyality7/Isecot/internal/device"
)

func ScanNetwork() []*device.Device {
    fmt.Println("Scanning network for IoT devices...")
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("Error getting network interfaces:", err)
        return nil
    }

    var wg sync.WaitGroup
    var devices []*device.Device
    var mutex sync.Mutex

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
                wg.Add(1)
                go func(ipnet *net.IPNet) {
                    defer wg.Done()
                    scanRange(ipnet, &devices, &mutex)
                }(ipnet)
            }
        }
    }
    wg.Wait()
    return devices
}

func scanRange(ipnet *net.IPNet, devices *[]*device.Device, mutex *sync.Mutex) {
    ip := ipnet.IP.Mask(ipnet.Mask)
    for {
        if d := tryConnection(ip.String()); d != nil {
            mutex.Lock()
            *devices = append(*devices, d)
            mutex.Unlock()
        }
        inc(ip)
        if !ipnet.Contains(ip) {
            break
        }
    }
}

func tryConnection(ip string) *device.Device {
    commonPorts := []int{80, 443, 8080, 22, 23, 554}
    openPorts := []int{}
    for _, port := range commonPorts {
        conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Millisecond*500)
        if err == nil {
            conn.Close()
            openPorts = append(openPorts, port)
        }
    }
    if len(openPorts) > 0 {
        d, _ := device.GetDeviceInfo(ip)
        d.OpenPorts = openPorts
        return d
    }
    return nil
}

func inc(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}


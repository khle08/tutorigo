////////////////////////////////////////////////////////////////////////

package tools

import (
    "fmt"
    "net"
)

////////////////////////////////////////////////////////////////////////


func GetLocalIP() (string, error) {
    // Get a list of all network interfaces on the system
    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    // Iterate through the network interfaces
    for _, iface := range interfaces {
        // Skip loopback interfaces and those that are down
        if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
            continue
        }

        // Get a list of all addresses assigned to the interface
        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }

        // Iterate through the addresses
        for _, addr := range addrs {
            ipNet, ok := addr.(*net.IPNet)
            if !ok || ipNet.IP.IsLoopback() {
                continue
            }

            if ipNet.IP.To4() != nil {
                return ipNet.IP.String(), nil
            }
        }
    }

    return "", fmt.Errorf("Unable to find local IP address")
}


func TestIP() {
    ip, err := GetLocalIP()
    if err != nil {
        fmt.Println("err: ", err)
    } else {
        fmt.Println("Local IP: ", ip)
    }
}


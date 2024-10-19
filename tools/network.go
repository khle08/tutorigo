////////////////////////////////////////////////////////////////////////

package tools

import (
    "fmt"
    "net"

    "regexp"
    "strings"
    "os/exec"
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


func GetWifiList(system string) ([][]string, error) {
    var wifiList [][]string

    cmd := exec.Command("./tools/scripts/wifi.sh", "list")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return wifiList, err
    }

    output := string(out)
    lines := strings.Split(output, "\n")

    for idx, line := range lines {
        if idx == 0 {
            continue
        }

        // Remove the space at beginning and end points
        line = strings.TrimSpace(line)
        // Split only those 2 or more spaces
        re := regexp.MustCompile(`\s{2,}`)
        elements := re.Split(line, -1)
        // elements := strings.Fields(line)

        if len(elements) > 6 {
            wifiList = append(wifiList, elements)            
        }
    }

    return wifiList, nil
}


func TestIP() {
    ip, err := GetLocalIP()
    if err != nil {
        fmt.Println("err: ", err)
    } else {
        fmt.Println("Local IP: ", ip)
    }


    wifiList, err := GetWifiList("linux")
    if err != nil {
        fmt.Printf("Failed to list wifi networks: %v", err)
    }
    fmt.Println(wifiList)

    // for _, elements := range wifiList {
    //     fmt.Println(len(elements))
    // }
}


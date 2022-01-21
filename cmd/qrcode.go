package main

import (
    "fmt"
    "net"


    // MIT
    "github.com/yeqown/go-qrcode/v2"
    "github.com/yeqown/go-qrcode/writer/standard"
)

func getIP() (string, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            return "", err
        }

        for _, addr := range addrs {
            if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                if ipnet.IP.To4() != nil {
                    return ipnet.IP.String(), nil
                }
            }
        }
    }

    return "", fmt.Errorf("Not found")
}

func generateQRCode(img string) error {
    ip, err := getIP()
    if err != nil {
        return err
    }

    url := fmt.Sprintf("http://%s", ip)

    qrc, err := qrcode.New(url)
    if err != nil {
        return fmt.Errorf("could not generate QRCode: %v", err)
    }

    w, err := standard.New(img)
    if err != nil {
        return fmt.Errorf("standard.New failed: %v", err)
    }

    if err = qrc.Save(w); err != nil {
        return fmt.Errorf("could not save image: %v", err)
    }

    return nil
}
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
import "ipcheck/datastructs"

func ipToInt(ip string) (uint32, bool) {
	parsedIp := net.ParseIP(ip).To4()
	if parsedIp == nil {
		return 0, false
	}
	return uint32(parsedIp[0])<<24 | uint32(parsedIp[1])<<16 | uint32(parsedIp[2])<<8 | uint32(parsedIp[3]), true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the filename as a command line argument.")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	biset := datastructs.NewBitSet(1 << 32)

	counter := 0
	for scanner.Scan() {
		ip := scanner.Text()
		ipInt, ok := ipToInt(ip)
		if !ok {
			continue
		}
		if !biset.IsBitExist(ipInt) {
			counter++
		}
		biset.Add(ipInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading the file:", err)
	}

	fmt.Printf("Unique IPs: %d\n", counter)
}

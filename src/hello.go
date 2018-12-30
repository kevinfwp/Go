package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	ping "github.com/sparrc/go-ping"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\nUsage:\tfping 192.168.1.\n\tfping 172.10.10.")
		return
	}
	j := os.Args[1]
	tip := "Type `exit` to exit"
	fmt.Println(tip)
	for i := 1; i < 255; i++ {

		addr := j + strconv.Itoa(i)
		pingger, err := ping.NewPinger(addr)
		if err != nil {
			panic(err)
		}
		pingger.Count = 1
		pingger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%s\t time=%v\n", pkt.IPAddr, pkt.Rtt)
		}
		go pingger.Run()
	}
	str := bufio.NewScanner(os.Stdin)
	for str.Scan() {
		if str.Text() != "exit" {
			fmt.Println(tip)
		} else {
			os.Exit(0)
		}
	}
}

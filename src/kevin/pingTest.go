package main

import (
	"fmt"

	ping "github.com/sparrc/go-ping"
)

func main() {
	pingger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}

	pingger.Count = 3
	pingger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s:  icmp_seq=%d time=%v\n", pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)

	}

	//pingger.OnFinish =
	pingger.Run()
	stats := pingger.Statistics()
	fmt.Println(stats.PacketsSent, " <==> ", stats.PacketsRecv)
}

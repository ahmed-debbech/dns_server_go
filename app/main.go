package main

import (
	"fmt"
	"net"
	"github.com/codecrafters-io/dns-server-starter-go/app/dns"
	"github.com/k0kubun/pp/v3"
	"os"
)

// Ensures gofmt doesn't remove the "net" import in stage 1 (feel free to remove this!)
var _ = net.ListenUDP

var f = func()(int){
	scheme := pp.ColorScheme{
		Integer:       pp.Green | pp.Bold,
		Float:         pp.Black | pp.BackgroundWhite | pp.Bold,
		String:        pp.Yellow,
	}
	
	// Register it for usage
	pp.SetColorScheme(scheme)
	return 0	
}()

func log(s string){
	fmt.Println(s)
}
func logf(f string, s interface{}){
	fmt.Printf(f, s)
}

func logc(f string, s interface{}){
	_, _= pp.Fprintf(os.Stdout, f, s)
}


func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()
	
	buf := make([]byte, 512)
	var IsFirst bool = false

	for {
		log("i'm listenniiiing")
		size, source, err := udpConn.ReadFromUDP(buf)
		if IsFirst == false{
			IsFirst = true
			continue;
		}
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
	
		
		//receivedData := string(buf[:size])**
		
		fmt.Println(size, source)

		buf = []byte(buf)
		fmt.Printf("ss %08b\n", buf)

		//parse packet
		var pdnsh dns.ParsedDnsHeader = dns.ParseHeader(buf[:12])
		label := dns.ParseQuestion(buf[12:])

		// Create a response
		var dnsh []byte = dns.BuildHeader(pdnsh)
		dnsq, dns_struct := dns.BuildQuestion(label)
		var dnsa []byte = dns.BuildAnswer(dns_struct)

		logf("HEADER; %08b\n", dnsh)
		logf("QUESTION: %08b\n", dnsq)
		logf("ANSWER: %08b\n", dnsa)
		finalByte := make([]byte, 0)
		finalByte = append(finalByte, dnsh...)
		finalByte = append(finalByte, dnsq...)
		finalByte = append(finalByte, dnsa...)
		logf("FINAL: %08b\n", finalByte)
		response := finalByte

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}

package dns

import(

	"fmt"
	"encoding/binary"
)
const dns_header_len = 12

type DnsHeader struct{
	Pid uint16
	Qr  uint8
	Opcode  uint8
	Aa  uint8
	Tc  uint8
	Rd  uint8
	Ra  uint8
	Z  uint8
	Rcode  uint8
	Qdcount  uint16
	Ancount  uint16
	Nscount  uint16
	Arcount  uint16
}

type ParsedDnsHeader struct {
	Pid uint16
	Flags uint16
	Qdcount  uint16
	Ancount  uint16
	Nscount  uint16
	Arcount  uint16	
}


func bitsToBytes(bits []int) []byte {
	var byteArray []byte
	var currentByte byte

	for i, bit := range bits {
		if bit != 0 {
			currentByte |= (1 << (7 - i%8)) // Set the bit at the correct position
		}

		// Every 8 bits, append the current byte to the result and reset it
		if (i+1)%8 == 0 {
			byteArray = append(byteArray, currentByte)
			currentByte = 0 // Reset for the next byte
		}
	}

	// If there are leftover bits that didn't complete a byte
	if len(bits)%8 != 0 {
		byteArray = append(byteArray, currentByte)
	}

	return byteArray
}

func BuildHeader(parsedHeader ParsedDnsHeader) []byte{

	fmt.Println("building header section...")
	
	var opreturned uint8 = func () uint8{
		var holder uint8
		var copy uint16 = parsedHeader.Flags
		holder = uint8(((copy & 0b0111100000000000) >> 11))
		return holder
	} ()

	var nRcode uint8
	if(opreturned == 0){
		nRcode = 0
	}else{
		nRcode = 4
	}

	dh := DnsHeader{
		Pid : parsedHeader.Pid,
		Qr : 1,
		Opcode : opreturned,
		Aa : 0,
		Tc: 0,
		Rd : func () uint8{
			var holder uint8
			var copy uint16 = parsedHeader.Flags
			holder = uint8((copy & 0b0000000100000000) >> 8)
			return holder
		} (),
		Ra : 0,
		Z : 0,
		Rcode : nRcode,
		Qdcount : 1,
		Ancount : 1,
		Nscount : 0,
		Arcount : 0,
	}

	binaryString := fmt.Sprintf("%016b", dh.Pid)
	binaryString += fmt.Sprintf("%0b", dh.Qr)
	binaryString += fmt.Sprintf("%04b", dh.Opcode)
	binaryString += fmt.Sprintf("%0b", dh.Aa)
	binaryString += fmt.Sprintf("%0b", dh.Tc)
	binaryString += fmt.Sprintf("%0b", dh.Rd)
	binaryString += fmt.Sprintf("%0b", dh.Ra)
	binaryString += fmt.Sprintf("%03b", dh.Z)
	binaryString += fmt.Sprintf("%04b", dh.Rcode)
	binaryString += fmt.Sprintf("%016b", dh.Qdcount)
	binaryString += fmt.Sprintf("%016b", dh.Ancount)
	binaryString += fmt.Sprintf("%016b", dh.Nscount)
	binaryString += fmt.Sprintf("%016b", dh.Arcount)


	var intArray = []int{}
	for _, char := range binaryString{
		if char == '0' {
			intArray = append(intArray, 0)
		} else if char == '1' {
			intArray = append(intArray, 1)
		}
	}

	fmt.Println("finish header section.")
	return bitsToBytes(intArray)
}	


func ParseHeader(bh []byte) ParsedDnsHeader{
	fmt.Println("start parsing header...")

	fmt.Printf("%08b\n", bh)

	pdh := ParsedDnsHeader{
		Pid : binary.BigEndian.Uint16(bh[:2]),
		Flags: binary.BigEndian.Uint16(bh[2:4]),
		Qdcount : binary.BigEndian.Uint16(bh[4:6]),
		Ancount : binary.BigEndian.Uint16(bh[6:8]),
		Nscount : binary.BigEndian.Uint16(bh[8:10]),
		Arcount : binary.BigEndian.Uint16(bh[10:12]),
	}


	fmt.Println("end parsing header.")
	return pdh 
}
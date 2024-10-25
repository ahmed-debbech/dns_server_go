package utils

import(

	//"fmt"

)

func Uint16ToUint8(m uint16) []uint8 {

	var buf []uint8 = make([]uint8, 0)

	var first uint8
	var second uint8

	for i:=0; i<=7; i++ {
		if(m & (1<<i) != 0){
			first |= (1 << i)
		}
	}
	for i:=8; i<=15; i++ {

		if(m & (1<<i) != 0){
			second |= (1 << (i-8))
		}
	}
	buf = append(buf, second)
	buf = append(buf, first)
	//fmt.Printf("result: %08b \n", buf)
	return buf
}

func Uint32ToUint8(m uint32) []uint8 {

	var buf []uint8 = make([]uint8, 0)

	var first uint8
	var second uint8
	var third uint8
	var fourth uint8

	for i:=0; i<=7; i++ {
		if(m & (1<<i) != 0){
			first |= (1 << i)
		}
	}
	for i:=8; i<=15; i++ {

		if(m & (1<<i) != 0){
			second |= (1 << (i-8))
		}
	}
	for i:=16; i<=23; i++ {

		if(m & (1<<i) != 0){
			third |= (1 << (i-16))
		}
	}
	for i:=24; i<=31; i++ {

		if(m & (1<<i) != 0){
			fourth |= (1 << (i-24))
		}
	}

	buf = append(buf, fourth)
	buf = append(buf, third)
	buf = append(buf, second)
	buf = append(buf, first)
	//fmt.Printf("result: %08b \n", buf)
	return buf
}
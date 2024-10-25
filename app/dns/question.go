package dns

import (
	"fmt"
	"reflect"
	"github.com/codecrafters-io/dns-server-starter-go/app/utils"
)

type DnsQuestion struct {
	Name []byte
	Type uint16
	Class uint16
}

func BuildQuestion(label []byte) ([]byte, DnsQuestion){
	fmt.Println("building question section...")

	dq := DnsQuestion{
		Name : label,
		Type : 1,
		Class : 1,
	}
	
	val := reflect.ValueOf(dq)

	buf := make([]byte, 0)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if(field.Kind() == reflect.Uint16){
			g := utils.Uint16ToUint8(field.Interface().(uint16))
			buf = append(buf, []byte(g)...)
		}else{
			buf = append(buf, []byte(field.Interface().([]byte))...)
		}
	}

	fmt.Println("finish question section.")
	return buf, dq
}


func ParseQuestion(buf []byte) []byte{
	buff := make([]byte, 0)
	for i:=0; i < len(buf) && buf[i] != 0x00; i++ {
		buff = append(buff, buf[i])
	}
	buff = append(buff, 0x00)
	return buff
}
package dns

import (
	"fmt"
	"reflect"
	"github.com/codecrafters-io/dns-server-starter-go/app/utils"
)

type DnsAnswer struct {
	Name []byte
	Type uint16
	Class uint16
	Ttl uint32
	Rdlength uint16
	Rdata []byte
}

func BuildAnswer(dq DnsQuestion) []byte{
	fmt.Println("building answer section..")

	var answer = []byte{0x08,0x08,0x08,0x08}
	da := DnsAnswer{
		Name : dq.Name,
		Type : dq.Type,
		Class : dq.Class,
		Ttl : 60,
		Rdlength : 4,
		Rdata : answer,
	}

	val := reflect.ValueOf(da)

	buf := make([]byte, 0)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if(field.Kind() == reflect.Uint16){
			g := utils.Uint16ToUint8(field.Interface().(uint16))
			buf = append(buf, []byte(g)...)
		}else{
			if(field.Kind() == reflect.Uint32){
				g := utils.Uint32ToUint8(field.Interface().(uint32))
				buf = append(buf, []byte(g)...)
			}else{
				buf = append(buf, []byte(field.Interface().([]byte))...)
			}
		}
		//fmt.Printf("%s: %v\n", fieldType.Name, field.Interface())
	}

	fmt.Println("finish answer section.")
	return buf
}
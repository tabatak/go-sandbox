package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//------------------------------------------------------------------
// binary parser no renshu
//------------------------------------------------------------------
type Value interface {
	String() string
}

type UInt struct {
	Value uint64
	Name  string
}

func (ui *UInt) String() string { return fmt.Sprintf(`"%s":%d`, ui.Name, ui.Value) }

type Boolean struct {
	Value bool
	Name  string
}

func (b *Boolean) String() string { return fmt.Sprintf(`"%s":%t`, b.Name, b.Value) }

type Str struct {
	Value string
	Name  string
}

func (s *Str) String() string { return fmt.Sprintf(`"%s":"%s"`, s.Name, s.Value) }

func readBin(data []byte) {

	// formatStr := "type1:char:4 type2:uint:8 TermType:uint:8 ICCID:uint:64 DataLength:uint:8"
	formatStr := "val1:char:4 val2:char:2 val3:uint:2"
	formats := strings.Split(formatStr, " ")

	// b := bufio.NewReader(bytes.NewReader(data))
	parsedData := make(map[string]Value)
	offset := 0

	for _, format := range formats {
		name := strings.Split(format, ":")[0]
		typ := strings.Split(format, ":")[1]
		lngth, _ := strconv.Atoi(strings.Split(format, ":")[2])
		log.Printf("name=%s, type=%s, length=%d", name, typ, lngth)

		switch typ {
		case "char":
			//char
			v := data[offset : offset+lngth]
			o := &Str{
				Name:  name,
				Value: string(v),
			}
			parsedData[name] = o

		case "uint":
			//uint
			var v uint64
			switch lngth {
			case 1:
				// v = uint64(data[offset : offset+lngth])
			case 2:
				v = uint64(binary.BigEndian.Uint16(data[offset : offset+lngth]))
			case 4:
				v = uint64(binary.BigEndian.Uint32(data[offset : offset+lngth]))
			case 8:
				v = uint64(binary.BigEndian.Uint64(data[offset : offset+lngth]))

			}
			o := &UInt{
				Name:  name,
				Value: v,
			}
			parsedData[name] = o

		}

		offset = offset + lngth
	}

	log.Printf("%+v", parsedData)

}

func main() {

	data := []byte{0x41, 0x42, 0x43, 0x44, 0x30, 0x31, 0xFF, 0xFF}
	readBin(data)
}

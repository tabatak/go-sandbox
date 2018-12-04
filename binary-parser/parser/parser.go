package parser

import (
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Value ...
type Value interface {
	String() string
}

// UInt ...
type UInt struct {
	Value uint64
	Name  string
}

func (ui *UInt) String() string { return fmt.Sprintf(`"%s":%d`, ui.Name, ui.Value) }

// Bool ...
type Bool struct {
	Value bool
	Name  string
}

func (b *Bool) String() string { return fmt.Sprintf(`"%s":%t`, b.Name, b.Value) }

// Str
type Str struct {
	Value string
	Name  string
}

func (s *Str) String() string { return fmt.Sprintf(`"%s":"%s"`, s.Name, s.Value) }

// Parse ...
func Parse(data []byte, format string) map[string]Value {
	formats := strings.Split(format, " ")

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

	return parsedData
}

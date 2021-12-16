package main

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"math"
	"strconv"
	"strings"
)

const (
	testLiteral   = `D2FE28`
	testOperator0 = `38006F45291200`
	testOperator1 = `EE00D40C823060`
	test1         = `8A004A801A8002F478`
	test2         = `620080001611562C8802118E34`
	test3         = `C0015000016115A2E0802F182340`
	test4         = `A0016C880162017C3686B18A3D4780`

	testp2 = `C200B40A82`

	in = `
C20D59802D2B0B6713C6B4D1600ACE7E3C179BFE391E546CC017F004A4F513C9D973A1B2F32C3004E6F9546D005840188C51DA298803F1863C42160068E5E37759BC4908C0109E76B00425E2C530DE40233CA9DE8022200EC618B10DC001098EF0A63910010D3843350C6D9A252805D2D7D7BAE1257FD95A6E928214B66DBE691E0E9005F7C00BC4BD22D733B0399979DA7E34A6850802809A1F9C4A947B91579C063005B001CF95B77504896A884F73D7EBB900641400E7CDFD56573E941E67EABC600B4C014C829802D400BCC9FA3A339B1C9A671005E35477200A0A551E8015591F93C8FC9E4D188018692429B0F930630070401B8A90663100021313E1C47900042A2B46C840600A580213681368726DEA008CEDAD8DD5A6181801460070801CE0068014602005A011ECA0069801C200718010C0302300AA2C02538007E2C01A100052AC00F210026AC0041492F4ADEFEF7337AAF2003AB360B23B3398F009005113B25FD004E5A32369C068C72B0C8AA804F0AE7E36519F6296D76509DE70D8C2801134F84015560034931C8044C7201F02A2A180258010D4D4E347D92AF6B35B93E6B9D7D0013B4C01D8611960E9803F0FA2145320043608C4284C4016CE802F2988D8725311B0D443700AA7A9A399EFD33CD5082484272BC9E67C984CF639A4D600BDE79EA462B5372871166AB33E001682557E5B74A0C49E25AACE76D074E7C5A6FD5CE697DC195C01993DCFC1D2A032BAA5C84C012B004C001098FD1FE2D00021B0821A45397350007F66F021291E8E4B89C118FE40180F802935CC12CD730492D5E2B180250F7401791B18CCFBBCD818007CB08A664C7373CEEF9FD05A73B98D7892402405802E000854788B91BC0010A861092124C2198023C0198880371222FC3E100662B45B8DB236C0F080172DD1C300820BCD1F4C24C8AAB0015F33D280
`
)

func main() {
	binary := toBinary(util.ReadInput(in, "\n")[0])
	p := packet(binary)

	fmt.Printf("first: %d\n", first(p))
	fmt.Printf("second: %d\n", second(p))
}

func first(p Packet) int64 {
	var sum int64
	next := []Packet{p}
	for len(next) > 0 {
		p, next = next[0], next[1:]
		switch pack := p.(type) {
		case LiteralPacket:
			sum += pack.Version
		case OperatorPacket:
			sum += pack.Version
			next = append(next, pack.Packets...)
		}
	}

	return sum
}

func second(p Packet) int64 {
	switch pack := p.(type) {
	case LiteralPacket:
		return btoi(pack.Value)

	case OperatorPacket:
		switch pack.TypeID {
		case 0:
			var sum int64
			for _, p := range pack.Packets {
				sum += second(p)
			}

			return sum

		case 1:
			var sum int64 = 1
			for _, p := range pack.Packets {
				sum *= second(p)
			}

			return sum

		case 2:
			var min int64 = math.MaxInt64
			for _, p := range pack.Packets {
				if v := second(p); v < min {
					min = v
				}
			}

			return min

		case 3:
			var max int64 = 0
			for _, p := range pack.Packets {
				if v := second(p); v > max {
					max = v
				}
			}

			return max

		case 5:
			v1, v2 := second(pack.Packets[0]), second(pack.Packets[1])
			if v1 > v2 {
				return 1
			} else {
				return 0
			}

		case 6:
			v1, v2 := second(pack.Packets[0]), second(pack.Packets[1])
			if v1 < v2 {
				return 1
			} else {
				return 0
			}

		case 7:
			v1, v2 := second(pack.Packets[0]), second(pack.Packets[1])
			if v1 == v2 {
				return 1
			} else {
				return 0
			}

		}
	}

	return 0
}

func toBinary(s string) string {
	var msg strings.Builder
	for _, c := range s {
		i, _ := strconv.ParseInt(string(c), 16, 64)
		msg.WriteString(fmt.Sprintf("%04s", strconv.FormatInt(i, 2)))
	}

	return msg.String()
}

type Packet interface {
	Length() int64
}

type LiteralPacket struct {
	Version int64
	TypeID  int64
	Value   string
	Bits    int64
}

type OperatorPacket struct {
	Version int64
	TypeID  int64
	Packets []Packet
	Bits    int64
}

func packet(s string) Packet {
	version := btoi(s[:3])
	typeID := btoi(s[3:6])

	switch typeID {
	case 4:
		return NewLiteral(version, typeID, s[6:])
	default:
		return NewOperator(version, typeID, s[6:])
	}
}

func NewLiteral(version, typeID int64, s string) LiteralPacket {
	var startBit uint8 = '1'
	var nr strings.Builder

	offset := 0
	for ; startBit != '0'; offset += 5 {
		next := s[offset : offset+5]
		startBit = next[0]
		bits := next[1:]
		nr.WriteString(bits)
	}

	return LiteralPacket{
		Version: version,
		TypeID:  typeID,
		Value:   nr.String(),
		Bits:    int64(offset + 6),
	}
}

func (n LiteralPacket) Length() int64 {
	return n.Bits
}

func NewOperator(version, typeID int64, s string) OperatorPacket {
	ltype := s[0]
	s = s[1:]
	var bitCount int64
	switch ltype {
	case '0':
		bitCount = 15
	case '1':
		bitCount = 11
	}

	var packets []Packet
	var offset int64
	if bitCount == 15 {
		length := btoi(s[:bitCount])

		for offset < length {
			nextp := packet(s[offset+bitCount:])
			offset += nextp.Length()
			packets = append(packets, nextp)
		}
	} else if bitCount == 11 {
		nr := btoi(s[:bitCount])

		for i := 0; i < int(nr); i++ {
			nextp := packet(s[offset+bitCount:])
			offset += nextp.Length()
			packets = append(packets, nextp)
		}
	}

	return OperatorPacket{
		Version: version,
		TypeID:  typeID,
		Packets: packets,
		Bits:    offset + bitCount + 7,
	}
}

func (n OperatorPacket) Length() int64 {
	return n.Bits
}

func btoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 2, 64)
	return i
}

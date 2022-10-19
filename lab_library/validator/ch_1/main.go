package main

import (
	"Labs/lab_library/validator/check"
	"fmt"
	"time"
)

var checker *check.Check

type StructChecked struct {
	RequiredString    string    `validate:"required"`
	RequiredNumber    int       `validate:"required"`
	RequiredMultiple  []string  `validate:"required"`
	LenString         string    `validate:"len=1"`
	LenNumber         float64   `validate:"len=1113.00"`
	LenMultiple       []string  `validate:"len=7"`
	MinString         string    `validate:"min=1"`
	MinNumber         float64   `validate:"min=1113.00"`
	MinMultiple       []string  `validate:"min=7"`
	MaxString         string    `validate:"max=3"`
	MaxNumber         float64   `validate:"max=1113.00"`
	MaxMultiple       []string  `validate:"max=7"`
	EqString          string    `validate:"eq=3"`
	EqNumber          float64   `validate:"eq=2.33"`
	EqMultiple        []string  `validate:"eq=7"`
	NeString          string    `validate:"ne="`
	NeNumber          float64   `validate:"ne=0.00"`
	NeMultiple        []string  `validate:"ne=0"`
	LtString          string    `validate:"lt=3"`
	LtNumber          float64   `validate:"lt=5.56"`
	LtMultiple        []string  `validate:"lt=2"`
	LtTime            time.Time `validate:"lt"`
	LteString         string    `validate:"lte=3"`
	LteNumber         float64   `validate:"lte=5.56"`
	LteMultiple       []string  `validate:"lte=2"`
	LteTime           time.Time `validate:"lte"`
	GtString          string    `validate:"gt=3"`
	GtNumber          float64   `validate:"gt=5.56"`
	GtMultiple        []string  `validate:"gt=2"`
	GtTime            time.Time `validate:"gt"`
	GteString         string    `validate:"gte=3"`
	GteNumber         float64   `validate:"gte=5.56"`
	GteMultiple       []string  `validate:"gte=2"`
	GteTime           time.Time `validate:"gte"`
	EqFieldString     string    `validate:"eqfield=MaxString"`
	Inner             Struct6
	EqCSFieldString   string  `validate:"eqcsfield=Inner.EqCSFieldString"`
	NeCSFieldString   string  `validate:"necsfield=Inner.NeCSFieldString"`
	GtCSFieldString   string  `validate:"gtcsfield=Inner.GtCSFieldString"`
	GteCSFieldString  string  `validate:"gtecsfield=Inner.GteCSFieldString"`
	LtCSFieldString   string  `validate:"ltcsfield=Inner.LtCSFieldString"`
	LteCSFieldString  string  `validate:"ltecsfield=Inner.LteCSFieldString"`
	NeFieldString     string  `validate:"nefield=EqFieldString"`
	GtFieldString     string  `validate:"gtfield=MaxString"`
	GteFieldString    string  `validate:"gtefield=MaxString"`
	LtFieldString     string  `validate:"ltfield=MaxString"`
	LteFieldString    string  `validate:"ltefield=MaxString"`
	AlphaString       string  `validate:"alpha"`
	AlphanumString    string  `validate:"alphanum"`
	NumericString     string  `validate:"numeric"`
	NumberString      string  `validate:"number"`
	HexadecimalString string  `validate:"hexadecimal"`
	HexColorString    string  `validate:"hexcolor"`
	RGBColorString    string  `validate:"rgb"`
	RGBAColorString   string  `validate:"rgba"`
	HSLColorString    string  `validate:"hsl"`
	HSLAColorString   string  `validate:"hsla"`
	Email             string  `validate:"email"`
	URL               string  `validate:"url"`
	URI               string  `validate:"uri"`
	Base64            string  `validate:"base64"`
	Contains          string  `validate:"contains=purpose"`
	ContainsAny       string  `validate:"containsany=!@#$"`
	Excludes          string  `validate:"excludes=text"`
	ExcludesAll       string  `validate:"excludesall=!@#$"`
	ExcludesRune      string  `validate:"excludesrune=☻"`
	ISBN              string  `validate:"isbn"`
	ISBN10            string  `validate:"isbn10"`
	ISBN13            string  `validate:"isbn13"`
	UUID              string  `validate:"uuid"`
	UUID3             string  `validate:"uuid3"`
	UUID4             string  `validate:"uuid4"`
	UUID5             string  `validate:"uuid5"`
	ASCII             string  `validate:"ascii"`
	PrintableASCII    string  `validate:"printascii"`
	MultiByte         string  `validate:"multibyte"`
	DataURI           string  `validate:"datauri"`
	Latitude          string  `validate:"latitude"`
	Longitude         string  `validate:"longitude"`
	SSN               string  `validate:"ssn"`
	IP                string  `validate:"ip"`
	IPv4              string  `validate:"ipv4"`
	IPv6              string  `validate:"ipv6"`
	CIDR              string  `validate:"cidr"`
	CIDRv4            string  `validate:"cidrv4"`
	CIDRv6            string  `validate:"cidrv6"`
	TCPAddr           string  `validate:"tcp_addr"`
	TCPAddrv4         string  `validate:"tcp4_addr"`
	TCPAddrv6         string  `validate:"tcp6_addr"`
	UDPAddr           string  `validate:"udp_addr"`
	UDPAddrv4         string  `validate:"udp4_addr"`
	UDPAddrv6         string  `validate:"udp6_addr"`
	IPAddr            string  `validate:"ip_addr"`
	IPAddrv4          string  `validate:"ip4_addr"`
	IPAddrv6          string  `validate:"ip6_addr"`
	UinxAddr          string  `validate:"unix_addr"`
	MAC               string  `validate:"mac"`
	IsColor           string  `validate:"iscolor"`
	StrPtrMinLen      *string `validate:"min=10"`
	StrPtrMaxLen      *string `validate:"max=1"`
	StrPtrLen         *string `validate:"len=2"`
	StrPtrLt          *string `validate:"lt=1"`
	StrPtrLte         *string `validate:"lte=1"`
	StrPtrGt          *string `validate:"gt=10"`
	StrPtrGte         *string `validate:"gte=10"`
	OneOfString       string  `validate:"oneof=red green"`
	OneOfInt          int     `validate:"oneof=5 63"`
	JsonString        string  `validate:"json"`
	LowercaseString   string  `validate:"lowercase"`
	UppercaseString   string  `validate:"uppercase"`
	Datetime          string  `validate:"datetime=2006-01-02"`
}

type Struct6 struct {
	EqCSFieldString  string
	NeCSFieldString  string
	GtCSFieldString  string
	GteCSFieldString string
	LtCSFieldString  string
	LteCSFieldString string
}

func init() {
	checker = check.NewCheck()
}

func main() {
	var test StructChecked
	// 必填
	test.RequiredString = ""
	err := checker.Struct(test)
	fmt.Println(fmt.Sprintf("RequiredString########################################：\n%v", err))
	test.RequiredString = "1"
	// 必填,数字类型不能为0
	test.RequiredNumber = 0
	err = checker.Struct(test)
	fmt.Println(fmt.Sprintf("RequiredNumber########################################：\n%v", err))
	test.RequiredNumber = 1

}

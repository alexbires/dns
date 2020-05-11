package header

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

const (
	//length of the id in the DNS header in bytes
	idlength = 2
)

//DNSQuestion is the baseline struct for the dns question
type DNSQuestion struct {
	qname  []label // the array of labels making up the query
	qtype  uint16  //16 bit value of the type of record being requested
	qclass uint16  //16 bit value of the class being requested typically 1
}

//DNSHeader contains all of the values necessary to construct a DNS header
type DNSHeader struct {
	ID     uint16 //16 bit value
	QR     uint8  //1 bit value: 0 if query 1 if reply
	Opcode uint8  //4 bits long
	flags  uint8  //4 bits long AA, TC, RD, RA flags
	Z      uint8  //3 bits long z flag
	RCode  uint8  //response code 4 bits long
	QDCNT  uint16 //number of queries in question sections
	ANCNT  uint16 //number of resource records in answer section
	NSCNT  uint16 //number of server name resource records
	ARCNT  uint16 // number of resource records in addtl records section
}

type label struct {
	length  uint8  //8 bit value for the length of the label
	content []byte //array of the content for the request
}

//SetID sets the ID field of the DNS header to a value that is generated with a CSPRNG
func (dh *DNSHeader) SetID() {
	headerBuffer := make([]byte, idlength)
	_, err := rand.Read(headerBuffer)
	if err != nil {
		return
	}
	u := binary.BigEndian.Uint16(headerBuffer)
	dh.ID = u
}

//SetQuery sets the appropriate query bit and Opcode bit for the DNS header
func (dh *DNSHeader) SetQuery(query bool) {
	if query {
		dh.QR = 0
		dh.Opcode = 0
		dh.flags = 0
		dh.RCode = 0
		dh.Z = 0
	} else {
		dh.QR = 1
		dh.Opcode = 1
	}
}

//SetNumberofQuestions sets the qdcount field for the number of questions in the packet
func (dh *DNSHeader) SetNumberofQuestions(questions uint16) {
	dh.QDCNT = questions
}

//SetNumberofNameServers sets the number of resource records in the Authority
//section.  (NSCOUNT in the protocol specification)
func (dh *DNSHeader) SetNumberofNameServers(ns uint16) {
	dh.NSCNT = ns
}

//SetNumberofAdditional sets the number of resource records in the
//additional section (ARCOUNT) of the DNS header.
func (dh *DNSHeader) SetNumberofAdditional(additional uint16) {
	dh.ARCNT = additional
}

//ToByteBuffer returns a byte buffer representation of the dns header
func (dh *DNSHeader) ToByteBuffer() *[]byte {
	ret := []byte{1}
	return &ret
}

func (dh *DNSHeader) String() string {
	toReturn := fmt.Sprintf("%016b\n", dh.ID)
	toReturn += fmt.Sprintf("%1b", dh.QR)
	toReturn += fmt.Sprintf("%04b", dh.Opcode)
	toReturn += fmt.Sprintf("%04b", dh.flags)
	toReturn += fmt.Sprintf("%03b", dh.Z)
	toReturn += fmt.Sprintf("%04b\n", dh.RCode)
	return toReturn
}

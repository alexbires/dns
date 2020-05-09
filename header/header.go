package header

import (
	"crypto/rand"
	"encoding/binary"
)

const (
	idlength = 2 //length of the id in the DNS header in bytes
)

type dnsQuestion struct {
	qname  []label // the array of labels making up the query
	qtype  uint16  //16 bit value of the type of record being requested
	qclass uint16  //16 bit value of the class being requested typically 1
}

type dnsHeader struct {
	ID     uint16 //16 bit value
	QR     uint8  //1 bit value: 0 if query 1 if reply
	Opcode uint8  //4 bits long
	flags  uint8  //4 bits long AA, TC, RD, RA flags
	Z      uint8  // z flag
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

//sets the ID field of the DNS header to a value that is generated with a CSPRNG
func (dh *dnsHeader) SetID() {
	headerBuffer := make([]byte, idlength)
	_, err := rand.Read(headerBuffer)
	if err != nil {
		return
	}
	u := binary.BigEndian.Uint16(headerBuffer)
	dh.ID = u
}

func (dh *dnsHeader) SetQuery(query bool) {
	if query {
		dh.QR = 0
	} else {
		dh.QR = 1
	}
}

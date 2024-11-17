package pkgsoap

import "encoding/xml"

// SOAPEnvelope represents the SOAP envelope
type SOAPEnvelope struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XMLNsSoap string   `xml:"xmlns:soap,attr"`
	Body      SOAPBody
}

// SOAPBody represents the SOAP body
type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Fault   *SOAPFault
	Content interface{}
}

// SOAPFault represents the SOAP fault
type SOAPFault struct {
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"faultcode"`
	String  string   `xml:"faultstring"`
	Detail  string   `xml:"detail"`
}

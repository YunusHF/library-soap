package pkgsoap

import (
	"encoding/xml"
	"log"
	"net/http"
)

func EncodeSOAPResponse(rw http.ResponseWriter, envelope SOAPEnvelope) {
	envelope.XMLNsSoap = "http://www.w3.org/2003/05/soap-envelope"

	enc := xml.NewEncoder(rw)
	enc.Indent("", "  ")
	if err := enc.Encode(envelope); err != nil {
		log.Printf("Error encoding fault response: %v", err)
	}
}

// sendSOAPError sends a SOAP fault response
func SendSOAPError(rw http.ResponseWriter, message, detail string) {
	fault := SOAPFault{
		Code:   "Server",
		String: message,
		Detail: detail,
	}

	envelope := SOAPEnvelope{
		Body: SOAPBody{
			Fault: &fault,
		},
	}

	rw.WriteHeader(http.StatusInternalServerError)

	EncodeSOAPResponse(rw, envelope)
}

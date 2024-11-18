package model

import "encoding/xml"

type (
	// BooksResponse represents the response structure
	BooksResponse struct {
		XMLName xml.Name `xml:"BooksResponse"`
		Books   []*Book  `xml:".any"`
	}

	// Book represents the book structure
	Book struct {
		XMLName       xml.Name `xml:"Book"`
		ID            int      `xml:"ID"`
		Title         string   `xml:"Title"`
		Author        string   `xml:"Author"`
		PublishedDate string   `xml:"PublishedDate"`
	}

	CreateOrderEnvelope struct {
		XMLName xml.Name        `xml:"Envelope"`
		Body    CreateOrderBody `xml:"Body"`
	}

	CreateOrderBody struct {
		XMLName            xml.Name           `xml:"Body"`
		CreateOrderRequest CreateOrderRequest `xml:"CreateOrderRequest"`
	}

	CreateOrderRequest struct {
		XMLName    xml.Name   `xml:"CreateOrderRequest"`
		CustomerID uint64     `xml:"CustomerID"`
		Products   []*Product `xml:"Products>Product"`
	}

	Product struct {
		ID       uint64 `xml:"ID"`
		Quantity uint64 `xml:"Quantity"`
	}

	// CreateOrderResponse represents the structure of the outgoing SOAP response
	CreateOrderResponse struct {
		XMLName xml.Name `xml:"CreateOrderResponse"`
		OrderID uint64   `xml:"OrderID"`
		Status  string   `xml:"Status"`
		Message string   `xml:"Message"`
	}

	OrderBooksRequestService struct {
		CustomerID uint64
		Products   []*Product
	}

	OrderBooksResponseService struct {
		OrderID uint64
		Status  string
		Message string
	}
)

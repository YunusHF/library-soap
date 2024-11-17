package model

import "encoding/xml"

// Book represents the book structure
type Book struct {
	XMLName       xml.Name `xml:"Book"`
	ID            int      `xml:"id"`
	Title         string   `xml:"title"`
	Author        string   `xml:"author"`
	PublishedDate string   `xml:"published_date"`
}

// BooksResponse represents the response structure
type BooksResponse struct {
	XMLName xml.Name `xml:"BooksResponse"`
	Books   []*Book  `xml:".any"`
}

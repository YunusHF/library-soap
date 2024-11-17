<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/4Y7IK9Q.jpg" alt="Library"></a>
</p>

<h3 align="center">Library SOAP</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/YunusHF/library-soap.svg)](https://github.com/YunusHF/library-soap/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/YunusHF/library-soap.svg)](https://github.com/YunusHF/library-soap/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> A simple Go application that provides a SOAP-based API endpoint.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

This app will:
- Retrieves book data (ID, Title, Author, Published Date) from a MySQL database
- Returns the book data in a SOAP envelope format
- Uses the Gorilla Mux router for the HTTP server

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them.

```
- Go programming language installed
- MySQL server running and accessible
```

### Installation

1. Clone the repository:
```
git clone https://github.com/YunusHF/library-soap.git
```

2. Navigate to the project directory:
```
cd book-soap-api
```

3. Install the required Go packages:
```
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
```

4. Update the database connection details in the `main.go` file:
```go
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/books_db")
```

5. Ensure the `books` table exists in your MySQL database with the following columns:
   - `id` (INT)
   - `title` (VARCHAR)
   - `author` (VARCHAR)
   - `published_date` (DATE)

### Usage

1. Start the server:
```
go run main.go
```

2. Send a GET request to `http://localhost:8080/books` to retrieve the book data in the SOAP envelope format.

Example response:
```xml
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
  <soap:Body>
    <BooksResponse>
      <Book>
        <ID>1</ID>
        <Title>Clean Code</Title>
        <Author>Robert C. Martin</Author>
        <PublishedDate>2008-08-01</PublishedDate>
      </Book>
      <Book>
        <ID>2</ID>
        <Title>The Pragmatic Programmer</Title>
        <Author>Andrew Hunt</Author>
        <PublishedDate>1999-10-30</PublishedDate>
      </Book>
    </BooksResponse>
  </soap:Body>
</soap:Envelope>
```

### Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

### License

This project is licensed under the [MIT License](LICENSE).
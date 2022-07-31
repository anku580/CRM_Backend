<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

This project is basic CRUD operation that is built using **Go**. It contains 5 APIs -

- Create Customer
- Fetch All Customers
- Update a Customer
- Delete a Customer
- Fetch Customer by ID

Currently when you start the application it have a mock data of 3 customers.

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

This project is built using Go. The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To run the project, please include this project in go workspace.

### Prerequisites

Need to have Go installed in your local machine.

### Installation

Install the packages before running `go run` command.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

### Create Customer

To create a customer using POST API - `http://localhost:3000/createCustomer`, please include the following object in the request body of the request.

```
{
    "name": "Anku",
    "role" : "user",
    "email": "ankugarg5@gmail.com",
    "phone": 1234,
    "conacted": false
}
```

### View Single Customer

To get a single customer, use GET API - `http://localhost:3000/getCustomer/{id}`, please specify the customer id. Since, it a GET API, no request body is required.

#### View List of Customers

To get the entire list of the customers, use GET API - `http://localhost:3000/getAllCustomers`.

### Updating a Customer

To update a customer using PUT API - `http://localhost:3000/updateCustomer/{id}`, please include the following object in the request body of the request.

```
{
    "name": "anku",
    "role" : "user",
    "email": "ankugarg580@gmail.com",
    "phone": 123412,
    "conacted": false
}
```

### Deleting a Customer

To delete a single customer, use DELETE API - `http://localhost:3000/deleteCustomer/{id}`, please specify the customer id. Since, it a GET API, no request body is required.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Your Name - [Anku Garg](https://www.linkedin.com/in/anku580/) - ankugarg580@gmail.com

<p align="right">(<a href="#top">back to top</a>)</p>

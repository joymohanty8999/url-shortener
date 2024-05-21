# URL Shortener

A URL shortener service built with Go, MongoDB, and Gorilla Mux. This service shortens long URLs and provides an API to retrieve the original URL and check the expiration status of the shortened URL.

## Features

- Shorten a long URL
- Retrieve the original URL from the shortened URL
- Check if a URL is stored in the database and if it is expired
- Store URLs in a MongoDB database with an expiration time

## Prerequisites

- Go 1.16+
- MongoDB 4.4+
- A running MongoDB instance

## Installation

### 1. Clone the repository

```sh
git clone https://github.com/your-username/url-shortener.git
cd url-shortener

### 2. Install GO dependencies

```sh
go mod tidy
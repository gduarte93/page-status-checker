# Page Status Checker

- Checks website pages to see if they are up, down, or slow to respond
- Server/API written in Go
- Front End written in HTML, CSS, and ES6 which uses fetch to communicate with the API

# How to Run

- Make sure you have go installed on your machine: https://golang.org/doc/install
- Then you can use `make`:
  - `make run` - runs the app without building
  - `make build` - packages the app into a single executable called `status-checker`
  - `make clean` - deletes the `status-checker` executable

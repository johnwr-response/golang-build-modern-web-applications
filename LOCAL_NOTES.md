# Building Modern Web Applications with Go (Golang)

## Section: Introduction

### Introduction: who I am, and what we're going to do

#### What will we do
- Learn the fundamentals of the Go programming language
- Learn how to build a web application in Go
- Learn how to deploy the app to a live server

### Why Go? Why not PHP, or Python, or Node.js, or whatever?
- Compiles to a single binary file
- No runtimes to worry about
- Statically typed, so no surprises runtime
- Object-oriented (sort of)
- Concurrency out of the box
- Cross-platform
- Excellent package management & testing built-in
- Incredibly concise and very, very easy to learn
- Modern and web-aware from the conception

### Why use Go? - System Resources
### Installing Go, an IDE, and writing a simple program
- Download Go [here](https://go.dev/dl/)
- Check version in a terminal
````shell
PS c:\golang-build-modern-web-applications> go version
go version go1.21.4 windows/amd64
````
- Download Visual Studio Code [here](https://code.visualstudio.com/Download) or use GoLand or similar IDE
- Make sure you IDE has appropriate GO Extensions installed
- Type Go: Install Update Tools in command window (<ctrl<>shift>P) and check all tools

### Getting help: How to ask questions
### Some Useful Resources

- The Go Standard Library documentation: [link](https://pkg.go.dev/std)
- GoDoc - a great place to find open source packages: [link](https://pkg.go.dev/)
- The Mozilla Developer Network's JavaScript documentation: [link](https://developer.mozilla.org/en-US/docs/Web/javascript)
- jsDelivr - a helpful content delivery network which hosts a lot of open source javascript and css packages: [link](https://www.jsdelivr.com/)
- cdnjs - another great CDN with lots of externally hosted javascript and css resources: [link](https://cdnjs.com/)
- The Mozilla Developer Network's documentation for Cascading Style Sheets (CSS) : [link](https://developer.mozilla.org/en-US/docs/Web/CSS)

## Section: Overview of the Go Language

### A note about the terminal on Windows: Git Bash
- Powershell can also do it all, if you don't want another tool basically doing the same

### Variables & Functions
- Every single go file starts with a package declaration
- You can call this whatever you like but the conventions of idiomatic go states to use a main package
- Every go program has to have at least one function named ``main``

### Pointers
- A pointer points to a specific location in memory and gives you a means to get that specific location
- Use asterisks to define as a pointer
- Put an ampersand in front a variable to pass a reference to that variable, a pointer if you will
- You can change the value of a scoped variable by passing a reference (or pointer) to it to some other function

### Types and Structs
- Scope and variable shadowing
- Declaring a func or variable with lowercase first letter, the func is only available within my package (protected)
- Declaring a func or variable with uppercase first letter, the func is available outside my package (public)

### Receivers: Structs with functions
- Structs can also have functions associated with them
- A pointer is created to the struct in front of the function

### Other data structures: Maps and Slices
- Create a map using the built-in make function
- A map can store anything
- Maps are extremely useful, very fast and are [**immutable**]. No need to pass as a pointer
- Maps are ***NOT SORTED*** and randomized by design, you **always** need to lookup with a key!
- In Golang, we almost never uses Arrays, instead we use Slices
- Create a Slice by adding [] in front of the type
- Add to the Slice by using the built-in append function
- Slices are also extremely useful and can be sorted by using the built-in sort functions in packaged sort
- Slices can also be defined using the shorthand syntax
- Slices like arrays start counting at 0
- A Slice can also store anything

### Decision Structures
- In Go, remember that switch statements break

### Loops and ranging over data
- To make things as simple as possible in Go, you only essentially need a for loop
- The blank identifier ("_") can also be used in loops
- Go can range over "anything"
- In Go, a string is actually a slice of bytes
- Strings in Go are actually immutable, so when yiu update it actually creates a new one and destroys the old

### Interfaces
- A receiver of pointer type is much faster and best practice
- It's easier to test interfaces than concrete types

### Test quiz
- In order for something to implement an interface, it must...
  - Implement the same functions as the interface in question




### Packages
### Channels
### Reading and Writing JSON
### Writing Tests in Go

## Section: Building a Basic Web Application 
## Section: Improved Routing & Middleware
## Section: State Management with Sessions 
## Section: Choosing a Project, and Working With Forms
## Section: Javascript & CSS
## Section: Converting our HTML to Go Templates, and Creating Handlers
## Section: Writing Tests
## Section: Improved Error Handling
## Section: Persisting Data with PostgresSQL
## Section: Designing the Database Structure
## Section: Connection our Application to the Database
## Section: Updating our Tests
## Section: Sending Mail using Go
## Section: Authentication
## Section: Setting up Secure Backend Administration
## Section: Updating our Application to Accept Command Line Parameters
## Section: Deploying our Application to a Server
## Section: Finishing Touches
## Section: Where to go Next

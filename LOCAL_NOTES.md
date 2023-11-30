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
- Go Modules is the way we use packages
  - Make sure you have [Enable Go modules integration] checked in settings of your IDE
  - Create a go module:  
    ```go mod init github.com/tsawler/myniceprogram```
  - Create a new sub-folder in project folder  
    ```md helpers```
  - Create a new go file in the helpers folder
    ```ni helpers/helpers.go -type file```

### Channels
- Setup:  
  ``` powershell
  go mod init github.com/someorg/helpers
  md helpers
  ni helpers/helpers.go -type file
  ```
- Channels are a means to sending information from one part of your program to another
- Channels are unique to Go
- Channels are made by calling the built-in make function
- Remember it's best practise to close the channel when done with it!
  - Use the defer keyword which tells Go that whatever comes after, execute it as soon as the current function is done
- Sending as a concurrent operation by running this routine in its own Go Routine as Go Routines runs concurrently
- Listening for the response of the channel

### Reading and Writing JSON
### Writing Tests in Go
- Setup:
  ``` powershell
  go mod init github.com/someorg/helpers
  ```
- Go makes testing remarkably easy
- In Go the tests themselves live right besides the code you are trying to test
- Create the test file the same as the go file, only with [`_test`] added as a postfix
- A test in Go is just a function with a name that starts with [`Test`]
- There are two ways to create tests:
  1. The manual way
  2. The Table Test
- To run tests simply execute the command: ```go test```
- To run tests verbose simply execute the command: ```go test -v```
- To run tests with coverage simply execute the command: ```go test -cover```
- To run tests with coverage and get a report in html: ```go test -coverprofile=coverage.out && go tool cover -html=coverage.out```

## Section: Building a Basic Web Application 

### How web applications work: the request/response lifecycle
### Making a "Hello, World" web application
### Making our application module-ready
- Setup:
  ``` powershell
  go mod init myapp
  ```

### Functions and handlers
- In order for a function to respond to a request from a web browser it has to handle two parameters:
  - A Response Writer
  - A (pointer to a) Request
- The convention in go is that the very first thing in a comment of a function is the function name

### Error checking
### Serving HTML Templates
- Note! To enable **Emmet** on `tmpl` files, add it to **Go template files** in ***Settings - Editor - File Types*** 
### A note for Windows Users
- Use ```go run .``` to run multiple Go files at the same time

### Reorganizing our code, and adding some basic styling to pages
- Note! Added Bootstrap 5

### Enabling Go Modules and refactoring our code to use packages
- When initializing a Go module, the name of a package should correspond to a git repository 
- Setup:
  ``` powershell
  go mod init github.com/myorg/myapp
  ```
- Run:
  ``` powershell
  go run .\cmd\web\.
  ```
### Working with Layouts
- Note! Consider using `.gohtml` as instead of `.tmpl`

### Building a simple template cache
### Building a more complex template cache
### A note about the next lectures
### Setting application wide configuration
### Why the application wide config is so useful
### Optimizing our template cache by using an application config
### A note about the next lecture
- VS Code will give a warning while Golang sometimes does not. This is warning about an `import cycle` which is deliberate in the next lecture.

### Sharing data with templates

## Section: Improved Routing & Middleware

### Using pat for routing
- Pat - A Sinatra style pattern muxer for Go's net/http library [link](https://github.com/bmizerany/pat)
- Setup:
  ```shell
  go get github.com/bmizerany/pat
  ```

### Using chi for routing
- Chi - A lightweight, idiomatic and composable router for building Go HTTP services. [link](https://github.com/go-chi/chi)
- Setup:
  ```shell
  go get -u github.com/go-chi/chi/v5
  ```
- It is good practice to tidy up after adding and/or removing modules:
  ```shell
  go mod tidy
  ```

### Developing our own middleware
- Chi - HTTP package for Go that helps you prevent Cross-Site Request Forgery attacks. [link](https://github.com/justinas/nosurf)
- Setup:
  ```shell
  go get github.com/justinas/nosurf
  ```

## Section: State Management with Sessions 

### Installing and setting up a sessions package
- SCS - HTTP Session Management for Go. [link](https://github.com/alexedwards/scs)
- Setup:
  ```shell
  go get github.com/alexedwards/scs/v2
  ```

### Experimenting with sessions






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

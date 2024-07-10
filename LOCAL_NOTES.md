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
- noSurf - HTTP package for Go that helps you prevent Cross-Site Request Forgery attacks. [link](https://github.com/justinas/nosurf)
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

### What are we going to build?
- Planning
  - Deciding what to build
  - Project Scope
  - Key Functionality
  - Requirements
- This project is going to build
  - A bookings and reservations website
  - For a Bed & Breakfast with two rooms
  - What more do we need to do?
- Key Functionality
  - Showcase the property
  - Allow for booking a room for one or more nights
  - Check a room's availability
  - Book the room
  - Notify guests, and notify property owner
  - Have a back end that the owner logs into
  - Review existing bookings
  - Show a calendar of bookings
  - Change or cancel a booking
- What will we need
  - An authentication system
  - Somewhere to store information (database)
  - A means of sending notifications (email/text)

### Setting up our project
- First copy the three folders from hello-world app: [`cmd`,`pkg`,`templates`]
- Then setup:
  ```shell
  go mod init github.com/johnwr-response/golang-build-modern-web-applications/15-bookings
  ```
- Change import statements to reflect new module above in these files
  - `./cmd/web/main.go`
  - ``./pkg/handlers/handlers.go``
  - ``./pkg/render/render.go``
- Then tidy up and try running the application:
  ```shell
  go mod tidy
  go run ./cmd/web/.
  ```

### Enabling static files
### Creating pages as HTML
### Creating a landing page
### Creating a page for each room
### Adding a form to search for availability
### Improving our form
### Creating the reservation page

## Section: Javascript & CSS
### What is Javascript, and why should I care?
### Making a better date picker
- Vanilla JS Datepicker - A vanilla JavaScript remake of bootstrap-datepicker for Bulma and other CSS frameworks
  [GitHub](https://github.com/mymth/vanillajs-datepicker)
### Custom alerts using Notie
- notie - a clean and simple notification, input, and selection suite for javascript, with no dependencies
  [GitHub](https://github.com/jaredreich/notie)
### Creating modals with SweetAlert
- sweetalert2 - A beautiful, responsive, highly customizable and accessible (WAI-ARIA) replacement for JavaScript's popup boxes. Zero dependencies.
  [GitHub](https://github.com/sweetalert2/sweetalert2)
### Implementing a Javascript module
### Adding custom alerts in our Javascript module
### Using our Javascript module on the "Book Now" button
### What is CSS, and how does it work?
- Remember: Order matters. Later overrides former.

## Section: Converting our HTML to Go Templates, and Creating Handlers
### Converting our pages to Go templates
### Creating handlers for our forms & adding CSRF Protection
### Creating a handler that return JSON
### Sending & processing an AJAX request
### Sending AJAX post and generalizing our custom function
### Refactoring to use internal packages
### Server-side form validation
### Server-side form validation II
### Server-side form validation III
### Server-side form validation IV
- goValidator - Package of validators and sanitizers for strings, numerics, slices and structs
  [GitHub](https://github.com/asaskevich/govalidator)
  `go get github.com/asaskevich/govalidator`
### Displaying a response to user after posting form data
- To avoid accidental submitting, we should always 303 redirect a post to a new page instead of just displaying a page.
### Finishing up our response to user, and adding alerts
### An aside: Alternate Templating Engines
- We are using the built-in html template package in go which works fine
- People coming from web development in other languages might find the syntax for templates are a little awkward
- But there are other template engines out there, one of them being `Jet Template Engine for Go`
  [GitHub](https://github.com/CloudyKit/jet)

## Section: Writing Tests
### Writing tests for our main package
- Syntax for running and verbose viewing test results in Powershell
  `go test -v ./cmd/web/.`
- Syntax for viewing coverage percent in Powershell
  `go test -cover ./cmd/web/.`
- Syntax for viewing full coverage report in Powershell
  `(go test '-coverprofile=coverage.out' ./cmd/web/.) -and (go tool cover '-html=coverage.out')`
### Writing tests for our GET handlers
- Syntax for running and verbose viewing test results for our handlers in Powershell
  `go test -v ./internal/handlers/.`
- Syntax for viewing coverage percent in Powershell
  `go test -cover ./internal/handlers/.`
- Syntax for viewing full coverage report in Powershell
  `(go test '-coverprofile=coverage.out' ./internal/handlers/.) -and (go tool cover '-html=coverage.out')`
### Writing tests for our POST handlers
### Writing tests for our Render package
### Writing tests for our Render package II
- Syntax for running and verbose viewing test results for our handlers in Powershell
  `go test -v ./internal/render/.`
- Syntax for viewing coverage percent in Powershell
  `go test -cover ./internal/render/.`
- Syntax for viewing full coverage report in Powershell
  `(go test '-coverprofile=coverage.out' ./internal/render/.) -and (go tool cover '-html=coverage.out')`
### Getting test coverage
- To add a command alias in Powershell, first open your profile editor
  `ise $profile`
- Then add the following to create an alias that displays coverage for the current folder
    ````
    // Set an alias for calling coverage tool on current folder
    Function go_coverage {(go test '-coverprofile=coverage.out') -and (go tool cover '-html=coverage.out')}
    Set-Alias -Name go-cov-curr-folder -Value go_coverage
    ````
### Exercise: Writing tests for the Forms package
### Solution to writing tests for the Forms package
### Making running our application easier
No longer needed, the Go toolchain now does not include the tests when run, thus this will still suffice:  
  `go run ./cmd/web/.`
## Section: Improved Error Handling
### Centralizing our error handling to a helpers package
### Using our ClientError and ServerError helper functions
### Updating our tests
- Run tests in current directory including all subdirectories
  `go test -v ./...`
- Run tests for viewing coverage percent including all subdirectories
  `go test -cover ./...`
- Run tests for viewing full coverage report including all subdirectories
  `(go test '-coverprofile=coverage.out' ./...) -and (go tool cover '-html=coverage.out')`

## Section: Persisting Data with PostgresSQL
### Installing PostgresSQL
- PostgresSql can be installed in many ways. As local binaries for one, but we will be using Docker.
  ````
  cd docker
  docker-compose up -d
  docker-compose down
  ````
- Postgres can be contacted on port 
### Connecting to the database with DBeaver on a Mac
### Connecting to the database with DBeaver on Windows
- [Adminer](http://localhost:8977/?pgsql=db&username=postgres)
- [CloudBeaver](http://localhost:8978/#/)
### Basic SQL syntax
- Walkthrough of simple `CREATE table` command
- Walkthrough of simple `SELECT` command with `WHERE` clause
- Walkthrough of simple `INSERT` command
- Walkthrough of simple `UPDATE` command with `WHERE` clause
- Walkthrough of simple `DELETE` command with `WHERE` clause
### More complex queries
- Walkthrough of simple `DROP TABLE` table command
- Adding sample data
  ````
  -- DROP TABLE public.people;
  -- DROP TABLE public.emails;
  -- DROP TABLE public.phones;
  CREATE TABLE public.people(id serial NOT NULL, first_name varchar(255) NOT NULL, last_name varchar(255) NOT NULL);
  CREATE TABLE public.emails(id serial NOT NULL, people_id int8 NOT NULL, email_address varchar(255) NOT NULL);
  CREATE TABLE public.phones(id serial NOT NULL, people_id int8 NOT NULL, phone_number varchar(255) NOT NULL);
  INSERT INTO public.people(first_name,last_name) VALUES('John', 'Smith');
  INSERT INTO public.people(first_name,last_name) VALUES('Mary', 'Jones');
  INSERT INTO public.emails(people_id,email_address) VALUES(1,'john@smith.ca');
  INSERT INTO public.emails(people_id,email_address) VALUES(1,'john@gmail.com');
  INSERT INTO public.emails(people_id,email_address) VALUES(2,'mary@jones.com');
  INSERT INTO public.phones(people_id,phone_number) VALUES(1,'555-555-1234');
  INSERT INTO public.phones(people_id,phone_number) VALUES(2,'555-555-4321');
  ````
- Selecting and joining data
  ````
  SELECT email_address FROM emails WHERE people_id = 1;
  SELECT email_address FROM emails WHERE people_id = 2;
  SELECT p.first_name, p.last_name, e.email_address FROM people p LEFT JOIN emails e ON (e.people_id = p.id);
  SELECT p.first_name, p.last_name, e.email_address FROM people p LEFT JOIN emails e ON (e.people_id = p.id) WHERE p.people_id = 1;
  SELECT p.first_name, p.last_name, e.email_address FROM people p LEFT JOIN emails e ON (e.people_id = p.id) WHERE p.first_name = 'John' and p.last_name = 'Smith';
  SELECT 
    p.first_name, p.last_name, e.email_address, p2.phone_number 
  FROM 
    people p 
    LEFT JOIN emails e ON (e.people_id = p.id) 
    LEFT JOIN phones p2 ON (p.id = p2.people_id) 
  WHERE 
    p.first_name = 'John' and p.last_name = 'Smith'
  ORDER BY
    p.last_name, e.email_address;
  ````
## Section: Designing the Database Structure
### Identifying database structure, and Entity Relationship Diagrams
- Tables
  - users : id, first_name, last_name, email, password, created_at, updated_at, access_level
  - reservations : id, first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at
  - rooms : id, room_name, created_at, updated_at
  - room_restrictions : id, start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id
  - restrictions : id, restriction_name
- Foreign keys
  - reservations.room_id = rooms.id 
  - room_restrictions.room_id = rooms.id 
  - room_restrictions.reservation_id = reservations.id 
  - room_restrictions.restriction_id = restrictions.id 
### Install Soda
- Database Migrations - An intelligent means of managing the structure of our database 
- Soda is part of the Buffalo framework (also called Pop)
- To install:  
  `go install github.com/gobuffalo/pop/v6/soda@latest`
### Creating the users table using migrations
- Generate a migration for our users table
  `soda generate fizz CreateUserTable`
- Add content to up and down files for the migration
- Run migration:
  `soda migrate`
- Revert migration:
  `soda migrate down`
### Creating the rest of our database using migrations
- Generate a migration for our reservation table
  `soda generate fizz CreateReservationsTable`  
- Generate a migration for our rooms table
  `soda generate fizz CreateRoomsTable`  
- Generate a migration for our restrictions table
  `soda generate fizz CreateRestrictionsTable`  
- Generate a migration for our room_restrictions table
  `soda generate fizz CreateRoomRestrictionsTable`  
- Remember to migrate in the end
  `soda migrate`
### Setting up a foreign key
- Generate a migration for foreign keys in our reservation table
  `soda generate fizz CreateFKForReservationsTable`
- Remember to migrate in the end
  `soda migrate`
- And revert if you like
  `soda migrate down`
### Setting up the rest of our foreign keys
- Generate a migration for foreign keys in our reservation table
  `soda generate fizz CreateFKForRoomRestricionsTable`
### Adding Indices & Exercise
- Generate a migration for unique email field in our users table
  `soda generate fizz CreateUniqueIndexForUsersTable`
- Generate a migration for indices in our room_restrictions table
  `soda generate fizz CreateIndicesOnRoomRestrictionsTable`
### Solution to Exercise
- Generate a migration for unique email field in our reservations table
  `soda generate fizz AddFKAndIndeciesToReservationsTable`
- Soda reset will run all down migrations and then all up migrations  
  `soda reset`  
  To do that you first have to stop all clients connected to the database
## Section: Connecting our Application to the Database
### How to connect a Go application to a database
````
CREATE DATABASE test_connect;
CREATE TABLE users(id serial, first_name varchar, last_name varchar);
INSERT INTO users(first_name,last_name) VALUES('John', 'Smith);
INSERT INTO users(first_name,last_name) VALUES('Mary', 'Jones);
````
- Then setup:
  ```shell
  cd 16-test-connect
  ni main.go -type file
  go mod init github.com/johnwr-response/golang-build-modern-web-applications/16-test-connect
  go get github.com/jackc/pgx/v5
  ```
- To run:
  ```shell
  go mod tidy
  go run main.go
  ```
### Creating a Driver package
- pgx - PostgresSQL driver and toolkit for Go
  [GitHub](https://github.com/jackc/pgx)
- Adding pgx to project  
  ```go get github.com/jackc/pgx/v5```
- NOTE! Also updating some outdated libraries at this point, then tidying up
  ```go mod tidy```
### Connecting to the database and adding the SQL connection to our Repository
- NOTE! Also updating handlers on the side to handle errors when rendering templates
### Setting up models
### Cleaning up our code
- NOTE! Tests are now broken, will be handled later
### A word about ORMs
- Libraries like Gorm or upper/DB are fine, but they add a lot of unnecessary complexity
- Most experienced GO developers tend to write raw sequel instead using go to remove that extra layer of complexity
- The code is then faster, easier to maintain and not that difficult to write
### Setting up database functions: inserting a reservation
- In GO, the standard reference datetime is set to:  
  `Mon Jan 2 15:04:05 MST 2006 (MST is GMT-0700`  
  Or put another (US) way to somehow make it easy to remember:  
  `01/02 03:04:05PM '06 -0700`  
  Or in ISO:  
  `2006-01-02 15:04:05`
### Testing our insert reservation function
````
INSERT INTO rooms(room_name, created_at, updated_at) VALUES('Generals Quarters', now(), now());
INSERT INTO rooms(room_name, created_at, updated_at) VALUES('Majors Suite', now(), now());
````
### Inserting Room Restrictions
````
INSERT INTO restrictions(restriction_name, created_at, updated_at) VALUES('Reservation', now(), now());
````
### Searching for availability by room
### Searching for availability for all rooms
````
INSERT INTO restrictions(restriction_name, created_at, updated_at) VALUES('Owner Block', now(), now());
````
- Generate a migration for removing NOT NULL on reservationID in room_restrictions table
  `soda generate fizz AddNotNullToReservationIDForRestrictions`
### Connecting our handlers to our new database functions
### Connecting search availability to the make reservation page
### Cleaning up our make reservation page and testing everything
### Cleaning up the reservation summary page and improving validation
### Searching for availability by Room
### Providing feedback when searching by room, and connecting to the reservation page
### Connecting the rooms page to the make reservation page
### Connecting the Major's Suite page, and extracting our javascript module
### Adding a migration for seeding rooms
- Generate a migration for seeding rooms in our room table
  `soda generate sql SeedRoomsTable`
### Adding a migration for seeding restrictions
- Generate a migration for seeding restrictions in our restrictions table
  `soda generate sql SeedRestrictionsTable`

## Section: Updating our Tests
### Creating a test database repository
### Updating our existing tests to handle sessions
### Improving our tests by handling multiple cases
- Run tests for viewing full coverage report including all subdirectories
  `(go test '-coverprofile=coverage.out' ./...) -and (go tool cover '-html=coverage.out')`
### Testing Post handlers
- Run tests for viewing full coverage report on handlers only
  `(go test '-coverprofile=coverage.out' ./internal/handlers/.) -and (go tool cover '-html=coverage.out')`
### Testing AvailabilityJSON
### Completed Handler tests
### Simplifying adding post parameters

## Section: Sending Mail using Go
### Sending email using the Standard Library
### Installing a mailer package and setting up a mail channel
- Go Simple Mail - Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP.
  [GitHub](https://github.com/xhit/go-simple-mail)
- Import into project
  ```go get github.com/xhit/go-simple-mail/v2```
### Installing MailHog on a Mac for testing purposes
### Installing MailHog on Windows for testing purposes
### Actually using Docker
- Starting and stopping mailHog
  ````
  cd docker
  docker-compose up mailHog -d
  docker-compose down mailHog 
  ````
- Usage:
  - [Sendmail](`localhost:1025`)
  - [Web interface](http://localhost:8025/ "MailHog web interface")
### Creating and sending mail notifications
### Solution to sending notification to property owner
### Sending nicely formatted email using Foundation
- Foundation for Emails 2 - Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP.
  [Website](https://get.foundation/emails.html)
### Updating our tests

## Section: Authentication
### Create the login screen
### Creating the authentication handlers for the login screen
### Creating the authentication and user database functions
### Creating our handler to log in
- Prevent `Session Fixation Attacks` by renewing tokens
### Writing Authentication Middleware
### Adding a user to the database
- A simple hashed password generator for use when manually inserting users
  [The Go Playground](https://go.dev/play/p/uKMMCzJWGsW)
### Testing login
### Checking to see if a user is logged in, and logging a user out
### Protecting our routes with authentication middleware

## Section: Setting up Secure Backend Administration
### Picking an admin template
- RoyalUI-Free-Bootstrap-Admin-Template - Free Bootstrap 4 Admin Template
  [GitHub](https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template)
### Convert the admin template into a Go template
### Important: A note on the admin.layout.tmpl file
### Solution to creating admin templates
### Create stub handlers for admin functionality
### Listing all reservations
- Simple-DataTables - DataTables but in TypeScript transpiled to Vanilla JS. Lightweight, extendable and dependency free
  [GitHub](https://github.com/fiduswriter/simple-datatables)
### Listing new reservations
- Generate a migration for seeding rooms in our room table
  `soda generate fizz AddProcessedToReservationsTable`
### Showing one reservation
### Database functions for editing a reservation
### Editing a reservation
### Marking a reservation as processed
### Deleting a reservation
### Showing the reservation calendar
### Reservation Calendar II
### Reservation Calendar III
### Reservation Calendar IV
### Handling Calendar changes I
### Handling Calendar changes II
### Handling Calendar changes III





## Section: Updating our Application to Accept Command Line Parameters
## Section: Deploying our Application to a Server
## Section: Finishing Touches
## Section: Where to go Next

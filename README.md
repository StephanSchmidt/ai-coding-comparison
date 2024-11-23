# AI Coding Comparison

How can we compare AI IDEs and their capabilities?

One way is to have a very simple application and compare how easy it is to
extend the application with new features.

- The idea is to have a very simple app and codebase that everyone can 
understand.

- Then have a list of common tasks on how to change and extend the codebase
with new features.

- Record working on the common tasks with different AI IDEs and compare their
capabilities.

## Quick Start

- Install Go https://go.dev/doc/install
- Run `make` to run the application

## Application

The application is a todo list application.

The application is a Go application with a simple REST API and a simple
frontend in JavaScript. 

The frontend is in the `internal/js` folder and is a simple single page
application that allows you to add todos to the todo list. The JavaScript
application is delivered with a Golang embedded filesystem.

The frontend uses [Alpine.js](https://alpinejs.dev/) as the JavaScript framework.

The Go backend is in the `internal/todo` folder and contains the following:

- `db`: The database models and the database access functions. 
- `usecase`: The use cases of the application.
- `web`: The HTTP handlers of the application.

The database is SQLite with [GORM](https://gorm.io/) as the database access library.

The *_module.go files are the entry points for each module.


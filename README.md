# Go-with-test

A personal learning repository for learning **Go (Golang)** through hands-on coding, testing, and experimentation.

The main goal of this repository is to build a solid understanding of Go fundamentals by writing small programs and accompanying tests rather than learning only through theory.

## 📁 Repository Structure

```text
Go-with-test/
│
├── go.mod
├── README.md
├── LICENSE
│
└── stable-code/
    │
    ├── hello-world/
    │   └── ...
    │
    └── integers/
        ├── adder.go
        └── adder_test.go
```

> Directory names are kept without spaces because Go package/module import paths should not contain spaces.

## 🧪 Testing

This repository uses Go's built-in [`testing`](https://pkg.go.dev/testing) package.

Tests follow Go's standard naming convention:

```go
func TestSomething(t *testing.T) {
    // test logic
}
```

For example:

```go
func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("Expected '%d' but got '%d'", expected, sum)
    }
}
```

### Run tests

From an individual package directory:

```bash
go test
```

To run all tests in the repository:

```bash
go test ./...
```

For more detailed output:

```bash
go test -v ./...
```

## 📚 Current Topics

The repository will gradually cover topics such as:

* Go syntax and fundamentals
* Variables and data types
* Functions
* Structs
* Methods
* Pointers
* Interfaces
* Packages and modules
* Error handling
* Unit testing
* Table-driven tests
* Test coverage
* Benchmarks
* Mocks and test doubles
* Concurrency testing
* Practical Go projects

## 🌱 Learning Approach

The learning process is based on:

1. Learn a concept.
2. Implement a small example.
3. Write tests for it.
4. Experiment with the code.
5. Refactor and improve it.
6. Move on to a more advanced concept.

The emphasis is on understanding **how Go works internally and why things are written the way they are**, rather than simply memorizing syntax.

## 🔧 Requirements

* [Go](https://go.dev/) installed
* Git
* A code editor such as VS Code

Check the installed Go version with:

```bash
go version
```

## 🔀 Git Workflow

The repository uses Git for version control.

The general workflow is:

```bash
git pull
git add .
git commit -m "Describe your changes"
git push
```

The `main` branch is intended to contain stable code, while development and experimentation can be done on a separate branch.

## 🎯 Goal

The long-term goal of this repository is to progress from basic Go syntax and testing to building **real-world Go applications and services**.

---

**Learning by building, testing, and breaking things.**

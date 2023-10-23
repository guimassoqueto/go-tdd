## Golang TDD


> In Test Driven Development Tests are created before actual code.
Than developers run the (UNIT) tests and Tests fail (because there is no code)
Than developer Write the code to pass the tests and Run the tests again.
Tests will PASS, against MVP(minimum viable product) code. 
lastly they REFACTOR YOUR CODE -> RE-DO ALL STEPS

<br/><br/>

> Write tests. Run the code: it fails. 
Because there is no code. Write code. 
Run the tests again: they will pass. Refactor, and re-do.


### Notes

To benchmarks run:
```shell
go test ./... -bench=.
```

To report coverage run:
```shell
go test -cover
```

To detect race conditions (on tests or in regular go files):
```shell
go run -race main.go
go test ./... -race
```

Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.

Vet is normally invoked through the go command. This command vets the package in the current directory: 
```shell
go vet
```

### Recommended Books
Concurrency in Go: Tools and Techniques for Developers - Katherine Cox-buday
Go Programming Language - Allan A.A. Donovan


### Important Concepts to Fix

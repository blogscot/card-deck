# A Deck of Cards

This Golang library is my implementation of [Gophercise Exercise 9](https://gophercises.com/), 'A Deck of Cards'.


## Developer Notes

The library contains generated files that produce string representations of the card suits, and values. These were generated using the [stringer library](https://godoc.org/golang.org/x/tools/cmd/stringer). To regenerated them use the following commands:

```
go get golang.org/x/tools/cmd/stringer
go generate
```
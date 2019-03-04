# First GO REST API

## Quick Start

### create project

``` bash
$ mkdir -p $GOPATH/src/github/{your username}/{project name}
```
### Write main.go
``` bash
# simply run
$ go run ./main.go
# or build then run ... 
$ go build
# run the execution file
$ ./project_name
```
### Install mux router
``` bash
go get -u github.com/gorilla/mux
```

### interesting delete item from array

```go
books = append(books[:index], books[index+1:]...)
```
## Version

1.0.0

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

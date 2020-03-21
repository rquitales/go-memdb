<!--
 Copyright (c) 2020 Ramon Quitales
 
 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

# go-memdb

A very simple in-memory key-value database implemented in Go for investigative purposes. Supports basic CRUD operations and nested transactions.

All CRUD operations occur at O(1) time complexity, as a hash map is utilised. A separate hash map (index) is created for the values stored, such that we can count the number of occurrences of a given value at O(1) time complexity as well.

:warning: WARNING: **Not** for production use!

## Documentation
The full documentation is available on [GoDoc](https://pkg.go.dev/github.com/rquitales/go-memdb).

## Getting Started

These instructions will inform you on how to run this project on a local/development machine for investigative purposes. Please do not use this in production. There is no error handling or tests implemented yet.

### Prerequisites
 * A working Go installation (v1.14 or above, preferred)

 An example CLI that integrate's this project's `memdb` APIs can be downloaded from the  [releases page](https://github.com/rquitales/go-memdb/releases).

### Installing/Importing
Fetch this package using:

```
go get github.com/rquitales/go-memdb/memdb
```

### Using This Package
Use within your own projects by importing `memdb`:

```{go}
import (
    "github.com/rquitales/go-memdb/memdb
)
```

Create a new database instance with:
```{go}
db := memdb.NewDB()
```

Please refer to the [GoDoc](https://pkg.go.dev/github.com/rquitales/go-memdb) page full documentation. A example CLI implementation of `memdb` can also be seen in the [main.go](main.go) file.

### Running The Example CLI
 1. Download the appropriate binary from your target OS from the [releases page](https://github.com/rquitales/go-memdb/releases), or compile from source with `go build -v .` after cloning this repository.
 2. Start the executable, eg: `./go-memdb`
 3. Type your commands (into STDIN)

 Supported statements are:
  * SET [key] [value]
  * GET [key]
  * DELETE [key]
  * COUNT [value]
  * END
  * BEGIN
  * ROLLBACK
  * COMMIT

The query/functions are case insensitive, but the key/values are case sensitive!




## Built With

* [Go >=v1.14](https://golang.org/dl/) - Programming language

## Authors

* **Ramon Quitales** - *Initial work* - [rquitales](https://github.com/rquitales)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

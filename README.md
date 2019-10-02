# mymod - sample Go modules package

For more details see https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/

To use the packages in this module, add an import in Go code and call it:

```
import "github.com/zh/mymod"
import "github.com/zh/mymod/clientlib"

// ... later
s := mymod.Config()
h := clientlib.Hello()
```

To use the `mymod-server` binary, do:

```
$ go get github.com/zh/mymod/cmd/mymod-server
$ mymod-server
```

All packages in this module are importable by other modules, except for packages
located in the `internal` directory. These can only be used from within the
module itself, but cannot be imported from the outside.

To run all tests in this module:

```
go test ./...
```

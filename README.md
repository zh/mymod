# mymod - sample Go modules package

To use the packages in this module, add an import in Go code and call it:

```
import "github.com/zh/mymod"

// ... later
s := mymod.Config()
h := mymod.Hello()
```

All packages in this module are importable by other modules, except for packages
located in the `internal` directory. These can only be used from within the
module itself, but cannot be imported from the outside.

To run all tests in this module:

```
go test ./...

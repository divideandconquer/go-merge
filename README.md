# go-merge
go-merge is a golang library for recursively merging data.  It will merge and overwrite maps, structs, and other basic golang types
as long as the base data and the override data are of a similar structure.

## Usage

To use this package first you must install it: 

```
go get github.com/divideandconquer/go-merge/merge
```

Then you can import it in your code and use it as follows:

```golang
package main

import (
	"log"
	"github.com/divideandconquer/go-merge/merge"
)

func main() {
	baseData := make(map[string]string)
	baseData["test"] = "foo"
	baseData["test2"] = "bar"

	overrideData := make(map[string]string)
	overrideData["test2"] = "override"

	result := merge.Merge(baseData, overrideData)

	log.Printf("%#v", result) // prints: map[string]string{"test":"foo", "test2":"override"}
}

```

For more examples of usage check the [unit tests](/merge/merge_test.go).

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Add tests for your changes
4. Test your changes (`go test merge/*.go`)
5. Commit your changes (`git commit -am 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new Pull Request

## License
This library is licensed using the Apache-2.0 [License](/LICENSE):

```
Copyright (c) 2015, Kyle Boorky
```
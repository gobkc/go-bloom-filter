# go-bloom-filter
A simple golang bloom filter


### Contributing
You can commit PR to this repository

### How to get it?
````
go get -u github.com/gobkc/go-bloom-filter
````

### Quick start
````
package main

import (
	"github.com/gobkc/go-bloom-filter"
	"fmt"
)

func main() {
    b := gobloom.NewBloom()
    b.Add("Alple")
    b.Add(123)
    fmt.Println(b.Has("Apple"))
    fmt.Println(b.Has(123))
    fmt.Println(b.Has("orange"))
}
````
result:
````
true
true
false
````

### 100 million data unit test

````
=== RUN   TestBloom_Has
=== RUN   TestBloom_Has/test_has_element_1
=== RUN   TestBloom_Has/test_has_element_2
=== RUN   TestBloom_Has/test_not_has_element_1
=== RUN   TestBloom_Has/test_not_has_element_2
--- PASS: TestBloom_Has (26.73s)
    --- PASS: TestBloom_Has/test_has_element_1 (2.43s)
    --- PASS: TestBloom_Has/test_has_element_2 (0.00s)
    --- PASS: TestBloom_Has/test_not_has_element_1 (0.00s)
    --- PASS: TestBloom_Has/test_not_has_element_2 (0.00s)
PASS
ok      gobloom 26.760s
````

### License
© Gobkc, 2023~time.Now

Released under the Apache License
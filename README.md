# SortedMap

[![Go Report Card](https://goreportcard.com/badge/github.com/tobshub/go-sortedmap)](https://goreportcard.com/report/github.com/tobshub/go-sortedmap) [![GoDoc](https://godoc.org/github.com/tobshub/go-sortedmap?status.svg)](https://godoc.org/github.com/tobshub/go-sortedmap)

SortedMap is a simple library that provides a value-sorted `map[K]V` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a section of stored values.

```sh
go get -u github.com/tobshub/go-sortedmap
```

### Complexity
Operation               | Worst-Case
------------------------|-----------
Has, Get                | ```O(1)```
Delete, Insert, Replace | ```O(n)```

## Example Usage

```go
package main

import (
  "fmt"
  "time"

  "github.com/tobshub/go-sortedmap"
  "github.com/tobshub/go-sortedmap/asc"
)

func main() {
  // Create an empty SortedMap with a size suggestion and a less than function:
  sm := sortedmap.New(4, asc.Time)

  // Insert example records:
  sm.Insert("OpenBSD",  time.Date(1995, 10, 18,  8, 37, 1, 0, time.UTC))
  sm.Insert("UnixTime", time.Date(1970,  1,  1,  0,  0, 0, 0, time.UTC))
  sm.Insert("Linux",    time.Date(1991,  8, 25, 20, 57, 8, 0, time.UTC))
  sm.Insert("GitHub",   time.Date(2008,  4, 10,  0,  0, 0, 0, time.UTC))

  // Set iteration options:
  reversed   := true
  lowerBound := time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC)
  upperBound := time.Now()

  // Select values > lowerBound and values <= upperBound.
  // Loop through the values, in reverse order:
  iterCh, err := sm.BoundedIterCh(reversed, lowerBound, upperBound)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer iterCh.Close()

  for rec := range iterCh.Records() {
    fmt.Printf("%+v\n", rec)
  }
}
```

Check out the [examples](https://github.com/tobshub/go-sortedmap/tree/master/examples), [documentation](https://godoc.org/github.com/tobshub/go-sortedmap), and test files, for more features and further explanations.

## Benchmarks

```sh
BenchmarkNew-8                               	    500	       88.52 ns/op

BenchmarkHas1of1CachedRecords-8              	    500	        7.754 ns/op
BenchmarkHas1of1Records-8                    	    500	       88.17 ns/op

BenchmarkGet1of1CachedRecords-8              	    500	       12.22 ns/op
BenchmarkGet1of1Records-8                    	    500	       78.72 ns/op

BenchmarkDelete1of1Records-8                 	    500	      156.5 ns/op

BenchmarkInsert1Record-8                     	    500	      835.6 ns/op
BenchmarkReplace1of1Records-8                	    500	      314.2 ns/op

BenchmarkDelete1of10Records-8                	    500	      567.0 ns/op
BenchmarkDelete1of100Records-8               	    500	     1895 ns/op
BenchmarkDelete1of1000Records-8              	    500	     2314 ns/op
BenchmarkDelete1of10000Records-8             	    500	    37247 ns/op

BenchmarkBatchDelete10of10Records-8          	    500	     2730 ns/op
BenchmarkBatchDelete100of100Records-8        	    500	    38437 ns/op
BenchmarkBatchDelete1000of1000Records-8      	    500	   586059 ns/op
BenchmarkBatchDelete10000of10000Records-8    	    500	 16578997 ns/op

BenchmarkBatchGet10of10Records-8             	    500	      907.9 ns/op
BenchmarkBatchGet100of100Records-8           	    500	     6139 ns/op
BenchmarkBatchGet1000of1000Records-8         	    500	    43334 ns/op
BenchmarkBatchGet10000of10000Records-8       	    500	   559872 ns/op

BenchmarkBatchHas10of10Records-8             	    500	      629.0 ns/op
BenchmarkBatchHas100of100Records-8           	    500	     5277 ns/op
BenchmarkBatchHas1000of1000Records-8         	    500	    35924 ns/op
BenchmarkBatchHas10000of10000Records-8       	    500	   510189 ns/op

BenchmarkBatchInsert10Records-8              	    500	     6093 ns/op
BenchmarkBatchInsert100Records-8             	    500	    64635 ns/op
BenchmarkBatchInsert1000Records-8            	    500	   861841 ns/op
BenchmarkBatchInsert10000Records-8           	    500	 19096605 ns/op

BenchmarkReplace1of10Records-8               	    500	      800.1 ns/op
BenchmarkReplace1of100Records-8              	    500	     2721 ns/op
BenchmarkReplace1of1000Records-8             	    500	     3647 ns/op
BenchmarkReplace1of10000Records-8            	    500	    56905 ns/op

BenchmarkBatchReplace10of10Records-8         	    500	     8055 ns/op
BenchmarkBatchReplace100of100Records-8       	    500	   104597 ns/op
BenchmarkBatchReplace1000of1000Records-8     	    500	  1580840 ns/op
BenchmarkBatchReplace10000of10000Records-8   	    500	 53922019 ns/op
```

The above benchmark tests were ran using a 4.0GHz Intel Core i7-4790K (Haswell) CPU.

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).

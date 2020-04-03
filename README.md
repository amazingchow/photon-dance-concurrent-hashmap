# comap

![comap](doc/comap.jpeg)

## Get Started

### Prerequisites

* OS: Linux
* Golang: v1.12+

### Usage

```golang
import (
    ...
    "github.com/amazingchow/comap"
    ...
)

...

cm := NewCoMap()
cm.Store("Bob", 15)
...
```

## Benchmark

```text
-----------------------------------------------------------------------
Item                                    Iter                TPS 
-----------------------------------------------------------------------
      CoMap Throughput Batch-1          5000000               376 ns/op
  GolangMap Throughput Batch-1          3000000               426 ns/op
     CoMap Throughput Batch-16           500000              3474 ns/op
 GolangMap Throughput Batch-16           300000              8354 ns/op
     CoMap Throughput Batch-32           200000              7016 ns/op
 GolangMap Throughput Batch-32           200000             16276 ns/op
     CoMap Throughput Batch-64           100000             13950 ns/op
 GolangMap Throughput Batch-64           100000             23842 ns/op
    CoMap Throughput Batch-128            50000             29551 ns/op
GolangMap Throughput Batch-128            30000             45355 ns/op
-----------------------------------------------------------------------
```

## Documentation

### Reference

* [How ConcurrentHashMap Works Internally in Java, by Arun Pandey](https://dzone.com/articles/how-concurrenthashmap-works-internally-in-java)

## Contributing

### Step 1

* 🍴 Fork this repo!

### Step 2

* 🔨 HACK AWAY!

### Step 3

* 🔃 Create a new PR using https://github.com/amazingchow/comap/compare!

## FAQ

* refer to [FAQ](FAQ.md).

## Support

* Reach out to me at <jianzhou42@163.com>.

## License

* This project is licensed under the MIT License - see the **[MIT license](http://opensource.org/licenses/mit-license.php)** for details.

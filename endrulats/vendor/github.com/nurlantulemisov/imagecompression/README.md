[![Go](https://github.com/nurlantulemisov/imagecompression/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/nurlantulemisov/imagecompression/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/nurlantulemisov/imagecompression/badge.svg)](https://coveralls.io/github/nurlantulemisov/imagecompression)
## SVD image compression

An implementation image compression using SVD decomposition on Go

### Built With

* [Go 1.17]()
* [Gonum](https://github.com/gonum/gonum)

## Compression examples

Compress ratio | Image
-------|-----------------------------------
| Original    | ![](fixtures/test10x10.jpeg)
| Ratio 10%   | ![](fixtures/test10x10_compressed10.jpeg)|
| Ratio 50%   | ![](fixtures/test10x10_compressed50.jpeg)
| Ratio 80%   | ![](fixtures/test10x10_compressed80.jpeg)|
| Ratio 90%   | ![](fixtures/test10x10_compressed90.jpeg)|
| Ratio 99%   | ![](fixtures/test10x10_compressed99.jpeg)|

## Getting Started

To get a local copy up and running follow these simple steps.

### Installation

Install vendor
   ```sh
   go get -u github.com/nurlantulemisov/imagecompression
   ```

<!-- USAGE EXAMPLES -->

### Simple usage

```Go
package main

import (
   compression "github.com/nurlantulemisov/imagecompression"
   "image/png"
   "log"
   "os"
)

func main() {
   file, err := os.Open("examples/simple_usage/tmp/test.png")

   if err != nil {
      log.Fatalf(err.Error())
   }

   img, err := png.Decode(file)

   if err != nil {
      log.Fatalf(err.Error())
   }
   
   compressing, _ := compression.New(95)
   compressingImage := compressing.Compress(img)

   f, err := os.Create("examples/simple_usage/tmp/compressed-test.png")
   if err != nil {
      log.Fatalf("error creating file: %s", err)
   }
   defer func(f *os.File) {
      err := f.Close()
      if err != nil {
         log.Fatalf(err.Error())
      }
   }(f)

   err = png.Encode(f, compressingImage)
   if err != nil {
      log.Fatalf(err.Error())
   }
}

```

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Roadmap

1. Compression of grayscale images
2. Provide Frobenius norm


<!-- CONTACT -->

## Contact

Tulemisov Nurlan - [@NurlanTulemisov](https://twitter.com/NurlanTulemisov)

Email - nurlan.tulemisov@gmail.com

Project Link: [https://github.com/nurlantulemisov/imagecompression](https://github.com/nurlantulemisov/imagecompression)

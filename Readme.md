## Image resize with Golang Test

This is image resizing testing with golang. Just for benchmarking.

#### Installation

```
go get -u github.com/disintegration/imaging
```

#### Usage

```
go run process.go

// Or

go build process.go
./process
```

If you run this command, script will be resize the `eg.jpeg` image (1500x1000) from `img` folder and then save the resized images to `output` folder.

Testing resize image size are,

- 100x100
- 200x200
- 300x300
- 400x400
- 500x500

#### Credits

- Example image for [https://unsplash.com](https://unsplash.com)
- Go image package for [https://github.com/disintegration/imaging](https://github.com/disintegration/imaging)


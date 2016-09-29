package main

import (
    "fmt"
    "time"
    "image"
    "strconv"
    "image/color"
    "github.com/disintegration/imaging"
)


type Size struct {
    w int
    h int
}

var all []Size

func main() {
    start := time.Now()

    all := []Size{
        Size{
            w: 100,
            h: 100,
        },
        Size{
            w: 200,
            h: 200,
        },
        Size{
            w: 300,
            h: 300,
        },
        Size{
            w: 400,
            h: 400,
        },
        Size{
            w: 500,
            h: 500,
        },
    }

    // input file
    file := "img/eg.jpeg"

    img, err := imaging.Open(file)
    if err != nil {
       panic(err)
    }

    // Generate all resized image sizes to output folder
    for _, s := range all {
        r := imaging.Resize(img, s.w, s.h, imaging.Lanczos)
        dst := imaging.New(s.w, s.h, color.NRGBA{0, 0, 0, 0})
        dst = imaging.Paste(dst, r, image.Pt(0, 0))
        err := imaging.Save(dst, "output/dst" + strconv.Itoa(s.w) + strconv.Itoa(s.h) +".jpg")
        if err != nil {
            panic(err)
        }
    }

    elapsed := time.Since(start)

    fmt.Println("Resize timing millisecond : " + strconv.Itoa(int(elapsed.Nanoseconds()/1000000)));
}
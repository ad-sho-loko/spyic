# spyic

spyic is a utility for spying on your variable.

## Install

`go get -u github.com/ad-sho-loko/spyic`

## Quick Start

![Demo](etc/demo.gif)

```
package main

import (
	"github.com/ad-sho-loko/spyic"
	"time"
)

func main(){
	array := []int{2,1,4,0,5}

	spy := spyic.NewSlice(array)
	spy.SetDuration(time.Millisecond * 500)

	spy.Start()
	for i:=0; i<len(array)-1; i++{
		for j:=len(array)-1; j>i; j--{
			spy.Step()
			if array[j-1] < array[j]{
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	spy.Finish()
}
```

## Author
Shogo Arakawa

## License
MIT
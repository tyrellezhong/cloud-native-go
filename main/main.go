package main

import (
	_chan "cloud_native_go/chan"
	"fmt"
)

func main() {
	_chan.ChanCap()
	// _chan.ChanReadNil()
	// _chan.ChanWriteNil()
	// _chan.ChanReadClosed()
	_chan.ChanWriteClosed()
	fmt.Println("excute success")

}

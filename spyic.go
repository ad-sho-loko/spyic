package spyic

import (
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

const Version = "0.1"

const (
	DefaultFormat = "%v"
)

type Spy struct {
	target      fmt.Stringer
	ticker      *time.Ticker
	writer      io.Writer
	counter     uint64
	onPrinted   func()
	format      string
	ignoreCount uint64
}

func New(target fmt.Stringer) *Spy {
	return &Spy{
		target:      target,
		writer:      os.Stderr,
		counter:     0,
		format:      DefaultFormat,
		onPrinted:   func() {},
		ignoreCount: 0,
	}
}

func NewSlice(target interface{}) *Spy {
	s := SliceStringer{Slice: target}
	return New(s)
}

func (o *Spy) SetDuration(duration time.Duration) {
	o.ticker = time.NewTicker(duration)
}

func (o *Spy) SetWriter(writer io.Writer) {
	o.writer = writer
}

func (o *Spy) SetFormat(format string) {
	o.format = format
}

func (o *Spy) SetIgnore(n uint64) {
	o.ignoreCount = n
}

func (o *Spy) Start() {
	o.counter = 0
}

func (o *Spy) Step() {
	atomic.AddUint64(&o.counter, 1)

	if o.ignoreCount > atomic.LoadUint64(&o.counter) {
		return
	}

	<-o.ticker.C
	o.print()
}

func (o *Spy) print() {
	_, _ = fmt.Fprintf(o.writer, "\r"+o.format, o.target.String())

	// hook
	if o.onPrinted != nil {
		o.onPrinted()
	}
}

func (o *Spy) OnPrinted(f func()) {
	if f == nil {
		o.onPrinted = func() {}
		return
	}

	o.onPrinted = f
}

func (o *Spy) Finish() {
	_, _ = fmt.Fprint(o.writer, "\n")
	o.ticker.Stop()
}

//Package passio provides wrappers to interfaces Read and Write and
//count the number of bytes read or write and the number of operations done

package paasio

import (
	"io"
	"sync"
)

type counter struct {
	bytes      int64
	operations int
	mux        *sync.RWMutex
}

func newCounter() counter {
	return counter{bytes: 0, operations: 0, mux: new(sync.RWMutex)}

}

func (rc *counter) getValues() (int64, int) {
	rc.mux.RLock()
	defer rc.mux.RUnlock()
	return rc.bytes, rc.operations
}
func (rc *counter) setValues(b int64) {
	if b > 0 {
		rc.mux.Lock()
		rc.operations++
		rc.bytes += b
		rc.mux.Unlock()
	}
}

type readCount struct {
	counter
	reader io.Reader
}

//Create a new struct that implement ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCount{reader: r, counter: newCounter()}
}

//Returns number of bytes and number of operation
func (r *readCount) ReadCount() (n int64, nops int) {
	return r.getValues()
}

//read a slice of bytes from an buffer
func (r *readCount) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	r.setValues(int64(n))
	return
}

type writeCount struct {
	counter
	writer io.Writer
}

//Create a new struct that implement WriteCounter
func NewWriteCounter(r io.Writer) WriteCounter {
	return &writeCount{writer: r, counter: newCounter()}
}

//Returns number of bytes and number of operation
func (r *writeCount) WriteCount() (n int64, nops int) {
	return r.getValues()
}

//Write a slice of bytes on an buffer
func (r *writeCount) Write(p []byte) (n int, err error) {
	n, err = r.writer.Write(p)
	r.setValues(int64(n))
	return
}

type readWriteCount struct {
	ReadCounter
	WriteCounter
}

//Create a new struct that implement ReadWriteCounter
func NewReadWriteCounter(r io.ReadWriter) ReadWriteCounter {
	return readWriteCount{
		ReadCounter:  NewReadCounter(r),
		WriteCounter: NewWriteCounter(r)}
}

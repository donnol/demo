// 流处理示例
package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

// Seek
func handleUploadWithSeek(u io.Reader) (err error) {
	// capture all bytes from upload
	b, err := ioutil.ReadAll(u)
	if err != nil {
		return
	}

	// wrap the bytes in a ReadSeeker
	r := bytes.NewReader(b)

	// process the meta data
	err = processMetaData(r)
	if err != nil {
		return
	}

	// rewind the reader back to the start
	r.Seek(0, 0)

	// upload the data
	err = uploadFile(r)
	if err != nil {
		return
	}

	return nil
}

// file
func handleUploadWithFile(u io.Reader) (err error) {
	// create a temporary file for the upload
	f, err := ioutil.TempFile("", "upload")
	if err != nil {
		return
	}

	// destroy the file once done
	defer func() {
		n := f.Name()
		f.Close()
		os.Remove(n)
	}()

	// transfer the bytes to the file
	_, err = io.Copy(f, u)
	if err != nil {
		return
	}

	// rewind the file
	f.Seek(0, 0)

	// process the meta data
	err = processMetaData(f)
	if err != nil {
		return
	}

	// rewind the file again
	f.Seek(0, 0)

	// upload the file
	err = uploadFile(f)
	if err != nil {
		return
	}

	return nil
}

// MultiReader
func handleUploadWithMultiReader(u io.Reader) (err error) {
	// read in the first two bytes
	b := make([]byte, 2)
	_, err = u.Read(b)
	if err != nil {
		return
	}

	// check that they match the JPEG header
	jpg := []byte{0xFF, 0xD8}
	if !bytes.Equal(b, jpg) {
		return errors.New("not a JPEG")
	}

	// glue those bytes back onto the reader
	r := io.MultiReader(bytes.NewReader(b), u)

	// upload the file
	err = uploadFile(r)
	if err != nil {
		return
	}

	return nil
}

// TeeReader && Pipe
func handleUploadWithPipe(u io.Reader) {
	// create the pipe and tee reader
	pr, pw := io.Pipe()
	tr := io.TeeReader(u, pw)

	// create channel to synchronize
	done := make(chan bool)
	defer close(done)

	go func() {
		// close the PipeWriter after the
		// TeeReader completes to trigger EOF
		defer pw.Close()

		// upload the original MP4 data
		uploadFile(tr)

		done <- true
	}()

	go func() {
		// transcode to WebM
		webmr := transcode(pr)

		// upload to storage
		uploadFile(webmr)

		done <- true
	}()

	// wait until both are done
	for c := 0; c < 2; c++ {
		<-done
	}
}

// Multiwriter
func handleUploadWithMultiWriter(u io.Reader) {
	// create the pipes
	mp4R, mp4W := io.Pipe()
	webmR, webmW := io.Pipe()
	oggR, oggW := io.Pipe()
	wavR, wavW := io.Pipe()

	// create channel to synchronize
	done := make(chan bool)
	defer close(done)

	// spawn all the task goroutines. These look identical to
	// the TeeReader example, but pulled out into separate
	// methods for clarity
	go uploadMP4(mp4R, done)
	go transcodeAndUploadWebM(webmR, done)
	go transcodeAndUploadOgg(oggR, done)
	go transcodeAndUploadWav(wavR, done)

	go func() {
		// after completing the copy, we need to close
		// the PipeWriters to propagate the EOF to all
		// PipeReaders to avoid deadlock
		defer mp4W.Close()
		defer webmW.Close()
		defer oggW.Close()
		defer wavW.Close()

		// build the multiwriter for all the pipes
		mw := io.MultiWriter(mp4W, webmW, oggW, wavW)

		// copy the data into the multiwriter
		io.Copy(mw, u)
	}()

	// wait until all are done
	for c := 0; c < 4; c++ {
		<-done
	}
}

func processMetaData(r io.Reader) error {
	return nil
}

func uploadFile(r io.Reader) error {
	return nil
}

func transcode(pr *io.PipeReader) (r io.Reader) {
	return
}

func uploadMP4(mp4R *io.PipeReader, done chan bool) {
	done <- true
}

func transcodeAndUploadWebM(webmR *io.PipeReader, done chan bool) {

	done <- true
}
func transcodeAndUploadOgg(oggR *io.PipeReader, done chan bool) {

	done <- true
}
func transcodeAndUploadWav(wavR *io.PipeReader, done chan bool) {

	done <- true
}

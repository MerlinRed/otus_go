package main

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotFound          = errors.New("file not found")
	ErrOpeningFile           = errors.New("file could not be opened")
	ErrCreatingFile          = errors.New("file could not be created")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Stat(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	fromFileSize := fromFile.Size()
	if offset > fromFileSize {
		return ErrOffsetExceedsFileSize
	}

	file, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileNotFound
		}
		return ErrOpeningFile
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("file not closed")
		}
	}(file)

	if limit == 0 || limit > fromFileSize {
		limit = fromFileSize
	}

	file.Seek(offset, io.SeekStart)
	buf := bufio.NewReaderSize(file, int(fromFileSize))
	newFile, err := os.Create(toPath)
	if err != nil {
		return ErrCreatingFile
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			panic("file not closed")
		}
	}(newFile)

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(buf)
	_, err = io.CopyN(newFile, barReader, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	bar.Finish()

	return nil
}

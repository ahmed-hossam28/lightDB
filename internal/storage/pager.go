package storage

import (
	"fmt"
	"io"
	"os"
)

const (
	PageSize      = 4096 // Size of each page in bytes
	TableMaxPages = 100  // Maximum number of pages
)

type Pager struct {
	File       *os.File
	FileLength uint64
	Page       [TableMaxPages][]byte //Cache Pages holding rows (each page is a byte slice)
}

func PagerOpen(filename string) *Pager {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	fileSize, err := file.Seek(0, io.SeekEnd)
	pager := &Pager{
		File:       file,
		FileLength: uint64(fileSize),
	}

	return pager
}

func (pager *Pager) GetPage(pageNumber uint32) []byte {
	if pageNumber > TableMaxPages {
		fmt.Println("Page number out of range", pageNumber)
		os.Exit(3)
	}
	if pager.Page[pageNumber] == nil {
		//cache miss.Allocate Memory
		page := make([]byte, PageSize)
		numberOfPages := uint32(pager.FileLength / PageSize)
		if pager.FileLength%PageSize != 0 {
			numberOfPages++
		}

		if pageNumber < numberOfPages {
			offset := int64(pageNumber * PageSize)
			_, err := pager.File.Seek(offset, io.SeekStart)
			if err != nil {
				fmt.Printf("Failed to seek to page number %d: %s\n", pageNumber, err)
				os.Exit(3)
			}

			_, err = pager.File.Read(page)
			if err != nil {
				fmt.Printf("Failed to read page number %d: %s\n", pageNumber, err)
				os.Exit(3)
			}

		}
		pager.Page[pageNumber] = page
	}

	return pager.Page[pageNumber]
}

func (pager *Pager) Flush(pageNumber, size uint32) error {
	if pager.Page[pageNumber] == nil {
		return fmt.Errorf("page number %d is nil", pageNumber)
	}
	offset := int64(pageNumber * PageSize)
	_, err := pager.File.Seek(offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to page number %d: %s", pageNumber, err)
	}
	page := pager.Page[pageNumber]

	bytesWritten, err := pager.File.Write(page[:size])
	if err != nil {
		return fmt.Errorf("failed to write page number %d: %s", pageNumber, err)
	}
	if bytesWritten != int(size) {
		return fmt.Errorf("incomplete write: wrote %d of %d bytes", bytesWritten, size)
	}
	return nil
}

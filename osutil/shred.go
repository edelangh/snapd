package osutil

import (
	"os"
	"github.com/edsrzf/mmap-go"
)

func Shred(filePath string) error {
	defer os.Remove(filePath)

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	file_size := stat.Size()

	mmap, _ := mmap.Map(file, mmap.RDWR, 0 )
	defer mmap.Unmap()

	rand, err := os.Open("/dev/urandom")
	if err != nil {
		return err
	}
	defer rand.Close()

	for i := 0; i < 10 ; i++ {
		for count := int64(0); count < file_size; {
			readed, err := rand.Read(mmap[count:])
			if err != nil {
				return err
			}
			count += int64(readed)
		}
	}
	return nil
}

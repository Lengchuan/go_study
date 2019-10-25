package file

import "os"

func main() {

}

//部分读取文件
func ReadFileOfPart(filename string, offset int64, length int64) (data []byte, err error) {
	var f *os.File
	if f, err = os.OpenFile(filename, os.O_RDONLY, 0444); err != nil {
		return
	}

	defer f.Close()

	data = make([]byte, length)

	_, err = f.ReadAt(data, offset)

	return
}

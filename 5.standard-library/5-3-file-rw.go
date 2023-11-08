package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//test()
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	test10()
}

// 复制文件
func test() {
	oriFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer oriFile.Close()
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, oriFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

//跳转到文件指定位置
func test1() {
	file, _ := os.Open("test.txt")
	defer file.Close()
	var offset int64 = 5
	// 0 = 文件开始位置 1 = 当前位置 2 = 文件结尾处
	var whence int = 0
	newPosition, err := file.Seek(offset,whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPosition)
	newPosition ,_ = file.Seek(-2, 1)
	 fmt.Println("Just moved back two:", newPosition)
	 // 使用下面的技巧得到当前的位置
  currentPosition, err := file.Seek(0, 1)
  fmt.Println("Current position:", currentPosition)
	newPosition, err = file.Seek(0, 0)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Position after seeking 0,0:", newPosition)
}

//写文件
func test2() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := []byte("Bytes!")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", byteWritten)
}

//快写文件
func test3() {
	err := ioutil.WriteFile("test.txt", []byte("hi"),0666)
	if err != nil {
		log.Fatal(err)
	}
}

//使用缓存写
func test4() {
	file, err := os.OpenFile("test.txt",os.O_WRONLY,0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferWriter := bufio.NewWriter(file)
	bytesWritten, err := bufferWriter.Write([]byte{65,66,67})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written: %d\n", bytesWritten)
	bytesWritten, err = bufferWriter.WriteString("Buffered string\n",)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// 检查缓存中的字节数
	unflushedBufferSize := bufferWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// 还有多少字节可用（未使用的缓存大小）
  bytesAvailable := bufferWriter.Available()
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Available buffer: %d\n", bytesAvailable)
	// 写内存buffer到硬盘
	bufferWriter.Flush()
	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
  // 当你想将缓存传给另外一个writer时有用
  bufferWriter.Reset(bufferWriter)
  bytesAvailable = bufferWriter.Available()
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Available buffer: %d\n", bytesAvailable)
	bufferWriter = bufio.NewWriterSize(bufferWriter, 8000)
	bytesAvailable = bufferWriter.Available()
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Available buffer: %d\n", bytesAvailable)
}

//读取最多 N 个字节
func test5() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 从文件中读取len(b)字节的文件。
  // 返回0字节意味着读取到文件尾了
  // 读取到文件会返回io.EOF的error
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
    log.Fatal(err)
  }
  log.Printf("Number of bytes read: %d\n", bytesRead)
  log.Printf("Data read: %s\n", byteSlice)
}

//读取正好 N 个字节
func test6() {
	// Open file for reading
  file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  // file.Read()可以读取一个小文件到大的byte slice中，
  // 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
  byteSlice := make([]byte, 2)
  numBytesRead, err := io.ReadFull(file, byteSlice)
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Number of bytes read: %d\n", numBytesRead)
  log.Printf("Data read: %s\n", byteSlice)
}

//读取至少 N 个字节
func test7() {
	file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }
	byteSlice := make([]byte, 512)
	minByts := 8
	numBytesRead, err := io.ReadAtLeast(file,byteSlice,minByts)
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Number of bytes read: %d\n", numBytesRead)
  log.Printf("Data read: %s\n", byteSlice)
}

//读取全部字节
func test8() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(file)
	}
	// os.File.Read(), io.ReadFull() 和
  // io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
  // 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex: %x\n", data)
  fmt.Printf("Data as string: %s\n", data)
  fmt.Println("Number of bytes read:", len(data))
}

//快读到内存
func test9() {
	data, err := ioutil.ReadFile("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Data read: %s\n", data)
}

//使用缓存读
func test10() {
	 // 打开文件，创建buffered reader
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)
	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)
	// 读取，指针同时移动
  numBytesRead, err := bufferedReader.Read(byteSlice)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)
	
}
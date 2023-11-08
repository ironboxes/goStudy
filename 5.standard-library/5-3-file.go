package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//test()
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}

// 创建空文件
func test() {
	newFile, err := os.Create("test.txt")
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(newFile)
}

// 裁剪
func test1() {
	err := os.Truncate("test.txt",100)
	if err !=nil {
		log.Fatal(err)
	}
}

//文件信息
func test2() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file name: ", fileInfo.Name())
	fmt.Println("size in bytes: ", fileInfo.Size())
	fmt.Println("Permission: ", fileInfo.Mode())
	fmt.Println("Last modify ", fileInfo.ModTime())
	fmt.Println("Is directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T \n", fileInfo.Sys())
	fmt.Printf("System info: %+v \n", fileInfo.Sys())
}

//重命名和移动
func test3() {
	// oriPath := "test.txt"
	// newPath := "test1.txt"
	// err := os.Rename(oriPath, newPath)
	err := os.Remove("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
}

//打开和关闭文件
func test4() {
	 // 简单地以只读的方式打开。下面的例子会介绍读写的例子。
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    // OpenFile提供更多的选项。

    // 最后一个参数是权限模式permission mode

    // 第二个是打开时的属性

    file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)

    if err != nil {

        log.Fatal(err)

    }

    file.Close()

    // 下面的属性可以单独使用，也可以组合使用。

    // 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：

    // os.O\_CREATE|os.O\_APPEND

    // 或者 os.O\_CREATE|os.O\_TRUNC|os.O_WRONLY

    // os.O_RDONLY // 只读

    // os.O_WRONLY // 只写

    // os.O_RDWR // 读写

    // os.O_APPEND // 往文件中添建（Append）

    // os.O_CREATE // 如果文件不存在则先创建

    // os.O_TRUNC // 文件打开时裁剪文件

    // os.O\_EXCL // 和O\_CREATE一起使用，文件不能存在

    // os.O_SYNC // 以同步I/O的方式打开
}

// 检查文件是否存在
func test5() {
	file , err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}


  log.Println("File does exist. File information:")
  log.Println(file)
}

// 检查读写权限
func test6() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	// 测试读权限
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()


}
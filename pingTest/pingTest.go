package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

func create(buf *bufio.Scanner,mq chan <- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		// fmt.Println(line)
		mq <- line
	}
}

func GoPing(mq <- chan string, pg *sync.WaitGroup)  {
	defer pg.Done()
	for {
		if len(mq) == 0 {
			break
		}
		ip := <- mq
		cmd := exec.Command("ping", "-l", "1", ip)
		_, err := cmd.CombinedOutput()
		if err != nil {
			//fmt.Printf("%s %s \r\n", ip, "false")
		} else {
			fmt.Printf("%s %s \r\n", ip, "success")
		}
	}
}

func main() {
	//var phone Phone
	timeUnix := time.Now().Unix()
	mq := make(chan string, 500)
	var wg sync.WaitGroup
	var pg sync.WaitGroup


	fp, err := os.Open("C:\\Users\\aa\\go\\src\\goTools\\pingTest\\addr.txt")
	if err != nil {
		fmt.Println(err) //打开文件错误
		return
	}
	buf := bufio.NewScanner(fp)


	for i :=0; i < 4; i++ {
		time.Sleep(1*time.Millisecond)
		wg.Add(1)
		go create(buf, mq, &wg)
	}

	wg.Wait()
	close(mq)

	for i :=0; i < 70; i++ {
		pg.Add(1)
		go GoPing(mq, &pg)
	}
	pg.Wait()

	timeUnix2 := time.Now().Unix()
	fmt.Println(timeUnix2 - timeUnix)
}

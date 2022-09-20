package main

import (
	"fmt"
	"sync"
	"time"
)

var SliceSrk []string
var status string


func syncProcess(syncSlice *[]string, processName string) {

	if status != "ok" {
		// 添加信号到全局slice
		*syncSlice = append(*syncSlice, processName)
	}
	fmt.Printf("【%s】当前通道信号为:%v\n", processName, *syncSlice)
	// 临时slice，用于存储除自己信号之外的其他信号
	var tmpOtherSlice []string
	// 判断除自己之外是否还存在其他信号
	for true {
		if status == "ok" {
			*syncSlice = []string{}
			break
		}

		for _, v := range *syncSlice {
			if v != processName {
				tmpOtherSlice = append(tmpOtherSlice, v)
			}
		}
		if len(tmpOtherSlice) == 1 {
			break
		}
		fmt.Printf("【%s】请等待对方订单成交\n", processName)
		time.Sleep(5*time.Second)
	}

	// 清除（消费）除自己之外的其他信号
	if len(tmpOtherSlice) != 0 {
		fmt.Printf("【%s】对方为：%s\n", processName, tmpOtherSlice[0])
		*syncSlice = func(elem string) []string {
			index := 0
			for _, v := range *syncSlice {
				if v != elem {
					(*syncSlice)[index] = v
					index++
				}
			}
			return (*syncSlice)[:index]
		}(tmpOtherSlice[0])
	}

	// 判断自己的信号是否被消除（消费）
	for true {
		if status == "ok" {
			break
		}
		var tmpSelfSlice []string
		for _, v := range *syncSlice {
			if v == processName {
				tmpSelfSlice = append(tmpSelfSlice, v)
			}
		}
		if len(tmpSelfSlice) == 0 {
			break
		}
		fmt.Printf("【%s】请等待对方处理通信通道\n", processName)
		time.Sleep(5*time.Second)
	}
}

func test1(wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 0; i < 3; i++{
		fmt.Println("111111111:", SliceSrk)
		syncProcess(&SliceSrk, "test1")
	}
	status = "ok"
}

func test2(wg *sync.WaitGroup)  {
	defer wg.Done()
	time.Sleep(10*time.Second)
	fmt.Println("22222222:", SliceSrk)
	syncProcess(&SliceSrk, "test2")
	status = "ok"
}


func main()  {
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go test1(&wg)
	//go test2(&wg)
	//
	//wg.Wait()
	//fmt.Println(SliceSrk)

	for a := "sss"; a != "aaa"
	{
		fmt.Println("ok")
		time.Sleep(1*time.Second)
		a = "aaa"
	}
}

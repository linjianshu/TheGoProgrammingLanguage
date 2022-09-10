package main

import (
	"fmt"
	"path"
)

func main() {
	//path包处理文件路径
	dir := path.Dir("‪E:\\康明斯\\康明斯项目0820\\#20项目文件\\MES程序清单_20201120_V1.1_ByCZQ.xlsx")
	fmt.Println(dir)

	join := path.Join("E:", "康明斯", "康明斯0820", "#20项目文件", "MES程序清单_20201120_V1.1_ByCZQ.xlsx")
	fmt.Println(join)

	split, file := path.Split(join)
	fmt.Println(split)
	fmt.Println(file)

	base := path.Base("‪E:\\康明斯\\康明斯项目0820\\#20项目文件\\MES程序清单_20201120_V1.1_ByCZQ.xlsx")
	fmt.Println(base)

	abs := path.IsAbs("‪E:\\康明斯\\康明斯项目0820\\#20项目文件\\MES程序清单_20201120_V1.1_ByCZQ.xlsx")
	fmt.Println(abs)
}

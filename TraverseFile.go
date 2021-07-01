package main
import (
    "io/ioutil"
	"log"
	"strings"
)
func main(){
  dir:="D:"
  level:=1
  dir_Traversal(dir,level)
}

// 遍历的文件夹
func dir_Traversal(dir string, level int) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, v:= range fileinfo {
		// 重复输出制表符，层级结构
		print(strings.Repeat("\t", level))

		// 判断是不是目录
		if v.IsDir() {
			println(`目录：`, v.Name())
			findDir(dir+`/`+fi.Name(), level+1)
		} else {
			println(`文件：`, v.Name())
		}
	}
}

__________________________________________________________________________________________________________________________
__________________________________________________________________________________________________________________________

<?php
/**
 * 将读取到的目录以数组的形式展现出来
 * @return array
 * opendir() 函数打开一个目录句柄，可由 closedir()，readdir() 和 rewinddir() 使用。
 * is_dir() 函数检查指定的文件是否是目录。
 * readdir() 函数返回由 opendir() 打开的目录句柄中的条目。
 * @param array $files 所有的文件条目的存放数组
 * @param string $file 返回的文件条目
 * @param string $dir 文件的路径
 * @param resource $handle 打开的文件目录句柄
 */
function my_scandir($dir)
{
    //定义一个数组
    $files = array();
    //检测是否存在文件
    if (is_dir($dir)) {
        //打开目录
        if ($handle = opendir($dir)) {
            //返回当前文件的条目
            while (($file = readdir($handle)) !== false) {
                //去除特殊目录
                if ($file != "." && $file != "..") {
                    //判断子目录是否还存在子目录
                    if (is_dir($dir . "/" . $file)) {
                        //递归调用本函数，再次获取目录
                        $files[$file] = my_scandir($dir . "/" . $file);
                    } else {
                        //获取目录数组
                        $files[] = $dir . "/" . $file;
                    }
                }
            }
            //关闭文件夹
            closedir($handle);
            //返回文件夹数组
            return $files;
        }
    }
}

echo "<pre>";
print_r(my_scandir("D:/work_pro")); //电脑的文件路径即可


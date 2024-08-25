package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	// "strconv"
)


func dirTree(out io.Writer,pathF string, sliceForGrafity []bool) error{
	// открытие файла
	openDir, err := os.Open(pathF)
	if err != nil {	return err }
	defer openDir.Close()


	mainDir, err := openDir.ReadDir(0)
	if err != nil { return err}

	sort.Slice(mainDir, func(i, j int) bool {
    return mainDir[i].Name() < mainDir[j].Name()
	})

	//распределим файлы по слайсам
	var dirs []string
	var files []string

	for _, file := range mainDir{
		if file.Name() != ".DS_Store" {
			if file.IsDir(){
				dirs = append(dirs, file.Name())
			} else {
				info, _ := file.Info()
				if strconv.FormatInt(info.Size(), 10) == "0"{
					files = append(files, file.Name() + " (empty)")
				} else {
					files = append(files, file.Name() + " (" + strconv.FormatInt(info.Size(), 10) + "b)")
				}
			}
		}
	}


	for idx, dir := range dirs{
		if idx == len(dirs) - 1{
			sliceForGrafity = append(sliceForGrafity, true)
		} else {
			sliceForGrafity = append(sliceForGrafity, false)
		}
		if len(files) > 0 {
			sliceForGrafity[len(sliceForGrafity)-1] = false
		}
		graf, _ := multiTabs(sliceForGrafity)
		fmt.Fprint(out, graf, dir,"\n")

		dirTree(out, pathF + "/" + dir, sliceForGrafity)
		sliceForGrafity = sliceForGrafity[:len(sliceForGrafity)-1]
	}


	for idx, file := range files{
		if idx == len(files) - 1{
			sliceForGrafity = append(sliceForGrafity, true)
		} else {
			sliceForGrafity = append(sliceForGrafity, false)
		}
		graf, _ := multiTabs(sliceForGrafity)
		fmt.Fprint(out, graf, file,"\n")

		dirTree(out, pathF + "/" + file, sliceForGrafity)
		sliceForGrafity = sliceForGrafity[:len(sliceForGrafity)-1]
	}
	return nil
}


//вывод графики
func multiTabs(slice []bool) (string, []bool){
	var str string
	for idx, item := range slice{
		if idx == len(slice)-1 {
			if !item{
				str += "├───"
			} else {
				str += "└───"
			}
		} else {
			if !item{
				str += "│ "
			} else {
				str += "\t"
			}
		}
	}
	return str, slice
} 


func main(){
	path := "./testdata"
	if err := dirTree(os.Stdout,path, nil); err != nil{
		log.Println(err)
	}

}
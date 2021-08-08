package main
import (
    "fmt"
)
func main(){
     s := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}

    fmt.Println(removeDuplicateElement(s))
}

func removeDuplicateElement(arr []string) []string{
    res := make([]string,0,len(arr))
    tempMap :=map[string]struct{}{}
    for _,item := range arr {
        if _,ok := tempMap[item]; !ok{
            tempMap[item] = struct{}{}
            res = append(res, item)
        }
    }
    return res
}


//使用map去重，空struct不占用空间

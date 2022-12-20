package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// 集合内的元素类型必须一致

func main() {
	// 插入字符串，set去重
	requiredClasses := mapset.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")
	fmt.Println(requiredClasses)

	scienceSlice := []interface{}{"Biology", "Chemistry"}
	scienceClasses := mapset.NewSetFromSlice(scienceSlice)

	electiveClasses := mapset.NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := mapset.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	//Show me all the available classes I can take  集合取并集
	allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses)
	fmt.Println(allClasses) //Set{Cooking, English, Math, Chemistry, Welding, Biology, Music, Automotive, Go Programming, Python Programming}

	//Is cooking considered a science class?  集合是否包含
	fmt.Println(scienceClasses.Contains("Cooking")) //false

	//Show me all classes that are not science classes, since I hate science  集合取补集.
	fmt.Println(allClasses.Difference(scienceClasses)) //Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	//Which science classes are also required classes? 集合取交集
	fmt.Println(scienceClasses.Intersect(requiredClasses)) //Set{Biology}

	//How many bonus classes do you offer?  集合容量
	fmt.Println(bonusClasses.Cardinality()) //2

	//Do you have the following classes? Welding, Automotive and English?  集合是否包含子集
	fmt.Println(allClasses.IsSuperset(mapset.NewSetFromSlice([]interface{}{"Welding", "Automotive", "English"}))) //true
}

package main

import (
	"fmt"
)

func main() {
	var conf map[uint64][][]uint64

	conf = map[uint64][][]uint64{
		264: [][]uint64{
			[]uint64{102},
		},
		265: [][]uint64{
			[]uint64{102},
		},
		266: [][]uint64{
			[]uint64{102},
		},
		311: [][]uint64{
			[]uint64{102},
		},
		491: [][]uint64{
			[]uint64{102},
		},
		580: [][]uint64{
			[]uint64{102},
		},
		584: [][]uint64{
			[]uint64{102},
		},
		685: [][]uint64{
			[]uint64{102},
		},
		686: [][]uint64{
			[]uint64{102},
		},
		1915: [][]uint64{
			[]uint64{15},
			[]uint64{21},
			[]uint64{27},
			[]uint64{28},
			[]uint64{100},
			[]uint64{101},
		},
	}
	fmt.Printf("%T\n", conf)
	fmt.Printf("%v\n", conf)

	pgListMap := map[uint64]struct{}{}
	var wechatPgSet []uint64
	for _, spec := range conf {
		for _, pgList := range spec {
			//fmt.Println("pgList", pgList)
			for _, v := range pgList {
				//fmt.Println("pgList", v)
				if v == 102 || v == 21 || v == 106 {
					if _, ok := pgListMap[v]; !ok {
						//fmt.Println("----", v)
						pgListMap[v] = struct{}{}
						wechatPgSet = append(wechatPgSet, v)
					}
				}
			}
		}
	}

	fmt.Println(wechatPgSet)

}

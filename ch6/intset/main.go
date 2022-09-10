package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(9)
	fmt.Println("X:", x.String())
	fmt.Println(x.Has(1))

	var y IntSet
	y.Add(2)
	y.Add(9)
	fmt.Println("Y:", y.String())
	//并集
	x.UnionWith(&y)
	fmt.Println("并集", x.String())

	//交集
	x.Intersection(&y)
	fmt.Println("交集", x.String())

	//对称差集
	x.Add(1)
	y.Add(10)
	fmt.Println(y.String())
	x.SymmetricDifference(&y)
	fmt.Println(x.String())

}

//把具体的int值按照64分组 每一组中有64个0 代表一个数值 例如intSet[1]=2 即 64*1+1  1是( 2的2进制位数1 也就是2<<1)
type IntSet struct {
	words []uint64
}

func (i *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	v := i.words[word]
	if v == 0 {
		return false
	} else {
		return v&(1<<bit) != 0
	}
}

func (i *IntSet) UnionWith(t *IntSet) {
	for key, value := range t.words {
		if key >= len(i.words) {
			i.words = append(i.words, value)
		}
		//old := i.words[key]

		for j := 0; j < 64; j++ {
			//直接合并
			i.words[key] |= value

			////如果他们的j位 相或结果是1  就说明应该合并
			//if old&1<<j|value&1<<j == 1 {
			//	//那么就把 这一位弄成1
			//	old = old & 1 << j
			//}
		}
	}
}

func (i *IntSet) Intersection(t *IntSet) {
	for key, value := range t.words {
		if key >= len(i.words) {
			continue
		}

		i.words[key] &= value
	}
}

func (i *IntSet) SymmetricDifference(t *IntSet) {
	for key, value := range t.words {
		if key >= len(i.words) {
			i.words = append(i.words, value)
			continue
		}

		for j := 0; j < 64; j++ {
			old := i.words[key] & (1 << j)
			new := value & (1 << j)

			if old^new != 0 {
				//异或不为1的 就是结果 更新一下
				i.words[key] |= 1 << j
			} else if old != 0 {
				//把都有的 那一位置为0 使用异或
				i.words[key] ^= 1 << j
			}
		}
	}
}
func (i *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(i.words) {
		i.words = append(i.words, 0)
	}
	i.words[word] |= 1 << bit
}

func (i *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range i.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

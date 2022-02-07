package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	max_uint := int(math.MaxUint32)
	sum_int := int(x) + int(y)
	sum := x + y
	is_owerflow := false
	if sum_int> max_uint {
		is_owerflow= true

		return sum, is_owerflow
	}
	return sum, is_owerflow
}

func CeilNumber(f float64) float64 {
	division := float64(int(f / float64(0.25)))
	if division * 0.25 != f {
		return (division + 1) * 0.25
	}
	return f
}

func AlphabetSoup(s string) string {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	str := ""

	for _, char := range r {
		str += string(char)
	}

	return str
}

func StringMask(s string, n uint) string {

	r_string := ""

	if len(s) == 0 {
		if n >= 0 {
			r_string = "*"
		}
	}else if len(s) == 1 && n == 1 {
		r_string = "*"
	}else if int(n) >= len(s) {
		for i := 0; i < len(s); i++ {
			r_string += "*"
		}
	}else {
		for i := 0; i < int(n); i++ {
			r_string += string(s[i])
		}
		for i := int(n); i < len(s); i++ {
			r_string += "*"
		}
	}

	return r_string
}

func WordSplit(arr [2]string) string {
	words := strings.Split(arr[1], ",")
	finished := false
	len_str := 0
	result_str := ""
	found := [2]string{}
	i := 0
	for  i < len(words) && !finished{
		if strings.Contains(arr[0], words[i]) {
			len_str += len(words[i])
			if strings.Index(arr[0], words[i]) == 0{
				found[0] = words[i]
			}else {
				found[1] = words[i]
			}
		}
		if len_str == len(arr[0]) {
			finished = true
		}
		i += 1
	}
	if !finished || len(arr[0]) == 0 {
		return "not possible"
	}else{
		result_str += found[0] + "," + found[1]
	}

	return result_str
}

func VariadicSet(i ...interface{}) []interface{} {
/*
	result := []interface{}{}

	for _, val := range i {
		is_ok := true
		for i := 0; i < len(result); i++ {
			if result[i] == val{
				is_ok = false
			}
		}
		if is_ok {
			result = append(result, val)
		}

	}



	return result*/
	return nil
}

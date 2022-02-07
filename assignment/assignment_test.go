package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/

	cases := []struct{
		input uint32
		input2 uint32
		output uint32
		output2 bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1 , 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6,  0, true},
		{4294967290, 10, 4, true},
	}

	for _, v := range cases {
		t.Run("v.output", func(t *testing.T) {
			output, output2 := AddUint32(v.input, v.input2)
			assert.Equal(t, v.output, output)
			assert.Equal(t, v.output2, output2)
		})
	}

	//sum, overflow := AddUint32(math.MaxUint32, 1)

	//assert.Equal(t, uint32(0), sum)
	//assert.True(t, overflow)
}

func TestCeilNumber(t *testing.T) {
	 cases := []struct{
	 	input float64
	 	output float64
	 }{
	 	{42.42, 42.50},
	 	{42, 42},
	 	{42.01, 42.25},
	 	{42.24, 42.25},
	 	{42.25, 42.25},
	 	{42.26, 42.50},
	 	{42.55, 42.75},
	 	{42.75, 42.75},
	 	{42.76, 43},
	 	{42.99, 43},
	 	{43.13, 43.25},
	 }
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	for _, v := range cases {
		t.Run("v.output", func(t *testing.T) {
			output := CeilNumber(v.input)
			assert.Equal(t, v.output, output)
		})
	}

	//point := CeilNumber(42.42)

	//assert.Equal(t, 42.50, point)
}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	cases := []struct{
		input string
		output string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, v := range cases {
		t.Run("v.output", func(t *testing.T) {
			output := AlphabetSoup(v.input)
			assert.Equal(t, v.output, output)
		})
	}

	//result := AlphabetSoup("hello")

	//assert.Equal(t, "ehllo", result)
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	cases := []struct{
		input string
		input2 uint
		output string
	}{
		{"!mysecret*", 2,"!m********"},
		{"", 1111,"*"},
		{"a", 1,"*"},
		{"string", 0,"******"},
		{"string", 3,"str***"},
		{"string", 5,"strin*"},
		{"string", 6, "******"},
		{"string", 7,"******"},
		{"s*r*n*", 3,"s*r***"},
	}

	for _, v := range cases {
		t.Run("v.output", func(t *testing.T) {
			output := StringMask(v.input, v.input2)
			assert.Equal(t, v.output, output)
		})
	}

	//result := StringMask("!mysecret*", 2)

	//assert.Equal(t, "!m********", result)
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/

	cases := []struct{
		input [2]string
		output string
	}{
		{[2]string{"hellocat",words}, "hello,cat"},
		{[2]string{"catbat",words},"cat,bat"},
		{[2]string{"yellowapple",words},"yellow,apple"},
		{[2]string{"",words},"not possible"},
		{[2]string{"notcat",words},"not possible"},
		{[2]string{"bootcamprocks!",words},"not possible"},
	}

	for _, v := range cases {
		t.Run("v.output", func(t *testing.T) {
			output := WordSplit(v.input)
			assert.Equal(t, v.output, output)
		})
	}

	//result := WordSplit([2]string{"hellocat", words})

	//assert.Equal(t, "hello,cat", result)
}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
/*
	cases := []struct{
		input []interface{}
		output []interface{}
	}{
		{interface{4,2,5,4,2,4}, []interface{4,2,5}},
		{[2]string{"catbat",words},"cat,bat"},
		{[2]string{"yellowapple",words},"yellow,apple"},
		{[2]string{"",words},"not possible"},
		{[2]string{"notcat",words},"not possible"},
		{[2]string{"bootcamprocks!",words},"not possible"},
	}

 */
	set := VariadicSet(4, 2, 5, 4, 2, 4)

	assert.Equal(t, []interface{}{4, 2, 5}, set)
}

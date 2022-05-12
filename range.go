package main


//range iterates over elements in a variety of data structures. 
//how to use range with some of the data structures we’ve already learned.

//Here we use range to sum the numbers in a slice. Arrays work like this too. 

import "fmt"

func main() {

// range : Permission to repeat on the structures we have created
//Here we use "range" to add numbers to a section.
//Arrays work the same way.

nums := []int{2, 3, 4}

sum := 0

for _, num := range nums {

sum += num

//range on arrays and slices provides both the index and value for each entry.
// Above we didn’t need the index, so we ignored it with the blank identifier _.
// Sometimes we actually want the indexes though.

}

fmt.Println("sum:", sum)

//'range' on arrays and slices provides both the index and value for each entry.
 //Above we didn't need the index, so we ignored it with the blank identifier '_'. 
 //Sometimes we actually want the indexes though.
 
 for i, num := range nums {

	if num == 3 {

	fmt.Println("index:", i)

	}

	}

// 'range' on map iterates over key/value pairs.

kvs := map[string]string{"a": "apple", "b": "banana"}

 for k, v := range kvs {

 fmt.Printf("%s -> %s\n", k, v)

 }

 // 'range' can also iterate over just the keys of a map.
 //range on strings iterates over Unicode code points.
 // The first value is the starting byte index of the rune and the second the rune itself
 
	fmt.Println("key:", k)

	}

	for i, c := range "go" {

		fmt.Println(i, c)

   
	}
	
}


// output :
//sum: 9
//index: 1
//a -> apple
//b -> banana
//key: a
//key: b
//0 103
//1 111



// more details: 
//In the first example, the variable nums, we performed iterate on a slice.

//In fact, the range keyword has two outputs, in arrays and slice,
// the first is the index and the second is the value of the house.
// In the first example we have the index variable with Put the identifier
// blank or _ we said do not extract because if we extract and do not use the Go compiler will give an error.

//In the kvs variable we created a map with the initial value and iterated it, 
//but the difference between the output range in the map is that the first
// parameter is the key and the second The parameter is the value.
//For maps, we can only get one output as a key from the range.
//With the help of range we can also iterate on strings, here is the first parameter of the index.
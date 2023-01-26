package main

import "fmt"

func Swap(arr []int, i int) {
	//variable temporal para almacenar el valor de i
	var temp int

	temp = arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = temp
}

func BubbleSort(arr []int) {
	fmt.Println("The list have a length of: ")
	fmt.Println(len(arr))

	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				Swap(arr, i)
			}
		}
	}
	fmt.Println("Sorted list: ")
	fmt.Println(arr)
}

func main() {
	// El algoritmo tiene que pedir al usuario que ingrese 10 numeros y ordenarlos de forma ascendente
	var list = make([]int, 10)
	var value int

	for i := 0; i < len(list); i++ {
		fmt.Printf("Enter %v value: ", i+1)
		fmt.Scan(&value)
		list[i] = value
	}

	BubbleSort(list)

}

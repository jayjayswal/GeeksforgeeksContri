package main

import "fmt"

/*Function to left Rotate arr[] by 1*/
func leftRotatebyOne(arr []int) {
    temp := arr[0]
    var i int
    for i = 0; i < len(arr)-1; i++ {
        arr[i] = arr[i+1]
    }
    arr[i] = temp
}

/*Function to left rotate arr[] by d*/
func leftRotate(arr []int, d int) {
    for i := 0; i < d; i++ {
        leftRotatebyOne(arr)
    }
}

/* Driver program to test above functions */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println(arr)
    leftRotate(arr, 2)
    fmt.Println(arr)
}

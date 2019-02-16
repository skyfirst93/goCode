package main
import "fmt"
func main(){
	var n int
	var arr [10]int
	fmt.Println("enter the number of array elements")
	fmt.Scanln(&n)
	fmt.Println("Enter the elements")
	for i:=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Elements are ")
	for i:=0;i<n;i++ {
		fmt.Println(arr[i])
	}
	sum:=0
	fmt.Print("average of the number is ")
	for i:=0;i<n;i++{
		sum=sum+arr[i]
	}
	fmt.Println(sum/n)
}

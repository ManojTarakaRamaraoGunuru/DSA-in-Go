package main
import "fmt"

func merge(l,m,r int, arr[]int){
	var temp []int
	i, j := l,m+1
	for i<=m && j<=r{
		if arr[i]<=arr[j]{
			temp = append(temp, arr[i])
			i+=1
		}else{
			temp = append(temp, arr[j])
			j+=1
		}
	}

	for i<=m {
		temp = append(temp, arr[i])
		i+=1
	}

	for j<=r{
		temp = append(temp, arr[j])
		j+=1
	}

	for i,v := range temp{
		arr[l+i] = v
	}
}

func mergeSort(l, r int, arr[]int){
	if l >= r{
		return
	}
	m := (l+r)/2
	mergeSort(l,m,arr)
	mergeSort(m+1,r,arr)
	merge(l,m,r,arr)
}



func main(){
	arr := []int{1,5,7,9,3,4,-1,0,2,8}	
	mergeSort(0,len(arr)-1,arr)
	for _,v:= range(arr){
		fmt.Println(v)
	}
}
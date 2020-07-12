package service

func Checkwin(a [16][16]int,m int,n int)(flag bool){
	total1 := context(a,m,n,1,0)
	total2 := context(a,m,n,0,1)
	total3 := context(a,m,n,1,1)
	total4 := context(a,m,n,1,-1)
	if total1 >= 5 || total2 >= 5 || total3 >= 5 || total4 >= 5{
		return true
	}else{
		return false
	}
}

func context(a [16][16]int,m int,n int,i int,j int)(total int){
	var min,max int
	if m > n {
		min = n
		max = m
	}else{
		min = m
		max = n
	}
	for t := 0;t <= min;t++{
		if a[m - i * t][n - j * t] == a[m][n] {
			total++
		}else{
			break
		}
	}
	for t := 1;t <= 15 - max;t++{
		if a[m + i * t][n + j * t] == a[m][n] {
			total++
		}else{
			break
		}
	}
	return total
}
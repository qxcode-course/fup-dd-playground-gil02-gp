package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    conta := 1
    maxqtd := 1

    for i := 1; i < n; i++ {
        if arr[i] == arr[i-1] {
            conta++
        } else {
            if conta > maxqtd {
                maxqtd = conta
            }
            conta = 1
        }
    }
    if conta > maxqtd {
        maxqtd = conta
    }

    fmt.Print("[ ")
    conta = 1
    for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			conta++
		} else {
			if conta == maxqtd {
				fmt.Print(arr[i-1], " ")
			}
			conta = 1
		}
	}
    if conta == maxqtd {
        fmt.Print(arr[n-1], " ")
    }
    fmt.Print("]\n")
}
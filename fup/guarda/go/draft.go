package main
import "fmt"

func main() {
    var wifi, login, admin int
    fmt.Scan(&wifi)
    fmt.Scan(&login)
    fmt.Scan(&admin)

    if(wifi == 0){
        fmt.Printf("you must connect to wifi\n")
    } else if(login == 0){
        fmt.Printf("you need to login first\n")
    }else if(admin == 0){
        fmt.Printf("you must to login as admin\n")
    } else{
        fmt.Printf("done\n")
    }
}

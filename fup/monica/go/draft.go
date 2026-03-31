package main
import "fmt"
func main() {
    var M, A, B int
    fmt.Scan(&M, &A, &B)

    C := M - (A + B)
    mais_velhos := A
    if B > A && B > C {
        mais_velhos = B
    } else if C > A && C > B {
        mais_velhos = C
    }
    fmt.Printf("%d\n", mais_velhos)
}

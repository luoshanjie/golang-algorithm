![banner](docs/images/golang-banner.jpg)
# golang-algorithm
无论是面试还是提高自身水平，练习算法都是一项比较给力的方式，值得我们下功夫。当前Github上练习的题目中大部分是单goroutine的，并发处理的很少。但在实际工作中这又很常见，为了避免脱离现实，所以下面的练习中将会并发融入其中

## 第一道开胃菜
**问题**: 2个Goroutine依次打印，一个打印数字，一个打印字母<br>
**结果**: 1A2B3C4D5E6F7G8H9I10J11K12L13M14N15O16P17Q18R19S20T21U22V23W24X25Y26Z
```go
package main

import (
    "fmt"
    "sync"
)

func number(wg *sync.WaitGroup, c1, c2 chan string) {
    defer wg.Done()

    for i := 0; i < 26; i++ {
        <- c1
        fmt.Printf("%d", i+1)
        c2 <- "continue"
    }
}

func letter(wg *sync.WaitGroup, c1, c2 chan string) {
    defer wg.Done()

    items := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for i := 0; i < 26; i++ {
        <- c1
        fmt.Printf("%c", items[i])
        c2 <- "continue"
    }
}

func main() {
    chA, chB := make(chan string, 1), make(chan string, 1)
    var wg sync.WaitGroup
    wg.Add(2)

    chA <- "start"
    go number(&wg, chA, chB)
    go letter(&wg, chB, chA)

    wg.Wait()
}
```
## 第二道开胃菜
**问题**: 3个Goroutine循环5次，按次序分别打印A,B,C<br>
**结果**:
```
goroutine(1) print A
goroutine(2) print B
goroutine(3) print C
...
goroutine(1) print A
goroutine(2) print B
goroutine(3) print C
```

```go
package main

import (
    "fmt"
    "sync"
)

const count = 5

func doing(wg *sync.WaitGroup, c1, c2 chan string, gid, text string) {
    defer wg.Done()

    for i := 0; i < count; i++ {
        <- c1
        fmt.Printf("gorutine[%s] print %s\n", gid, text)
        c2 <- "continue"
    }
}

func main() {
    chA, chB, chC := make(chan string, 1), make(chan string, 1), make(chan string, 1)

    var wg sync.WaitGroup
    wg.Add(3)

    chA <- "start"
    go doing(&wg, chA, chB, "1", "A")
    go doing(&wg, chB, chC, "2", "B")
    go doing(&wg, chC, chA, "3", "C")

    wg.Wait()
}
```

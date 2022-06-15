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
## 第三道开胃菜
**问题**: n个Goroutine循环依次打印从1到m的数<br>
**结果**: 这里预设4个Goroutine打印从1到17的数
```
goroutine(1) print 1
goroutine(2) print 2
goroutine(3) print 3
goroutine(4) print 4
goroutine(1) print 5
goroutine(2) print 6
goroutine(3) print 7
goroutine(4) print 8
goroutine(1) print 9
goroutine(2) print 10
goroutine(3) print 11
goroutine(4) print 12
goroutine(1) print 13
goroutine(2) print 14
goroutine(3) print 15
goroutine(4) print 16
goroutine(1) print 17
```

```go
package main

import (
    "fmt"
    "sync"
)

const (
    total = 17
    num   = 4
)

func doing(wg *sync.WaitGroup, c1, c2 chan int, gid int) {
    defer wg.Done()

    for {
        value := <-c1
        if value <= total {
            fmt.Printf("goroutine(%d) print %d\n", gid, value)
            value++
            c2 <- value
        } else {
            c2 <- value
            break
        }
    }
}

func main() {
    channels := make([]chan int, 0)
    for i := 0; i < num; i++ {
        ch := make(chan int, 1)
        channels = append(channels, ch)
    }

    var wg sync.WaitGroup
    wg.Add(num)
    channels[0] <- 1
    for i := 0; i < num; i++ {
        if i == num-1 {
            go doing(&wg, channels[i], channels[0], i+1)
        } else {
            go doing(&wg, channels[i], channels[i+1], i+1)
        }
    }

    wg.Wait()
}
```
这背后的理论依据是：**CSP(Communicating Sequential Process)**，又名通信顺序进程。这种并发编程模型用于描述两个独立的并发实体，通过共享的通讯channel(管道)进行通信。

## 正菜开始
### 排序
- [冒泡](internal/sort/bubble)

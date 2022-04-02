<!--
 * @Author: NorthCity1984
 * @LastEditTime: 2022-04-02 14:30:28
 * @Description: 
 * @Website: https://grimoire.cn
 * Copyright (c) NorthCity1984 All rights reserved.
-->
# stl-go

一个封装了一些简单的数据结构的提供泛型支持的stl库 (go >= 1.18)

# usage

## Dequeue

```go
package main

import (
	"fmt"

	"gitee.com/NorthCityChen/stl-go/dequeue"
)

func main() {
	q := dequeue.Init[int]()
	q.LPush(12)
	q.LPush(23)
	q.LPush(34)
	q.Clear()
	q.LPush(34)
	q.RPush(44)
	fmt.Println(q.LPop())
	fmt.Println(q.LPop())
	fmt.Println(q.LPop())
	fmt.Println(q.LPop())
}
```

## Queue

```go
import (
	"fmt"

	"gitee.com/NorthCityChen/stl-go/queue"
)

func main() {
	q := queue.Init[int]()
	q.Push(12)
	q.Push(23)
	q.Push(44)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

```

# RoadMap

- [x] 双端队列
- [x] 队列
- [x] 栈
- [x] 数学
- [ ] 优先队列
- [ ] 红黑树
- [ ] 集合
- [ ] 二分查找
- [ ] 下一个排序组合
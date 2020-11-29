# 学习笔记
```go
// dfs模板
res := returnValue
var dfs func(...)
dfs = func(...) {
    if returnCondition {
        res = append(res, ...)
    }
    thisLevelDo
    dfs(...)
}
dfs(.)
return returnValue
```

```go
// bfs模板
res := returnValue
queue := []type{firstElem}
for len(queue) != 0 {
    size := len(queue)
    levelValueInit
    for i := 0; i < size; i++ {
        curElem := queue[0]
        queue[1:]
        DoSomething with curElem {
            queue = append(queue, ...)
        }
    }
    levelValueHandler
}
return returnValue
```

```go
// 二分
left, right := 0, len(list) - 1
for left <= right {
    mid := left + (right - left) / 2
    if mid == target {
        return mid
    } else mid < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}
```

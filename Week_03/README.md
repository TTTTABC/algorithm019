# 学习笔记
```go
// 模板
result = []
var dfs(路径，选择列表)
func dfs(路径，选择列表) {
	if 满足结束条件 {
        result = append(result,路径...) 
	}
	return

	for 选择 in 选择列表 {
		做选择
		dfs(路径，选择列表)
		撤销选择
	}
}
dfs()
return result
```

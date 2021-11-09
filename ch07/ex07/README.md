# ex07

ヘルプメッセージの表示には `flag.Value.String()` が使用されており、
`celsiusFlag.String()` は `"%g°C"` という形式で値を表示するため。

```go
// String implements the flag.Value
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```
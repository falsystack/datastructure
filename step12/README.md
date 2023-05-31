# Graph
- NodeとEdgeで構成された全てのデータ構造をGraphと呼ぶ
- Linked list, TreeもGraphの一種である。
- 全てびRelationshipはGraphで表せる。
  - Network
  - Geographical
  - 人間関係
  - 等々

## Edge
- 片方向(Directional Edge)
- 双方向(Non-Directional Edge)

```go
type Node[T any] struct {
	value T
	nodes []*Node[T]
}
```
構造は簡単だがそれにより派生される問題は複雑だ。

## Cycle Detection
- ノード上Cycle（循環）が発生するかを検査する
- Floyd's Cycle Finding Algorithm
- Linkが多数ある場合はDFSで探す

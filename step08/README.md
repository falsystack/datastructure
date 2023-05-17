# Sort

## Quick Sort
- Worst : O(n^2), BEST : O(Nlog2N)
- 一番多く使われるソートアルゴリズム
- Pivot（基準点）を中心として左には小さい値を右には大きい値を配置
  - ソートされるまで繰り返し
## Merge Sort
- Worst : O(Nlog2N), BEST : O(Nlog2N)
  - BigO法でみると同じO(Nlog2N)だがquick sortが早い
  - だがMergeSortはWorstでもO(Nlog2N)
- フォン・ノイマンが開発した古典的なアルゴリズム
- Divide & Conquerのフォーマル
  - 並列環境でいい
- 半分ずつ分けて順番に合うようにマージする
## Binary Insert Sort
- とんでもなく遅い O(N^2)
- 追加される時は O(logN)
- すでにソートされている配列に値を挿入する時に流用なソート。
- 値が常に追加されるデータ構造に流用。
- Sorted Mapで使用される。

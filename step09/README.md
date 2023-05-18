# Map
- key-valueの形で保存するデータ構造
- keyを用いて照会、追加、削除、巡回することができる。
- Mapは大きく二つに分けられる
  - Unordered Map
    - Hash Map
  - Ordered Map
    - Sorted Map

## Hash Map
- Goの基本MapもHash Mapである
- Hash関数を基準として保存
- 順番がない
- 速度が速い
- Sparse(<-> Dense)なデータ構造
  - Cashと親しくない

### Hash Mapの速度
- 追加: `O(1)`
- 照会: `O(1)`
- 削除: `O(1)`

### Hash関数
- 細かく砕いてまとめる。
- データの検証、暗号化、仮想貨幣等々広い領域で使用される。
- 片方向関数 -> 逆関数が存在しない
  - `f(x) = x % 2`
- 損失圧縮(代表的な損失圧縮；`JPG`)
  - 複合できない。
- 簡単なCheckSum用度で使用される関数
  - `crc32.ChecksumIEEE([]byte(key))`
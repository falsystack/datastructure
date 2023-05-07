# Buffer
## 緩衝の役割
- FIFO, BufferはQueueの構造
- あるデバイスから別のデバイスにデータを送信するときに発生する時間の差やデータフローの速度差を調整するために一時的にデータを記憶するデバイス
- 2つのモジュール間でデータを送受信する際の時間差によって一時データを保存しておくストア

## GoでBufferを作る方法
- Slice buffer
  - Slice bufferの問題点は配列の長さがしょっちゅう変わるのでデータを書き換えるコストが掛かる
- Ring buffer
  - Ringとは一番後ろと一番前がつながるデータ構造
  - 固定された長さをもつ配列を使う
    - `WritePoint - ReadPoint`
  - 固定された配列の長さを超えた場合の式
    - `配列サイズ - ReadPoint + WritePoint`
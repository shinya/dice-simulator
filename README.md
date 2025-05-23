# サイコロシミュレーター

このプログラムは、n個のサイコロを振った時の目の合計数の分布とゾロ目の出現確率を計算します。

## 機能

- サイコロの数をコマンドライン引数として受け取ります
- 各合計値が出る組み合わせの数（通り数）を計算します
- ゾロ目（全てのサイコロの目が同じ）が出る確率を計算します
- 結果を見やすい形式で表示します
- マルチスレッド・マルチコア処理による高速な計算を実現します

## 使用方法

### Go環境がある場合

```
go run main.go <サイコロの数>
```

例：
```
go run main.go 2
```

### Dockerを使用する場合

#### Dockerイメージを直接使用

```
# イメージをビルド
docker build -t dice-simulator .

# 実行（例：サイコロ2個の場合）
docker run dice-simulator 2
```

#### Docker Composeを使用

```
# サイコロ2個の場合（デフォルト）
docker-compose up

# サイコロの数を指定する場合（例：サイコロ3個）
DICE_COUNT=3 docker-compose up
```

## 出力例

サイコロを2個振った場合：
```
サイコロを2個振った時の情報：
- 2: 1通り（2.78%）
- 3: 2通り（5.56%）
- 4: 3通り（8.33%）
- 5: 4通り（11.11%）
- 6: 5通り（13.89%）
- 7: 6通り（16.67%）
- 8: 5通り（13.89%）
- 9: 4通り（11.11%）
- 10: 3通り（8.33%）
- 11: 2通り（5.56%）
- 12: 1通り（2.78%）

ゾロ目が出る確率： 1/6（16.67%）
```

## パフォーマンス

このプログラムはマルチスレッド・マルチコア処理を活用して計算効率を高めています：

- 利用可能なCPUコア数を自動検出して最大限活用
- サイコロの数が多い場合に並列処理を実行（6つの並列処理を使用）
- 小さな入力（サイコロ2個以下）では順次処理を使用（オーバーヘッド回避）
- 処理時間の計測と表示

### パフォーマンス例

- サイコロ5個: 約300マイクロ秒
- サイコロ10個: 約170ミリ秒

通常の再帰的な実装と比較して、特に大きな入力に対して大幅なパフォーマンス向上を実現しています。

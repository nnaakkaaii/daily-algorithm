# 0000 : puzzle

15パズルは1つの空白を含む4x4のマス上に15枚のパネルが配置され、空白を使ってパネルを上下左右にスライドさせ、絵柄を揃えるパズル。

次のように空白を0、各パネルを1から15の番号でパズルを表す。

```
 1  2  3  4
 6  7  8  0
 5 10 11 12
 9 13 14 15
```

1回の操作で空白の方向に1つのパネルを移動することができ、ゴールは次のようなパネルの配置とする。

```
 1  2  3  4
 5  6  7  8
 9 10 11 12
13 14 15  0
```

15パズルの初期状態が与えられるので、ゴールまでの最短手数を求めるプログラムを作成すること。

## 入力

パネルの数字あるいは空白を表す4x4個の整数。空白で区切られた4つの整数が4行で与えられる。

## 出力

最短手数を1行に出力すること。

## 実行例

```shell
$ go run main.go -search simple-search -list priority-queue < data/1.in
```

```shell
$ go run main.go -search iterative-deepening -list stack < data/2.in
```

## 関連キーワード

* [A<sup>*</sup>](../docs/a-star/README.md)
* [IDA<sup>*</sup>](../docs/ida-star/README.md)
* [Iterative Deepening](../docs/iterative-deepening/README.md)

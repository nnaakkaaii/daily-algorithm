# 0000 : puzzle / pkg

Searchインターフェースを次の4つの構造体で実装する。

(ただし、デフォルトでヒューリスティックに基づく枝刈りを行う)

* SimpleSearch : データ構造から取り出したノードを次に探索する
* IterativeDeepening : 深さの制限を増やしながらSimpleSearchを行う

さらに、これら構造体の引数に与えるデータ構造であるListインターフェースを次の3つの構造体で実装する

* Stack
* Queue
* PriorityQueue

これらの探索手法とデータ構造の組み合わせにより、以下のアルゴリズムが達成される

| 探索手法               | Stack                | Queue      | PriorityQueue        |
|--------------------|----------------------|------------|----------------------|
| SimpleSearch       | (DFS) (limitまで止まらない) | BFS        | A<sup>*</sup>        |
| IterativeDeepening | IDA<sup>*</sup>      | (BFS) (無駄) | (A<sup>*</sup>) (無駄) |

## TODO

* 深化を線形探索で行っているが、二分探索で実装できないか

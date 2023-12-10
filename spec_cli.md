# ルール

標準的なオセロのルールに従う
- 初期配置
- 黒から始める
- ひっくり返す場所にしか石を置けない
- 石を置けない場合は Pass
- 両者石を置けない場合はゲームが終了

# 勝利判定

ゲーム途中で Give up したら、 Give up した Player の負け
ゲームが終了したときに自分の石が多い方が勝ち

# UI

単一の CLI での実行とする

## 通常時

- o: 白
- x: 黒

```
Player 1:
   1 2 3 4 5 6 7 8
A: . . . . . . . .
B: . . . . . . . .
C: . . . . . . . .
D: . . . o x . . .
E: . . . x o . . .
F: . . . . . . . .
G: . . . . . . . .
H: . . . . . . . .

[0] Put Stone
[1] Give Up
```

0 を押したら、
```
Input the point to put your disk:
(ex. E5)
```

正しい入力であれば相手のターンになり、不適切な入力であれば、
```
Error: <不適切である理由>
Please try again.

Input the point to put your disk:
(ex. E5)
```
を表示する。

## 石が置けないとき

```
Player 2:
. . (盤面) . .

[0] Pass
[1] Give Up
```

## ゲーム終了時

```
. . (盤面) . .
Player 1 WIN! (もしくは Drow)

[0] Continue
[1] End
```

1 を押した場合
```
Shut down the game...
```
を表示して終了する

## (Advanced) 途中で中断する

memo: non-blocking IO を使う必要があり面倒そう

Ctrl-C を押した場合、
```
Shut down the game...
```
を表示して終了する。

(Advanced) 中断時にデータを保存し、次回立ち上げ時に途中から始めるか最初から始めるかを選択可能にする。

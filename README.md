# go-monkey-Interpreter
Implement code that "Writing An Interpreter In Go" by Thorsten Ball

## Abstract

[Go言語でつくるインタプリタ](https://www.oreilly.co.jp/books/9784873118222/) の実装です。

## Monkey Language Syntax

今回実装するmonkey言語のsyntax.

```monkey
# 代入
let x = 5 + "5";

# 関数
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
}

let result = add(five, ten)
```


## Memo

* 字句解析
  * ソースコード --(字句解析)--> トークン列 ---(構文解析)--> 抽象構文木
  * 字句解析＝lexer
* Monkeylangによる式のsyntax
  * -5
  * !true
  * !false
  * 5+5
  * foo == bar
  * ((5+5) * 5) * 5
  * add(add(2,3), add(5, 10))
  * foo * bar /foobar
  * let add = fn(x, y) {return x+y};
  * 関数呼び出し
    * fn(x,y){return x+y}(5, 5)
  * if式がある
    * let result = if (10 > 5) {true} else {false};
* 用語（p50）
  * 前置演算子（prefix operator）--5
  * 後置演算子(postfix operator) foobar++
  * 中間演算子(infix operator) 5*8
  * 優先順 5+5*10

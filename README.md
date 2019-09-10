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

## Directory layout





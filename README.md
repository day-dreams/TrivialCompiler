整个编译器玩玩。

* [参考](https://www.freecodecamp.org/news/write-a-compiler-in-go-quick-guide-30d2f33ac6e0/)

* [gocc](https://github.com/goccmack/gocc)

* 基本思路
    1. 定义语法，形成bnf文件
    2. 用`gocc`生成lexer和parser
    3. 定义AST
    4. 借助lexer和parse，实现从源代码到AST的转换
    5. *可选项* 实现一个类型检查器，检查源代码转换成的AST是否合法
    6. 实现一个代码生成器，根据AST生成目标代码
    7. 去执行目标代码

* 嫌6 7麻烦，略过，直接写代码去解释AST。然后就变成了解释器。以后有需要了，再来丰富这个玩具。


* usage
```bash
make test
make launch

$ 1+2/3;
(1 + (2 / 3)) = 1.66666667
$ 12345+6789/10;16500*17;
(12345 + (6789 / 10)) = 13023.90000000
(16500 * 17) = 280500.00000000
$ 17750*15;
(17750 * 15) = 266250.00000000
$ 17750*17;
(17750 * 17) = 301750.00000000
$
```

* another usage

```bash
make test
make build
make install
tcompiler -src=./test/gostruct.txt # 生成user.go和user.proto文件
```
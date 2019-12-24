整个编辑器玩玩。

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

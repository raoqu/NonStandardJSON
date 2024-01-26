### 环境配置

根据ANTLR4 [官方文档](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md) 的指引进行安装（推荐）。

或者下载antlr4的jar包： [下载链接](https://www.antlr.org/download.html)

```
antlr4='java -Xmx500M -cp "/Users/raoqu/mylab/_lib/antlr-4.12.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
```

### 语法测试

```shell
 antlr4-parse NJson.g4 "{ \"a\":123}" -tree
```

### parser 生成

```shell
go get github.com/antlr4-go/antlr/v4

antlr4 -Dlanguage=Go NJson.g4
```

如果生成的parser代码有报错，需要将依赖 `github.com/antlr/antlr4/runtime/Go/antlr/v4` 修改为 `github.com/antlr4-go/antlr/v4`。

其他可能生成报错的代码参考以下方法修改：

修改前：

```Go
localctx = NewObjectContext(p.GetParserRuleContext(), p.GetState())
```

修改后：

```Go
localctx = NewObjectContext(p.BaseParser, p.GetParserRuleContext(), p.GetState())
```

### 常见问题

1. 最常见的问题是包路径问题 <p>生成的代码是 `github.com/antlr/antlr4/runtime/Go/antlr/v4`，而实际使用过程中需要修改成：`github.com/antlr4-go/antlr/v4`</p>
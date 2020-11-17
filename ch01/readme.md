# 第1章 命令行工具

## 1.1 准备工作
1. 安装java jdk
2. 安装go
3. 创建目录结构

## 1.2 java命令
JVM的工作是运行java应用程序。那么java应用程序也需要一个入口点，这个入口点就是我们所熟知的main()方法，这个类就是我们用来启动java应用程序，我们称之为主类。

那么jvm怎么知道我们要从哪个类启动应用程序呢。对此并没有明确规定。这由jvm自行决定。

### java命令有如下4中形式
```
java [-options] class [args]
java [-options] -jar jarfile [args]
javaw [-options] class [args]
javaw [-options] -jar jarfile [args]
```
可以向java命令传递三组参数：选项、主类名（或者JAR文件名） 和main()方法参数。选项由减号（–）开头。
通常，第一个非选项参数给出主类的完全限定名（fully qualified class name）。但是如果用户提供了–jar选项，则第一个非选项参数表示JAR文件名，java命令必须从这个JAR文件中寻找主类。
javaw命令和java命令几乎一样，唯一的差别在于，javaw命令不显示命令行窗口，因此特别适合用于启动GUI（图形用户界面）应用程序。

>[完整java命令用法请参考](http://docs.oracle.com/javase/8/docs/technotes/tools/windows/java.html)

## 1.3 编写命令行工具
在Java语言中，API一般以类库的形式提供。在Go语言中，API 则是以包（package）的形式提供。包可以向用户提供常量、变量、结构体以及函数等。

Java内置了丰富的类库，Go也同样内置了功能强大的包。本章将用到fmt、os和flag包。
- os包定义了一个Args变量，其中存放传递给命令行的全部参数。
- 如果直接处理os.Args变量，需要写很多代码。还好Go语言内置了flag包，这个包可以帮助我们处理命令行选项。有了flag包，我们的工作就简单了很多。

### go的一些特性
1. Go源文件一般以.go作为后缀，文件名全部小写，多个单词之间用下划线分隔。Go语言规范要求Go源文件必须使用UTF-8编码，详见https://golang.org/ref/spec。
2. Go语言有函数（Function）和方法（Method）之分，方法调用需要receiver，函数调用则不需要。

## 1.4 测试本章代码
main.go

```
go install ./ch01
```

然后就可以在go的workspace下bin路径下找到相应的可执行文件

```
➜  bin ch01 -h  
Usage: ch01 [-options] class [args...]
```

```
➜  bin ch01 -v  
version 0.0.1 by ymk 2020
```

```
➜  bin ch01 -cp ./test className -xms
classpath:[./test] class:[className] args:[-xms]
```

## 1.5 本章小结

学习了java命令的基本用法，并且编写了一个简化版的命令行工具。
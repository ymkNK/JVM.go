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
可以向java命令传递三组参数：选项、主类名（或者JAR文件名） 和main（）方法参数。选项由减号（–）开头。
通常，第一个非选项参数 给出主类的完全限定名（fully qualified class name）。但是如果用户 提供了–jar选项，则第一个非选项参数表示JAR文件名，java命令必 须从这个JAR文件中寻找主类。
javaw命令和java命令几乎一样，唯 一的差别在于，javaw命令不显示命令行窗口，因此特别适合用于启 动GUI（图形用户界面）应用程序。

>[完整java命令用法请参考](http://docs.oracle.com/javase/8/docs/technotes/tools/windows/java.html)
# 第2章 搜索class文件

ch01介绍了java命令的用法以及如何启动java应用程序：
- 首先启动java虚拟机
- 加载主类
- 加载主类的main方法

但是我们知道，哪怕是最简单的hello world也是无法独自运行的

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, world!");
    }
}
```

- 在加载HelloWorld类之前，我们要加载它的超类，也就是`java.lang.Object`.
- 在调用main方法之前，因为虚拟机需要准备好参数数组，因此需要加载 `java.lang.String` 和`java.lang.String[]`类。
- 把字符打印到控制台，还需要加载`java.lang.System`类 

本章节，就是学习和实现如何去找到这些类

## 2.1 类路径
Java虚拟机规范并没有规定虚拟机应该从哪里寻找类，因此不同的虚拟机实现可以采用不同的方法。Oracle的Java虚拟机实现根据类路径（classpath）来搜索类。按照搜索的先后顺序，类路径可以分为以下3个部分：
- 启动类路径（bootstrap classpath）
- 扩展类路径（extension classpath）
- 用户类路径（user classpath）

　　启动类路径默认对应jre\lib目录，Java标准库（大部分在rt.jar里）位于该路径。

　　扩展类路径默认对应jre\lib\ext目录，使用Java扩展机制的类位于这个路径。我们自己实现的类，以及第三方类库则位于用户类路径。可以通过-Xbootclasspath选项修改启动类路径，不过通常并不需要这样做，所以这里就不详细介绍了。

　　用户类路径的默认值是当前目录，也就是“.”。可以设置CLASSPATH环境变量来修改用户类路径，但是这样做不够灵活，所以不推荐使用。更好的办法是给java命令传递-classpath（或简写为cp）选项。-classpath/-cp选项的优先级更高，可以覆盖CLASSPATH环境变量设置。第1章简单介绍过这个选项，这里再详细解释一下。

　　-classpath/-cp选项既可以指定目录，也可以指定JAR文件或者ZIP文件
```
java -cp path\to\classes ... 
java -cp path\to\lib1.jar ... 
java -cp path\to\lib2.zip ...
```

　　还可以同时指定多个目录或文件，用分隔符分开即可。分隔符因操作系统而异。在Windows系统下是分号，在类UNIX（包括Linux、MacOSX等）系统下是冒号。

```
java -cp path\to\classes ... 
java -cp path\to\lib1.jar ... 
java -cp path\to\lib2.zip ...
```

　　从Java6开始，还可以使用通配符（*）指定某个目录下的所有JAR文件，格式如下：
```
java -cp classes;lib\* ...
```  

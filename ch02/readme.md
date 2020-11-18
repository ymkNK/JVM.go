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

## 2.2 准备工作

将第一章的代码复制一份，增加一个classpath路径。

- 增加一个非标准选项-Xjre
- 更新相应代码

## 2.3 实现类路径
可以把类路径想象成一个大的整体，它由启动类路径、扩展类路径和用户类路径三个小路径构成。三个小路径又分别由更小的路径构成。是不是很像组合模式（composite pattern）？没错，本节就套用组合模式来设计和实现类路径。

### 2.3.1 Entry接口

在ch02\classpath\entry.go文件

- 常量pathListSeparator是string类型，存放路径分隔符
- readClass()方法负责寻找和加载class文件
- String()相当于Java中的toString()
- newEntry()函数根据参数创建不同类型的Entry实例（有四种）

readClass()方法的参数是class文件的相对路径，路径之间用斜线（/）分隔，文件名有.class后缀。比如要读取java.lang.Object类，传入的参数应该是java/lang/Object.class。返回值是读取到的字节数据、最终定位到class文件的Entry，以及错误信息。Go的函数或方法允许返回多个值，按照惯例，可以使用最后一个返回值作为错误信息。

### 2.3.2 DirEntry  
目录形式的类路径
- 只有一个字段，用于存放目录的绝对路径

### 2.3.3 ZipEntry
用于表示ZIP或者JAR文件心事的类路径

### 2.3.4 CompositeEntry
如前所述，CompositeEntry由更小的Entry组成，正好可以表示成[]Entry。
- 在Go语言中，数组属于比较低层的数据结构，很少直接使用。大部分情况下，使用更便利的slice类型。
- 构造函数把参数（路径列表）按分隔符分成小路径，然后把每个小路径都转换成具体的Entry实例
- readClass():依次调用每一个子路径的readClass方法，如果成功读取到class数据，返回数据即可;如果收到错误信息则继续;如果遍历结束还没有找到class文件，则返回错误。

### 2.3.5 WildcardEntry
WildcardEntry实际上也是CompositeEntry，所以就不再定义新的类型了
- 首先把路径末尾的星号去掉，得到baseDir，然后调用filepath包的Walk（）函数遍历baseDir创建ZipEntry。
- Walk（）函数的第二个参数也是一个函数，了解函数式编程的读者应该一眼就可以认出这种用法（即函数可作为参数）。
- 在walkFn中，根据后缀名选出JAR文件，并且返回SkipDir跳过子目录（通配符类路径不能递归匹配子目录下的JAR文件）。

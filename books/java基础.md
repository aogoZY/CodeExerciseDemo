
## 概念

 - 对象：对象是类的一个实例，有状态和行为。例如，一条狗是一个对象，它的状态有：颜色、名字、品种；行为有：摇尾巴、叫、吃等。
 - 类：类是一个模板，它描述一类对象的行为和状态。
 - 方法：方法就是行为，一个类可以有很多方法。逻辑运算、数据修改以及所有动作都是在方法中完成的。
 - 实例变量：每个对象都有独特的实例变量，对象的状态由这些实例变量的值决定。

##  hello world

```
public class HelloWorld {
    /* 第一个Java程序
     * 它将输出字符串 Hello World
     */
    public static void main(String[] args) {
        System.out.println("Hello World"); // 输出 Hello World
    }
}
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20210421203114754.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3pob3Vib2tl,size_16,color_FFFFFF,t_70)

## 基本语法

编写 Java 程序时，应注意以下几点：

- **大小写敏感**：Java 是大小写敏感的，这就意味着标识符 Hello 与 hello 是不同的。
- **类名**：对于所有的类来说，类名的首字母应该大写。如果类名由若干单词组成，那么每个单词的首字母应该大写，例如 **MyFirstJavaClass** 。
- **方法名**：所有的方法名都应该以小写字母开头。如果方法名含有若干单词，则后面的每个单词首字母大写。
- **源文件名**：源文件名必须和类名相同。当保存文件的时候，你应该使用类名作为文件名保存（切记 Java 是大小写敏感的），文件名的后缀为 **.java**。（如果文件名和类名不相同则会导致编译错误）。
- **主方法入口**：所有的 Java 程序由 **public static void main(String[] args)** 方法开始执行。

## Java修饰符

像其他语言一样，Java可以使用修饰符来修饰类中方法和属性。主要有两类修饰符：

- 访问控制修饰符 : default, public , protected, private
- 非访问控制修饰符 : final, abstract, static, synchronized

## Java 变量

Java 中主要有如下几种类型的变量

- 局部变量
- 类变量（静态变量）
- 成员变量（非静态变量）

## 构造方法

每个类都有构造方法。如果没有显式地为类定义构造方法，Java 编译器将会为该类提供一个默认构造方法。

在创建一个对象的时候，至少要调用一个构造方法。构造方法的名称必须与类同名，一个类可以有多个构造方法。

下面是一个构造方法示例：

```
public class Puppy{
    public Puppy(){
    }
 
    public Puppy(String name){
        // 这个构造器仅有一个参数：name
    }
}
```

## 创建对象

对象是根据类创建的。在Java中，使用关键字 new 来创建一个新的对象。创建对象需要以下三步：

- **声明**：声明一个对象，包括对象名称和对象类型。
- **实例化**：使用关键字 new 来创建一个对象。
- **初始化**：使用 new 创建对象时，会调用构造方法初始化对象。

## Employee.java 文件代码：

```
package test;

import org.omg.Messaging.SYNC_WITH_TRANSPORT;

public class Employ {
    String name;
    int age;
    String destination;
    double salary;

    // Employee 类的构造器
    public Employ(String name) {
        this.name = name;
    }

    // 设置age的值
    public void setAge(int age) {
        this.age = age;
    }

    /* 设置salary的值*/
    public void setSalary(double salary) {
        this.salary = salary;
    }

    /* 设置designation的值*/
    public void setDestination(String destination) {
        this.destination = destination;
    }

    /* 打印信息 */
    public void printEmploy(){
        System.out.println("名字"+this.name);
        System.out.println("年龄"+this.age);
        System.out.println("薪水"+this.salary);
        System.out.println("职称"+this.destination);
    }
}

```

## EmployeeTest.java 文件代码：

```
package test;

public class EmployeeTest {
    public static void main(String[] args) {
        /* 使用构造器创建两个对象 */
        Employ employ1 = new Employ("aogo");
        Employ employ2 = new Employ("amao");

        // 调用这两个对象的成员方法
        employ1.setAge(35);
        employ1.setDestination("高级程序员");
        employ1.setSalary(10000);
        employ1.printEmploy();
        System.out.println();

        employ2.setAge(20);
        employ2.setDestination("菜鸟程序员");
        employ2.setSalary(500);

        employ2.printEmploy();
    }
}

```

```
名字aogo
年龄35
薪水10000.0
职称高级程序员

名字amao
年龄20
薪水500.0
职称菜鸟程序员
```

# Java 变量类型

Java语言支持的变量类型有：

- 类变量：独立于方法之外的变量，用 static 修饰。
- 实例变量：独立于方法之外的变量，不过没有 static 修饰。
- 局部变量：类的方法中的变量。

```
public class Variable {

    static int alllicks;    // 类变量
    String str ="hello";    // 实例变量

    public void methodOne(){
        int i = 0;          // 局部变量
    }
}
```


## 1、Go语言的特色：

- 没有继承多态的面向对象
- 强一致类型
- interface不需要显式声明(Duck Typing)
- 没有异常处理(Error is value)
- 基于首字母的可访问特性
- 不用的import或者变量引起编译错误
- 完整而卓越的标准库包
- Go内置runtime（作用是性能监控、垃圾回收等）
- 部署方便
- 简单的并发
- 稳定性--Go拥有强大的编译检查、严格的编码规范和完整的软件生命周期工具，具有很强的稳定性，稳定压倒一切。Go提供了软件生命周期（开发、测试、部署、维护等等）的各个环节的工具，如go tool、gofmt、go test。

Go语言中面向对象编程的核心是组合和方法(function)

java、go、python不同语言之间的区别

|      | golang | java | python |
| ---- | ------ | ---- | ------ |
| 类型 | 静态   | 静态 | 动态   |
|      |        |      |        |
|      |        |      |        |

java类型

```java
package com.example.offer;

//抽象出来的人
abstract class human {
    protected String sex;
    protected String name;

    public void setSex(String sex) {
        this.sex = sex;
    }
    public String getSex() {
        return this.sex;
    }
    public void setName(String name) {
        this.name = name;
    }
    public String getName() {
        return this.name;
    }
    abstract void ShakeHand(); //抽象的方法

}

//学习接口
interface study {
    public abstract void learnEnglish();
}

//具象的男性，继承抽象类实现study接口
class Man extends human implements study {
    public Man() {
        this.sex = "男";
    }
    //    实现的方法
    public void ShakeHand() {
        System.out.println(this.name + this.sex + "左");
    }
    //    实现的接口
    public void learnEnglish() {
        System.out.println(this.name + ":how are u");
    }
}

class Woman extends human {
    public Woman() {
        this.sex = "女";
    }
    public void ShakeHand() {
        System.out.println(this.name + this.sex + "右");
    }
    public void learnEnglish() {
        System.out.println(this.name + ":i am fine,thank u");
    }
}

public class Main {
    public static void main(String[] args) {
        Man lilei = new Man();
        lilei.setName("李磊");
        System.out.println(lilei.getName() + " " + lilei.getSex() + " " + "出场");

        Woman hanmeimei = new Woman();
        hanmeimei.setName("韩梅梅");
        System.out.println(hanmeimei.getName() + " " + hanmeimei.getSex() + " " + "出场");

        lilei.ShakeHand();
        hanmeimei.ShakeHand();

        lilei.learnEnglish();
        hanmeimei.learnEnglish();
    }
}
```

python类型

```python
#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Human:
def __init__(self):
    self.__name = ""
    self.__sex = ""

def setName(self, name):
    self.__name = name

def getName(self):
    return self.__name

def setSex(self, sex):
    self.__sex = sex

def getSex(self):
    return self.__sex

name = property(getName, setName) # 就像java中的POJO
sex = property(getSex, setSex) # 就像java中的POJO


if __name__ == '__main__':

lilei = Human()
lilei.sex = "男"
lilei.name = "李磊"
print "%s %s 出场" % (lilei.name, lilei.sex)

hanmeimei = Human()
hanmeimei.sex = "女"
hanmeimei.name = "韩梅梅"
print "%s %s 出场" % (hanmeimei.name, hanmeimei.sex)


# Pee的方法
def shakeHand(self, how):
    print "%s %s %s撒尿" % (self.name, self.sex, how)

Human.shankeHand = shakeHand #动态绑定方法

lilei.doPee("站着")
hanmeimei.doPee("蹲着")

# 学习的方法
def doLearning(self, learn):
    print "%s: %s" % (self.name, learn)

Human.doLearning = doLearning #动态绑定方法

lilei.doLearning("How are you?")
lilei.doLearning("I'm fine, thank you!")
```



go

```go
package main

import (
	"fmt"
)

// 接口 学生
type Student interface {
	learningEnglish(string)
}

// 结构
type Human struct {
	Name string
	Sex  string
}

// 学习英语方法，绑定于Human
func (student Human) learningEnglish(learning string) {
	fmt.Printf("%s: %s\n", student.Name, learning)
}

// 结构 男人
// go没有继承这个概念，这里是嵌入
type Man struct {
	Human "嵌入字段"
}

type Woman struct {
	Human
}

// 方法, 绑定到了Human结构
func (this Human) Pee(how string) {
	fmt.Printf("%s %s %s撒尿\n", this.Name, this.Sex, how)
}

// 学习
func doLearning(learning Student, learing string) {
	learning.learningEnglish(learing)
}

// Pee
func doPee(human interface {}) {
	switch sex := human.(type){
	case Man:
		sex.Pee("站着")
	case Woman:
		sex.Pee("蹲着")
	}
}

func main() {
	lilei := Man{}
	lilei.Name = "李雷"
	lilei.Sex = "男"
	fmt.Printf("%s %s 出场\n", lilei.Name, lilei.Sex)

	hanmeimei := Woman{}
	hanmeimei.Name = "韩梅梅"
	hanmeimei.Sex = "女"
	fmt.Printf("%s %s 出场\n", hanmeimei.Name, hanmeimei.Sex)

	doPee(lilei)
	doPee(hanmeimei)

	doLearning(lilei, "How are you?")
	doLearning(hanmeimei, "I'm fine, thank you!")
}
```

 
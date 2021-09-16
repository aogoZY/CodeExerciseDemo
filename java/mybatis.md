mybatis

### 一、概述

1、框架介绍
框架是我们使用软件开发中的一套解决方案，不同的框架能解决不同的问题，在框架中封装了很多的细节，使开发者可以使用极为简便的方式实现功能，大大的提升了开发的效率。

2、三层架构
表现层：用于展现数据
业务层：用于处理业务需求
持久层：用于和数据库交互
3、MyBatis 框架简介
mybatis是一个优秀的基于 java 的持久层框架，它内部封装了 jdbc，使开发者只需要关注 sql语句本身， 而不需要花费精力去处理加载驱动、创建连接、创建 statement 等繁杂的过程，它使用了ORM思想实现了结果集的封装。

ORM：Object Relational Mapping（对象关系映射），即：把数据库表和实体类及实体类的属性对应起来，让开发人员可以操作实体类就可以实现对数据库表的操作

### 二、IntelliJ IDEA使用MyBatis框架1、mybatis的环境搭建（查询）

#### 【1】前期准备

在搭建之前先创建了mybatis数据库，并创建了user表，填入 id、username、birthday、sex、address 字段相关的数据

![image-20210913080726859](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913080726859.png)

```
DROP TABLE IF EXISTS user;
 
CREATE TABLE user
(
	id BIGINT(20) AUTO_INCREMENT  NOT NULL COMMENT '主键ID',
	username VARCHAR(30) NULL DEFAULT NULL COMMENT '姓名',
	birthday datetime DEFAULT NULL COMMENT '生日',
	gender INT  DEFAULT NULL COMMENT '性别',
	address VARCHAR(50) NULL DEFAULT NULL COMMENT '地址',

	PRIMARY KEY (id)
);
```



#### 【2】创建 maven 工程

![image-20210913082242488](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913082242488.png)

全部创建好后目录结构如下：

![image-20210913083250651](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913083250651.png)

#### 【3】导入坐标

在 pom.xml 文件中添加 Mybatis3.4.5 坐标和相关坐标

```html
<dependencies>
        <!--导入mybatis的jar包-->
        <dependency>
            <groupId>org.mybatis</groupId>
            <artifactId>mybatis</artifactId>
            <version>3.4.5</version>
        </dependency>
        <!--导入SQL相关jar包-->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>5.1.32</version>
        </dependency>
        <!--导入日志相关jar包-->
        <dependency>
            <groupId>log4j</groupId>
            <artifactId>log4j</artifactId>
            <version>1.2.17</version>
        </dependency>
        <!--导入测试相关jar包-->
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <version>4.12</version>
        </dependency>
    </dependencies>
```

#### 【4】编写 User 实体类

```java
public class User {
    private Long id;
    private String userName;
    private Date birthday;
    private Integer gender;
    private String address;
 		@Override
    public String toString() {
        return "User" + "id:" + id + ", username:" + userName + ", birthday:" + birthday + " ,gender" + gender + ", address" + address;
    }
}
```

####  【5】编写持久层接口 IUserDao

```java
public interface IUserDao {
    //    查询所有
    List<User> queryAll();
}
```

#### 【6】编写持久层接口的映射文件 IUserDao.xml

- 创建位置：必须和持久层接口在相同的包中
- 名称：必须以持久层接口名称命名，拓展名为 .xml
- 需要加上resultType 返回类型

```html
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">

<mapper namespace="cn.aogo.dao.IUserDao">
<!--    配置查询所有操作-->
    <select id ="queryAll" resultType="cn.aogo.domain.User">
        select * from user
    </select>
</mapper>

```

#### 【7】创建 SqlMapConifg.xml 配置文件

```html
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-config.dtd">
<!--mybatis的主配置文件-->
<configuration>
    <!--环境配置-->
    <environments default="mysql">
        <!--配置mysql的环境-->
        <environment id="mysql">
            <!--配置事务的类型-->
            <transactionManager type="JDBC"></transactionManager>
            <!--配置数据源（连接池）-->
            <dataSource type="POOLED">
                <!--配置数据库连接的四个基本信息-->
                <property name="driver" value="com.mysql.jdbc.Driver"/>
                <property name="url" value="jdbc:mysql://localhost:3306/mybatis"/>
                <property name="username" value="root"/>
                <property name="password" value="123456"/>
            </dataSource>
        </environment>
    </environments>
    <!--指定映射配置文件的位置，映射配置文件指的是每个dao独立的配置文件-->
    <mappers>
        <mapper resource="cn.aogo.dao.IUserDao.xml"/>
    </mappers>
</configuration>
```

#### 【8】测试类

Mybatis 在测试类中使用的模式：

SqlSessionFactory factory = builder.build(in);
构建者模式：把对象的创建细节隐藏，使用者直接调用方法即可拿到对象
factory.openSession();
工厂模式：降低了类之间的依赖关系
session.getMapper(IUserDao.class);
代理模式：不改变源码的基础上对已有的方法进行增强

##### 1、查询所有

```
public class mybatisTest {
    @Test
    public void test() throws IOException {
//        1、读取配置文件
        InputStream in = Resources.getResourceAsStream("SqlMapConfig.xml");
//        2、创建sqlSessionFactory的构建对象
        SqlSessionFactoryBuilder factoryBuilder = new SqlSessionFactoryBuilder();
//        3、使用构造者创建工厂对象 sqlSessionFactory
        SqlSessionFactory factory = factoryBuilder.build(in);
//        4、使用selSessionFactory生产 sqlSession 对象
        SqlSession sqlsession = factory.openSession();
//        5、使用sqlSession创建dao接口的代理对象
        IUserDao iUserDao = sqlsession.getMapper(IUserDao.class);

//        6、使用代理对象查询执行所有的方法
        List<User> users = iUserDao.queryAll();
        for (User user : users) {
            System.out.println(user);
        }

//        7、释放资源
        in.close();
    }
}
```



##### 2、插入数据

持久层添加insert方法

```
public interface IUserDao {
    //    查询所有
    List<User> queryAll();

    //    插入数据
    void insertUser(User user);
}
```





### 踩坑记

##### 1、A query was run and no Result Maps were found for the Mapped Statement

![image-20210915084145216](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210915084145216.png)

```
<mapper namespace="cn.aogo.dao.IUserDao">
<!--    配置查询所有操作-->
    <select id ="queryAll" resultType="cn.aogo.domain.User">
        select * from user
    </select>
</mapper>
```



##### 2、com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException

pom.xm将mysql-connector-java换为8.0.23：

![image-20210915084540213](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210915084540213.png)

```
      <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>8.0.23</version>
        </dependency>
```



##### 3、Could not find resource com/XXX/dao/IUserDao.xml



```
直接IUserDao.xml将放在resources文件夹下，同时修改为制定配置为
    <!--指定映射配置文件的位置，映射配置文件指的是每个dao独立的配置文件-->
    <mappers>
        <mapper resource="IUserDao.xml"/>
    </mappers>

```

![image-20210915084744099](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210915084744099.png)

https://blog.csdn.net/u010648555/article/details/70880425

##### 4、log4j:WARN Please initialize the log4j system properly.

![image-20210916074613247](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210916074613247.png)

日志未能正确输出，log4j.properties加上如下内容

```
log4j.rootLogger=debug, stdout, R

log4j.appender.stdout=org.apache.log4j.ConsoleAppender
log4j.appender.stdout.layout=org.apache.log4j.PatternLayout

log4j.appender.stdout.layout.ConversionPattern=%5p - %m%n

log4j.appender.R=org.apache.log4j.RollingFileAppender
log4j.appender.R.File=firestorm.log

log4j.appender.R.MaxFileSize=100KB
log4j.appender.R.MaxBackupIndex=1

log4j.appender.R.layout=org.apache.log4j.PatternLayout
log4j.appender.R.layout.ConversionPattern=%p %t %c - %m%n

log4j.logger.com.codefutures=DEBUG
```

##### 5、mybatis自动提交失败，Setting autocommit to false on JDBC Connection]

打印sql语句显示update有值但是db并未真正插入

![image-20210916083202351](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210916083202351.png)

Mybatis默认情况下开启的是手动提交， 我们需要开启自动提交，改动如下：

```
//        4、使用selSessionFactory生产 sqlSession 对象
        SqlSession sqlsession = factory.openSession(true);
```

![image-20210916083453026](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210916083453026.png)

6、插入空指针

```
    @Before
    public void init() throws IOException {
//        1、读取配置文件
         in = Resources.getResourceAsStream("SqlMapConfig.xml");
//        2、创建sqlSessionFactory的构建对象
        SqlSessionFactoryBuilder factoryBuilder = new SqlSessionFactoryBuilder();
//        3、使用构造者创建工厂对象 sqlSessionFactory
        SqlSessionFactory factory = factoryBuilder.build(in);
//        4、使用selSessionFactory生产 sqlSession 对象
         sqlSession = factory.openSession();
//        5、使用sqlSession创建dao接口的代理对象
         userDao = sqlSession.getMapper(IUserDao.class);

    }
```



对比一下

```
public class mybatisTest {
    private InputStream in;
    private SqlSession sqlSession;
    private IUserDao userDao;

    @Before
    public void init() throws IOException {
//        1、读取配置文件
        InputStream in = Resources.getResourceAsStream("SqlMapConfig.xml");
//        2、创建sqlSessionFactory的构建对象
        SqlSessionFactoryBuilder factoryBuilder = new SqlSessionFactoryBuilder();
//        3、使用构造者创建工厂对象 sqlSessionFactory
        SqlSessionFactory factory = factoryBuilder.build(in);
//        4、使用selSessionFactory生产 sqlSession 对象
        SqlSession sqlsession = factory.openSession();
//        5、使用sqlSession创建dao接口的代理对象
        IUserDao userDao = sqlsession.getMapper(IUserDao.class);
        sqlsession.commit();
        System.out.println("aaaa");
    }
    }
```


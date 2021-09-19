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

##### 【1】持久层添加insert方法

```
public interface IUserDao {
    //    查询所有
    List<User> queryAll();

    //    插入数据
    void insertUser(User user);
}
```

##### 【2】配置映射配置文件

```html
    <!--    插入数据-->
    <insert id="insertUser" parameterType="cn.aogo.domain.User">
 insert  into  user(userName,birthday,gender,address) value (#{userName},#{birthday},#{gender},#{address})
    </insert>
```

##### 【3】添加测试方法

![image-20210917070459623](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917070459623.png)

```java
package cn.aogo.test;

import cn.aogo.dao.IUserDao;
import cn.aogo.domain.User;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;


import java.io.IOException;
import java.io.InputStream;
import java.util.Date;
import java.util.List;

public class mybatisTest {
    private InputStream in;
    private SqlSession sqlSession;
    private IUserDao userDao;

    @Before
    public void init() throws IOException {
//        1、读取配置文件
        in = Resources.getResourceAsStream("SqlMapConfig.xml");
//        2、创建sqlSessionFactory的构建对象
        SqlSessionFactoryBuilder factoryBuilder = new SqlSessionFactoryBuilder();
//        3、使用构造者创建工厂对象 sqlSessionFactory
        SqlSessionFactory factory = factoryBuilder.build(in);
//        4、使用selSessionFactory生产 sqlSession 对象
        sqlSession = factory.openSession(true);
//        5、使用sqlSession创建dao接口的代理对象
        userDao = sqlSession.getMapper(IUserDao.class);

    }

    @After
    public void destroy() throws Exception {
        sqlSession.close();
        in.close();
    }


    @Test
    public void queryAll() throws IOException {
        List<User> users = userDao.queryAll();
        for (User user : users) {
            System.out.println(user);
        }

    }

    @Test
    public void addTest() throws IOException {
        User user = new User();
        user.setUserName("aogo");
        user.setGender(1);
        user.setBirthday(new Date());
        user.setAddress("广州");

        Integer affected = userDao.insertUser(user);
        System.out.println("affected:" + affected);

    }
}

```

##### 3、更新数据

##### 【1】持久层添加update方法

```
   //    更新数据
    Integer updateUser(User user);
```

##### 【2】配置update的映射配置文件

```
  <!--    更新数据-->
    <insert id="updateUser" parameterType="cn.aogo.domain.User">
            update user set userName =#{userName},birthday =#{birthday},gender =#{gender},address =#{address} where id =#{id}
    </insert>
```

##### 【3】添加update测试方法

```
 //    这里相当于是全量覆盖 如果某列没有set值是会被置为null的
    @Test
    public void updateTest() throws IOException {
        User user = new User();
        user.setUserName("amao");
        user.setAddress("北京");
        user.setBirthday(new Date());
        user.setId(23L);
        user.setGender(1);

        Integer affected = userDao.updateUser(user);
        System.out.println(affected);
    }
```

##### 4、删除数据

##### 【1】持久层添加delete方法

```
   //    删除数据
    Integer deleteUser(Integer id);
```

##### 【2】配置delete的映射配置文件

```
    <!--    删除数据-->
    <delete id="deleteUser" parameterType="Integer">
            delete from user where id = #{id}
    </delete>
```

##### 【3】添加delete测试方法

```
 @Test
    public void deleteTest() throws IOException {
        Integer affected = userDao.deleteUser(22);
        System.out.println(affected);
    }
```

##### 5、模糊查询

##### 【1】持久层添加findByName方法

```
    //    根据名字模糊查询
    List<User> findByName(String name);
```

##### 【2】配置findByName的映射配置文件

```
   <!--    根据名字模糊查询-->
    <select id="findByName" parameterType="String" resultType="cn.aogo.domain.User">
        select  * from user where userName like #{userName}
    </select>
```

##### 【3】添加findByName测试方法

```
  //    根据名称模糊查询
    @Test
    public void findByNameTest() {
        List<User> userList = userDao.findByName("%星%");
        for (User user : userList) {
            System.out.println(user);
        }
    }
```

##### 6、queryVO查询

##### 【1】创建vo类

```
package cn.aogo.domain;

import java.io.Serializable;

public class QueryVo implements Serializable {
    private User user;

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
```

##### 【2】持久层添加queryByVo的方法

```
    //    根据vo查询
    List<User> queryByVo(QueryVo vo);
```

##### 【3】配置queryByVo的映射配置文件

```
<!--    根据vo查询-->
    <select id="queryByVo" parameterType="cn.aogo.domain.QueryVo" resultType="cn.aogo.domain.User">
   select * from user where userName like #{user.userName}
    </select>
```

##### 【4】添加queryByVo测试方法

```
   //    根据queryVo中查询条件查询
    @Test
    public void queryByVo() {
        User user = new User();
        user.setUserName("%五%");
        QueryVo vo = new QueryVo();
        vo.setUser(user);
        List<User> userList = userDao.queryByVo(vo);
        for (User u : userList) {
            System.out.println(u);
        }
    }
```



### 三、MyBatis 中的连接池和事务控制

Mybatis 连接池采用的是自己的连接池技术，在 Mybatis 的 SQLMapConfig.xml 配置文件中，通过 <dataSource type="pooled"> 来实现 Mybatis 中连接池的配置。

![image-20210917075722900](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917075722900.png)

Mybatis 将它自己的数据源 DataSource 分为三类： 

- UNPOOLED：不使用连接池的数据源

- POOLED：使用连接池的数据源

- JNDI：使用 JNDI 实现的数据源
  相应地，MyBatis 内部分别定义了实现了 java.sql.DataSource 接口的 UnpooledDataSource， PooledDataSource 类来表示 UNPOOLED、POOLED 类型的数据源。 

  

MyBatis 在初始化时，根据<dataSource>的 type 属性来创建相应类型的的数据源 DataSource，即：

type=”POOLED”：MyBatis 会创建 PooledDataSource 

type=”UNPOOLED”  MyBatis 会创建 UnpooledDataSource 

type=”JNDI”：MyBatis 会从 JNDI 服务上查找 DataSource 实例，然后返回使用 

当我们需要创建 SqlSession 对象并需要执行 SQL 语句时，这时候 MyBatis 才会去调用 dataSource 对象 来创建java.sql.Connection对象。也就是说，java.sql.Connection对象的创建一直延迟到执行SQL语句 的时候。

只有当第 4句sqlSession.selectList("findUserById")，才会触发MyBatis 在底层执行下面这个方 法来创建 java.sql.Connection 对象。

只有在要用到的时候，才去获取并打开连接，当我们用完了就再 立即将数据库连接归还到连接池中。 
![image-20210917080200602](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917080200602.png)

### 四、Mybatis 的动态 SQL 语句

##### 1、if 标签

##### 【1】编写持久层接口

```
    //    根据if标签查询
    List<User> findByCondition(User user);
```

##### 【2】持久层映射配置

```html
    <!--    根据条件查询 if标签-->
    <select id="findByCondition" parameterType="cn.aogo.domain.User" resultType="cn.aogo.domain.User">
        select * from user where 1=1
        <if test="userName!=null">
            and userName = #{userName}
        </if>
    </select>
```

##### 【3】测试类

```
    //    根据if标签做条件查询
    @Test
    public void queryByIfCondition() {
        User user = new User();
        user.setUserName("aogou");

        List<User> userList = userDao.findByCondition(user);
        for (User u : userList) {
            System.out.println(u);
        }
    }

```

![image-20210917083548053](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917083548053.png)

##### 2、where>标签

##### 【1】编写持久层接口

```
  //    根据where标签查询
    List<User> findByWhereCondition(User user);
```



##### 【2】持久层映射配置

```
 <!--    根据条件查询 where标签-->
    <select id="findByWhereCondition" parameterType="cn.aogo.domain.User" resultType="cn.aogo.domain.User">
        select * from user
        <where>
            <if test="userName!=null">
                and userName =#{userName}
            </if>
        </where>
    </select>

```

##### 【3】测试类

```
//    根据where 标签做条件查询
    @Test
    public void queryByWhereCondition() {
        User user = new User();
        user.setUserName("星期五");
        List<User> userList = userDao.findByWhereCondition(user);
        for (User u : userList) {
            System.out.println(u);
        }
    }
```

##### 3、foreach标签

##### 【1】在queryVo中加入一个list做封装函数

```
public class QueryVo implements Serializable {
    private List<Integer> ids;
    public List<Integer> getIds() {
        return ids;
    }
    public void setIds(List<Integer> ids) {
        this.ids = ids;
    }
}
```

##### 【2】编写持久层接口

```
  //    根据foreach标签查询
    List<User> findByForeachCondition(QueryVo vo);
```

##### 【3】持久层映射配置

- foreach用于遍历集合，属性：

  - collection：代表要遍历的集合元素
  - open：代表语句的开始部分
  - close：代表结束部分
  - item：代表遍历集合的每个元素，生成的变量名
  - sperator：代表分隔符

  ```
      <!--    根据条件查询 foreach标签-->
      <select id="findByForeachCondition" parameterType="cn.aogo.domain.QueryVo" resultType="cn.aogo.domain.User">
          select * from user
          <where>
              <if test="ids!=null and ids.size()>0">
                  <foreach collection="ids" open="and id in (" close=")" item="id" separator=",">
                      #{id}
  
                  </foreach>
              </if>
          </where>
      </select>
  ```

【4】测试类

```
  //    根据foreach 标签做条件查询
    @Test
    public void queryByForeachCondition() {
        QueryVo vo = new QueryVo();
        List<Integer> ids = new ArrayList<Integer>();
        ids.add(1);
        ids.add(2);
        ids.add(3);
        vo.setIds(ids);
        List<User> users = userDao.findByForeachCondition(vo);
        for (User user : users) {
            System.out.println(user);
        }
    }
```



# 五、Mybatis 多表查询

## 1、数据准备

```
CREATE TABLE user
(
	id int(11) AUTO_INCREMENT  NOT NULL COMMENT '主键ID',
	username VARCHAR(30) NULL DEFAULT NULL COMMENT '姓名',
	birthday datetime DEFAULT NULL COMMENT '生日',
	gender INT  DEFAULT NULL COMMENT '性别',
	address VARCHAR(50) NULL DEFAULT NULL COMMENT '地址',

	PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

```
CREATE TABLE `account` (
  `ID` int(11) NOT NULL COMMENT '编号',
  `UID` int(11) default NULL COMMENT '用户编号',
  `MONEY` double default NULL COMMENT '金额',
  PRIMARY KEY  (`ID`),
  KEY `FK_Reference_8` (`UID`),
  CONSTRAINT `FK_Reference_8` FOREIGN KEY (`UID`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

创建实体类

```
package cn.aogo.domain;

import java.io.Serializable;

public class Account implements Serializable {
    private Integer id;
    private Long uid;
    private double money;
    
    @Override
    public String toString() {
        return "{Account:" + "id:" + id + ", uid:" + uid + ", money" + money + " }";
    }
}

```

## 2、一对一查询（通过子类方式）

eg：需要查询所有账户信息及账户对应的用户信息，包括用户名及地址。

#### 【1】定义accountUser类

为了能够封装SQL的查询结果，accountUser类继承至account，同时为了包含用户信息需要添加user类的username和address属性。

```
public class AccountUser extends Account {
    private String username;
    private String address;

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    @Override
    public String toString() {
        return super.toString() + "AccountUser:{username:" + username + ", address" + address + "}";
    }
}

```



#### 【2】编写持久层接口

新建IAccountDao的接口文件

```
public interface IAccountDao {
    //    查询所有用账户信息，并带有用户名称和地址
    List<AccountUser> findAll();
}
```



#### 【3】持久层映射配置

新增account的配置文件，IAccountDao.xml

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">

<mapper namespace="cn.aogo.dao.IAccountDao">
    <select id="findAll" resultType="cn.aogo.domain.AccountUser">
        select a.*,u.username,u.address from account a,user u where a.uid=u.id
    </select>
</mapper>
```

在SqlMapConfig.xml引入IAccountDao.xml

```
 <mappers>
        <!--        <mapper resource="IUserDao.xml"/>-->
        <!--        user表的配置-->
        <mapper resource="cn/aogo/dao/IUserDao.xml"/>
        <!--        account表的配置-->
        <mapper resource="cn/aogo/dao/IAccountDao.xml"/>
    </mappers>

```

目录结构如图：![image-20210919081525244](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210919081525244.png)



#### 【4】编写测试类

```
    //    查询所有account信息及其对应的用户名和地址
    @Test
    public void queryAccountAllIncludeUser() {
        List<AccountUser> accountUserList = accountDao.findAll();
        for (AccountUser accountUser : accountUserList) {
            System.out.println(accountUser);
        }
    }
```

![image-20210919081709735](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210919081709735.png)

## 3、一对一查询（通过建立实体类方式）

通过建立实体类的方式，使用 resultMap，定义专门的 resultMap 用于映射一对一查询结果，所以可以在 Account 类中加入 User 类的对象作为 Account 类的一个属性。

【1】修改Account实体类，添加user对象作为Account类的一个属性

```
public class Account implements Serializable {
    private User user;
    public User getUser() {
        return user;
    }
    public void setUser(User user) {
        this.user = user;
    }
}
```

【2】修改IAccountDao中的接口方法

```
    //查询所有账户信息及其用户名称和地址 使用实体类的方式 account类包含了user类
    List<Account> findAllByBean();
```

【3】重新定义 IAccountDao.xml 文件

```
 <select id="findAllByBean" resultMap="accountUserMap">
        select u.*,a.id as aid,a.uid,a.money from user u,account a where u.id=a.uid
    </select>

    <resultMap id="accountUserMap" type="cn.aogo.domain.Account">
        <id property="id" column="aid" ></id>
        <result property="money" column="money"></result>
        <result property="uid" column="uid"></result>

        <!--        一对一的关系映射 配置封装user的内容 -->
        <association property="user" column="uid" javaType="cn.aogo.domain.User">
            <id column="id" property="id"></id>
            <result property="userName" column="userName"></result>
            <result property="address" column="address"></result>
            <result property="gender" column="gender"></result>
            <result property="birthday" column="birthday"></result>
        </association>
    </resultMap>

```

【4】编写测试类

```

    //    查询所有account信息及其对应用户地址 使用account类包含user类的形式
    @Test
    public void findAllByBean() {
        List<Account> accountList = accountDao.findAllByBean();
        for (Account account : accountList) {
            System.out.println(account);
            System.out.println(account.getUser());
        }
    }
```

![image-20210919090540548](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210919090540548.png)

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

法一：

直接IUserDao.xml将放在resources文件夹下，同时修改为制定配置为

```
    <!--指定映射配置文件的位置，映射配置文件指的是每个dao独立的配置文件-->
    <mappers>
        <mapper resource="IUserDao.xml"/>
    </mappers>

```

![image-20210915084744099](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210915084744099.png)

https://blog.csdn.net/u010648555/article/details/70880425

法二：修改SqlMapConfig.xml的mappers的配置文件

```
    <mappers>
        <mapper resource="cn/aogo/dao/IUserDao.xml"/>
    </mappers>
```

![image-20210917065817969](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917065817969.png)

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

##### 6、插入空指针

![image-20210917070039993](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917070039993.png)

解决方案：在init函数不要加类名

![image-20210917070202515](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210917070202515.png)

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

##### 7、foreign key constraint 'FK_Reference_8' are incompatible

```
关联表的主键类型得保持一致
```

![image-20210918082322180](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210918082322180.png)

![image-20210918082344892](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210918082344892.png)
springboot-mybatis

```
本文将对MyBatis-Plus进行学习，主要是对简单的增删改查、分页查询、自动填充、乐观锁、逻辑删除、性能分析进行学习。
```





一、MyBatis-Plus介绍
1. Mybatis-Plus简介
2. MyBatis-Plus特性
3. 支持数据库
二、MyBatis-Plus入门
1. 入门案例
1.1 创建并初始化数据库
1.1.1 创建数据库
1.1.2 创建User表
1.2 创建springboot工程
1.2.1 导入依赖
1.2.2 修改配置文件
1.3 代码编写
1.3.1 编写User实体类
1.3.2 编写UserMapper接口
1.3.3 编写测试类（简单增删改查）
2. 提升案例
2.1 分页查询
2.1.1 配置插件
2.1.2 测试分页
2.2 自动填充
2.2.1 添加字段
2.2.2 实现元对象处理器接口
2.3 mybatisplus实现乐观锁
2.3.1 添加乐观锁版本号字段
2.3.2 配置乐观锁插件
2.3.3 测试乐观锁
2.4 逻辑删除
2.4.1 添加字段
2.4.2 配置逻辑删除插件
2.4.3 逻辑删除测试
2.5 性能分析
2.5.1 配置插件
2.5.2 性能分析测试



## 一、MyBatis-Plus介绍

MyBatis-Plus官网：https://mp.baomidou.com/

MyBatis-Plus（简称 MP）是一个 MyBatis 的增强工具，在 MyBatis 的基础上只做增强不做改变，为简化开发、提高效率而生。

## 二、Mybatis-Plus入门--基本的增删改查

### 1.1 创建并初始化数据库

#### 1.1.1 创建数据库

这里采用Navicat可视化工具创建数据库，如下：

- 数据库名：mybatis_plus
- 字符集：常用为utf8
- 排序规则：这里选utf8_general_ci

![image-20210912234341350](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210912234341350.png)

#### 1.1.2 创建User表

```
DROP TABLE IF EXISTS user;
 
CREATE TABLE user
(
	id BIGINT(20) NOT NULL COMMENT '主键ID',
	NAME VARCHAR(30) NULL DEFAULT NULL COMMENT '姓名',
	age INT(11) NULL DEFAULT NULL COMMENT '年龄',
	email VARCHAR(50) NULL DEFAULT NULL COMMENT '邮箱',
	PRIMARY KEY (id)
);
```

插入测试数据：

```
INSERT INTO user (id,name,age,email) VALUES
(1,'onestar',18,'onestar@136.com'),
(2,'twostar',18,'twostar@136.com'),
(3,'threestar',18,'threestar@136.com'),
(4,'fourstar',18,'fourstar@136.com'),
(5,'fivestar',18,'fivestar@136.com');
```



### 1.2 创建springboot工程

![image-20210912234817560](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210912234817560.png)

![image-20210912234728322](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210912234728322.png)

#### 1.2.1 导入依赖

【1】在pom.xml文件里面添加SpringBoot的起步依赖，SpringBoot要求，项目要继承SpringBoot的起步依赖spring-boot-starter-parent

```
 <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.0.1.RELEASE</version>
    </parent>
```

 【2】在pom.xml文件里面添加web的启动依赖，SpringBoot要集成SpringMVC进行Controller的开发，所以项目要导入web的启动依赖

```
 <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
```

【3】在pom.xml文件里面添加mybatis-plus、MySQL依赖导入

```
 <!--mybatis-plus-->
        <dependency>
            <groupId>com.baomidou</groupId>
            <artifactId>mybatis-plus-boot-starter</artifactId>
            <version>3.4.0</version>
        </dependency>
        <!--mysql-->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
        </dependency>
        <!--lombok用来简化实体类，不用写get、set、toString方法-->
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
        </dependency>
```

#### 1.2.2 修改配置文件

创建配置文件resources文件夹下创建application.yml，并添加以下两个配置，如下：

1. 数据库配置：driver、url、username、password
2. mybatisplus日置配置：用来在终端显示数据库执行的详细信息

```
server:
  port: 8080

spring:
  datasource:
    #    type: com.alibaba.druid.pool.DruidDataSource
    url: jdbc:p6spy:mysql://localhost:3306/mybatis_plus?useUnicode=true&characterEncoding=utf-8&useSSL=false&serverTimezone=UTC
    username: root
    password: 123456
    driver-class-name: com.p6spy.engine.spy.P6SpyDriver

mybatis-plus:
  configuration:
    log-impl: org.apache.ibatis.logging.stdout.StdOutImpl
```

### 1.3 代码编写

在com.star包下创建两个包，并创建User实体类和UserMapper接口，目录结构如下：

- entity包：用来放实体类

- mapper包：用来放持久层接口
  需要让springboot扫描到mapper接口，在mySpringBootApplication.java类中添加注解@MapperScan("com.star.mapper")

  ```
  @SpringBootApplication //申明该类是一个springboot引导类
  @MapperScan("cn.aogo.mapper")
  public class mySpringBootApplication {
      public static void main(String[] args) {
  
          SpringApplication.run(mySpringBootApplication.class,args);
      }
  }
  
  ```

  ![image-20210912235729568](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210912235729568.png)

#### 1.3.1 编写User实体类

```
@Data   //Data注解：用于简化实体类，可以不用写get、set等方法
public class User {
    @TableId(type = IdType.AUTO)  //主键策略
    private Long id;
    private String name;
    private Integer age;
    private String email;
    }
```

#### 1.3.2 编写UserMapper接口

 ```
@Repository
public interface UserMapper extends BaseMapper<User> {
}
 ```

#### 1.3.3 编写测试类（简单增删改查）

在test/java/新建测试类，mySpringBootApplicationTest.java

##### 查询case

```
@RunWith(SpringRunner.class)
@SpringBootTest(classes = mySpringBootApplication.class)
public class mySpringBootApplicationTest {
    @Autowired
    private UserMapper userMapper;

    //    查询所有
    @Test
    public void queryAll() {
        List<User> usersList = userMapper.selectList(null);
        System.out.println(usersList);
    }
    
    //    根据id查询
    @Test
    public void queryById() {
        User user = userMapper.selectById(1);
        System.out.println(user);
    }

    //    通过多个id批量查询
    @Test
    public void batchQueryByIds() {
        //    数组转化成List集合
        List<User> userList = userMapper.selectBatchIds(Arrays.asList(1, 2, 3));
        userList.forEach(System.out::println);

    }

    //简单的查询条件
    @Test
    public void queryByMap() {
        HashMap<String, Object> map = new HashMap<String, Object>();
        map.put("name", "onestar");
        map.put("age", 18);
        List<User> userList = userMapper.selectByMap(map);
        userList.forEach(System.out::println);
    }
  }
```

##### 添加case

```
  @Test
    public void addUser() {
        User user = new User();
        user.setName("13star");
        user.setAge(6);
        user.setEmail("thirteen@136.com");
        int affexted = userMapper.insert(user);
        if (affexted > 0) {
            System.out.println("insert success");
        } else {
            System.out.println("insert failed");
        }


    }
```

##### 更新case

```
  //    通过id更新user
    @Test
    public void updateUserById() {
        User user = new User();
        user.setAge(8);
        user.setEmail("aogo15@qq.com");
        user.setId(17L);
        int affected = userMapper.updateById(user);
        if (affected > 0) {
            System.out.println("update success");
        } else {
            System.out.println("update failed");
        }
    }
```

##### 删除case

```
//  根据id删除
    @Test
    public void deleteById() {
        int affected = userMapper.deleteById(14L);
        if (affected > 0) {
            System.out.println("delete success");
        } else {
            System.out.println("delete failed");
        }
    }

    //    批量删除
    @Test
    public void batchDeleteByIds() {
        int affected = userMapper.deleteBatchIds(Arrays.asList(5L, 6L));
        if (affected > 0) {
            System.out.println("delete success");
        } else {
            System.out.println("delete failed");
        }

    }

    //    根据条件删除
    @Test
    public void deleteByMap() {
        HashMap<String, Object> map = new HashMap<String, Object>();
        map.put("name", "fourstar");
        map.put("age", 18);
        int affected = userMapper.deleteByMap(map);
        if (affected > 0) {
            System.out.println("deleteByMap success");
        } else {
            System.out.println("deleteByMap failed");
        }
    }
```

## 三、 提升案例

进阶案例主要讲以下几个知识点：



- 分页查询
- 自动填充
- mybatisplus实现乐观锁

- 逻辑删除

- 性能分析

- 复杂条件查询

  

### 1.1 分页查询

Mybatis-Plus是自带了分页查询功能的，直接使用自集成的插件进行分页查询，在使用之前要配置插件，可以专门创建一个配置类来配置插件

#### 1.1.1 配置插件

在com.star包下创建config包，创建MpConfig配置类，添加分页插件

```
@EnableTransactionManagement  //事务管理
@Configuration
public class MpConfig {
    //    分页插件
    @Bean
    public PaginationInterceptor paginationInterceptor() {
        return new PaginationInterceptor();
    }
}
```

#### 1.1.2 测试分页

```
   //    分页查询
    @Test
    public void selectByPage() {
//        创建page对象，当前页（2）、每页显示对象（4）
        Page<User> page = new Page<>(2, 4);
        userMapper.selectPage(page, null);
//        获取当前页
        System.out.println("currentPage:" + page.getCurrent());
//        每页数据list集合
        System.out.println("records:" + page.getRecords());
//        获取总量
        System.out.println("total:" + page.getTotal());
        //每页显示记录数
        System.out.println("size:" + page.getSize());
//        是否有下一页
        System.out.println("next:" + page.hasNext());
//        是否有当前页
        System.out.println("previous:" + page.hasPrevious());
    }
```

![image-20210913000823832](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913000823832.png)

### 1.2 自动填充

在平时开发中，会有些数据需要自动填充，比如创建时间、更新时间等，这就可以使用MybatisPlus的自动填充功能，这里就以创建时间和更新时间为例进行演示

#### 1.2.1 添加字段

User表中添加字段：

- 创建时间：create_time

- 更新时间：update_time

  

  实体类添加属性和注解

```
		@TableField(fill = FieldFill.INSERT)
    private Date createTime;

    @TableField(fill = FieldFill.INSERT_UPDATE)
    private Date updateTime;
```

#### 1.2.2 实现元对象处理器接口

这里专门创建一个hander包来访处理器接口，在com.star包下创建hander包，创建MyMetaObjectHandler接口

```
@Slf4j
@Component  //交给spring管理
public class MyMetaObjectHandler implements MetaObjectHandler {
    //    使用mybatis-plus实现添加操作，执行该方法
    @Override
    public void insertFill(MetaObject metaObject) {
        this.setFieldValByName("createTime", new Date(), metaObject);
        this.setFieldValByName("updateTime", new Date(), metaObject);
    }

    //    使用mybatis-plus实现更新操作，执行该方法
    @Override
    public void updateFill(MetaObject metaObject) {
        this.setFieldValByName("updateTime", new Date(), metaObject);
    }

}

```

### 1.3 mybatisplus实现乐观锁

#### 1.3.1 乐观锁和悲观锁

悲观锁，顾名思义，就是很悲观，每次去拿数据的时候都认为别人会修改，所以每次在拿数据的时候都会上锁，这样别人想拿这个数据就会block直到它拿到锁。悲观锁适用于并发竞争很厉害，写比较多的操作。

乐观锁，就是很乐观，每次去拿数据的时候都认为别人不会修改，所以不会上锁，但是在提交更新的时候会判断一下在此期间别人有没有去更新这个数据。乐观锁适用于读多写少的应用场景，这样可以提高吞吐量。

这里只讲乐观锁，为了解决写数据时丢失更新问题而出现的一种解决方法。当多个人同时修改同一条记录，最后提交的把之前的提交数据覆盖，这就是丢失更新，为了防止出现这种情况，就可以采用乐观锁，在提交更新的时候会判断一下在此期间别人有没有去更新这个数据，如12306抢票

MybatisPlus实现原理：通过添加version字段来判断是否对数据进行了修改，修改将version加一，比较新的version和原有的version是不是一样，一样的version才进行更新操作，更新完成后，version就会+1，这时候另外一个拿到数据想要更新的人，在比较version那里就会不同从而更新失败。

#### 1.3.2 添加乐观锁版本号字段

添加字段
ALTER TABLE `user` ADD COLUMN `version` INT

#### 1.3.3 实体类添加属性和注解

```
  @Version
    @TableField(fill = FieldFill.INSERT_UPDATE)
    private Integer version;
```

#### 1.3.4 添加操作初始化version

```
//使用MybatisPlus实现添加操作，执行该方法
@Override
public void insertFill(MetaObject metaObject) {
    ......
    this.setFieldValByName("version", 1, metaObject);
}
```

#### 1.3.5 配置乐观锁插件

在MpConfig配置文件中添加乐观锁配置插件

```
    @Bean
//    乐观锁
    public OptimisticLockerInterceptor optimisticLockerInterceptor() {
        return new OptimisticLockerInterceptor();
    }
```

#### 1.3.6 测试乐观锁

- ##### 乐观锁模拟更新失败

```
   //    乐观锁 模拟更新失败
    @Test
    public void optimisticLocker() {
        User user = userMapper.selectById(17L);
        System.out.println(user);

//        修改数据
        user.setAge(1);

//        模拟数据库中的version比取出来的值大，即其他线程更改了数据
        user.setVersion(user.getVersion() - 1);

        int affected = userMapper.updateById(user);
        if (affected > 0) {
            System.out.println("update success");
        } else {
            System.out.println("update failed");
        }

    }
```

- ##### 乐观锁模拟更新成功

```
  //    乐观锁 模拟更新成功 注意：这种更新version需要先查在update  不然会造成version为null
    @Test
    public void lockerSuccess() {
        User user = userMapper.selectById(16L);
        user.setAge(16);
        user.setEmail("16@163.com");
        int affected = userMapper.updateById(user);
        if (affected > 0) {
            System.out.println("update success");
        } else {
            System.out.println("update failed");
        }
    }
```

### 1.4 逻辑删除

删除可以分为两种，一种是物理删除，一种是逻辑删除

物理删除：真实删除，将对应数据从数据库中删除，之后查询不到此条删除数据
逻辑删除：假删除，将对应数据中代表是否被删除字段状态修改为“被删除状态”，之后在数据库中仍能看到此数据，通过标志位字段来实现

#### 1.4.1 添加字段

添加deleted标志位字段
ALTER TABLE `user` add COLUMN is_delete INT;

#### 1.4.2 实体类添加属性和注解

```
  @TableLogic
    @TableField(fill = FieldFill.INSERT)
    private Integer isDelete;
```

#### 1.4.3 添加初始化deleted

```
//使用MybatisPlus实现添加操作，执行该方法
@Override
public void insertFill(MetaObject metaObject) {
    ......
    this.setFieldValByName("deleted", 0, metaObject);
}
```

#### 1.4.4 配置逻辑删除插件

在MpConfig配置类中添加逻辑删除插件

```
//    逻辑删除
//    3.1.1后的版本不需要这一块
//    @Bean
//    public ISqlInjector sqlInjector() {
//        return new LogicSqlInjector();
//    }
```

#### 1.4.5 逻辑删除测试

先添加一条数据，可以看到添加数据的deleted字段默认为0，然后逻辑删除该数据，执行删除后，可以看到数据库中该数据任然在，只是deleted字段由0变成了1，逻辑删除成功

```
    //    逻辑删除 将is_delete字段置为1  0：未删除  1：已删除
//    删除语句变成了 UPDATE user SET is_delete=1 WHERE id=? AND is_delete=0
//    查询语句加了is_delete=0的限制
    @Test
    public void logicDelete() {
        User user = new User();
        user.setAge(100);
        user.setEmail("100@163.com.cn");
        user.setName("one zero zero");
        userMapper.insert(user);
        System.out.println("插入的数据：" + user);
        Long id = user.getId();
        System.out.println("id：" + id);
        Integer affected = userMapper.deleteById(id);
        if (affected > 0) {
            System.out.println("logic delete success");
        } else {
            System.out.println("logic delete failed");
        }
    }
```

### 1.5 性能分析

性能分析是记录每条SQL语句执行的时间，用来帮助开发者判断某些部分是否还需要优化，有助于发现问题，需要配置插件

#### 1.5.1 配置插件

在MpConfig配置类中添加性能分析配置插件

```
//    3.2版本以后弃用
    //    @Bean
//    @Profile({"dev", "test"})
//    public PerformanceInterceptor performanceInterceptor() {
//        PerformanceInterceptor performanceInterceptor = new PerformanceInterceptor();
//        performanceInterceptor.setMaxTime(100);//ms，超过此处设置的ms则sql不执行
//        performanceInterceptor.setFormat(true);
//        return performanceInterceptor;

//
//    }
```

由于mybaisplus 3.2版本以上，移除了PerformanceInterceptor sql性能分析插件，所以我们用了MyBatis-Plus官网推荐的方式。

pom.xml新加配置。

```XML
 <!--        sql性能分析第三方包-->
        <dependency>
            <groupId>p6spy</groupId>
            <artifactId>p6spy</artifactId>
            <version>3.9.1</version>
        </dependency>
```

修改配置文件`application.yml`，将driver-class-name和url都加上p6spy

```
    url: jdbc:p6spy:mysql://localhost:3306/mybatis_plus?useUnicode=true&characterEncoding=utf-8&useSSL=false&serverTimezone=UTC
    driver-class-name: com.p6spy.engine.spy.P6SpyDriver

```

再在`resources`文件夹下创建`p6spy`的配置文件`spy.properties`：

```
#3.2.1以上使用
modulelist=com.baomidou.mybatisplus.extension.p6spy.MybatisPlusLogFactory,com.p6spy.engine.outage.P6OutageFactory
#3.2.1以下使用或者不配置
#modulelist=com.p6spy.engine.logging.P6LogFactory,com.p6spy.engine.outage.P6OutageFactory
# 自定义日志打印
logMessageFormat=com.baomidou.mybatisplus.extension.p6spy.P6SpyLogger
#日志输出到控制台
appender=com.baomidou.mybatisplus.extension.p6spy.StdoutLogger
# 使用日志系统记录 sql
#appender=com.p6spy.engine.spy.appender.Slf4JLogger
# 设置 p6spy driver 代理
deregisterdrivers=true
# 取消JDBC URL前缀
useprefix=true
# 配置记录 Log 例外,可去掉的结果集有error,info,batch,debug,statement,commit,rollback,result,resultset.
excludecategories=info,debug,result,commit,resultset
# 日期格式
dateformat=yyyy-MM-dd HH:mm:ss
# 实际驱动可多个
#driverlist=org.h2.Driver
# 是否开启慢SQL记录
outagedetection=true
# 慢SQL记录标准 2 秒
outagedetectioninterval=2

```

#### 1.5.2 性能分析测试

随便执行增删改查中的一个测试，便可以看到控制台会有SQL的执行时间

 打印出对应sql的语句执行耗时![image-20210913002440923](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913002440923.png)

#### 1.6 知识补充--查询的more

##### 查询notNull、between

version和id不为空的，年龄为18-19岁之间的数据

```
   //    根据wrapper查询 between & not null
    @Test
    public void wrapperQuery() {
        QueryWrapper<User> queryWrapper = new QueryWrapper<>();
        queryWrapper.isNotNull("version").isNotNull("id").between("age", 18, 19);
        userMapper.selectObjs(queryWrapper).forEach(System.out::println);
    }
```

![image-20210913003412001](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913003412001.png)

##### 查询eq、notBetween

name等于aogo并且年龄不在17-18之间的

```
  //    根据wrapper查询  等于 、not between
    @Test
    public void wrapper() {
        QueryWrapper<User> wrappper = new QueryWrapper<>();
        wrappper.eq("name", "aogo").notBetween("age", 17, 18);
        userMapper.selectList(wrappper).forEach(System.out::println);
    }
```

![image-20210913003505676](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913003505676.png)

##### 查询like

email没有qq、且name以t开头、以star结尾的数据

```
 //    模糊匹配查询
    @Test
    public void likeQuery() {
        QueryWrapper<User> wrapper = new QueryWrapper<>();
        wrapper.notLike("email", "qq").likeRight("name", "t").likeLeft("name", "star");

        userMapper.selectList(wrapper).forEach(System.out::println);
        userMapper.selectObjs(wrapper).forEach(System.out::println);
    }
```

![image-20210913003611525](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913003611525.png)

##### insql子查询

id在子查询中查询出来

```
  //    连接查询
    @Test
    public void connectQuery() {
        QueryWrapper<User> wrapper = new QueryWrapper<>();
        wrapper.inSql("id", "select id from user where id <4");
        List<Object> userList = userMapper.selectObjs(wrapper);
        userList.forEach(System.out::println);

    }
```

![image-20210913003659363](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913003659363.png)

##### 查询排序

```
  //    查询结果排序
    @Test
    public void orderQuery() {
        QueryWrapper<User> wrapper = new QueryWrapper<>();
        wrapper.orderByDesc("id");
        List<User> userList = userMapper.selectList(wrapper);
        userList.forEach(System.out::println);

    }
```

![image-20210913003743997](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210913003743997.png)
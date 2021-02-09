## 六种关联查询

- 交叉连接（CROSS JOIN）

- 内连接（INNER JOIN）

- 外连接（LEFT JOIN/RIGHT JOIN）

- 联合查询（UNION与UNION ALL）

- 全连接（FULL JOIN）

- 交叉连接（CROSS JOIN）

  

### 内连接分为三类

- 等值连接：ON A.id=B.id
- 不等值连接：ON A.id > B.id
- 自连接：SELECT * FROM A T1 INNER JOIN A T2 ON T1.id=T2.pid

### 外连接（LEFT JOIN/RIGHT JOIN）

- 左外连接：LEFT OUTER JOIN, 以左表为主，先查询出左表，按照ON后的关联条件匹配右表，没有匹配到的用NULL填充，可以简写成LEFT JOIN
- 右外连接：RIGHT OUTER JOIN, 以右表为主，先查询出右表，按照ON后的关联条件匹配左表，没有匹配到的用NULL填充，可以简写成RIGHT JOIN

## 实战

有2张表，1张R、1张S，R表有ABC三列，S表有CD两列，表中各有三条记录。

R表

| A    | B    | C    |
| ---- | ---- | ---- |
| a1   | b1   | c1   |
| a2   | b2   | c2   |
| a3   | b3   | c3   |

S表

| C    | D    |
| ---- | ---- |
| c1   | d1   |
| c2   | d2   |
| c4   | d3   |

##### 1、交叉连接(笛卡尔积)

select r.`*`,s.`*` from r,s

| A    | B    | C    | C    | D    |
| ---- | ---- | ---- | ---- | ---- |
| a1   | b1   | c1   | c1   | d1   |
| a2   | b2   | c2   | c1   | d1   |
| a3   | b3   | c3   | c1   | d1   |
| a1   | b1   | c1   | c2   | d2   |
| a2   | b2   | c2   | c2   | d2   |
| a3   | b3   | c3   | c2   | d2   |
| a1   | b1   | c1   | c4   | d3   |
| a2   | b2   | c2   | c4   | d3   |
| a3   | b3   | c3   | c4   | d3   |

##### 2、内连接

select r.`*`,s.`*` from r inner join s on r.c=s.c

| A    | B    | C    | C    | D    |
| ---- | ---- | ---- | ---- | ---- |
| a1   | b1   | c1   | c1   | d1   |
| a2   | b2   | c2   | c2   | d2   |

##### 3、左连接结果

select r.`*`,s.`*` from r left join s on r.c=s.c

| A    | B    | C    | C    | D    |
| ---- | ---- | ---- | ---- | ---- |
| a1   | b1   | c1   | c1   | d1   |
| a2   | b2   | c2   | c2   | d2   |
| a3   | b3   | c3   |      |      |

##### 4、右连接结果

select r.`*`,s.`*` from r right join s on r.c=s.c

|      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- |
| A    | B    | C    | C    | D    |
| a1   | b1   | c1   | c1   | d1   |
| a2   | b2   | c2   | c2   | d2   |
|      |      |      | c4   | d3   |

##### 5、全表连接的结果（MySql不支持，Oracle支持）

select r.`*`,s.`*` from r full join s on r.c=s.c

| A    | B    | C    | C    | D    |
| ---- | ---- | ---- | ---- | ---- |
| a1   | b1   | c1   | c1   | d1   |
| a2   | b2   | c2   | c2   | d2   |
| a3   | b3   | c3   |      |      |
|      |      |      | c4   | d3   |

### 
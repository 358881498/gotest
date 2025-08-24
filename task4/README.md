# 个人博客后端系统
## 项目结构

```
.
├── config/
│   └── logger.go          # 日志配置
├── controller/
│   ├── base.go            # 公共控制器
│   ├── comment.go         # 评论控制器
│   ├── login.go           # 用户登录注册控制器
│   └── post.go            # 文章控制器
├── database/
│   └── database.go        # 数据库
├── models/
│   ├── user.go            # 用户模型
│   ├── post.go            # 文章模型
│   └── comment.go         # 评论模型
├── routes/
│   └── routes.go          # 路由定义
├── app.log                # 日志记录
├── go.mod                 # Go 模块定义
├── go.sum                 # Go 模块校验和
├── main.go                # 程序入口
└── README.md              # 项目说明文档
```
## 主要功能

- 用户的注册和登录
- 文章的创建、读取、更新、删除（CRUD）
- 评论的创建、读取
- 权限控制（只有作者可以更新、删除自己的文章）
- 统一的错误处理、响应格式
- 统一日志记录

## 技术栈
- **语言**：Go 1.24.6
- **Web框架**：Gin
- **ORM**：GORM
- **数据库**：MySql
- **认证功能**：JWT
- **加密功能**：bcrypt

## 配置和运行

### 创建数据库
需要在MySql中创建数据库：
```sql
CREATE DATABASE blog CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```


### 运行项目
```sybase
cd task4
go run main.go
```
项目会在`http://localhost:8080`启动

## API文档

### 基础信息

- **认证方式**: Bearer Token (JWT)
- **响应格式**: JSON

### 认证相关

#### 用户注册

```http
POST /login/register
Content-Type: application/json
{
    "username": "Baby",
    "password": "123456",
    "email" : "Baby@sina.com"
}
```
**响应示例**:
```json
{
   "data": {
      "message": "注册成功"
   }
}
```

#### 用户登录

```http
POST /login
Content-Type: application/json

{"username":"admin12","password":"123456"}
```

**响应示例**:
```json
{
   "data": {
      "message": "登录成功",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYxMjUwMzEsImlkIjo4LCJ1c2VybmFtZSI6ImFkbWluMTIifQ.HU0DXRUhNpNTvNHrYZREoSzhNGFeLMjKqgjFJcbUSgk"
   }
}
```

### 文章管理

#### 获取文章列表

```http
GET /post/list
```
返回文章列表及文章发布人名称
```json
{
   "data": [
      {
         "ID": 0,
         "CreatedAt": "0001-01-01T00:00:00Z",
         "UpdatedAt": "0001-01-01T00:00:00Z",
         "DeletedAt": null,
         "Title": "0",
         "Content": "00",
         "UserID": 0,
         "User": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Username": "admin",
            "Email": ""
         },
         "Comments": null
      }
   ]
}
```
#### 获取文章详情&评论

```http
GET /post/index?id={id}
```
返回文章信息与文章发布人名称&文章的所有评论和评论人名称
#### 创建文章（需要认证）

```http
POST /post/add
Authorization: Bearer <登录接口返回的token>
Content-Type: application/json

{
  "title": "春晓",
  "content": "春眠不觉晓..."
}
```

#### 更新文章（需要认证）

```http
PUT /post/update?id={id}
Authorization: Bearer <登录接口返回的token>
Content-Type: application/json
{"title":"更新的标题","content":"更新的内容"}
```


#### 删除文章同时删除文章下的评论（需要认证）

```http
DELETE /post/delete?id={id}
Authorization: Bearer <登录接口返回的token>
```


### 评论管理



#### 添加评论（需要认证）

```http
POST /comment/add
Authorization: Bearer <登录接口返回的token>
Content-Type: application/json
{"content":"评论评论评论","post_id":1}
```
## 使用 Postman 测试

## 错误处理

系统使用统一的错误响应格式：

```json
{
  "data": [...],
}
```
```json
{
  "error": "错误描述信息"
}
```
## 数据库

项目使用MySQL数据库，需要预先创建数据库。数据库包含以下表：

- `users`: 用户表
- `posts`: 文章表
- `comments`: 评论表



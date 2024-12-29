0# IRT考试系统 API文档

## API 基础信息

- 基础URL: `/api/v1`
- 认证方式: Bearer Token
- 响应格式: JSON
- API版本: v1
- 默认字符集: UTF-8

## 通用规范

### API 版本控制
- 版本号在URL中体现: `/api/v1/`
- 主版本号变更代表不兼容的API变化
- 次版本号变更代表向后兼容的功能性新增
- 修订号变更代表向后兼容的问题修正

### 速率限制
- 默认限制: 1000次/小时/IP
- 认证用户限制: 5000次/小时/用户
- 特定接口限制:
  - 登录接口: 10次/分钟/IP
  - 考试提交: 60次/小时/用户
  - 文件上传: 100次/小时/用户

响应头包含以下信息：
```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1640995200
```

### 分页参数
所有列表接口都支持以下分页参数：
```json
{
    "page": "当前页码，默认1",
    "page_size": "每页条数，默认20，最大100",
    "sort": "排序字段",
    "order": "排序方向(asc/desc)"
}
```

分页响应格式：
```json
{
    "data": [],
    "meta": {
        "total": "总记录数",
        "page": "当前页码",
        "page_size": "每页条数",
        "total_pages": "总页数"
    }
}
```

### 错误处理
所有错误响应统一格式：
```json
{
    "code": "错误代码",
    "message": "错误描述",
    "details": "详细信息(可选)",
    "timestamp": "错误发生时间"
}
```

## API 端点

### 1. 认证相关

#### 1.1 登录
- 路径: `POST /auth/login`
- 描述: 用户登录
- 速率限制: 10次/分钟/IP
- 请求体:
```json
{
    "username": "string",
    "password": "string",
    "captcha": "string(可选)"
}
```
- 响应体:
```json
{
    "token": "string",
    "user": {
        "id": "string",
        "username": "string",
        "role": "string",
        "permissions": ["string"]
    },
    "expires_in": 3600
}
```

### 2. 考试管理

#### 2.1 获取考试列表
- 路径: `GET /exams`
- 描述: 获取可用考试列表
- 认证: 必需
- 分页: 支持
- 查询参数:
  - status: 考试状态
  - subject_id: 科目ID
  - start_date: 开始日期
  - end_date: 结束日期
- 响应体:
```json
{
    "data": [
        {
            "id": "string",
            "title": "string",
            "duration": "number",
            "total_questions": "number",
            "total_score": "number",
            "pass_score": "number",
            "status": "string",
            "start_time": "datetime",
            "end_time": "datetime",
            "subject_id": "number",
            "subject_name": "string"
        }
    ],
    "meta": {
        "total": "number",
        "page": "number",
        "page_size": "number",
        "total_pages": "number"
    }
}
```

### 3. 能力值管理

#### 3.1 获取用户能力值
- 路径: `GET /ability/user/{userId}/subject/{subjectId}`
- 描述: 获取用户在特定科目的能力值
- 认证: 必需
- 响应体:
```json
{
    "user_id": "number",
    "subject_id": "number",
    "ability": "number",
    "standard_error": "number",
    "updated_at": "datetime",
    "history": [
        {
            "ability": "number",
            "standard_error": "number",
            "timestamp": "datetime"
        }
    ]
}
```

### 4. 知识点管理

#### 4.1 获取知识点树
- 路径: `GET /knowledge/subject/{subjectId}`
- 描述: 获取科目的知识点树结构
- 认证: 必需
- 响应体:
```json
{
    "data": [
        {
            "id": "number",
            "name": "string",
            "description": "string",
            "parent_id": "number",
            "children": [],
            "mastery_level": "number",
            "question_count": "number"
        }
    ]
}
```

### 5. WebSocket接口

#### 5.1 考试实时状态
- 路径: `ws://api/v1/ws/exam/{examId}`
- 描述: 考试进行时的实时状态更新
- 事件类型:
  - exam_start: 考试开始
  - time_warning: 时间警告
  - exam_end: 考试结束
  - answer_submit: 答案提交确认
- 消息格式:
```json
{
    "type": "事件类型",
    "data": {
        "message": "消息内容",
        "timestamp": "时间戳",
        "details": {}
    }
}
```

## 安全性要求

### 1. 认证要求
- 所有非公开API都需要通过Bearer Token认证
- Token有效期为24小时
- 支持Token刷新机制
- 敏感操作需要二次认证

### 2. 数据加密
- 所有API通信必须使用HTTPS
- 密码必须使用bcrypt加密存储
- 敏感数据传输使用AES加密

### 3. 访问控制
- 基于角色的访问控制(RBAC)
- 资源级别的权限控制
- 操作审计日志记录
- IP白名单控制
# IRT考试系统项目结构说明

## 项目概述
本项目是一个基于项目反应理论(IRT)的自适应考试系统，采用DDD(领域驱动设计)架构模式实现。

## 技术栈
- 后端：Go 1.22.0
- 数据库：PostgreSQL 12+
- 缓存：Redis 6+
- Web框架：Gin
- ORM：GORM
- 认证：JWT
- API文档：Swagger
- 测试：Go testing

## 项目结构

```
backend/
├── cmd/                    # 应用程序入口
│   └── server/            # HTTP服务器
├── configs/               # 配置文件
├── docs/                  # 文档
│   ├── swagger/          # Swagger API文档
│   └── migrations/       # 数据库迁移脚本
├── internal/             # 内部包
│   ├── domain/          # 领域层
│   │   ├── models/     # 领域模型
│   │   └── repositories/ # 仓储接口
│   ├── application/    # 应用层
│   │   └── services/   # 应用服务
│   ├── infrastructure/ # 基础设施层
│   │   ├── persistence/ # 持久化实现
│   │   └── repositories/ # 仓储实现
│   └── interfaces/     # 接口层
│       └── api/       # API接口
│           ├── handlers/  # 处理器
│           ├── middleware/ # 中间件
│           └── dto/      # 数据传输对象
├── pkg/                 # 公共包
│   ├── auth/           # 认证相关
│   ├── errors/         # 错误处理
│   └── utils/          # 工具函数
├── scripts/            # 脚本文件
├── tests/              # 测试文件
└── web/               # Web资源
    └── swagger/       # Swagger UI

```

## 分层说明

### 领域层 (Domain Layer)
- 包含核心业务逻辑和规则
- 定义领域模型和仓储接口
- 与技术实现无关的纯业务代码

主要组件：
- 领域模型（models）：User, Exam, Question, KnowledgePoint等
- 仓储接口（repositories）：定义数据访问接口
- 领域服务：复杂的领域逻辑

### 应用层 (Application Layer)
- 协调领域对象完成用户用例
- 处理事务和业务流程
- 不包含业务规则

主要组件：
- 应用服务：ExamService, AbilityService, KnowledgeService等
- 事务处理
- 用例实现

### 基础设施层 (Infrastructure Layer)
- 提供技术实现
- 实现仓储接口
- 提供外部服务集成

主要组件：
- 仓储实现
- 数据库访问
- 缓存实现
- 外部服务集成

### 接口层 (Interfaces Layer)
- 处理外部请求
- 数据转换
- 请求验证

主要组件：
- API处理器
- 中间件
- DTO对象
- 参数验证

## 关键模块

### 考试模块
- 试卷管理
- 考试进行
- 成绩评估
- 答题分析

### IRT模块
- 能力值估计
- 题目参数计算
- 试题推荐
- 难度分析

### 知识点模块
- 知识图谱
- 掌握度评估
- 学习路径推荐
- 知识点关联

### 用户模块
- 用户管理
- 角色权限
- 认证授权
- 用户画像

## 开发规范

### 代码规范
- 遵循Go语言规范
- 使用gofmt格式化代码
- 遵循项目既定的命名规范
- 编写单元测试

### API规范
- RESTful设计
- 版本控制
- 统一错误处理
- 请求参数验证

### 文档规范
- 及时更新API文档
- 编写技术文档
- 注释关键代码
- 维护变更日志

### 安全规范
- 数据加密
- 访问控制
- 输入验证
- 日志记录
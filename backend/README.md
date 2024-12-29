# IRT考试系统

基于项目反应理论(IRT)的自适应考试系统，采用DDD架构设计，使用Go语言开发。

## 功能特点

- 基于IRT的自适应题目推荐
- 实时能力值估计
- 知识点掌握度分析
- 学习路径推荐
- 完整的考试管理功能
- 详细的数据分析报告

## 技术栈

- Go 1.22.0
- PostgreSQL 12+
- Redis 6+
- Gin Web框架
- GORM ORM框架
- JWT认证
- Swagger API文档

## 快速开始

### 环境要求

- Go 1.22.0+
- PostgreSQL 12+
- Redis 6+
- Make

### 安装步骤

1. 克隆项目
```bash
git clone https://github.com/your-username/irt-exam-system.git
cd irt-exam-system/backend
```

2. 安装依赖
```bash
go mod download
```

3. 配置环境变量
```bash
cp .env.example .env
# 编辑.env文件，配置数据库等信息
```

4. 初始化数据库
```bash
make migrate
```

5. 运行项目
```bash
make run
```

### 开发模式

```bash
make dev
```

### 运行测试

```bash
make test
```

### 构建项目

```bash
make build
```

## 项目结构

详见 [项目结构文档](docs/project_structure.md)

## API文档

启动项目后访问：http://localhost:8080/swagger/index.html

详细API文档见：[API文档](docs/swagger_api_documentation.md)

## 数据库设计

详见 [数据库设计文档](docs/data_structure_documentation.md)

## 架构设计

详见 [架构设计文档](docs/architecture/README.md)

## 开发规范

### 代码规范

- 使用gofmt格式化代码
- 遵循Go语言规范
- 编写单元测试
- 注释关键代码

### Git提交规范

格式：`<type>(<scope>): <subject>`

type:
- feat: 新功能
- fix: 修复bug
- docs: 文档更新
- style: 代码格式
- refactor: 重构
- test: 测试
- chore: 构建过程或辅助工具的变动

### 分支管理

- main: 主分支，稳定版本
- develop: 开发分支
- feature/*: 功能分支
- bugfix/*: 修复分支
- release/*: 发布分支

## 部署

### Docker部署

1. 构建镜像
```bash
docker build -t irt-exam-system .
```

2. 运行容器
```bash
docker-compose up -d
```

### 手动部署

1. 构建项目
```bash
make build
```

2. 配置服务
```bash
sudo cp deploy/irt-exam-system.service /etc/systemd/system/
sudo systemctl enable irt-exam-system
sudo systemctl start irt-exam-system
```

## 监控

- 健康检查：`/health`
- Metrics：`:9090/metrics`
- 性能监控：集成Prometheus + Grafana

## 贡献指南

1. Fork项目
2. 创建功能分支
3. 提交代码
4. 创建Pull Request

## 许可证

MIT License

## 联系方式

- 作者：Your Name
- 邮箱：your.email@example.com
- 项目地址：https://github.com/your-username/irt-exam-system 
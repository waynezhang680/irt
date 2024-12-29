# IRT考试系统数据库结构文档

## 数据库概述

本系统使用PostgreSQL数据库，版本要求12.0及以上。主要包含以下模块的数据表：
- 用户管理
- 考试管理
- 题目管理
- 知识点管理
- IRT模型

## 表结构说明

### 1. 用户相关表

#### 1.1 users (用户表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | 用户ID |
| username | varchar(50) | UNIQUE, NOT NULL | 用户名 |
| password | varchar(100) | NOT NULL | 密码(bcrypt加密) |
| email | varchar(100) | UNIQUE | 邮箱 |
| role | varchar(20) | NOT NULL | 角色(student/teacher/admin) |
| status | varchar(20) | NOT NULL | 状态(active/inactive/blocked) |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |
| deleted_at | timestamptz | | 删除时间 |

索引：
- idx_users_username (username)
- idx_users_email (email)
- idx_users_role (role)

### 2. 考试相关表

#### 2.1 exam_papers (试卷表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | 试卷ID |
| title | text | NOT NULL | 试卷标题 |
| subject_id | bigint | FK, NOT NULL | 科目ID |
| description | text | | 试卷说明 |
| time_limit | bigint | NOT NULL | 考试时长(秒) |
| total_score | numeric | NOT NULL | 总分 |
| pass_score | numeric | NOT NULL | 及格分数 |
| status | varchar(20) | NOT NULL | 状态 |
| start_time | timestamptz | NOT NULL | 开始时间 |
| end_time | timestamptz | NOT NULL | 结束时间 |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |
| deleted_at | timestamptz | | 删除时间 |

索引：
- idx_exam_papers_subject (subject_id)
- idx_exam_papers_status (status)
- idx_exam_papers_time (start_time, end_time)

#### 2.2 exam_records (考试记录表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | 记录ID |
| user_id | bigint | FK, NOT NULL | 用户ID |
| exam_paper_id | bigint | FK, NOT NULL | 试卷ID |
| start_time | timestamptz | NOT NULL | 开始时间 |
| end_time | timestamptz | | 结束时间 |
| score | numeric | | 得分 |
| status | varchar(20) | NOT NULL | 状态 |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |

索引：
- idx_exam_records_user (user_id)
- idx_exam_records_paper (exam_paper_id)
- idx_exam_records_status (status)

### 3. 题目相关表

#### 3.1 questions (题目表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | 题目ID |
| type | varchar(20) | NOT NULL | 题目类型 |
| content | text | NOT NULL | 题目内容 |
| answer | text | NOT NULL | 标准答案 |
| analysis | text | | 解析 |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |
| deleted_at | timestamptz | | 删除时间 |

索引：
- idx_questions_type (type)

#### 3.2 question_options (选项表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | 选项ID |
| question_id | bigint | FK, NOT NULL | 题目ID |
| content | text | NOT NULL | 选项内容 |
| is_correct | boolean | NOT NULL | 是否正确答案 |
| label | varchar(10) | NOT NULL | 选项标签(A/B/C/D) |
| order | bigint | NOT NULL | 排序 |

索引：
- idx_question_options_question (question_id)

### 4. IRT相关表

#### 4.1 user_abilities (用户能力值表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | ID |
| user_id | bigint | FK, NOT NULL | 用户ID |
| subject_id | bigint | FK, NOT NULL | 科目ID |
| ability | numeric | NOT NULL | 能力值θ |
| standard_error | numeric | NOT NULL | 标准误 |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |

索引：
- idx_user_abilities_user (user_id)
- idx_user_abilities_subject (subject_id)

#### 4.2 question_parameters (题目参数表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | ID |
| question_id | bigint | FK, NOT NULL | 题目ID |
| difficulty | numeric | NOT NULL | 难度参数b |
| discrimination | numeric | NOT NULL | 区分度参数a |
| guessing | numeric | NOT NULL | 猜测参数c |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |

索引：
- idx_question_parameters_question (question_id)

### 5. 知识点相关表

#### 5.1 knowledge_points (知识点表)
| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | bigserial | PK | 知识点ID |
| subject_id | bigint | FK, NOT NULL | 科目ID |
| name | text | NOT NULL | 知识点名称 |
| description | text | | 知识点描述 |
| parent_id | bigint | FK | 父知识点ID |
| created_at | timestamptz | NOT NULL | 创建时间 |
| updated_at | timestamptz | NOT NULL | 更新时间 |
| deleted_at | timestamptz | | 删除时间 |

索引：
- idx_knowledge_points_subject (subject_id)
- idx_knowledge_points_parent (parent_id)

## 数据库约束

### 外键约束
- exam_papers.subject_id -> subjects.id
- exam_records.user_id -> users.id
- exam_records.exam_paper_id -> exam_papers.id
- question_options.question_id -> questions.id
- user_abilities.user_id -> users.id
- user_abilities.subject_id -> subjects.id
- question_parameters.question_id -> questions.id
- knowledge_points.subject_id -> subjects.id
- knowledge_points.parent_id -> knowledge_points.id

### 唯一约束
- users.username
- users.email
- question_parameters.question_id

### 检查约束
- exam_papers.total_score >= 0
- exam_papers.pass_score >= 0
- exam_papers.time_limit > 0
- question_parameters.discrimination > 0
- question_parameters.guessing >= 0 AND question_parameters.guessing <= 1

## 数据库索引策略

### 主键索引
所有表都使用bigserial类型的id作为主键，自动创建主键索引

### 外键索引
所有外键字段都创建索引以提高关联查询性能

### 复合索引
- (user_id, subject_id) 用于用户能力值查询
- (exam_paper_id, status) 用于考试记录查询
- (start_time, end_time) 用于考试时间查询

### 全文索引
- questions.content 用于题目内容搜索
- knowledge_points.name 用于知识点搜索

## 数据库优化建议

1. 分区策略
- exam_records表按时间分区
- question_parameters表按科目分区

2. 物化视图
- 用户成绩统计
- 知识点掌握度统计
- 题目难度分析

3. 缓存策略
- 用户能力值缓存
- 试卷信息缓存
- 知识点树缓存

4. 性能优化
- 定期VACUUM
- 定期更新统计信息
- 监控慢查询
- 优化查询计划
#!/usr/bin/env python3
import os
import psycopg2
from dotenv import load_dotenv

# 加载环境变量
load_dotenv()

# 数据库连接配置
DB_CONFIG = {
    'dbname': os.getenv('DB_NAME', 'irt_exam_system'),
    'user': os.getenv('DB_USER', 'postgres'),
    'password': os.getenv('DB_PASSWORD', 'Waynez0625@wh'),
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'sslmode': 'disable'
}

def create_tables():
    """创建数据库表"""
    try:
        # 连接数据库
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()
        
        # 创建用户表
        cur.execute("""
            CREATE TABLE IF NOT EXISTS users (
                id SERIAL PRIMARY KEY,
                username VARCHAR(50) NOT NULL UNIQUE,
                email VARCHAR(100) NOT NULL UNIQUE,
                password VARCHAR(255) NOT NULL,
                role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'teacher', 'student')),
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            )
        """)
        
        # 创建试题表
        cur.execute("""
            CREATE TABLE IF NOT EXISTS questions (
                id SERIAL PRIMARY KEY,
                subject VARCHAR(50) NOT NULL,
                title VARCHAR(500) NOT NULL,
                content TEXT NOT NULL,
                question_type VARCHAR(20) NOT NULL CHECK (question_type IN ('single_choice', 'multiple_choice', 'true_false')),
                difficulty DECIMAL(5,4) NOT NULL CHECK (difficulty >= 0 AND difficulty <= 1),
                options JSONB NOT NULL,
                correct_answer TEXT NOT NULL,
                created_by INTEGER REFERENCES users(id),
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            )
        """)
        
        # 创建试卷表
        cur.execute("""
            CREATE TABLE IF NOT EXISTS exam_papers (
                id SERIAL PRIMARY KEY,
                title VARCHAR(200) NOT NULL,
                subject VARCHAR(50) NOT NULL,
                description TEXT,
                duration INTEGER NOT NULL CHECK (duration > 0),
                total_score DECIMAL(6,2) NOT NULL CHECK (total_score > 0),
                pass_score DECIMAL(6,2) NOT NULL CHECK (pass_score <= total_score),
                status VARCHAR(20) NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'published', 'expired')),
                start_time TIMESTAMP WITH TIME ZONE,
                end_time TIMESTAMP WITH TIME ZONE,
                created_by INTEGER REFERENCES users(id),
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                CHECK (end_time > start_time)
            )
        """)
        
        # 创建试卷-试题关联表
        cur.execute("""
            CREATE TABLE IF NOT EXISTS exam_paper_questions (
                id SERIAL PRIMARY KEY,
                exam_paper_id INTEGER REFERENCES exam_papers(id) ON DELETE CASCADE,
                question_id INTEGER REFERENCES questions(id),
                score DECIMAL(5,2) NOT NULL CHECK (score > 0),
                question_order INTEGER NOT NULL,
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                UNIQUE (exam_paper_id, question_id),
                UNIQUE (exam_paper_id, question_order)
            )
        """)
        
        # 创建考试记录表
        cur.execute("""
            CREATE TABLE IF NOT EXISTS exam_records (
                id SERIAL PRIMARY KEY,
                user_id INTEGER REFERENCES users(id),
                exam_paper_id INTEGER REFERENCES exam_papers(id),
                start_time TIMESTAMP WITH TIME ZONE NOT NULL,
                end_time TIMESTAMP WITH TIME ZONE,
                score DECIMAL(6,2),
                status VARCHAR(20) NOT NULL DEFAULT 'in_progress' CHECK (status IN ('not_started', 'in_progress', 'completed', 'graded')),
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            )
        """)
        
        # 创建答题记录表
        cur.execute("""
            CREATE TABLE IF NOT EXISTS exam_responses (
                id SERIAL PRIMARY KEY,
                exam_record_id INTEGER REFERENCES exam_records(id) ON DELETE CASCADE,
                question_id INTEGER REFERENCES questions(id),
                answer TEXT NOT NULL,
                is_correct BOOLEAN NOT NULL,
                time_spent INTEGER NOT NULL CHECK (time_spent >= 0),
                score DECIMAL(5,2) NOT NULL DEFAULT 0,
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            )
        """)
        
        # 提交事务
        conn.commit()
        print("数据库表创建成功！")
        
        cur.close()
        conn.close()
        
    except Exception as e:
        conn.rollback()
        print(f"创建数据库表时发生错误: {str(e)}")
        raise

if __name__ == "__main__":
    create_tables() 
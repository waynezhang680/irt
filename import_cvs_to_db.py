import pandas as pd
from sqlalchemy import create_engine, text, MetaData, Table, Column, Integer, String, Float, DateTime, Text, JSON, Boolean
from urllib.parse import quote_plus
import psycopg2
import os
import json
from sqlalchemy.orm import sessionmaker
from datetime import datetime
from sqlalchemy.schema import CreateTable

def create_tables(engine):
    """创建数据库表结构"""
    try:
        # 先删除所有相关表
        with engine.begin() as conn:
            conn.execute(text("""
                DROP TABLE IF EXISTS exam_responses CASCADE;
                DROP TABLE IF EXISTS exam_records CASCADE;
                DROP TABLE IF EXISTS exam_paper_questions CASCADE;
                DROP TABLE IF EXISTS exam_papers CASCADE;
                DROP TABLE IF EXISTS question_knowledge_points CASCADE;
                DROP TABLE IF EXISTS question_options CASCADE;
                DROP TABLE IF EXISTS questions CASCADE;
                DROP TABLE IF EXISTS knowledge_points CASCADE;
                DROP TABLE IF EXISTS subjects CASCADE;
                DROP TABLE IF EXISTS user_abilities CASCADE;
                DROP TABLE IF EXISTS users CASCADE;
                DROP TABLE IF EXISTS roles CASCADE;
            """))
            print("已删除旧表结构")

        metadata = MetaData()

        # 角色表
        roles = Table('roles', metadata,
            Column('id', Integer, primary_key=True),
            Column('name', String(50), nullable=False, unique=True),
            Column('description', Text),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 用户表
        users = Table('users', metadata,
            Column('id', Integer, primary_key=True),
            Column('username', String(50), nullable=False, unique=True),
            Column('password_hash', String(255), nullable=False),
            Column('role_id', Integer, nullable=False),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 科目表
        subjects = Table('subjects', metadata,
            Column('id', Integer, primary_key=True),
            Column('name', String(255), nullable=False, unique=True),
            Column('description', Text),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 知识点表
        knowledge_points = Table('knowledge_points', metadata,
            Column('id', Integer, primary_key=True),
            Column('subject_id', Integer, nullable=False),
            Column('name', String(255), nullable=False),
            Column('description', Text),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 试题表
        questions = Table('questions', metadata,
            Column('id', Integer, primary_key=True),
            Column('subject_id', Integer, nullable=False),
            Column('title', String(255), nullable=False),
            Column('content', Text, nullable=False),
            Column('type', String(50), nullable=False),
            Column('difficulty', Float, nullable=False, default=0.5),
            Column('discrimination', Float, nullable=False, default=0.0),
            Column('guessing', Float, nullable=False, default=0.0),  # IRT 3PL模型的猜测参数
            Column('options', Text, nullable=True),  # JSON格式存储选项
            Column('correct_answer', String(10), nullable=False),
            Column('reference', Text),
            Column('author', String(100)),
            Column('notes', Text),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 试题知识点关联表
        question_knowledge_points = Table('question_knowledge_points', metadata,
            Column('id', Integer, primary_key=True),
            Column('question_id', Integer, nullable=False),
            Column('knowledge_point_id', Integer, nullable=False),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 试卷表
        exam_papers = Table('exam_papers', metadata,
            Column('id', Integer, primary_key=True),
            Column('title', String(255), nullable=False),
            Column('subject_id', Integer, nullable=False),
            Column('description', Text),
            Column('time_limit', Integer, nullable=False),  # 考试时长（分钟）
            Column('total_score', Float, nullable=False),
            Column('pass_score', Float, nullable=False),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 试卷题目关联表
        exam_paper_questions = Table('exam_paper_questions', metadata,
            Column('id', Integer, primary_key=True),
            Column('exam_paper_id', Integer, nullable=False),
            Column('question_id', Integer, nullable=False),
            Column('score', Float, nullable=False),
            Column('order', Integer, nullable=False),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 考试记录表
        exam_records = Table('exam_records', metadata,
            Column('id', Integer, primary_key=True),
            Column('user_id', Integer, nullable=False),
            Column('exam_paper_id', Integer, nullable=False),
            Column('start_time', DateTime(timezone=True), nullable=False),
            Column('end_time', DateTime(timezone=True)),
            Column('score', Float),
            Column('status', String(20), nullable=False),  # 进行中/已完成/已交卷
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 答题记录表（用于IRT分析）
        exam_responses = Table('exam_responses', metadata,
            Column('id', Integer, primary_key=True),
            Column('exam_record_id', Integer, nullable=False),
            Column('question_id', Integer, nullable=False),
            Column('user_answer', String(255), nullable=False),
            Column('is_correct', Boolean, nullable=False),
            Column('response_time', Integer, nullable=False),  # 答题用时（秒）
            Column('score', Float, nullable=False),
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 用户能力值表（IRT估计）
        user_abilities = Table('user_abilities', metadata,
            Column('id', Integer, primary_key=True),
            Column('user_id', Integer, nullable=False),
            Column('subject_id', Integer, nullable=False),
            Column('ability', Float, nullable=False),  # θ值
            Column('standard_error', Float, nullable=False),  # 标准误
            Column('created_at', DateTime(timezone=True), default=datetime.utcnow),
            Column('updated_at', DateTime(timezone=True), default=datetime.utcnow)
        )

        # 创建表
        metadata.create_all(engine)
        print("数据库表创建成功！")

        # 创建外键约束
        with engine.begin() as conn:
            # 用户-角色外键
            conn.execute(text("""
                ALTER TABLE users 
                ADD CONSTRAINT fk_users_roles 
                FOREIGN KEY (role_id) REFERENCES roles(id)
            """))

            # 知识点-科目外键
            conn.execute(text("""
                ALTER TABLE knowledge_points 
                ADD CONSTRAINT fk_knowledge_points_subjects 
                FOREIGN KEY (subject_id) REFERENCES subjects(id)
            """))

            # 试题-科目外键
            conn.execute(text("""
                ALTER TABLE questions 
                ADD CONSTRAINT fk_questions_subjects 
                FOREIGN KEY (subject_id) REFERENCES subjects(id)
            """))

            # 试题知识点关联表的外键
            conn.execute(text("""
                ALTER TABLE question_knowledge_points 
                ADD CONSTRAINT fk_question_knowledge_points_question 
                FOREIGN KEY (question_id) REFERENCES questions(id)
            """))

            conn.execute(text("""
                ALTER TABLE question_knowledge_points 
                ADD CONSTRAINT fk_question_knowledge_points_knowledge_point 
                FOREIGN KEY (knowledge_point_id) REFERENCES knowledge_points(id)
            """))

            # 试卷表的外键
            conn.execute(text("""
                ALTER TABLE exam_papers 
                ADD CONSTRAINT fk_exam_papers_subject 
                FOREIGN KEY (subject_id) REFERENCES subjects(id)
            """))

            # 试卷题目关联表的外键
            conn.execute(text("""
                ALTER TABLE exam_paper_questions 
                ADD CONSTRAINT fk_exam_paper_questions_exam_paper 
                FOREIGN KEY (exam_paper_id) REFERENCES exam_papers(id)
            """))

            conn.execute(text("""
                ALTER TABLE exam_paper_questions 
                ADD CONSTRAINT fk_exam_paper_questions_question 
                FOREIGN KEY (question_id) REFERENCES questions(id)
            """))

            # 考试记录表的外键
            conn.execute(text("""
                ALTER TABLE exam_records 
                ADD CONSTRAINT fk_exam_records_user 
                FOREIGN KEY (user_id) REFERENCES users(id)
            """))

            conn.execute(text("""
                ALTER TABLE exam_records 
                ADD CONSTRAINT fk_exam_records_exam_paper 
                FOREIGN KEY (exam_paper_id) REFERENCES exam_papers(id)
            """))

            # 答题记录表的外键
            conn.execute(text("""
                ALTER TABLE exam_responses 
                ADD CONSTRAINT fk_exam_responses_exam_record 
                FOREIGN KEY (exam_record_id) REFERENCES exam_records(id)
            """))

            conn.execute(text("""
                ALTER TABLE exam_responses 
                ADD CONSTRAINT fk_exam_responses_question 
                FOREIGN KEY (question_id) REFERENCES questions(id)
            """))

            # 用户能力值表的外键
            conn.execute(text("""
                ALTER TABLE user_abilities 
                ADD CONSTRAINT fk_user_abilities_user 
                FOREIGN KEY (user_id) REFERENCES users(id)
            """))

            conn.execute(text("""
                ALTER TABLE user_abilities 
                ADD CONSTRAINT fk_user_abilities_subject 
                FOREIGN KEY (subject_id) REFERENCES subjects(id)
            """))

            print("所有外键约束创建成功！")

            # 修改 options 列为 JSONB 类型
            conn.execute(text("ALTER TABLE questions ALTER COLUMN options TYPE JSONB USING options::jsonb"))
            print("修改 options 列类型为 JSONB 成功！")

            # 插入默认角色
            conn.execute(text("""
                INSERT INTO roles (id, name, description) VALUES 
                (1, 'admin', '系统管理员'),
                (2, 'teacher', '教师'),
                (3, 'student', '学生')
                ON CONFLICT (id) DO NOTHING
            """))

            # 插入默认科目
            conn.execute(text("""
                INSERT INTO subjects (id, name, description) 
                VALUES (1, '默认科目', '默认科目描述')
                ON CONFLICT (id) DO NOTHING
            """))

            print("默认数据创建成功！")

    except Exception as e:
        print(f"创建表结构时出错: {str(e)}")
        raise

def import_csv_to_db():
    try:
        # 处理数据库密码中的特殊字符
        password = quote_plus("Waynez0625@wh")
        db_url = f"postgresql://postgres:{password}@localhost:5432/irt_exam_system"
        
        # 创建数据库连接
        engine = create_engine(db_url)
        
        # 创建表结构
        create_tables(engine)

        Session = sessionmaker(bind=engine)
        session = Session()

        # 测试数据库连接
        try:
            with engine.connect() as conn:
                conn.execute(text("SELECT 1"))
                print("数据库连接成功！")
        except Exception as e:
            print(f"数据库连接失败: {str(e)}")
            return

        # 读取 CSV 文件
        csv_path = "试题样例.csv"
        if not os.path.exists(csv_path):
            print(f"错误: 找不到 CSV 文件: {csv_path}")
            return

        print("读取 CSV 文件...")
        df = pd.read_csv(csv_path)
        print("CSV 文件读取成功！")

        # 重置索引作为新的 ID
        df = df.reset_index(drop=True)
        df.index += 1  # 从1开始
        df['序号'] = df.index  # 使用新的索引作为序号

        # CSV列名与数据库表字段的映射
        column_mapping = {
            '序号': 'id',
            '难度': 'difficulty',
            '试题主题': 'title',
            '题型': 'type',
            '试题正文': 'content',
            '试题选项': 'options',
            '试题答案': 'correct_answer',
            '依据 出处': 'reference',
            '出题人': 'author',
            '备注': 'notes'
        }

        # 重命名���
        df = df.rename(columns=column_mapping)

        # 处理试题选项，转换为JSON格式
        def parse_options(options_str):
            try:
                if pd.isna(options_str):
                    return '{}'
                
                options_dict = {}
                options = str(options_str).split(';')
                
                for i, opt in enumerate(options):
                    opt = opt.strip()
                    if opt:
                        key = chr(65 + i)  # A, B, C, D...
                        options_dict[key] = opt.replace('$', '').strip()

                return json.dumps(options_dict, ensure_ascii=False)
            except Exception as e:
                print(f"解析选项时出错: {str(e)}, 选项内容: {options_str}")
                return '{}'

        print("处理试题选项...")
        df['options'] = df['options'].apply(parse_options)

        # 处理难度值
        def convert_difficulty(diff_str):
            difficulty_map = {
                '易': 0.3,
                '中': 0.5,
                '难': 0.7
            }
            return difficulty_map.get(str(diff_str).strip(), 0.5)

        df['difficulty'] = df['difficulty'].apply(convert_difficulty)

        # 添加必要的字段
        df['discrimination'] = 0.0
        df['subject_id'] = 1
        current_time = pd.Timestamp.now(tz='UTC')
        df['created_at'] = current_time
        df['updated_at'] = current_time

        # 确保所有必需的列都存在并且数据类型正确
        df['title'] = df['title'].fillna('')
        df['content'] = df['content'].fillna('')
        df['type'] = df['type'].fillna('单选题')
        df['correct_answer'] = df['correct_answer'].fillna('')
        df['reference'] = df['reference'].fillna('')
        df['author'] = df['author'].fillna('')
        df['notes'] = df['notes'].fillna('')

        # 选择需要的列并按正确顺序排列
        columns_to_import = [
            'id', 'title', 'content', 'type', 'difficulty', 
            'discrimination', 'subject_id', 'options', 
            'correct_answer', 'reference', 'author', 'notes',
            'created_at', 'updated_at'
        ]

        df = df[columns_to_import]

        print("开始导入数据到数据库...")
        # 手动插入数据
        with engine.begin() as connection:
            # 先删除现有数据
            connection.execute(text("DELETE FROM questions"))
            print("已清空现有数据")
            
            # 添加 guessing 字段（IRT 3PL模型的猜测参数）
            df['guessing'] = 0.0  # 默认值
            
            for _, row in df.iterrows():
                try:
                    data = row.to_dict()
                    print(f"正在处理第 {data['id']} 条数据")
                    
                    query = text("""
                        INSERT INTO questions (
                            id, title, content, type, difficulty, 
                            discrimination, subject_id, options, guessing,
                            correct_answer, reference, author, notes,
                            created_at, updated_at
                        ) VALUES (
                            :id, :title, :content, :type, :difficulty,
                            :discrimination, :subject_id, cast(:options as jsonb), :guessing,
                            :correct_answer, :reference, :author, :notes,
                            :created_at, :updated_at
                        )
                    """)
                    
                    # 打印要插入的数据
                    print("插入数据:", data)
                    
                    connection.execute(query, data)
                    print(f"第 {data['id']} 条数据导入成功")
                except Exception as e:
                    print(f"插入第 {data['id']} 条数据时出错: {str(e)}")
                    raise
            
            # 重置序列
            connection.execute(text("""
                SELECT setval('questions_id_seq', (SELECT MAX(id) FROM questions))
            """))
            
            print("数据导入成功！")

    except Exception as e:
        print(f"导入数据时出错: {str(e)}")
        print("请检查数据库连接信息和CSV文件格式")
        raise  # 抛出异常以查看完整的错误堆栈
    finally:
        if 'session' in locals():
            session.close()

if __name__ == "__main__":
    try:
        import_csv_to_db()
    except Exception as e:
        print(f"程序执行出错: {str(e)}")
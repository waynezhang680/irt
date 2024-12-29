import psycopg2
from psycopg2.extras import Json
import bcrypt
from datetime import datetime, timedelta
import os
from dotenv import load_dotenv

# 加载环境变量
load_dotenv()

# 数据库连接配置
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'irt_exam_system'),
    'user': os.getenv('DB_USER', 'postgres'),
    'password': os.getenv('DB_PASSWORD', '')
}

def connect_db():
    """连接数据库"""
    return psycopg2.connect(**DB_CONFIG)

def hash_password(password):
    """密码哈希"""
    salt = bcrypt.gensalt()
    return bcrypt.hashpw(password.encode('utf-8'), salt).decode('utf-8')

def clear_data(conn):
    """清空所有表数据"""
    with conn.cursor() as cur:
        # 禁用外键约束
        cur.execute("SET CONSTRAINTS ALL DEFERRED;")
        
        # 清空表（按照依赖关系倒序）
        tables = [
            'exam_responses',
            'exam_records',
            'exam_paper_questions',
            'exam_papers',
            'questions',
            'knowledge_points',
            'subjects',
            'users',
            'roles'
        ]
        
        for table in tables:
            cur.execute(f"TRUNCATE TABLE {table} CASCADE;")
        
        conn.commit()

def init_test_data(conn):
    """初始化测试数据"""
    with conn.cursor() as cur:
        # 1. 创建角色
        roles = [
            ('admin', '系统管理员'),
            ('teacher', '教师'),
            ('student', '学生')
        ]
        role_ids = {}  # 用于存储角色ID
        for role in roles:
            cur.execute(
                """
                INSERT INTO roles (name, description, created_at, updated_at) 
                VALUES (%s, %s, NOW(), NOW()) 
                RETURNING id;
                """,
                role
            )
            role_ids[role[0]] = cur.fetchone()[0]

        # 2. 创建用户
        users = [
            ('admin', hash_password('admin123'), role_ids['admin']),      # 管理员
            ('teacher1', hash_password('teacher123'), role_ids['teacher']),  # 教师
            ('student1', hash_password('student123'), role_ids['student'])   # 学生
        ]
        
        for username, password_hash, role_id in users:
            cur.execute(
                """
                INSERT INTO users (username, password, role_id, created_at, updated_at) 
                VALUES (%s, %s, %s, NOW(), NOW())
                """,
                (username, password_hash, role_id)
            )

        # 提交事务以确保角色和用户创建成功
        conn.commit()

        # 3. 创建科目
        subjects = [
            ('高等数学', '大学理科高等数学课程'),
            ('线性代数', '大学理科线性代数课程'),
            ('概率论', '概率论与数理统计')
        ]
        for subject in subjects:
            cur.execute(
                """
                INSERT INTO subjects (name, description, created_at, updated_at) 
                VALUES (%s, %s, NOW(), NOW()) 
                RETURNING id;
                """,
                subject
            )
            if subject[0] == '高等数学':
                math_subject_id = cur.fetchone()[0]

        # 4. 创建知识点
        knowledge_points = [
            (math_subject_id, '函数极限', '函数极限的概念和计算'),
            (math_subject_id, '导数', '导数的概念和计算'),
            (math_subject_id, '积分', '积分的概念和计算')
        ]
        for point in knowledge_points:
            cur.execute(
                "INSERT INTO knowledge_points (subject_id, name, description) VALUES (%s, %s, %s);",
                point
            )

        # 5. 创建试题
        questions = [
            (math_subject_id, '求极限', '计算极限：lim(x→0) sin(x)/x', 
             'multiple_choice', 0.5, 0.8, 0.25,
             Json({'A': '0', 'B': '1', 'C': '2', 'D': '不存在'}), 'B'),
            (math_subject_id, '求导数', '计算 f(x)=x² 在 x=2 处的导数',
             'multiple_choice', 0.4, 0.7, 0.25,
             Json({'A': '2', 'B': '4', 'C': '8', 'D': '1'}), 'B'),
            (math_subject_id, '计算积分', '计算 ∫x dx (从0到1)',
             'multiple_choice', 0.6, 0.75, 0.25,
             Json({'A': '0', 'B': '1', 'C': '0.5', 'D': '2'}), 'C')
        ]
        for q in questions:
            cur.execute("""
                INSERT INTO questions (subject_id, title, content, type, 
                                     difficulty, discrimination, guessing,
                                     options, correct_answer)
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)
                RETURNING id;
            """, q)
            if q[1] == '求极限':
                q1_id = cur.fetchone()[0]
            elif q[1] == '求导数':
                q2_id = cur.fetchone()[0]
            elif q[1] == '计算积分':
                q3_id = cur.fetchone()[0]

        # 6. 创建试卷
        cur.execute("""
            INSERT INTO exam_papers (title, subject_id, description, time_limit, total_score, pass_score)
            VALUES (%s, %s, %s, %s, %s, %s)
            RETURNING id;
        """, ('高等数学测试', math_subject_id, '高等数学基础知识测试', 60, 100, 60))
        exam_paper_id = cur.fetchone()[0]

        # 7. 试卷题目关联
        exam_questions = [
            (exam_paper_id, q1_id, 30, 1),
            (exam_paper_id, q2_id, 30, 2),
            (exam_paper_id, q3_id, 40, 3)
        ]
        for eq in exam_questions:
            cur.execute("""
                INSERT INTO exam_paper_questions (exam_paper_id, question_id, score, "order")
                VALUES (%s, %s, %s, %s);
            """, eq)

        # 8. 创建考试记录（模拟已完成的考试）
        start_time = datetime.now() - timedelta(days=1)
        end_time = start_time + timedelta(minutes=45)
        cur.execute("""
            INSERT INTO exam_records (user_id, exam_paper_id, start_time, end_time, score, status)
            VALUES 
            ((SELECT id FROM users WHERE username = 'student1'), %s, %s, %s, 85, 'completed');
        """, (exam_paper_id, start_time, end_time))

        # 9. 创建答题记录
        responses = [
            (1, q1_id, 'B', True, 180, 30),
            (1, q2_id, 'B', True, 240, 30),
            (1, q3_id, 'B', False, 300, 25)
        ]
        for resp in responses:
            cur.execute("""
                INSERT INTO exam_responses (exam_record_id, question_id, user_answer, 
                                         is_correct, response_time, score)
                VALUES (%s, %s, %s, %s, %s, %s);
            """, resp)

        conn.commit()

def main():
    """主函数"""
    try:
        conn = connect_db()
        print("Connected to database successfully!")

        print("Clearing existing data...")
        clear_data(conn)
        print("Data cleared successfully!")

        print("Initializing test data...")
        init_test_data(conn)
        print("Test data initialized successfully!")

    except Exception as e:
        print(f"Error: {e}")
    finally:
        if conn:
            conn.close()
            print("Database connection closed.")

if __name__ == "__main__":
    main() 
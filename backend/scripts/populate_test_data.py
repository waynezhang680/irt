#!/usr/bin/env python3
import os
import sys
import random
import datetime
import bcrypt
import psycopg2
from psycopg2.extras import Json
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

class TestDataGenerator:
    def __init__(self):
        self.conn = psycopg2.connect(**DB_CONFIG)
        self.cur = self.conn.cursor()
        
    def __del__(self):
        self.cur.close()
        self.conn.close()

    def clean_all_tables(self):
        """清理所有表数据"""
        tables = [
            'ability_estimations',
            'exam_responses',
            'exam_records',
            'exam_paper_questions',
            'exam_papers',
            'question_parameters',
            'question_options',
            'question_knowledge_points',
            'questions',
            'knowledge_points',
            'user_abilities',
            'users',
            'subjects'
        ]
        
        print("开始清理数据...")
        for table in tables:
            self.cur.execute(f"TRUNCATE TABLE {table} CASCADE")
        self.conn.commit()
        print("数据清理完成")

    def generate_roles(self):
        """生成角色数据"""
        roles = [
            ('admin', '系统管理员'),
            ('teacher', '教师'),
            ('student', '学生')
        ]
        
        print("正在生成角色数据...")
        for role_name, description in roles:
            self.cur.execute(
                "INSERT INTO roles (name, description) VALUES (%s, %s) RETURNING id",
                (role_name, description)
            )
        self.conn.commit()
        print("角色数据生成完成")

    def generate_users(self, count=50):
        """生成用户数据"""
        print(f"正在生成{count}个用户...")
        
        # 获取所有角色ID
        self.cur.execute("SELECT id, name FROM roles")
        roles = dict(self.cur.fetchall())
        
        # 确保每个角色至少有一个用户
        for role_id, role_name in roles.items():
            username = f"{role_name}_1"
            password = bcrypt.hashpw("password123".encode(), bcrypt.gensalt()).decode()
            self.cur.execute(
                """INSERT INTO users (username, password, role_id, email)
                VALUES (%s, %s, %s, %s)""",
                (username, password, role_id, f"{username}@example.com")
            )

        # 生成剩余的随机用户
        for i in range(count - len(roles)):
            username = f"user_{i+1}"
            password = bcrypt.hashpw("password123".encode(), bcrypt.gensalt()).decode()
            role_id = random.choice(list(roles.keys()))
            self.cur.execute(
                """INSERT INTO users (username, password, role_id, email)
                VALUES (%s, %s, %s, %s)""",
                (username, password, role_id, f"{username}@example.com")
            )
        
        self.conn.commit()
        print("用户数据生成完成")

    def generate_subjects(self):
        """生成科目数据"""
        subjects = [
            ("数学", "高等数学课程"),
            ("物理", "大学物理课程"),
            ("化学", "普通化学课程"),
            ("计算机基础", "计算机基础知识"),
            ("英语", "大学英语课程")
        ]
        
        print("正在生成科目数据...")
        for name, description in subjects:
            self.cur.execute(
                """INSERT INTO subjects (name, description)
                VALUES (%s, %s)""",
                (name, description)
            )
        self.conn.commit()
        print("科目数据生成完成")

    def generate_knowledge_points(self):
        """生成知识点数据"""
        print("正在生成知识点数据...")
        
        # 获取所有科目
        self.cur.execute("SELECT id, name FROM subjects")
        subjects = self.cur.fetchall()
        
        knowledge_points = {
            "数学": ["极限", "导数", "积分", "线性代数", "概率论"],
            "物理": ["力学", "热学", "电磁学", "光学", "量子力学"],
            "化学": ["无机化学", "有机化学", "物理化学", "分析化学", "生物化学"],
            "计算机基础": ["程序设计", "数据结构", "计算机网络", "操作系统", "数据库"],
            "英语": ["听力", "阅读", "写作", "翻译", "口语"]
        }
        
        for subject_id, subject_name in subjects:
            if subject_name in knowledge_points:
                parent_id = None
                for point in knowledge_points[subject_name]:
                    self.cur.execute(
                        """INSERT INTO knowledge_points (subject_id, name, description, parent_id)
                        VALUES (%s, %s, %s, %s)""",
                        (subject_id, point, f"{point}的基础知识", parent_id)
                    )
        
        self.conn.commit()
        print("知识点数据生成完成")

    def generate_questions(self, count_per_subject=20):
        """生成试题数据"""
        print("正在生成试题数据...")
        
        # 获取所有科目
        self.cur.execute("SELECT id FROM subjects")
        subject_ids = [row[0] for row in self.cur.fetchall()]
        
        question_types = ['single_choice', 'multiple_choice', 'true_false']
        
        for subject_id in subject_ids:
            for i in range(count_per_subject):
                # 生成试题基本信息
                question_type = random.choice(question_types)
                options = self._generate_options(question_type)
                correct_answer = self._generate_correct_answer(question_type, options)
                
                self.cur.execute(
                    """INSERT INTO questions (
                        subject_id, title, content, type, difficulty,
                        discrimination, guessing, options, correct_answer
                    ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s) RETURNING id""",
                    (
                        subject_id,
                        f"题目 {i+1}",
                        f"这是第 {i+1} 道测试题的内容",
                        question_type,
                        random.uniform(0.3, 0.9),  # difficulty
                        random.uniform(0.5, 1.5),  # discrimination
                        random.uniform(0.1, 0.3),  # guessing
                        Json(options),
                        correct_answer
                    )
                )
                
                question_id = self.cur.fetchone()[0]
                
                # 生成选项
                for j, option in enumerate(options):
                    is_correct = option == correct_answer if question_type == 'single_choice' else option in correct_answer.split(',')
                    self.cur.execute(
                        """INSERT INTO question_options (
                            question_id, content, is_correct, label, "order"
                        ) VALUES (%s, %s, %s, %s, %s)""",
                        (question_id, option, is_correct, chr(65+j), j+1)
                    )
                
                # 生成题目参数
                self.cur.execute(
                    """INSERT INTO question_parameters (
                        question_id, difficulty, discrimination, guessing
                    ) VALUES (%s, %s, %s, %s)""",
                    (
                        question_id,
                        random.uniform(0.3, 0.9),  # difficulty
                        random.uniform(0.5, 1.5),  # discrimination
                        random.uniform(0.1, 0.3)   # guessing
                    )
                )
                
                # 关联知识点
                self.cur.execute(
                    "SELECT id FROM knowledge_points WHERE subject_id = %s ORDER BY RANDOM() LIMIT 2",
                    (subject_id,)
                )
                knowledge_point_ids = [row[0] for row in self.cur.fetchall()]
                
                for point_id in knowledge_point_ids:
                    self.cur.execute(
                        """INSERT INTO question_knowledge_points (question_id, knowledge_point_id)
                        VALUES (%s, %s)""",
                        (question_id, point_id)
                    )
        
        self.conn.commit()
        print("试题数据生成完成")

    def _generate_options(self, question_type):
        """生成选项"""
        if question_type == 'true_false':
            return ['True', 'False']
        elif question_type in ['single_choice', 'multiple_choice']:
            return ['选项A', '选项B', '选项C', '选项D']
        return []

    def _generate_correct_answer(self, question_type, options):
        """生成正确答案"""
        if question_type == 'true_false':
            return random.choice(['True', 'False'])
        elif question_type == 'single_choice':
            return random.choice(options)
        elif question_type == 'multiple_choice':
            return ','.join(random.sample(options, random.randint(1, len(options))))
        return ''

    def generate_exam_papers(self, count=10):
        """生成试卷数据"""
        print(f"正在生成{count}份试卷...")
        
        # 获取所有科目
        self.cur.execute("SELECT id FROM subjects")
        subject_ids = [row[0] for row in self.cur.fetchall()]
        
        for i in range(count):
            subject_id = random.choice(subject_ids)
            
            # 生成试卷基本信息
            self.cur.execute(
                """INSERT INTO exam_papers (
                    title, subject_id, description, time_limit,
                    total_score, pass_score, status
                ) VALUES (%s, %s, %s, %s, %s, %s, %s) RETURNING id""",
                (
                    f"测试试卷 {i+1}",
                    subject_id,
                    f"这是第 {i+1} 份测试试卷",
                    120,  # 时长120分钟
                    100,  # 总分100分
                    60,   # 及格分60分
                    'published'
                )
            )
            
            exam_paper_id = self.cur.fetchone()[0]
            
            # 为试卷添加试题
            self.cur.execute(
                "SELECT id FROM questions WHERE subject_id = %s ORDER BY RANDOM() LIMIT 20",
                (subject_id,)
            )
            question_ids = [row[0] for row in self.cur.fetchall()]
            
            for order, question_id in enumerate(question_ids, 1):
                self.cur.execute(
                    """INSERT INTO exam_paper_questions (
                        exam_paper_id, question_id, score, "order"
                    ) VALUES (%s, %s, %s, %s)""",
                    (exam_paper_id, question_id, 5, order)  # 每题5分
                )
        
        self.conn.commit()
        print("试卷数据生成完成")

    def generate_exam_records(self, count=100):
        """生成考试记录数据"""
        print(f"正在生成{count}份考试记录...")
        
        # 获取学生用户
        self.cur.execute("SELECT id FROM users WHERE role_id = (SELECT id FROM roles WHERE name = 'student')")
        student_ids = [row[0] for row in self.cur.fetchall()]
        
        # 获取已发布的试卷
        self.cur.execute("SELECT id FROM exam_papers WHERE status = 'published'")
        exam_paper_ids = [row[0] for row in self.cur.fetchall()]
        
        for i in range(count):
            student_id = random.choice(student_ids)
            exam_paper_id = random.choice(exam_paper_ids)
            
            # 生成考试记录
            start_time = datetime.datetime.now() - datetime.timedelta(days=random.randint(1, 30))
            end_time = start_time + datetime.timedelta(minutes=random.randint(30, 120))
            
            self.cur.execute(
                """INSERT INTO exam_records (
                    user_id, exam_paper_id, start_time, end_time,
                    score, status
                ) VALUES (%s, %s, %s, %s, %s, %s) RETURNING id""",
                (
                    student_id,
                    exam_paper_id,
                    start_time,
                    end_time,
                    random.uniform(0, 100),  # 随机分数
                    'completed'
                )
            )
            
            exam_record_id = self.cur.fetchone()[0]
            
            # 生成答题记录
            self.cur.execute(
                """SELECT q.id, q.correct_answer FROM questions q
                JOIN exam_paper_questions epq ON q.id = epq.question_id
                WHERE epq.exam_paper_id = %s""",
                (exam_paper_id,)
            )
            questions = self.cur.fetchall()
            
            for question_id, correct_answer in questions:
                is_correct = random.choice([True, False])
                user_answer = correct_answer if is_correct else self._generate_wrong_answer(correct_answer)
                
                self.cur.execute(
                    """INSERT INTO exam_responses (
                        exam_record_id, question_id, user_answer,
                        is_correct, response_time, score
                    ) VALUES (%s, %s, %s, %s, %s, %s)""",
                    (
                        exam_record_id,
                        question_id,
                        user_answer,
                        is_correct,
                        random.randint(30, 300),  # 答题时间30-300秒
                        5 if is_correct else 0  # 答对得5分，答错0分
                    )
                )
            
            # 生成能力值估计
            self.cur.execute(
                "SELECT subject_id FROM exam_papers WHERE id = %s",
                (exam_paper_id,)
            )
            subject_id = self.cur.fetchone()[0]
            
            self.cur.execute(
                """INSERT INTO ability_estimations (
                    user_id, subject_id, exam_record_id, ability,
                    standard_error, method
                ) VALUES (%s, %s, %s, %s, %s, %s)""",
                (
                    student_id,
                    subject_id,
                    exam_record_id,
                    random.uniform(-2, 2),  # 能力值
                    random.uniform(0.1, 0.5),  # 标准误
                    'MLE'  # 最大似然估计
                )
            )
        
        self.conn.commit()
        print("考试记录数据生成完成")

    def _generate_wrong_answer(self, correct_answer):
        """生成错误答案"""
        if ',' in correct_answer:  # 多选题
            options = ['选项A', '选项B', '选项C', '选项D']
            return ','.join(random.sample([opt for opt in options if opt not in correct_answer.split(',')], 2))
        elif correct_answer in ['True', 'False']:  # 判断题
            return 'True' if correct_answer == 'False' else 'False'
        else:  # 单选题
            options = ['选项A', '选项B', '选项C', '选项D']
            return random.choice([opt for opt in options if opt != correct_answer])

    def generate_all(self):
        """生成所有测试数据"""
        try:
            self.clean_all_tables()
            self.generate_roles()
            self.generate_users()
            self.generate_subjects()
            self.generate_knowledge_points()
            self.generate_questions()
            self.generate_exam_papers()
            self.generate_exam_records()
            print("\n所有测试数据生成完成！")
        except Exception as e:
            self.conn.rollback()
            print(f"生成测试数据时发生错误: {str(e)}")
            raise

def main():
    try:
        generator = TestDataGenerator()
        generator.generate_all()
    except Exception as e:
        print(f"程序执行失败: {str(e)}")
        sys.exit(1)

if __name__ == "__main__":
    main()
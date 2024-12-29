-- 创建数据库
DROP DATABASE IF EXISTS irt_exam_system;
CREATE DATABASE irt_exam_system;

\c irt_exam_system;

import csv
import psycopg2

# 数据库连接信息
conn = psycopg2.connect(
    dbname='irt_exam_system',
    user='postgres',
    password='Waynez0625@wh',
    host='localhost'
)
cursor = conn.cursor()

# 读取 CSV 文件
with open('试题样例.csv', 'r', encoding='utf-8') as file:
    reader = csv.reader(file)
    next(reader)  # 跳过表头
    for row in reader:
        # 插入到 questions 表
        cursor.execute("""
            INSERT INTO questions (difficulty, score, time_limit, subject, question_type, stem, correct_answer, source, created_by, remarks, status)
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) RETURNING id;
        """, (row[1], row[2], row[3], row[4], row[5], row[6], row[8], row[9], 1, row[11], 'active'))
        
        question_id = cursor.fetchone()[0]  # 获取插入的题目 ID

        # 插入选项到 question_options 表
        options = row[7].split(';')  # 选项分隔
        for i, option in enumerate(options):
            is_correct = (option == row[8])  # 判断是否为正确答案
            cursor.execute("""
                INSERT INTO question_options (question_id, option_label, option_content, is_correct)
                VALUES (%s, %s, %s, %s);
            """, (question_id, chr(65 + i), option, is_correct))  # chr(65 + i) 生成 A, B, C, D

# 提交更改并关闭连接
conn.commit()
cursor.close()
conn.close()

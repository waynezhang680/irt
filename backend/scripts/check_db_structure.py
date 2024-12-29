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

def check_table_structure():
    """检查数据库表结构"""
    try:
        # 连接数据库
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()
        
        # 获取所有表名
        cur.execute("""
            SELECT table_name 
            FROM information_schema.tables 
            WHERE table_schema = 'public'
            AND table_type = 'BASE TABLE'
        """)
        tables = cur.fetchall()
        
        print("数据库中的表：")
        for table in tables:
            table_name = table[0]
            print(f"\n表名: {table_name}")
            
            # 获取表结构
            cur.execute(f"""
                SELECT column_name, data_type, character_maximum_length, 
                       is_nullable, column_default, udt_name
                FROM information_schema.columns 
                WHERE table_schema = 'public' 
                AND table_name = '{table_name}'
                ORDER BY ordinal_position
            """)
            columns = cur.fetchall()
            
            print("字段结构：")
            for col in columns:
                col_name, data_type, max_length, nullable, default, udt_name = col
                print(f"  - {col_name}: {udt_name}", end='')
                if max_length:
                    print(f"({max_length})", end='')
                print(f" {'NULL' if nullable == 'YES' else 'NOT NULL'}", end='')
                if default:
                    print(f" DEFAULT {default}", end='')
                print()
            
            # 获取表的约束
            cur.execute(f"""
                SELECT c.conname, c.contype, pg_get_constraintdef(c.oid)
                FROM pg_constraint c
                JOIN pg_namespace n ON n.oid = c.connamespace
                WHERE conrelid = '{table_name}'::regclass
                AND n.nspname = 'public'
            """)
            constraints = cur.fetchall()
            
            if constraints:
                print("约束：")
                for con_name, con_type, con_def in constraints:
                    print(f"  - {con_name} ({con_def})")
        
        cur.close()
        conn.close()
        
    except Exception as e:
        print(f"检查数据库结构时发生错误: {str(e)}")
        raise

if __name__ == "__main__":
    check_table_structure() 
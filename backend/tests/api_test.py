import requests
import json
import time
import random
from typing import Dict, Optional

class APITester:
    def __init__(self,base_url: str = "http://localhost:8080/api/v1"):
        self.base_url = base_url
        self.token = None
        self.headers = {
            "Content-Type": "application/json"
        }
        self.current_user = None

    def update_auth_header(self, token: str):
        """更新认证头"""
        self.token = token
        self.headers["Authorization"] = f"Bearer {token}"
        print(f"Token updated: {token[:20]}...")

    def cleanup_test_users(self):
        """清理测试用户"""
        print("\n=== 清理测试用户 ===")
        
        # 1. 先登录管理员
        admin_login = self.login_user("admin", "admin123")
        if admin_login.get("code") != 200:
            print("管理员登录失败，无法清理测试数据")
            return
        
        # 2. 获取所有用户列表
        url = f"{self.base_url}/users"
        response = requests.get(url, headers=self.headers)
        if response.status_code == 200:
            users = response.json().get("data", [])
            for user in users:
                username = user.get("username")
                if username and username.startswith("test_"):
                    # 删除测试用户
                    delete_url = f"{self.base_url}/users/{username}"
                    delete_response = requests.delete(delete_url, headers=self.headers)
                    print(f"删除用户 {username}: {delete_response.status_code}")
        else:
            print(f"获取用户列表失败: {response.status_code}")
            if response.text:
                try:
                    print(f"错误详情: {response.json()}")
                except:
                    print(f"响应内容: {response.text}")

    def create_test_exam(self):
        """创建测试考试（需要教师权限）"""
        print("\n=== 创建测试考试 ===")
        
        # 1. 创建测试题目
        questions = self.create_test_questions()  # 确保题目先被创建
        if not questions:
            raise AssertionError("没有可用的测试题目")
        
        # 2. 登录教师账号
        login_result = self.login_user("teacher", "teacher123")
        if login_result.get("code") != 200:
            raise AssertionError(f"教师登录失败: {login_result}")
        
        # 3. 创建考试
        url = f"{self.base_url}/exams"
        exam_data = {
            "title": "测试考试",
            "subject_id": 1,
            "description": "用于API测试的考试",
            "time_limit": 60,
            "total_score": 100,
            "pass_score": 60,
            "questions": [
                {
                    "question_id": q["id"],  # 使用实际的题目 ID
                    "score": 50,
                    "order": i + 1
                } for i, q in enumerate(questions)
            ]
        }
        
        response = requests.post(url, json=exam_data, headers=self.headers)
        print(f"创建考试响应: {response.status_code}")
        
        try:
            result = response.json()
            print(f"创建考试结果: {result}")
            return result
        except Exception as e:
            print(f"解析响应失败: {response.text}")
            raise

    def register_user(self, username: str, password: str) -> Dict:
        """测试用户注册"""
        url = f"{self.base_url}/register"
        data = {
            "username": username,
            "password": password
        }
        response = requests.post(url, json=data, headers=self.headers)
        print(f"\n测试用户注册 - {username}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def login_user(self, username: str, password: str) -> Dict:
        """测试用户登录"""
        url = f"{self.base_url}/login"
        data = {
            "username": username,
            "password": password
        }
        response = requests.post(url, json=data, headers=self.headers)
        print(f"\n测试用户登录 - {username}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        
        if response.status_code == 200:
            token = response.json().get("data", {}).get("token")  # 从 data 中获取 token
            if token:
                self.update_auth_header(token)
            else:
                print("警告: 登录响应中没有找到 token")
        
        return response.json()

    def get_available_exams(self) -> Dict:
        """测试获取可用考试列表"""
        url = f"{self.base_url}/exams"
        response = requests.get(url, headers=self.headers)
        print("\n测试获取可用考试列表")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def start_exam(self, exam_id: int) -> Dict:
        """测试开始考试"""
        url = f"{self.base_url}/exams/{exam_id}/start"
        response = requests.post(url, headers=self.headers)
        print(f"\n测试开始考试 - ID: {exam_id}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def get_next_question(self, exam_id: int) -> Dict:
        """测试获取下一题"""
        url = f"{self.base_url}/exams/{exam_id}/next"
        response = requests.get(url, headers=self.headers)
        print(f"\n测试获取下一题 - 考试ID: {exam_id}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def submit_answer(self, exam_id: int, question_id: int, answer: str, time_spent: int) -> Dict:
        """测试提交答案"""
        url = f"{self.base_url}/exams/{exam_id}/submit"
        data = {
            "question_id": question_id,
            "answer": answer,
            "time_spent": time_spent
        }
        response = requests.post(url, json=data, headers=self.headers)
        print(f"\n测试提交答案  考试ID: {exam_id}, 题目ID: {question_id}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def finish_exam(self, exam_id: int) -> Dict:
        """测试结束考试"""
        url = f"{self.base_url}/exams/{exam_id}/finish"
        response = requests.post(url, headers=self.headers)
        print(f"\n测试结束考试 - ID: {exam_id}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def get_exam_results(self, exam_id: int) -> Dict:
        """测试获取考试结果"""
        url = f"{self.base_url}/exams/{exam_id}/result"
        response = requests.get(url, headers=self.headers)
        print(f"\n测试获取考试结果 - ID: {exam_id}")
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.json()

    def create_test_questions(self) -> list:
        """创建测试题目"""
        print("\n=== 创建测试题目 ===")
        
        # 1. 登录教师账号
        login_result = self.login_user("teacher", "teacher123")
        if login_result.get("code") != 200:
            raise AssertionError(f"教师登录失败: {login_result}")
        
        questions = [
            {
                "subject_id": 1,
                "content": "测试题目1",
                "type": "single_choice",
                "options": ["A", "B", "C", "D"],
                "correct_answer": "A",
                "difficulty": 0.7,
                "discrimination": 0.8,
                "guessing": 0.25
            },
            {
                "subject_id": 1,
                "content": "测试题目2",
                "type": "single_choice",
                "options": ["A", "B", "C", "D"],
                "correct_answer": "B",
                "difficulty": 0.6,
                "discrimination": 0.7,
                "guessing": 0.25
            }
        ]
        
        created_questions = []
        url = f"{self.base_url}/questions"
        
        for q in questions:
            response = requests.post(url, json=q, headers=self.headers)
            print(f"创建题目响应: {response.status_code}")
            if response.status_code == 200:
                result = response.json()
                print(f"创建题目结果: {result}")
                if result.get("code") == 200:
                    created_questions.append(result["data"])
        
        return created_questions

    def cleanup_test_questions(self):
        """清理测试题目"""
        print("\n=== 清理测试题目 ===")
        
        # 1. 登录教师账号
        login_result = self.login_user("teacher", "teacher123")
        if login_result.get("code") != 200:
            print("教师登录失败，无法清理测试题目")
            return
        
        # 2. 获取题目列表
        url = f"{self.base_url}/questions"
        response = requests.get(url, headers=self.headers)
        if response.status_code == 200:
            questions = response.json().get("data", [])
            for q in questions:
                if q["content"].startswith("测试题目"):
                    delete_url = f"{self.base_url}/questions/{q['id']}"
                    delete_response = requests.delete(delete_url, headers=self.headers)
                    print(f"删除题目 {q['id']}: {delete_response.status_code}")

def run_full_test():
    """运行完整的测试流程"""
    tester = APITester()
    
    print("\n=== 开始测试 ===")
    print("Base URL:", tester.base_url)
    
    try:
        # 1. 清理旧的测试数据
        print("\n=== 清理旧数据 ===")
        tester.cleanup_test_users()
        time.sleep(1)  # 等待清理完成
        
        # 2. 测试用户注册
        test_users = [
            {"username": "test_student1", "password": "test123"},
            {"username": "test_student2", "password": "test123"}
        ]
        
        print("\n=== 注册测试用户 ===")
        for user in test_users:
            register_result = tester.register_user(user["username"], user["password"])
            print(f"注册结果 ({user['username']}): {register_result}")
            if register_result.get("code") != 200:
                print(f"警告: 用户 {user['username']} 注册失败")
                continue
        
        # 3. 测试用户登录
        print("\n=== 测试用户登录 ===")
        login_result = tester.login_user(test_users[0]["username"], test_users[0]["password"])
        if login_result.get("code") != 200:
            raise AssertionError(f"登录失败: {login_result}")
        
        print(f"当前 Headers: {tester.headers}")
        time.sleep(1)  # 等待 token 生效
        
        # 4. 测试获取考试列表
        print("\n=== 获取考试列表 ===")
        exams = tester.get_available_exams()
        if exams.get("code") != 200:
            raise AssertionError(f"获取考试列表失败: {exams}")
        
        # 如果没有考试，创建测试考试
        if len(exams.get("data", [])) == 0:
            print("\n=== 创建测试考试 ===")
            exam_result = tester.create_test_exam()
            if exam_result.get("code") != 200:
                raise AssertionError(f"创建考试失败: {exam_result}")
            exam_id = exam_result["data"]["id"]
        else:
            exam_id = exams["data"][0]["id"]
        
        print(f"使用考试 ID: {exam_id}")
        
        # 重新登录学生账号
        tester.login_user(test_users[0]["username"], test_users[0]["password"])
        
        # 5. 开始考试
        print("\n=== 开��考试流程 ===")
        start_result = tester.start_exam(exam_id)
        assert start_result.get("code") == 200, f"开始考试失败: {start_result}"
        
        # 6. 模拟答题流程
        print("\n=== 答题流程 ===")
        question_count = 0
        total_score = 0
        
        while True:
            # 获取下一题
            question = tester.get_next_question(exam_id)
            if question.get("code") != 200:
                print(f"获取题目结束: {question}")
                break
            
            question_count += 1
            print(f"\n--- 第 {question_count} 题 ---")
            
            # 获取题目信息
            question_data = question.get("data", {})
            question_id = question_data.get("id")
            correct_answer = question_data.get("correct_answer", "A")  # 获取正确答案
            
            # 模拟思考时间
            think_time = random.randint(10, 60)  # 10-60秒随机思考时间
            time.sleep(1)  # 实际测试时只等待1秒
            
            # 随机决定是否答对
            is_correct = random.choice([True, False])
            answer = correct_answer if is_correct else "B"
            
            # 提交答案
            submit_result = tester.submit_answer(exam_id, question_id, answer, think_time)
            assert submit_result.get("code") == 200, f"提交答案失败: {submit_result}"
            
            # 累计分数
            score = submit_result.get("data", {}).get("score", 0)
            total_score += score
            
            print(f"答题结果: {'正确' if is_correct else '错误'}, 得分: {score}")
            
            if question_count >= 5:  # 最多答5题
                break
        
        # 7. 结束考试
        print("\n=== 结束考试 ===")
        finish_result = tester.finish_exam(exam_id)
        assert finish_result.get("code") == 200, f"结束考试失败: {finish_result}"
        
        # 8. 获取考试结果
        print("\n=== 获取考试结果 ===")
        result = tester.get_exam_results(exam_id)
        assert result.get("code") == 200, f"获取考试结果失败: {result}"
        
        # 打印考试统计信息
        result_data = result.get("data", {})
        print("\n考试统计:")
        print(f"总题数: {result_data.get('total_questions')}")
        print(f"正确题数: {result_data.get('correct_count')}")
        print(f"总分: {result_data.get('score')}")
        print(f"用时: {result_data.get('time_spent')} 秒")
        print(f"能力值: {result_data.get('ability')}")
        
        # 打印详细答题记录
        print("\n答题详情:")
        for resp in result_data.get("responses", []):
            print(f"题目 {resp['question_id']}: "
                  f"{'正确' if resp['is_correct'] else '错误'}, "
                  f"得分: {resp['score']}, "
                  f"用时: {resp['response_time']}秒")

    except Exception as e:
        print(f"\n❌ 测试过程中出错: {str(e)}")
        raise
    
    finally:
        # 清理测试数据
        print("\n=== 最终清理 ===")
        tester.cleanup_test_questions()  # 先清理题目
        tester.cleanup_test_users()      # 再清理用户

if __name__ == "__main__":
    try:
        run_full_test()
        print("\n✅ 所有测试通过!")
    except AssertionError as e:
        print(f"\n❌ 测试失败: {str(e)}")
    except Exception as e:
        print(f"\n❌ 测试出错: {str(e)}")
        import traceback
        traceback.print_exc()
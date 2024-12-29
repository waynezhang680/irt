import requests

BASE_URL = "http://localhost:8080/api/v1"
HEADERS = {
    "Content-Type": "application/json"
}

def user_exists(username):
    """检查用户是否存在"""
    url = f"{BASE_URL}/users/{username}"
    response = requests.get(url, headers=HEADERS)
    return response.status_code == 200

def delete_user(username):
    """删除用户"""
    url = f"{BASE_URL}/users/{username}"
    response = requests.delete(url, headers=HEADERS)
    print(f"删除用户 {username} - 状态码: {response.status_code}, 响应: {response.json()}")

def create_user(username, password):
    """创建用户"""
    if user_exists(username):
        print(f"用户 {username} 已存在，准备删除并重新创建。")
        delete_user(username)  # 删除已存在的用户

    url = f"{BASE_URL}/register"
    data = {
        "username": username,
        "password": password
    }
    response = requests.post(url, json=data, headers=HEADERS)
    try:
        response_data = response.json()
    except ValueError:
        response_data = {"error": "Invalid JSON response", "text": response.text}
    
    print(f"创建用户 {username} - 状态码: {response.status_code}, 响应: {response_data}")

def subject_exists(subject_name):
    """检查科目是否存在"""
    url = f"{BASE_URL}/subjects"
    response = requests.get(url, headers=HEADERS)
    if response.status_code == 200:
        subjects = response.json().get("data", [])
        return any(subject['name'] == subject_name for subject in subjects)
    return False

def delete_subject(subject_name):
    """删除科目"""
    url = f"{BASE_URL}/subjects/{subject_name}"  # 假设有删除科目的 API
    response = requests.delete(url, headers=HEADERS)
    print(f"删除科目 {subject_name} - 状态码: {response.status_code}, 响应: {response.json()}")

def create_subject(subject_name):
    """创建科目"""
    if subject_exists(subject_name):
        print(f"科目 {subject_name} 已存在，准备删除并重新创建。")
        delete_subject(subject_name)  # 删除已存在的科目

    url = f"{BASE_URL}/subjects"
    data = {
        "name": subject_name
    }
    response = requests.post(url, json=data, headers=HEADERS)
    print(f"创建科目 {subject_name} - 状态码: {response.status_code}, 响应: {response.json()}")

def create_question(subject_id, content, correct_answer):
    """创建题目"""
    url = f"{BASE_URL}/questions"
    data = {
        "subject_id": subject_id,
        "content": content,
        "type": "single_choice",
        "options": ["A", "B", "C", "D"],
        "correct_answer": correct_answer,
        "difficulty": 0.5,
        "discrimination": 0.5,
        "guessing": 0.25
    }
    response = requests.post(url, json=data, headers=HEADERS)
    print(f"创建题目 '{content}' - 状态码: {response.status_code}, 响应: {response.json()}")

def create_exam(title, subject_id, questions):
    """创建考试"""
    url = f"{BASE_URL}/exams"
    data = {
        "title": title,
        "subject_id": subject_id,
        "description": "测试考试",
        "time_limit": 60,
        "total_score": 100,
        "pass_score": 60,
        "questions": questions  # 这里填充题目 ID
    }
    response = requests.post(url, json=data, headers=HEADERS)
    print(f"创建考试 '{title}' - 状态码: {response.status_code}, 响应: {response.json()}")

def main():
    # 创建用户
    create_user("admin", "admin123")
    create_user("teacher", "teacher123")
    create_user("test_student1", "test123")
    create_user("test_student2", "test123")

    # 创建科目
    create_subject("数学")
    create_subject("英语")

    # 创建题目
    create_question(1, "测试题目1", "A")
    create_question(1, "测试题目2", "B")

    # 创建考试
    questions = [
        {"question_id": 1, "score": 50, "order": 1},
        {"question_id": 2, "score": 50, "order": 2}
    ]
    create_exam("测试考试", 1, questions)

if __name__ == "__main__":
    main()
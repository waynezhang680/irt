# IRT Exam System API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication
All API endpoints except login and register require JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

## Error Response Format
```json
{
    "message": "Error message",
    "status": 400
}
```

## Endpoints

### Authentication

#### Login
```
POST /auth/login
```
Request:
```json
{
    "username": "string",
    "password": "string"
}
```
Response:
```json
{
    "token": "string",
    "user": {
        "id": "string",
        "username": "string",
        "role": "string"
    }
}
```

#### Register
```
POST /auth/register
```
Request:
```json
{
    "username": "string",
    "password": "string",
    "email": "string"
}
```
Response:
```json
{
    "message": "Registration successful",
    "userId": "string"
}
```

### Exams

#### Get Exam List
```
GET /exams
```
Response:
```json
{
    "exams": [
        {
            "id": "string",
            "title": "string",
            "duration": "number",
            "totalQuestions": "number",
            "difficulty": "string",
            "status": "string"
        }
    ]
}
```

#### Get Exam Details
```
GET /exams/{examId}
```
Response:
```json
{
    "id": "string",
    "title": "string",
    "duration": "number",
    "questions": [
        {
            "id": "number",
            "question": "string",
            "options": ["string"],
            "difficulty": "number"
        }
    ]
}
```

#### Submit Answer
```
POST /exams/{examId}/answers
```
Request:
```json
{
    "questionId": "number",
    "answer": "number",
    "timeSpent": "number"
}
```
Response:
```json
{
    "message": "Answer submitted successfully"
}
```

#### Submit Exam
```
POST /exams/{examId}/submit
```
Request:
```json
{
    "examId": "string",
    "answers": {
        "questionId": "number"
    },
    "timeSpent": "number"
}
```
Response:
```json
{
    "message": "Exam submitted successfully",
    "resultId": "string"
}
```

### Results

#### Get Exam Result
```
GET /results/{examId}
```
Response:
```json
{
    "examId": "string",
    "title": "string",
    "score": "number",
    "totalQuestions": "number",
    "correctAnswers": "number",
    "incorrectAnswers": "number",
    "timeTaken": "string",
    "difficulty": "string",
    "submittedAt": "string"
}
```

## Status Codes

- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Data Types

### Difficulty Levels
- "Easy"
- "Medium"
- "Hard"

### Exam Status
- "Available"
- "In Progress"
- "Completed"

### User Roles
- "student"
- "teacher"
- "admin"

## Notes

1. All timestamps are in ISO 8601 format
2. All IDs are strings unless specified otherwise
3. Question IDs and answers are numbers
4. Time values are in seconds
5. Difficulty values are between 0 and 1 for IRT calculations 
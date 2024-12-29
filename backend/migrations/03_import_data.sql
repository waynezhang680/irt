/* 为已有的题目创建选项 */
INSERT INTO question_options (
    question_id, 
    content, 
    is_correct, 
    created_at, 
    updated_at
)
SELECT 
    q.id as question_id,
    value as content,
    key = q.correct_answer as is_correct,
    q.created_at,
    q.updated_at
FROM 
    questions q,
    jsonb_each_text(q.options) as opt(key, value)
WHERE 
    q.options IS NOT NULL; 
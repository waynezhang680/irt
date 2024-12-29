-- 创建 question_options 表
CREATE TABLE IF NOT EXISTS question_options (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    question_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL,
    CONSTRAINT fk_question_options_question 
        FOREIGN KEY (question_id) 
        REFERENCES questions(id) 
        ON DELETE CASCADE
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_question_options_deleted_at ON question_options(deleted_at);
CREATE INDEX IF NOT EXISTS idx_question_options_question_id ON question_options(question_id);

-- 为已有的题目创建选项
WITH RECURSIVE options_expanded AS (
    SELECT 
        q.id as question_id,
        q.correct_answer,
        q.created_at,
        q.updated_at,
        key as option_key,
        value as option_content
    FROM 
        questions q,
        jsonb_each_text(q.options) as opt(key, value)
    WHERE 
        q.options IS NOT NULL
)
INSERT INTO question_options (
    question_id, 
    content, 
    is_correct, 
    created_at, 
    updated_at
)
SELECT 
    question_id,
    option_content,
    option_key = correct_answer,
    created_at,
    updated_at
FROM 
    options_expanded;

-- 创建触发器以保持选项和答案的一致性
CREATE OR REPLACE FUNCTION update_question_correct_answer()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.is_correct THEN
        WITH option_key AS (
            SELECT key 
            FROM questions q,
                 jsonb_each_text(q.options) as opt(key, value)
            WHERE q.id = NEW.question_id 
            AND opt.value = NEW.content
            LIMIT 1
        )
        UPDATE questions 
        SET correct_answer = (SELECT key FROM option_key)
        WHERE id = NEW.question_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER question_option_correct_answer_trigger
AFTER INSERT OR UPDATE ON question_options
FOR EACH ROW
EXECUTE FUNCTION update_question_correct_answer(); 
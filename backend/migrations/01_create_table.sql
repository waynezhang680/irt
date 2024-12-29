/* 创建 question_options 表 */
CREATE TABLE question_options (
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

/* 创建索引 */
CREATE INDEX idx_question_options_deleted_at ON question_options(deleted_at);
CREATE INDEX idx_question_options_question_id ON question_options(question_id); 
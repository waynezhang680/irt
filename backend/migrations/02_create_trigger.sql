/* 创建触发器函数 */
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

/* 创建触发器 */
CREATE TRIGGER question_option_correct_answer_trigger
AFTER INSERT OR UPDATE ON question_options
FOR EACH ROW
EXECUTE FUNCTION update_question_correct_answer(); 
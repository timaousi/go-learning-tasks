-- 插入
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');

-- 查询年龄 > 18
SELECT * FROM students WHERE age > 18;

-- 更新张三的年级
UPDATE students SET grade = '四年级' WHERE name = '张三';

-- 删除年龄小于 15 的学生
DELETE FROM students WHERE age < 15;

-- 假设 A 的 ID 是 1，B 的 ID 是 2
START TRANSACTION;

-- 检查余额是否足够
SELECT balance FROM accounts WHERE id = 1 FOR UPDATE;

-- 假设 balance >= 100
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;

-- 插入交易记录
INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (1, 2, 100);

COMMIT;
-- 如果余额不足则 ROLLBACK;
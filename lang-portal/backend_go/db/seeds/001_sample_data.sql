-- Insert sample words
INSERT INTO words (spanish, english, parts) VALUES
('hola', 'hello', '{"type": "greeting"}'),
('gracias', 'thank you', '{"type": "courtesy"}'),
('por favor', 'please', '{"type": "courtesy"}'),
('buenos días', 'good morning', '{"type": "greeting"}'),
('buenas noches', 'good night', '{"type": "greeting"}');

-- Insert sample groups
INSERT INTO groups (name) VALUES
('Basic Greetings'),
('Courtesy Words'),
('Time Expressions');

-- Link words to groups
INSERT INTO words_groups (word_id, group_id) VALUES
(1, 1), -- hola -> Basic Greetings
(2, 2), -- gracias -> Courtesy Words
(3, 2), -- por favor -> Courtesy Words
(4, 1), -- buenos días -> Basic Greetings
(4, 3), -- buenos días -> Time Expressions
(5, 1), -- buenas noches -> Basic Greetings
(5, 3); -- buenas noches -> Time Expressions

-- Insert a sample study session
INSERT INTO study_sessions (group_id, study_activity_id, created_at) VALUES
(1, 1, datetime('now', '-1 day')),
(2, 1, datetime('now', '-2 days'));

-- Insert sample word reviews
INSERT INTO word_review_items (word_id, study_session_id, correct, created_at) VALUES
(1, 1, 1, datetime('now', '-1 day')),
(2, 1, 1, datetime('now', '-1 day')),
(3, 1, 0, datetime('now', '-1 day')),
(4, 2, 1, datetime('now', '-2 days')),
(5, 2, 1, datetime('now', '-2 days'));

INSERT INTO User (email, username, password_hash) VALUES
('john.doe@example.com', 'john_doe', 'hashed_password_123'),
('jane.smith@example.com', 'jane_smith', 'hashed_password_456'),
('alice.wonder@example.com', 'alice_wonder', 'hashed_password_789');

INSERT INTO Category (name) VALUES
('Technology'),
('Lifestyle'),
('Education'),
('Health'),
('Entertainment');

INSERT INTO Post (user_id, title, content, category_id) VALUES
(1, 'The Future of AI', 'Artificial intelligence is rapidly advancing...', 1),
(2, 'Healthy Living Tips', 'Here are my top tips for living a healthier life...', 4),
(3, 'Learning Online', 'Online learning has become more prevalent...', 3);

INSERT INTO Comment (post_id, user_id, content) VALUES
(1, 2, 'Great insights on AI!'),
(2, 3, 'Thanks for sharing these tips!'),
(3, 1, 'I totally agree about online learning.');

INSERT INTO LikeDislike (user_id, post_id, comment_id, type) VALUES
(2, 1, NULL, 1), -- User 2 likes Post 1
(3, 2, NULL, 1), -- User 3 likes Post 2
(1, NULL, 1, 1), -- User 1 likes Comment 1
(2, NULL, 2, -1); -- User 2 dislikes Comment 2

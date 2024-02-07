INSERT INTO user (email_address, password_hash) VALUES 
('system@xost-tech.com', '$2a$10$WG1O6xWxsgjWEO3w7N1xl.Bdlf6AeZR7fresrHm11Ypl0m4MFLkJW'),
('xost.technologies@gmail.com', '$2a$10$jLovMw5HoG55oAbXXgByMOih1Z1gSnZpYPYbiR1w8H/a/6Cd8augK');

INSERT INTO category (user_id, category_name) VALUES
(1, 'Personal'),
(1, 'Work'),
(1, 'Study'),
(1, 'Health'),
(1, 'Shopping'),
(1, 'Wishlist'),
(1, 'Birthday');

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `address` text,
  `email` varchar(100) UNIQUE NOT NULL,
  `password` text NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `recipes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `likes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `recipe_id` int,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `comments` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `recipe_id` int,
  `content` text NOT NULL,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `save_recipes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `recipe_id` int,
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `recipes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`);

ALTER TABLE `save_recipes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `save_recipes` ADD FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`);

CREATE DATABASE IF NOT EXISTS service_pace;

USE service_pace;

CREATE TABLE IF NOT EXISTS `notes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` text NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE IF NOT EXISTS `users` (
  `id` CHAR(36) PRIMARY KEY,
  `name` varchar(100) NOT NULL,
  `address` text,
  `email` varchar(100) UNIQUE NOT NULL,
  `password` text NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS `recipes` (
  `id` CHAR(36) PRIMARY KEY,
  `user_id` CHAR(36),
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `deleted_at` timestamp DEFAULT null
);

CREATE TABLE IF NOT EXISTS `likes` (
  `id` CHAR(36) PRIMARY KEY,
  `user_id` CHAR(36),
  `recipe_id` CHAR(36),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS `comments` (
  `id` CHAR(36) PRIMARY KEY,
  `user_id` CHAR(36),
  `recipe_id` CHAR(36),
  `content` text NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS `save_recipes` (
  `id` CHAR(36) PRIMARY KEY,
  `user_id` CHAR(36),
  `recipe_id` CHAR(36),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS `notifications` (
  `id` CHAR(36) PRIMARY KEY,
  `user_id` CHAR(36),
  `type` ENUM ('like', 'comment', 'save') NOT NULL DEFAULT 'like',
  `message` text NOT NULL,
  `created_at` timestamp DEFAULT (now())
);

ALTER TABLE `recipes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`);

ALTER TABLE `save_recipes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `save_recipes` ADD FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

CREATE TABLE `countries` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `iso` char(2) NOT NULL,
  `name` varchar(80) NOT NULL,
  `nicename` varchar(80) NOT NULL,
  `iso3` char(3) DEFAULT NULL,
  `numcode` smallint DEFAULT NULL,
  `phonecode` int NOT NULL
);

CREATE TABLE `roles` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(50) UNIQUE NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  `description` VARCHAR(255),
  `created_at` bigint(20) NOT NULL,
  `createdBy` bigint(20) NOT NULL
);

CREATE TABLE `organizations` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200),
  `description` VARCHAR(255),
  `website` VARCHAR(50),
  `address1` VARCHAR(255),
  `address2` VARCHAR(255),
  `country_id` bigint(20) NOT NULL,
  `status` int(1),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `users` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `firstName` VARCHAR(100),
  `lastName` VARCHAR(100) NOT NULL,
  `password` VARCHAR(50) NOT NULL,
  `email` VARCHAR(255),
  `username` VARCHAR(255),
  `phone` VARCHAR(50) NOT NULL,
  `organization_id` bigint(20) NOT NULL,
  `country_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  `joinedDate` bigint(20) NOT NULL,
  `lastOnline` bigint(20) NOT NULL,
  `address1` VARCHAR(255),
  `address2` VARCHAR(255),
  `profileUrl` VARCHAR(255),
  `status` int(1),
  `inviteBy` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `teams` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(150) NOT NULL,
  `description` VARCHAR(255),
  `mage_url` VARCHAR(255),
  `role_id` bigint(20) NOT NULL,
  `organization_id` bigint(20) NOT NULL,
  `status` int(1),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `teams_users` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `team_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `categories` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(150) NOT NULL,
  `description` VARCHAR(255),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `tags` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(150) NOT NULL,
  `description` VARCHAR(255),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `courses` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` longtext COLLATE utf8mb4_unicode_ci,
  `long_description` longtext COLLATE utf8mb4_unicode_ci,
  `image_url` mediumtext COLLATE utf8mb4_unicode_ci,
  `duration` bigint(20) NOT NULL DEFAULT '0',
  `author_id` bigint(20), /* TODO: Make it not null */
  `category_id` bigint(20),
  `published_at` bigint(20) DEFAULT NULL,
  `organization_id` bigint(20), /* TODO: Make it not null */
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'DRAFT',
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `enrollments` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT
);

CREATE TABLE `courses_users_enrollments` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` bigint(20) NOT NULL,
  `userID` bigint(20) NOT NULL,
  `status` int(1)
);

CREATE TABLE `courses_teams_enrollments` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` bigint(20) NOT NULL,
  `team_id` bigint(20) NOT NULL,
  `status` int(1)
);

CREATE TABLE `courses_tags` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `course_id` bigint(20) NOT NULL,
  `tag_id` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `lessons_tags` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `lesson_id` bigint(20) NOT NULL,
  `tag_id` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `lessons` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(256),
  `description` longtext,
  `course_id` bigint(20) NOT NULL,
  `order` int(3),
  `status` int(1),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `contents` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(256),
  `description` longtext,
  `content` longtext,
  `fileheader` varchar(60),
  `embed_url` varchar(256),
  `image_url` varchar(256),
  `lesson_id` bigint(20) NOT NULL,
  `size` int,
  `caption` varchar(256),
  `order` int(3),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `attachments` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(256),
  `description` longtext,
  `name` varchar(256),
  `size` int,
  `type` varchar(60),
  `course_id` bigint(20) NULL,
  `status` int(1),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE;

ALTER TABLE `users` ADD FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE;

ALTER TABLE `teams` ADD FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE;

ALTER TABLE `teams` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE;

ALTER TABLE `teams_users` ADD FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`) ON DELETE CASCADE;

ALTER TABLE `teams_users` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `organizations` ADD FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE;

ALTER TABLE `users` ADD FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses_tags` ADD FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses_tags` ADD FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE;

ALTER TABLE `attachments` ADD FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE;

ALTER TABLE `lessons` ADD FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE;

ALTER TABLE `lessons_tags` ADD FOREIGN KEY (`lesson_id`) REFERENCES `lessons` (`id`) ON DELETE CASCADE;

ALTER TABLE `lessons_tags` ADD FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE;

ALTER TABLE `contents` ADD FOREIGN KEY (`lesson_id`) REFERENCES `lessons` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses` ADD FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses_users_enrollments` ADD FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses_users_enrollments` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses_teams_enrollments` ADD FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE;

ALTER TABLE `courses_teams_enrollments` ADD FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`) ON DELETE CASCADE;

CREATE INDEX `index_on_country_id` ON `organizations` (`country_id`);

CREATE INDEX `index_on_country_id` ON `users` (`country_id`);

CREATE INDEX `index_on_role_id` ON `users` (`role_id`);

CREATE INDEX `index_on_organization_id` ON `users` (`organization_id`);

CREATE INDEX `index_on_role_id` ON `teams` (`role_id`);

CREATE INDEX `index_on_organization_id` ON `teams` (`organization_id`);

CREATE INDEX `index_on_team_id` ON `teams_users` (`team_id`);

CREATE INDEX `index_on_user_id` ON `teams_users` (`user_id`);

CREATE INDEX `index_on_author_id` ON `courses` (`author_id`);

CREATE INDEX `index_on_category_id` ON `courses` (`category_id`);
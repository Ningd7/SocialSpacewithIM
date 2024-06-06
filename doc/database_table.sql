SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- Table `comments`
-- 用于存储评论信息
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
                            `id` int(16) NOT NULL AUTO_INCREMENT,
                            `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
                            `userId` int(16) NOT NULL,
                            `postId` int(16) NOT NULL,
                            PRIMARY KEY (`id`),
                            INDEX `comment_post_idx` (`postId`),
                            INDEX `comment_user_idx` (`userId`),
                            CONSTRAINT `fk_comment_user` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE,
                            CONSTRAINT `fk_comment_post` FOREIGN KEY (`postId`) REFERENCES `posts`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论表';

-- Table `likes`
-- 用于存储点赞信息
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
                         `id` int(16) NOT NULL AUTO_INCREMENT,
                         `userId` int(16) NOT NULL,
                         `postId` int(16) NOT NULL,
                         PRIMARY KEY (`id`),
                         INDEX `like_user_idx` (`userId`),
                         INDEX `like_post_idx` (`postId`),
                         CONSTRAINT `fk_like_user` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE,
                         CONSTRAINT `fk_like_post` FOREIGN KEY (`postId`) REFERENCES `posts`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=21 CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='点赞表';

-- Table `posts`
-- 用于存储帖子信息
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
                         `id` int(16) NOT NULL AUTO_INCREMENT,
                         `description` varchar(280) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `img` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `userId` int(16) NOT NULL,
                         `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         INDEX `post_user_idx` (`userId`),
                         CONSTRAINT `fk_post_user` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=15 CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='帖子表';

-- Table `relationships`
-- 用于存储用户之间的关注关系
DROP TABLE IF EXISTS `relationships`;
CREATE TABLE `relationships` (
                                 `id` int(16) NOT NULL AUTO_INCREMENT,
                                 `followerUserId` int(16) NOT NULL,
                                 `followedUserId` int(16) NOT NULL,
                                 PRIMARY KEY (`id`),
                                 INDEX `follower_idx` (`followerUserId`),
                                 INDEX `followed_idx` (`followedUserId`),
                                 CONSTRAINT `fk_follower_user` FOREIGN KEY (`followerUserId`) REFERENCES `users`(`id`) ON DELETE CASCADE,
                                 CONSTRAINT `fk_followed_user` FOREIGN KEY (`followedUserId`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户关系表';

-- Table `stories`
-- 用于存储用户的故事信息
DROP TABLE IF EXISTS `stories`;
CREATE TABLE `stories` (
                           `id` int(16) NOT NULL AUTO_INCREMENT,
                           `img` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                           `userId` int(16) NOT NULL,
                           PRIMARY KEY (`id`),
                           INDEX `story_user_idx` (`userId`),
                           CONSTRAINT `fk_story_user` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='故事表';

-- Table `users`
-- 用于存储用户的信息
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int(16) NOT NULL AUTO_INCREMENT,
                         `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `gender` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `password` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `coverPic` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `profilePic` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `city` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `website` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;

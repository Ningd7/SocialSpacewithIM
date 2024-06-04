SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- Table `comments`
-- 用于存储评论信息
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
                            `id` int(16) NOT NULL AUTO_INCREMENT, -- 评论的唯一标识符
                            `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL, -- 评论内容
                            `createdAt` datetime(6) NULL DEFAULT NULL, -- 评论创建时间
                            `userId` int(16) NOT NULL, -- 评论的用户ID
                            `postId` int(16) NOT NULL, -- 评论所属帖子ID
                            PRIMARY KEY (`id`) USING BTREE, -- 主键
                            UNIQUE INDEX `commentUserId` (`userId`) USING BTREE, -- 用户ID的唯一索引
                            INDEX `postId` (`postId`) USING BTREE, -- 帖子ID的索引
                            CONSTRAINT `commentUserId` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE, -- 用户ID的外键约束
                            CONSTRAINT `postId` FOREIGN KEY (`postId`) REFERENCES `posts`(`id`) ON DELETE CASCADE -- 帖子ID的外键约束
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT='评论表';

-- Table `likes`
-- 用于存储点赞信息
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
                         `id` int(16) NOT NULL AUTO_INCREMENT, -- 点赞的唯一标识符
                         `userId` int(16) NOT NULL, -- 点赞的用户ID
                         `postId` int(16) NOT NULL, -- 点赞所属帖子ID
                         PRIMARY KEY (`id`) USING BTREE, -- 主键
                         UNIQUE INDEX `id_INDEX` (`id`) USING BTREE, -- 唯一索引
                         INDEX `likeUserId` (`userId`) USING BTREE, -- 用户ID的索引
                         INDEX `likePostId` (`postId`) USING BTREE, -- 帖子ID的索引
                         CONSTRAINT `likeUserId` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE, -- 用户ID的外键约束
                         CONSTRAINT `likePostId` FOREIGN KEY (`postId`) REFERENCES `posts`(`id`) ON DELETE CASCADE -- 帖子ID的外键约束
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT='点赞表';

-- Table `posts`
-- 用于存储帖子信息
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
                         `id` int(16) NOT NULL AUTO_INCREMENT, -- 帖子的唯一标识符
                         `desc` varchar(280) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL, -- 帖子内容
                         `img` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL, -- 帖子图片
                         `userId` int(16) NOT NULL, -- 发帖用户ID
                         `createdAt` datetime(6) NULL DEFAULT NULL, -- 帖子创建时间
                         PRIMARY KEY (`id`) USING BTREE, -- 主键
                         UNIQUE INDEX `id_INDEX` (`id`) USING BTREE, -- 唯一索引
                         INDEX `postUserId` (`userId`) USING BTREE, -- 用户ID的索引
                         CONSTRAINT `postUserId` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE -- 用户ID的外键约束
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT='帖子表';

-- Table `relationships`
-- 用于存储用户之间的关注关系
DROP TABLE IF EXISTS `relationships`;
CREATE TABLE `relationships` (
                                 `id` int(16) NOT NULL AUTO_INCREMENT, -- 关系的唯一标识符
                                 `followerUserId` int(16) NOT NULL COMMENT '关注者', -- 关注者用户ID
                                 `followedUserId` int(16) NOT NULL COMMENT '被关注者', -- 被关注者用户ID
                                 PRIMARY KEY (`id`) USING BTREE, -- 主键
                                 UNIQUE INDEX `id_INDEX` (`id`) USING BTREE, -- 唯一索引
                                 INDEX `followerUser` (`followerUserId`) USING BTREE, -- 关注者用户ID的索引
                                 INDEX `followedUser` (`followedUserId`) USING BTREE, -- 被关注者用户ID的索引
                                 CONSTRAINT `followerUser` FOREIGN KEY (`followerUserId`) REFERENCES `users`(`id`) ON DELETE CASCADE, -- 关注者用户ID的外键约束
                                 CONSTRAINT `followedUser` FOREIGN KEY (`followedUserId`) REFERENCES `users`(`id`) ON DELETE CASCADE -- 被关注者用户ID的外键约束
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT='用户关系表';

-- Table `stories`
-- 用于存储用户的故事信息
DROP TABLE IF EXISTS `stories`;
CREATE TABLE `stories` (
                           `id` int(16) NOT NULL AUTO_INCREMENT, -- 故事的唯一标识符
                           `img` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL, -- 故事图片
                           `userId` int(16) NOT NULL, -- 故事的用户ID
                           PRIMARY KEY (`id`) USING BTREE, -- 主键
                           UNIQUE INDEX `id_INDEX` (`id`) USING BTREE, -- 唯一索引
                           INDEX `storyUserId` (`userId`) USING BTREE, -- 用户ID的索引
                           CONSTRAINT `storyUserId` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE -- 用户ID的外键约束
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT='故事表';

-- Table `users`
-- 用于存储用户的信息
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int(16) NOT NULL AUTO_INCREMENT, -- 用户的唯一标识符
                         `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL, -- 用户名
                         `gender` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL, -- 性别
                         `email` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL, -- 用户邮箱
                         `password` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL, -- 用户密码
                         `coverPic` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL, -- 封面图片
                         `profilePic` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL, -- 头像图片
                         `city` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL, -- 用户所在城市
                         `website` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL, -- 用户网站
                         PRIMARY KEY (`id`) USING BTREE, -- 主键
                         UNIQUE INDEX `id_INDEX` (`id`) USING BTREE -- 唯一索引
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;

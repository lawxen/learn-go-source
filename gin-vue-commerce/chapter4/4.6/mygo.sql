/*
 Navicat Premium Data Transfer

 Source Server         : MyDB
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : mygo

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 12/08/2023 16:10:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for programs
-- ----------------------------
DROP TABLE IF EXISTS `programs`;
CREATE TABLE `programs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `online` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of programs
-- ----------------------------
INSERT INTO `programs` VALUES (1, '西游记', '2023-08-12 12:03:05.084');
INSERT INTO `programs` VALUES (2, '红楼梦', '2023-08-12 12:03:05.102');
INSERT INTO `programs` VALUES (3, '三国A', '2023-08-08 12:02:41.000');
INSERT INTO `programs` VALUES (4, '水浒传A', '2023-08-12 12:03:05.108');
INSERT INTO `programs` VALUES (5, '封神榜', '2023-08-12 12:03:05.112');
INSERT INTO `programs` VALUES (6, '姜子牙A', '2023-08-12 12:03:05.115');
INSERT INTO `programs` VALUES (7, '哪吒', NULL);
INSERT INTO `programs` VALUES (8, '大闹天宫', '2023-08-12 12:02:41.035');
INSERT INTO `programs` VALUES (9, '流浪地球', '2023-08-12 12:02:41.035');
INSERT INTO `programs` VALUES (10, '西游记', '2023-08-12 12:03:05.100');

-- ----------------------------
-- Table structure for user_program
-- ----------------------------
DROP TABLE IF EXISTS `user_program`;
CREATE TABLE `user_program`  (
  `user_id` bigint UNSIGNED NOT NULL,
  `program_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`user_id`, `program_id`) USING BTREE,
  INDEX `fk_user_program_program`(`program_id`) USING BTREE,
  CONSTRAINT `fk_user_program_program` FOREIGN KEY (`program_id`) REFERENCES `programs` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_program_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_program
-- ----------------------------
INSERT INTO `user_program` VALUES (1, 1);
INSERT INTO `user_program` VALUES (2, 8);
INSERT INTO `user_program` VALUES (2, 9);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `nationality` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'Tom', 'China');
INSERT INTO `users` VALUES (2, 'Tim', 'China');

SET FOREIGN_KEY_CHECKS = 1;

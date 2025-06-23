/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80035
 Source Host           : localhost:3306
 Source Schema         : babys

 Target Server Type    : MySQL
 Target Server Version : 80035
 File Encoding         : 65001

 Date: 14/11/2023 15:51:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'p', 'admin', '/auth/', 'GET', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for jwts
-- ----------------------------
DROP TABLE IF EXISTS `jwts`;
CREATE TABLE `jwts`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `token` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `expire` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwts
-- ----------------------------
INSERT INTO `jwts` VALUES (1, '2023-11-14 15:12:13.558', '2023-11-14 15:31:50.605', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoi5aWl5Yqb57uZIn0.8fOAjkx_OI4yueFXzcUieNDOqXbhO0rZMK2gZJtiPH0', '2023-11-14 16:01:50.604');
INSERT INTO `jwts` VALUES (2, '2023-11-14 15:12:25.744', '2023-11-14 15:12:25.744', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoi5aWl5Yqb57uZIn0.8fOAjkx_OI4yueFXzcUieNDOqXbhO0rZMK2gZJtiPH0', '2023-11-14 15:42:25.744');
INSERT INTO `jwts` VALUES (3, '2023-11-14 15:33:46.191', '2023-11-14 15:33:46.191', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoi5aWl5Yqb57uZIn0.8fOAjkx_OI4yueFXzcUieNDOqXbhO0rZMK2gZJtiPH0', '2023-11-14 16:03:46.190');

SET FOREIGN_KEY_CHECKS = 1;

/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80035
 Source Host           : localhost:3306
 Source Schema         : baby

 Target Server Type    : MySQL
 Target Server Version : 80035
 File Encoding         : 65001

 Date: 25/04/2024 16:48:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `quantity` bigint(0) NULL DEFAULT NULL,
  `commodity_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_carts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of carts
-- ----------------------------
INSERT INTO `carts` VALUES (1, '2024-04-25 16:37:47.701', '2024-04-25 16:37:47.701', NULL, 1, 1, 1);

-- ----------------------------
-- Table structure for commodities
-- ----------------------------
DROP TABLE IF EXISTS `commodities`;
CREATE TABLE `commodities`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `types` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `price` double NULL DEFAULT NULL,
  `discount` double NULL DEFAULT NULL,
  `stock` bigint(0) NULL DEFAULT NULL,
  `sold` bigint(0) NULL DEFAULT NULL,
  `likes` bigint(0) NULL DEFAULT NULL,
  `created` datetime(3) NULL DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `details` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_commodities_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of commodities
-- ----------------------------
INSERT INTO `commodities` VALUES (1, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '婴儿衣服加厚连体衣', '粉色', '童装', 199, 188, 1314, 1666, 666, '2020-02-24 00:00:00.000', 'static/imgs/p1.jpg', 'static/details/p1_details.jpg');
INSERT INTO `commodities` VALUES (2, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '儿童长袖加厚防寒加绒衫', '玫红', '童装', 121, 66, 1234, 2111, 599, '2020-02-24 00:00:00.000', 'static/imgs/p2.jpg', 'static/details/p2_details.jpg');
INSERT INTO `commodities` VALUES (3, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '婴儿实木摇摇床', '原木色', '婴儿床', 1099, 999, 2346, 1322, 333, '2020-02-24 00:00:00.000', 'static/imgs/p3.jpg', 'static/details/p3_details.jpg');
INSERT INTO `commodities` VALUES (4, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '婴儿防寒针织帽', '蓝色', '配饰', 50, 39, 2347, 4521, 902, '2020-02-24 00:00:00.000', 'static/imgs/p4.jpg', 'static/details/p4_details.jpg');
INSERT INTO `commodities` VALUES (5, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '婴儿秋冬连体衣冬装衣服', '粉色', '童装', 145, 111, 1341, 3412, 2356, '2020-02-24 00:00:00.000', 'static/imgs/p5.jpg', 'static/details/p5_details.jpg');
INSERT INTO `commodities` VALUES (6, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '男女宝宝秋冬新年套装', '红色', '童装', 123, 119, 2342, 232, 1233, '2020-02-24 00:00:00.000', 'static/imgs/p6.jpg', 'static/details/p6_details.jpg');
INSERT INTO `commodities` VALUES (7, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '外出服抱衣加厚冬季', '粉色', '童装', 166, 111, 213, 2341, 1233, '2020-02-24 00:00:00.000', 'static/imgs/p7.jpg', 'static/details/p7_details.jpg');
INSERT INTO `commodities` VALUES (8, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '宝宝衣服动物造型熊猫哈衣', '黑白色', '童装', 124, 121, 531, 1345, 879, '2020-02-24 00:00:00.000', 'static/imgs/p8.jpg', 'static/details/p8_details.jpg');
INSERT INTO `commodities` VALUES (9, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '澳洲进口Aptamil爱他美', '1段', '进口奶粉', 399, 366, 1233, 1231, 666, '2020-02-24 00:00:00.000', 'static/imgs/p9.jpg', 'static/details/p9_details.jpg');
INSERT INTO `commodities` VALUES (10, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '超薄干爽绿帮纸尿裤', '大码', '纸尿裤', 209, 159, 1234, 4321, 3335, '2020-02-24 00:00:00.000', 'static/imgs/p10.jpg', 'static/details/p10_details.jpg');
INSERT INTO `commodities` VALUES (11, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '婴儿面条宝宝辅食无添加', '面条', '宝宝辅食', 30, 20, 3211, 1231, 2152, '2020-02-24 00:00:00.000', 'static/imgs/p11.jpg', 'static/details/p11_details.jpg');
INSERT INTO `commodities` VALUES (12, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '宝宝零食儿童营养磨牙饼干', '饼干', '宝宝辅食', 66, 38, 1543, 1845, 3245, '2020-02-24 00:00:00.000', 'static/imgs/p12.jpg', 'static/details/p12_details.jpg');
INSERT INTO `commodities` VALUES (13, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '宝宝芝麻粉辅食调味拌饭', '辅食调味', '宝宝辅食', 59, 39, 1234, 3453, 2321, '2020-02-24 00:00:00.000', 'static/imgs/p13.jpg', 'static/details/p13_details.jpg');
INSERT INTO `commodities` VALUES (14, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '纽曼思海藻油DHA软胶囊', 'DHA软胶囊', '营养品', 499, 399, 3231, 3412, 1234, '2020-02-24 00:00:00.000', 'static/imgs/p14.jpg', 'static/details/p14_details.jpg');
INSERT INTO `commodities` VALUES (15, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '婴儿推车可坐可躺', '蓝色', '婴儿车', 888, 439, 1234, 1245, 2353, '2020-02-24 00:00:00.000', 'static/imgs/p15.jpg', 'static/details/p15_details.jpg');
INSERT INTO `commodities` VALUES (16, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '儿童安全座椅增高垫', '粉色', '安全座椅', 688, 588, 3421, 3644, 6245, '2020-02-24 00:00:00.000', 'static/imgs/p16.jpg', 'static/details/p16_details.jpg');
INSERT INTO `commodities` VALUES (17, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '80片纯水湿纸巾婴儿', '湿纸巾', '婴儿湿巾', 49.9, 29.9, 1235, 5674, 2317, '2020-02-24 00:00:00.000', 'static/imgs/p17.jpg', 'static/details/p17_details.jpg');
INSERT INTO `commodities` VALUES (18, '2024-04-25 15:16:05.000', '2024-04-25 15:16:05.000', NULL, '儿童球类玩具室内特大号户外亲子', '10瓶2球', '儿童玩具', 69, 29.9, 666, 1986, 1569, '2020-02-24 00:00:00.000', 'static/imgs/p18.jpg', 'static/details/p18_details.jpg');

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

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `price` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pay_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `state` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_orders_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orders
-- ----------------------------

-- ----------------------------
-- Table structure for records
-- ----------------------------
DROP TABLE IF EXISTS `records`;
CREATE TABLE `records`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `commodity_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of records
-- ----------------------------

-- ----------------------------
-- Table structure for types
-- ----------------------------
DROP TABLE IF EXISTS `types`;
CREATE TABLE `types`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `firsts` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `seconds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_types_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of types
-- ----------------------------
INSERT INTO `types` VALUES (1, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '奶粉辅食', '进口奶粉');
INSERT INTO `types` VALUES (2, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '奶粉辅食', '宝宝辅食');
INSERT INTO `types` VALUES (3, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '奶粉辅食', '营养品');
INSERT INTO `types` VALUES (4, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童用品', '纸尿裤');
INSERT INTO `types` VALUES (5, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童用品', '婴儿湿巾');
INSERT INTO `types` VALUES (6, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童用品', '婴儿车');
INSERT INTO `types` VALUES (7, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童用品', '婴儿床');
INSERT INTO `types` VALUES (8, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童用品', '安全座椅');
INSERT INTO `types` VALUES (9, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童教育', '儿童玩具');
INSERT INTO `types` VALUES (10, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童教育', '早教书籍');
INSERT INTO `types` VALUES (11, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童教育', '育儿书籍');
INSERT INTO `types` VALUES (12, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童服饰', '童装');
INSERT INTO `types` VALUES (13, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童服饰', '童鞋');
INSERT INTO `types` VALUES (14, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '儿童服饰', '配饰');
INSERT INTO `types` VALUES (15, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '孕妈专区', '孕妇装');
INSERT INTO `types` VALUES (16, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '孕妈专区', '孕妇护肤');
INSERT INTO `types` VALUES (17, '2024-04-25 15:15:30.000', '2024-04-25 15:15:30.000', NULL, '孕妈专区', '孕妇用品');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_staff` bigint(0) NULL DEFAULT 0,
  `last_login` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2024-04-25 16:37:40.509', '2024-04-25 16:37:40.509', NULL, 'admin', '21232f297a57a5a743894a0e4a801fc3', 1, '2024-04-25 16:37:40.505');

SET FOREIGN_KEY_CHECKS = 1;

/*
 Navicat Premium Data Transfer

 Source Server         : laipeng
 Source Server Type    : MySQL
 Source Server Version : 80037 (8.0.37)
 Source Host           : localhost:3306
 Source Schema         : file_manager

 Target Server Type    : MySQL
 Target Server Version : 80037 (8.0.37)
 File Encoding         : 65001

 Date: 31/07/2025 11:45:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gc_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `gc_admin_menu`;
CREATE TABLE `gc_admin_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `meta` json NULL COMMENT '元数据',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '组件',
  `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '别名',
  `parent_id` int NULL DEFAULT NULL COMMENT '上级id',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '路径',
  `redirect` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '重定向uri',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_admin_menu_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1019 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_admin_menu
-- ----------------------------
INSERT INTO `gc_admin_menu` VALUES (1, '2022-10-19 18:27:57.555', '2022-10-19 20:53:45.507', NULL, '{\"icon\": \"&#xe721;\", \"title\": \"仪表盘\", \"keepAlive\": false}', 'RoutesAlias.Home', 'Dashboard', 0, 0, '/dashboard', '');
INSERT INTO `gc_admin_menu` VALUES (2, '2022-10-19 18:27:59.191', '2022-10-19 18:27:59.191', NULL, '{\"icon\": \"&#xe86e;\", \"title\": \"用户管理\", \"keepAlive\": false}', 'RoutesAlias.Home', 'User', 0, 0, '/user', '');
INSERT INTO `gc_admin_menu` VALUES (3, '2022-10-19 18:28:00.009', '2022-10-19 18:28:00.009', NULL, '{\"icon\": \"&#xe8a4;\", \"title\": \"菜单管理\", \"keepAlive\": false}', 'RoutesAlias.Home', 'Menu', 0, 0, '/menu', '');
INSERT INTO `gc_admin_menu` VALUES (4, '2022-10-19 18:28:00.966', '2022-10-19 19:08:13.304', NULL, '{\"icon\": \"&#xe7b9;\", \"title\": \"系统管理\", \"keepAlive\": false}', 'RoutesAlias.Home', 'System', 0, 0, '/system', '');
INSERT INTO `gc_admin_menu` VALUES (5, '2022-10-21 08:44:55.217', '2022-11-06 10:23:28.533', NULL, '{\"icon\": \"&#xe816;\", \"title\": \"文件管理\", \"keepAlive\": false}', 'RoutesAlias.Home', 'Safeguard', 0, 0, '/safeguard', '');
INSERT INTO `gc_admin_menu` VALUES (101, '2022-10-19 19:15:38.766', '2022-10-19 19:15:38.766', NULL, '{\"title\": \"电子商务\", \"keepAlive\": true}', 'RoutesAlias.Dashboard', 'Console', 1, 0, 'console', '');
INSERT INTO `gc_admin_menu` VALUES (102, '2022-10-19 19:15:43.514', '2022-10-19 20:53:55.312', NULL, '{\"title\": \"分析页\", \"keepAlive\": true}', 'RoutesAlias.Analysis', 'Analysis', 1, 0, 'analysis', '');
INSERT INTO `gc_admin_menu` VALUES (103, '2022-10-19 20:28:24.523', '2022-10-19 20:45:36.911', NULL, '{\"icon\": \"\", \"sort\": 0, \"title\": \"电子商务\", \"isMenu\": true, \"authList\": [{\"icon\": \"1\", \"sort\": 0, \"title\": \"新增\", \"auth_mark\": \"1\"}, {\"icon\": \"\", \"sort\": 0, \"title\": \"编辑\", \"auth_mark\": \"up\"}], \"isEnable\": true, \"isHidden\": false, \"isIframe\": false, \"keepAlive\": true}', 'RoutesAlias.Home', 'Ecommerce', 1, 0, 'ecommerce', '');
INSERT INTO `gc_admin_menu` VALUES (201, '2022-10-21 08:57:26.565', '2022-11-06 12:00:03.758', NULL, '{\"title\": \"账号管理\", \"keepAlive\": true}', 'RoutesAlias.Account', 'Account', 2, 10, 'account', '');
INSERT INTO `gc_admin_menu` VALUES (202, '2022-10-21 09:05:47.388', '2022-11-05 11:40:52.867', NULL, '{\"title\": \"部门管理\", \"keepAlive\": false}', 'RoutesAlias.Department', 'Department', 2, 0, 'department', '');
INSERT INTO `gc_admin_menu` VALUES (203, '2022-10-24 10:51:20.317', '2022-11-06 10:22:18.934', NULL, '{\"title\": \"角色权限\", \"keepAlive\": true}', 'RoutesAlias.Role', 'Role', 2, 1, 'role', '');
INSERT INTO `gc_admin_menu` VALUES (204, '2022-11-02 16:40:14.723', '2022-11-06 12:01:13.097', NULL, '{\"title\": \"个人中心\", \"isHide\": true, \"isHideTab\": true, \"keepAlive\": true}', 'RoutesAlias.UserCenter', 'UserCenter', 2, 3, 'user', '');
INSERT INTO `gc_admin_menu` VALUES (301, '2022-11-05 11:42:16.723', '2022-11-06 12:01:18.893', NULL, '{\"icon\": \"&#xe8a4;\", \"title\": \"菜单权限\", \"authList\": [{\"id\": 3011, \"title\": \"新增\", \"auth_mark\": \"add\"}, {\"id\": 3012, \"title\": \"编辑\", \"auth_mark\": \"edit\"}, {\"id\": 3013, \"title\": \"删除\", \"auth_mark\": \"delete\"}], \"keepAlive\": true}', 'RoutesAlias.Menu', 'Menus', 3, 3, 'menu', '');
INSERT INTO `gc_admin_menu` VALUES (302, '2022-11-05 11:44:48.945', '2022-11-06 12:01:24.837', NULL, '{\"icon\": \"&#xe831;\", \"title\": \"权限控制\", \"authList\": [{\"id\": 3021, \"title\": \"新增\", \"auth_mark\": \"add\"}, {\"id\": 3022, \"title\": \"编辑\", \"auth_mark\": \"edit\"}, {\"id\": 3023, \"title\": \"删除\", \"auth_mark\": \"delete\"}], \"keepAlive\": true, \"showTextBadge\": \"new\"}', 'RoutesAlias.Permission', 'Permission', 3, 2, 'permission', '');
INSERT INTO `gc_admin_menu` VALUES (303, '2022-11-05 11:45:05.993', '2022-11-06 12:01:31.869', NULL, '{\"icon\": \"&#xe676;\", \"title\": \"嵌套菜单\", \"keepAlive\": true}', '', 'Nested', 3, 1, 'nested', '');
INSERT INTO `gc_admin_menu` VALUES (401, '2022-11-06 10:18:06.707', '2022-11-06 10:21:56.690', NULL, '{\"title\": \"系统设置\", \"keepAlive\": true}', 'RoutesAlias.Setting', 'Setting', 4, 1, 'setting', '');
INSERT INTO `gc_admin_menu` VALUES (402, '2022-11-06 10:18:16.309', '2022-11-06 10:21:49.434', NULL, '{\"title\": \"api管理\", \"keepAlive\": true}', 'RoutesAlias.Api', 'Api', 4, 1, 'api', '');
INSERT INTO `gc_admin_menu` VALUES (403, '2022-11-06 10:18:23.074', '2022-11-06 10:21:07.173', NULL, '{\"title\": \"系统日志\", \"keepAlive\": true}', 'RoutesAlias.Log', 'Log', 4, 1, 'log', '');
INSERT INTO `gc_admin_menu` VALUES (501, '2022-11-06 10:23:01.143', '2022-11-06 10:24:16.708', NULL, '{\"title\": \"文件组管理\", \"keepAlive\": true}', 'RoutesAlias.Server', 'Server', 5, 1, 'server', '');
INSERT INTO `gc_admin_menu` VALUES (502, '2025-04-14 10:32:49.000', '2025-04-14 10:32:51.000', NULL, '{\"title\": \"文件上传\", \"keepAlive\": true}', 'RoutesAlias.File', 'File', 5, 2, 'file', NULL);

-- ----------------------------
-- Table structure for gc_admin_user
-- ----------------------------
DROP TABLE IF EXISTS `gc_admin_user`;
CREATE TABLE `gc_admin_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '真实名称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `gender` tinyint NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_gc_admin_user_name`(`name` ASC) USING BTREE,
  INDEX `idx_gc_admin_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 94 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_admin_user
-- ----------------------------
INSERT INTO `gc_admin_user` VALUES (54, '2022-10-10 15:01:35.024', '2025-04-08 13:17:51.072', NULL, '', '凡人修仙', '$2a$04$2hDcVIpxphRSdqoK4hlibeL6nBcpNQQf43sc2d2iOubQsZUIb2JGS', '2195143506@qq.com', 'test', 'https://www.lpmyblog.cn/applet/static/init1.png', '18582921019', 1);
INSERT INTO `gc_admin_user` VALUES (57, '2022-10-10 19:52:34.901', '2024-12-31 13:14:19.813', NULL, '', '斗罗大陆', '$2a$04$m0BD57ApLej9WQ7g20fIAushtu5BAuW/ZyhyyBUTfhlJsAiyO1Zka', '2195143506@qq.com', 'admin', 'https://www.lpmyblog.cn/applet/static/init1.png', '18582921019', 2);
INSERT INTO `gc_admin_user` VALUES (92, '2025-04-08 19:04:52.202', '2025-04-08 19:05:31.837', NULL, '', 'tourist', '$2a$04$YilBpgAebm1SlrmDW6AthOFssMCgDel2w3lPGwkZgJ4aFGB/FBqjW', '2195143506@qq.com', 'tourist', 'https://www.lpmyblog.cn/applet/static/init1.png', '18582921019', 1);
INSERT INTO `gc_admin_user` VALUES (93, '2025-04-21 10:29:22.993', '2025-04-21 10:29:22.993', NULL, '', '赖鹏', '$2a$04$KuRkbPtsUcHwfeGffJBzT.G2ABAcHbLcuqQfaxJ7vES2.KquxJGEu', '2195143506@qq.com', 'laipeng', 'https://www.lpmyblog.cn/applet/static/init1.png', '18582921019', 1);

-- ----------------------------
-- Table structure for gc_department
-- ----------------------------
DROP TABLE IF EXISTS `gc_department`;
CREATE TABLE `gc_department`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT NULL,
  `sort` bigint NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_department_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_department
-- ----------------------------
INSERT INTO `gc_department` VALUES (1, '人力资源部', 1, 1, '2025-03-31 16:58:38', '2025-03-31 16:58:40', NULL);
INSERT INTO `gc_department` VALUES (2, '财务部', 1, 2, '2025-03-31 16:59:01', '2025-03-31 16:59:03', NULL);
INSERT INTO `gc_department` VALUES (3, '技术部', 1, 8, '2025-03-31 19:27:26', '2025-03-31 20:10:56', NULL);

-- ----------------------------
-- Table structure for gc_file_group
-- ----------------------------
DROP TABLE IF EXISTS `gc_file_group`;
CREATE TABLE `gc_file_group`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件组名',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '文件组上级id',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_file_group_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_file_group
-- ----------------------------

-- ----------------------------
-- Table structure for gc_files
-- ----------------------------
DROP TABLE IF EXISTS `gc_files`;
CREATE TABLE `gc_files`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件名',
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件路径',
  `file_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件网络链接',
  `type` bigint NULL DEFAULT NULL COMMENT '1图片，2视频，3html',
  `uploader` bigint NULL DEFAULT NULL COMMENT '上传者',
  `group_id` bigint NULL DEFAULT NULL COMMENT '分组_id',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_files_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1998 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_files
-- ----------------------------

-- ----------------------------
-- Table structure for gc_menu_api_list
-- ----------------------------
DROP TABLE IF EXISTS `gc_menu_api_list`;
CREATE TABLE `gc_menu_api_list`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '关键字',
  `url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地址',
  `menu_id` bigint NULL DEFAULT NULL,
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_menu_api_list_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_menu_api_list
-- ----------------------------
INSERT INTO `gc_menu_api_list` VALUES (1, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'user.all', '/system/user/all', 201, '获取用户列表');
INSERT INTO `gc_menu_api_list` VALUES (2, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'user.list', '/system/user/list', 201, '获取用户分页列表');
INSERT INTO `gc_menu_api_list` VALUES (3, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'user.add', '/system/user/add', 201, '添加用户');
INSERT INTO `gc_menu_api_list` VALUES (4, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'user.update', '/system/user/update', 201, '更新用户信息');
INSERT INTO `gc_menu_api_list` VALUES (5, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'user.dels', '/system/user/del', 201, '批量删除用户');
INSERT INTO `gc_menu_api_list` VALUES (6, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'user.del', '/system/user/:id', 201, '删除单个用户');
INSERT INTO `gc_menu_api_list` VALUES (7, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'role.group', '/system/role/group', 203, '获取用户组');
INSERT INTO `gc_menu_api_list` VALUES (8, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'role.list', '/system/role/list', 203, '获取角色列表');
INSERT INTO `gc_menu_api_list` VALUES (9, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'role.add', '/system/role/add', 203, '添加角色');
INSERT INTO `gc_menu_api_list` VALUES (10, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'role.update', '/system/role/update', 203, '更新角色信息');
INSERT INTO `gc_menu_api_list` VALUES (11, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'role.del', '/system/role/del', 203, '删除角色');
INSERT INTO `gc_menu_api_list` VALUES (12, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'role.upMenu', '/system/role/upMenu', 203, '更新角色菜单权限');
INSERT INTO `gc_menu_api_list` VALUES (13, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'menu.add', '/system/menu/add', 301, '添加菜单');
INSERT INTO `gc_menu_api_list` VALUES (14, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'menu.update', '/system/menu/update', 301, '修改菜单');
INSERT INTO `gc_menu_api_list` VALUES (15, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'menu.list', '/system/menu/list', 301, '获取菜单列表');
INSERT INTO `gc_menu_api_list` VALUES (16, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'menu.dels', '/system/menu/del', 301, '批量删除菜单');
INSERT INTO `gc_menu_api_list` VALUES (17, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'menu.del', '/system/menu/:id', 301, '删除单个菜单');
INSERT INTO `gc_menu_api_list` VALUES (18, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.list', '/system/depart/list', 202, '获取部门列表');
INSERT INTO `gc_menu_api_list` VALUES (19, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.add', '/system/depart/add', 202, '添加部门');
INSERT INTO `gc_menu_api_list` VALUES (20, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.update', '/system/depart/update', 202, '修改部门信息');
INSERT INTO `gc_menu_api_list` VALUES (21, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.dels', '/system/depart/del', 202, '批量删除部门');
INSERT INTO `gc_menu_api_list` VALUES (22, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.del', '/system/depart/:id', 202, '删除单个部门');
INSERT INTO `gc_menu_api_list` VALUES (23, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.upAdmin', '/system/depart/upAdmin', 202, '更新部门管理员');
INSERT INTO `gc_menu_api_list` VALUES (24, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'depart.addAdmin', '/system/depart/addAdmin', 202, '添加部门管理员');
INSERT INTO `gc_menu_api_list` VALUES (25, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'api.add', '/system/api/add', 402, '添加API');
INSERT INTO `gc_menu_api_list` VALUES (26, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'api.update', '/system/api/update', 402, '修改API');
INSERT INTO `gc_menu_api_list` VALUES (27, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'api.list', '/system/api/list', 402, '获取API列表');
INSERT INTO `gc_menu_api_list` VALUES (28, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'api.dels', '/system/api/del', 402, '批量删除API');
INSERT INTO `gc_menu_api_list` VALUES (29, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'api.del', '/system/api/:id', 402, '删除单个API');
INSERT INTO `gc_menu_api_list` VALUES (30, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'system.upload', '/system/system/common/upload', 401, '文件上传');
INSERT INTO `gc_menu_api_list` VALUES (31, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'system.operationLog.list', '/system/system/operationLog/list', 403, '获取操作日志列表');
INSERT INTO `gc_menu_api_list` VALUES (32, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'system.menu.my', '/system/system/menu/my', 3, '获取我的菜单权限');
INSERT INTO `gc_menu_api_list` VALUES (33, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.start', '/system/timer/start', 4, '启动定时任务管理器');
INSERT INTO `gc_menu_api_list` VALUES (34, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.stop', '/system/timer/stop', 4, '停止定时任务管理器');
INSERT INTO `gc_menu_api_list` VALUES (35, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.status', '/system/timer/status', 4, '获取定时任务状态');
INSERT INTO `gc_menu_api_list` VALUES (36, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.list', '/system/timer/task/list', 4, '获取任务列表');
INSERT INTO `gc_menu_api_list` VALUES (37, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.create', '/system/timer/task/create', 4, '创建定时任务');
INSERT INTO `gc_menu_api_list` VALUES (38, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.update', '/system/timer/task/update', 4, '更新定时任务');
INSERT INTO `gc_menu_api_list` VALUES (39, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.delete', '/system/timer/task/delete/:id', 4, '删除定时任务');
INSERT INTO `gc_menu_api_list` VALUES (40, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.get', '/system/timer/task/get/:id', 4, '获取单个任务详情');
INSERT INTO `gc_menu_api_list` VALUES (41, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.execute', '/system/timer/task/execute', 4, '执行定时任务');
INSERT INTO `gc_menu_api_list` VALUES (42, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.test', '/system/timer/task/test', 4, '测试定时任务');
INSERT INTO `gc_menu_api_list` VALUES (43, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.toggle', '/system/timer/task/toggle/:id', 4, '切换任务状态');
INSERT INTO `gc_menu_api_list` VALUES (44, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.task.logs', '/system/timer/task/logs', 4, '获取任务日志');
INSERT INTO `gc_menu_api_list` VALUES (45, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'timer.cron.examples', '/system/timer/cron/examples', 4, '获取Cron表达式示例');
INSERT INTO `gc_menu_api_list` VALUES (46, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'fileGroup.index', '/system/fileGroup/:id', 502, '获取文件组详情');
INSERT INTO `gc_menu_api_list` VALUES (47, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'fileGroup.list', '/system/fileGroup/list', 502, '获取文件组列表');
INSERT INTO `gc_menu_api_list` VALUES (48, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'fileGroup.add', '/system/fileGroup/add', 502, '添加文件组');
INSERT INTO `gc_menu_api_list` VALUES (49, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'fileGroup.update', '/system/fileGroup/update', 502, '修改文件组');
INSERT INTO `gc_menu_api_list` VALUES (50, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'fileGroup.delete', '/system/fileGroup/:id', 502, '删除文件组');
INSERT INTO `gc_menu_api_list` VALUES (51, '2025-07-30 14:27:30.000', '2025-07-30 14:27:30.000', NULL, 'fileGroup.check', '/system/fileGroup/check/:id', 502, '检查文件组');

-- ----------------------------
-- Table structure for gc_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `gc_operation_log`;
CREATE TABLE `gc_operation_log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` int NULL DEFAULT NULL COMMENT '用户ID',
  `user_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '访问路径',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'IP',
  `method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方式',
  `path_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求名称',
  `do_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '处理数据',
  `user_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '未知' COMMENT '用户名',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_operation_log_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 946 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_operation_log
-- ----------------------------

-- ----------------------------
-- Table structure for gc_role
-- ----------------------------
DROP TABLE IF EXISTS `gc_role`;
CREATE TABLE `gc_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `alias` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `status` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_role_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_role
-- ----------------------------
INSERT INTO `gc_role` VALUES (1, '2022-11-02 16:46:16.253', '2022-11-06 10:17:36.789', NULL, 'administrator', '超级管理员', '超级管理员', 2, 1);
INSERT INTO `gc_role` VALUES (2, '2022-11-02 16:49:23.268', '2025-04-21 10:27:26.407', NULL, 'testRole', '测试角色名称', '', 10, 1);
INSERT INTO `gc_role` VALUES (3, '2025-01-02 10:50:26.766', '2025-04-09 20:19:45.356', NULL, 'tourist', '游客', '', 4, 1);
INSERT INTO `gc_role` VALUES (19, '2025-04-21 10:28:03.653', '2025-04-21 10:28:10.663', NULL, 'filer', '文件用户', '文件上传操作员', 10, 1);

-- ----------------------------
-- Table structure for gc_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `gc_role_menu`;
CREATE TABLE `gc_role_menu`  (
  `role_id` bigint UNSIGNED NOT NULL,
  `admin_menu_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`, `admin_menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_role_menu
-- ----------------------------
INSERT INTO `gc_role_menu` VALUES (1, 1);
INSERT INTO `gc_role_menu` VALUES (1, 2);
INSERT INTO `gc_role_menu` VALUES (1, 3);
INSERT INTO `gc_role_menu` VALUES (1, 4);
INSERT INTO `gc_role_menu` VALUES (1, 5);
INSERT INTO `gc_role_menu` VALUES (1, 101);
INSERT INTO `gc_role_menu` VALUES (1, 102);
INSERT INTO `gc_role_menu` VALUES (1, 103);
INSERT INTO `gc_role_menu` VALUES (1, 201);
INSERT INTO `gc_role_menu` VALUES (1, 202);
INSERT INTO `gc_role_menu` VALUES (1, 204);
INSERT INTO `gc_role_menu` VALUES (1, 301);
INSERT INTO `gc_role_menu` VALUES (1, 302);
INSERT INTO `gc_role_menu` VALUES (1, 303);
INSERT INTO `gc_role_menu` VALUES (1, 401);
INSERT INTO `gc_role_menu` VALUES (1, 402);
INSERT INTO `gc_role_menu` VALUES (1, 403);
INSERT INTO `gc_role_menu` VALUES (1, 501);
INSERT INTO `gc_role_menu` VALUES (2, 5);
INSERT INTO `gc_role_menu` VALUES (2, 501);
INSERT INTO `gc_role_menu` VALUES (2, 502);
INSERT INTO `gc_role_menu` VALUES (3, 1);
INSERT INTO `gc_role_menu` VALUES (3, 4);
INSERT INTO `gc_role_menu` VALUES (3, 5);
INSERT INTO `gc_role_menu` VALUES (3, 101);
INSERT INTO `gc_role_menu` VALUES (3, 102);
INSERT INTO `gc_role_menu` VALUES (3, 103);
INSERT INTO `gc_role_menu` VALUES (3, 401);
INSERT INTO `gc_role_menu` VALUES (3, 402);
INSERT INTO `gc_role_menu` VALUES (3, 403);
INSERT INTO `gc_role_menu` VALUES (3, 501);
INSERT INTO `gc_role_menu` VALUES (19, 5);
INSERT INTO `gc_role_menu` VALUES (19, 501);
INSERT INTO `gc_role_menu` VALUES (19, 502);

-- ----------------------------
-- Table structure for gc_timer_task_logs
-- ----------------------------
DROP TABLE IF EXISTS `gc_timer_task_logs`;
CREATE TABLE `gc_timer_task_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `task_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '任务ID',
  `status` tinyint NULL DEFAULT NULL COMMENT '执行状态 1:成功 0:失败',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '执行结果',
  `duration` bigint NULL DEFAULT NULL COMMENT '执行时长(毫秒)',
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应内容',
  `status_code` bigint NULL DEFAULT NULL COMMENT 'HTTP状态码',
  `run_time` datetime(3) NULL DEFAULT NULL COMMENT '执行时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_timer_task_logs_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_timer_task_logs
-- ----------------------------

-- ----------------------------
-- Table structure for gc_timer_tasks
-- ----------------------------
DROP TABLE IF EXISTS `gc_timer_tasks`;
CREATE TABLE `gc_timer_tasks`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务名称',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务描述',
  `cron_expression` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'cron表达式',
  `target_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目标接口URL',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方法',
  `headers` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求头',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求体',
  `status` tinyint NULL DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
  `last_run_time` datetime(3) NULL DEFAULT NULL COMMENT '最后执行时间',
  `next_run_time` datetime(3) NULL DEFAULT NULL COMMENT '下次执行时间',
  `run_count` bigint NULL DEFAULT 0 COMMENT '执行次数',
  `success_count` bigint NULL DEFAULT 0 COMMENT '成功次数',
  `fail_count` bigint NULL DEFAULT 0 COMMENT '失败次数',
  `timeout` bigint NULL DEFAULT 30 COMMENT '超时时间(秒)',
  `retry_count` bigint NULL DEFAULT 0 COMMENT '重试次数',
  `retry_interval` bigint NULL DEFAULT 60 COMMENT '重试间隔(秒)',
  `creator` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gc_timer_tasks_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_timer_tasks
-- ----------------------------

-- ----------------------------
-- Table structure for gc_user_department
-- ----------------------------
DROP TABLE IF EXISTS `gc_user_department`;
CREATE TABLE `gc_user_department`  (
  `admin_user_id` bigint NOT NULL,
  `department_id` bigint NOT NULL,
  PRIMARY KEY (`admin_user_id`, `department_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_user_department
-- ----------------------------
INSERT INTO `gc_user_department` VALUES (54, 1);
INSERT INTO `gc_user_department` VALUES (54, 3);
INSERT INTO `gc_user_department` VALUES (57, 2);

-- ----------------------------
-- Table structure for gc_user_role
-- ----------------------------
DROP TABLE IF EXISTS `gc_user_role`;
CREATE TABLE `gc_user_role`  (
  `admin_user_id` bigint UNSIGNED NOT NULL,
  `role_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`admin_user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of gc_user_role
-- ----------------------------
INSERT INTO `gc_user_role` VALUES (54, 2);
INSERT INTO `gc_user_role` VALUES (57, 1);
INSERT INTO `gc_user_role` VALUES (92, 3);
INSERT INTO `gc_user_role` VALUES (93, 19);

SET FOREIGN_KEY_CHECKS = 1;

CREATE DATABASE  `sports` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
grant all PRIVILEGES on sports.* to sports@'%' identified by '123456';
flush privileges;
use sports;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '密码',
  `phone` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '手机号',
  `email` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮箱',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
INSERT INTO `sports`.`users` (`id`, `user_name`, `password`, `phone`, `email`, `create_time`)
  VALUES ('1', 'test', '123456', '12345678910', '123@qq.com', '2021-11-11 01:22:07');

CREATE DATABASE `micro_book_mall` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;

CREATE TABLE `inventory`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `book_id`      int(10) unsigned NOT NULL COMMENT '书id',
    `unit_price`   int(10) unsigned NOT NULL COMMENT '单价',
    `stock`        int(10) unsigned NOT NULL COMMENT '总数',
    `version`      int(10) unsigned NOT NULL COMMENT '版本号',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='库存表';

INSERT INTO micro_book_mall.inventory (id, book_id, unit_price, stock, version)
VALUES (1, 1, 20, 9, 1);

CREATE TABLE `inventory_history`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `book_id`      int(10) unsigned NOT NULL COMMENT '书id',
    `user_id`      int(10) unsigned NOT NULL COMMENT '单价',
    `state`        int(10) unsigned NOT NULL COMMENT '未出库1，出库2',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='订单表历史';

CREATE TABLE `orders`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned          DEFAULT NULL COMMENT '用户id',
    `book_id`      int(10)          NOT NULL COMMENT '书id',
    `inv_his_id`   int(10)          NOT NULL COMMENT '库存历史记录id',
    `state`        tinyint(1)       NOT NULL COMMENT '1未支付，2支付',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='订单表';

CREATE TABLE `payment`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned          DEFAULT NULL COMMENT '用户id',
    `book_id`      int(10)          NOT NULL COMMENT '书id',
    `order_id`     int(10)          NOT NULL COMMENT '订单id',
    `inv_his_id`   int(10)          NOT NULL COMMENT '库存历史id',
    `state`        tinyint(1)       NOT NULL COMMENT '1：未支付，2：支付',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='支付表';

CREATE TABLE `user`
(
    `id`           int(10) unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned                                                       DEFAULT NULL COMMENT '用户id',
    `user_name`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户名',
    `pwd`          varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
    `created_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_user_name_uindex` (`user_name`),
    UNIQUE KEY `user_user_id_uindex` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户表';
CREATE TABLE `inventory`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `book_id`      int(10) unsigned NOT NULL COMMENT '书id',
    `unit_price`   int(10) unsigned NOT NULL COMMENT '单价',
    `stock`        int(10) unsigned NOT NULL COMMENT '总数',
    `version`      int(10) unsigned NOT NULL COMMENT '版本号',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='库存表';

CREATE TABLE `inventory_history`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `book_id`      int(10) unsigned NOT NULL COMMENT '书id',
    `user_id`      int(10) unsigned NOT NULL COMMENT '单价',
    `state`        int(10) unsigned NOT NULL COMMENT '未出库1，出库2',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='订单表历史';

CREATE TABLE `orders`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned          DEFAULT NULL COMMENT '用户id',
    `book_id`      int(10)          NOT NULL COMMENT '书id',
    `inv_his_id`   int(10)          NOT NULL COMMENT '库存历史记录id',
    `state`        tinyint(1)       NOT NULL COMMENT '1未支付，2支付',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='订单表';

CREATE TABLE `payment`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned          DEFAULT NULL COMMENT '用户id',
    `book_id`      int(10)          NOT NULL COMMENT '书id',
    `order_id`     int(10)          NOT NULL COMMENT '订单id',
    `inv_his_id`   int(10)          NOT NULL COMMENT '库存历史id',
    `state`        tinyint(1)       NOT NULL COMMENT '1：未支付，2：支付',
    `created_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='支付表';

CREATE TABLE `user`
(
    `id`           int(10) unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned                                                       DEFAULT NULL COMMENT '用户id',
    `user_name`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户名',
    `pwd`          varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
    `created_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_user_name_uindex` (`user_name`),
    UNIQUE KEY `user_user_id_uindex` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户表';

INSERT INTO micro_book_mall.user (id, user_id, user_name, pwd)
VALUES (1, 10001, 'micro', '1234');
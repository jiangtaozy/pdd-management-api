/*
 * Maintained by jemo from 2020.9.11 to now
 * Created by jemo on 2020.9.11 15:20:14
 * Dy Access Token
 */

package database

const createDyAccessToken = `
  CREATE TABLE IF NOT EXISTS dyAccessToken (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    accessToken VARCHAR(36) NOT NULL COMMENT 'access_token',
    expiresIn INTEGER UNSIGNED NOT NULL COMMENT '超时时间，单位（秒），默认有效期：7天',
    refreshToken VARCHAR(36) NOT NULL COMMENT '刷新access_token的刷新令牌（有效期：14 天）',
    scope VARCHAR(64) NOT NULL COMMENT '授权作用域，使用逗号,分隔。预留字段',
    shopId VARCHAR(7) NOT NULL COMMENT '店铺ID',
    shopName VARCHAR(16) NOT NULL COMMENT '店铺名称',
    createdAt TIMESTAMP NOT NULL COMMENT '创建时间'
  );
`

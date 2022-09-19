
-- +migrate Up
CREATE TABLE IF NOT EXISTS favorite_blogs (
  id bigint(20) AUTO_INCREMENT,
  user_id bigint(20) NOT NULL,
  blog_id bigint(20) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS favorite_blogs;
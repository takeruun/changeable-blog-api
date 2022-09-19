
-- +migrate Up
CREATE TABLE IF NOT EXISTS related_blogs (
  id bigint(20) AUTO_INCREMENT,
  blog_id bigint(20) NOT NULL,
  related_blog_id bigint(20) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS related_blogs;
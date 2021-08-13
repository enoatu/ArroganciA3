CREATE TABLE IF NOT EXISTS tweet (
  id                INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  tweet_id          BIGINT UNSIGNED NOT NULL                                                 COMMENT 'tweetのID',
  search_word_id    INT UNSIGNED NOT NULL                                                    COMMENT '検索ワードID',
  body              TEXT NOT NULL                                                            COMMENT 'ツイート本文',
  user_id           BIGINT UNSIGNED NOT NULL                                                 COMMENT 'ツイートユーザーID',
  user_name         VARCHAR(255) NOT NULL                                                    COMMENT 'ユーザーネーム',
  user_screen_name  VARCHAR(255) NOT NULL                                                    COMMENT 'ユーザースクリーンネーム',
  optimize_level    TINYINT NOT NULL                                                         COMMENT '最適化レベル',
  created_at        DATETIME NOT NULL                                                        COMMENT 'ツイート作成日時',
  created_on        DATETIME NOT NULL                                                        COMMENT '作成日時',
  modified_on       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY tweet_index01 (search_word_id, optimize_level),
  KEY tweet_index02 (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='ツイート';

CREATE TABLE IF NOT EXISTS search_word (
  id               INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  word             VARCHAR(255) NOT NULL                                                    COMMENT '検索ワード',
  created_on       DATETIME NOT NULL                                                        COMMENT '作成日時',
  modified_on      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='検索ワード';

CREATE TABLE IF NOT EXISTS deny_twitter_user (
  id               INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  user_id          BIGINT UNSIGNED NOT NULL                                                 COMMENT '拒否tweetユーザーID',
  created_on       DATETIME NOT NULL                                                        COMMENT '作成日時',
  modified_on      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  UNIQUE KEY deny_twitter_user_unique_index01 (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='拒否Twitterユーザー';

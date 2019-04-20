CREATE TABLE `mk_posts`
(
  `id`      int(10) unsigned NOT NULL,
  `title`   varchar(255) DEFAULT '',
  `content` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

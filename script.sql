-- lp.favorite definition

CREATE TABLE `favorite` (
  `user_id` bigint(20) unsigned NOT NULL,
  `music_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`music_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- lp.music definition

CREATE TABLE `music` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `singer` varchar(100) NOT NULL,
  `duration` varchar(100) DEFAULT NULL,
  `album` varchar(100) DEFAULT NULL,
  `release_year` char(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;


-- lp.`user` definition

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `hobby` varchar(64) DEFAULT NULL,
  `gender` varchar(64) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_UN` (`username`),
  KEY `users_username_IDX` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=latin1;


INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(1, 'All i need', 'Jacob Collier', '00:03:00', 'Djesse Vol. 3', '2020');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(2, 'Isn''t she', 'Stevie Wonder', '00:03:00', 'Songs in the Key of Life', '1976');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(3, 'Stay With Me', 'Miki Matsubara', '00:03:00', 'Pocket Park', '1980');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(4, 'Love theory', 'Kirk Franklin', '00:03:00', 'Long, Live, Love', '2019');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(5, 'Smile', 'Kirk Franklin', '00:03:00', 'Hello Fear', '2011');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(6, 'Mungkin Nanti', 'Ariel', '00:03:00', 'Bintang di Surga', '2004');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(7, 'OverJoyed', 'Stevie Wonder', '00:03:00', 'In Square Circle', '1985');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(8, 'I just called to say i love you', 'Stevie Wonder', '00:03:00', 'The Woman in Red', '1984');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(9, 'Januari', 'Glenn Fredly', '00:03:00', 'Selamat Pagi, Dunia!', '2002');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(10, 'Akhir cerita cinta', 'Glenn Fredly', '00:03:00', 'Selamat Pagi, Dunia!', '2002');
INSERT INTO lp.music (id, title, singer, duration, album, release_year) VALUES(11, 'Terserah', 'Glenn Fredly', '00:03:00', 'Private Collection', '2008');

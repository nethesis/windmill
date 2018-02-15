CREATE TABLE IF NOT EXISTS `logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `session_id` varchar(64) NOT NULL,
  `operator_id` varchar(64) NOT NULL,
  `session_created` varchar(64) NOT NULL,
  `session_connected` varchar(64) NOT NULL,
  `session_disconnected` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
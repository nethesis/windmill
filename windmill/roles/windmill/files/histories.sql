CREATE TABLE IF NOT EXISTS `histories` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `session_id` varchar(64) NOT NULL,
  `server_id` varchar(64) NOT NULL,
  `started` varchar(64) NOT NULL,
  `bytes_received` varchar(64) DEFAULT NULL,
  `bytes_sent` varchar(64) DEFAULT NULL,
  `ifconfig_pool_local_ip` varchar(64) DEFAULT NULL,
  `ifconfig_pool_netmask` varchar(64) DEFAULT NULL,
  `ifconfig_pool_remote_ip` varchar(64) DEFAULT NULL,
  `time_duration` varchar(64) DEFAULT NULL,
  `trusted_ip` varchar(64) DEFAULT NULL,
  `trusted_port` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `session_id` (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
CREATE SCHEMA `go-elevators` ;

CREATE TABLE `go-elevators`.`elevators` (
  `elevator_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `max_size` int(11) unsigned NOT NULL DEFAULT '1',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`elevator_id`),
  UNIQUE KEY `uniq_elevators_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `go-elevators`.`requests` (
  `request_id` int(11) unsigned NOT NULL AUTO_INCREMENT,  
  `person_name` varchar(32) NOT NULL,
  `initial_floor` int(11) unsigned NOT NULL,
  `destination_floor` int(11) unsigned NOT NULL,
  `current_floor` int(11) unsigned NOT NULL,
  `elevator_id` int(11) unsigned,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`request_id`),  
  UNIQUE KEY `uniq_requests_person_name` (`person_name`),
  UNIQUE KEY `uniq_requests_elevator_person_floor` (`elevator_id`, `person_name`, `initial_floor`,  `destination_floor`),
  CONSTRAINT `fk_requests_elevator_id` FOREIGN KEY (`elevator_id`) REFERENCES `elevators` (`elevator_id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `go-elevators`.`operations` (
  `operation_id` int(11) unsigned NOT NULL AUTO_INCREMENT,  
  `elevator_id` int(11) unsigned NOT NULL,
  `is_going_up`  tinyint(4) NOT NULL DEFAULT '1',
  `current_floor` int(11) unsigned NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`operation_id`),  
  UNIQUE KEY `uniq_operations_elevator_floor` (`elevator_id`,`current_floor`),
  CONSTRAINT `fk_operations_elevator_id` FOREIGN KEY (`elevator_id`) REFERENCES `elevators` (`elevator_id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `go-elevators`.`configurations` (
  `conf_key` varchar(32) NOT NULL,
  `conf_value` varchar(32) NOT NULL,
  `description` varchar(50) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`conf_key`),
  UNIQUE KEY `uniq_configurations_key` (`conf_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
  
INSERT INTO `go-elevators`.`elevators` (`name`,  `max_size`,  `status`)
VALUES
  ('Elevator1', 5, 1),
  ('ElevatorOFF', 8, 0);

INSERT INTO `go-elevators`.`configurations` ( `conf_key`,  `conf_value`,  `description`)
VALUES
  ('floor_count', 5, 'Quantity of floors available'),
  ('elevator_delay', 2, 'Time between elevators updates. In seconds'),
  ('request_delay', 15, 'Time between requests calls. In seconds');
  
  
  select * from `go-elevators`.`configurations`;

CREATE TABLE `tbCommunity` (
	`community_id` INT NOT NULL AUTO_INCREMENT,
	`created_date` DATETIME NOT NULL,
	`created_by` INT NOT NULL,
	`community_name` VARCHAR(80) NOT NULL,
	`community_description` TEXT NOT NULL,
	`community_avatar` VARCHAR(100) NOT NULL,
	PRIMARY KEY (`community_id`)
);
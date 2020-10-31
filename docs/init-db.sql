SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema coucoo
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `coucoo` DEFAULT CHARACTER SET utf8 ;
USE `coucoo` ;

-- -----------------------------------------------------
-- Table `coucoo`.`coupons`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `coucoo`.`coupons` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `discount` VARCHAR(45) NOT NULL,
  `discount_type` VARCHAR(45) NOT NULL,
  `valid_from` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
  `valid_until` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `coucoo`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `coucoo`.`users` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(45) NOT NULL,
  `password` CHAR(64) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `coucoo`.`users_has_coupons`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `coucoo`.`users_has_coupons` (
  `users_id` INT(11) NOT NULL,
  `coupons_id` INT(11) NOT NULL,
  PRIMARY KEY (`users_id`, `coupons_id`),
  INDEX `fk_users_has_coupons_coupons1_idx` (`coupons_id` ASC) VISIBLE,
  INDEX `fk_users_has_coupons_users_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_users_has_coupons_users`
    FOREIGN KEY (`users_id`)
    REFERENCES `coucoo`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_coupons_coupons1`
    FOREIGN KEY (`coupons_id`)
    REFERENCES `coucoo`.`coupons` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


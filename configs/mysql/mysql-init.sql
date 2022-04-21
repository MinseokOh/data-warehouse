DROP TABLE IF EXISTS `SDK`;

CREATE TABLE `SDK` (
                           `ID` VARCHAR(20) NOT NULL,
                           `NAME` VARCHAR(20) DEFAULT NULL,
                           `VERSION` VARCHAR(20) DEFAULT NULL,
                           PRIMARY KEY (`ID`)
) ENGINE=INNODB DEFAULT CHARSET=utf8
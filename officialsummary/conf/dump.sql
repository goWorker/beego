-- MySQL dump 10.13  Distrib 8.0.15, for osx10.14 (x86_64)
--
-- Host: 127.0.0.1    Database: officialsummary
-- ------------------------------------------------------
-- Server version	8.0.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `caseLog`
--

DROP TABLE IF EXISTS `caseLog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `caseLog` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `relaseVersion` varchar(128) NOT NULL COMMENT 'release version',
  `jobName` varchar(128) NOT NULL COMMENT 'jobName',
  `caseName` varchar(128) NOT NULL COMMENT 'caseName',
  `tag` varchar(128) NOT NULL DEFAULT '' COMMENT 'caseName',
  `build` varchar(255) NOT NULL DEFAULT '',
  `executeTime` int(11) NOT NULL DEFAULT '0' COMMENT 'case运行时间',
  `executeDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '0 or 1',
  `log` varchar(255) DEFAULT NULL,
  `jira` varchar(255) DEFAULT NULL,
  `comment` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_caselog` (`relaseVersion`,`jobName`,`caseName`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `caseLog`
--

LOCK TABLES `caseLog` WRITE;
/*!40000 ALTER TABLE `caseLog` DISABLE KEYS */;
INSERT INTO `caseLog` VALUES (1,'6.1.0','spring','spring111','test','NorthStar610',33,'2020-06-18 16:30:07',1,NULL,NULL,NULL),(2,'6.1.0','basic','basic222','test11','NorthStar610',40,'2020-06-18 16:30:07',0,NULL,NULL,NULL),(3,'6.1.0','basic','basic111','test','NorthStar610',10,'2020-06-18 18:26:30',1,NULL,NULL,NULL);
/*!40000 ALTER TABLE `caseLog` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-18 18:35:10

select a.id, a.relaseVersion, a.jobName, a.caseName, b.status from job2case as a left join caseLog as b on (a.jobName=b.jobName and a.caseName=b.caseName and a.relaseVersion=b.relaseVersion and b.executeDate> '2020-06-17 00:00:00' and b.executeDate<'2020-06-19 00:00:00');
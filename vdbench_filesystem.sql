/*
Navicat MariaDB Data Transfer

Source Server         : prometheus-grafana
Source Server Version : 50560
Source Host           : 10.121.9.23:3306
Source Database       : ceph

Target Server Type    : MariaDB
Target Server Version : 50560
File Encoding         : 65001

Date: 2019-06-11 13:29:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for vdbench_filesystem
-- ----------------------------
DROP TABLE IF EXISTS `vdbench_filesystem`;
CREATE TABLE `vdbench_filesystem` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Datetime` datetime DEFAULT NULL,
  `Outputinterval` int(255) DEFAULT NULL,
  `ReqstdOpsRate` double(255,3) DEFAULT NULL,
  `ReqstdOpsResp` double(255,3) DEFAULT NULL,
  `CpuTotal` double(255,3) DEFAULT NULL,
  `CpuSys` double(255,3) DEFAULT NULL,
  `ReadPct` double(255,3) DEFAULT NULL,
  `ReadRate` double(255,3) DEFAULT NULL,
  `ReadResp` double(255,3) DEFAULT NULL,
  `WriteRate` double(255,3) DEFAULT NULL,
  `WriteResp` double(255,3) DEFAULT NULL,
  `MbSecRead` double(255,3) DEFAULT NULL,
  `MbSecWrite` double(255,3) DEFAULT NULL,
  `MbSecTotal` double(255,3) DEFAULT NULL,
  `XferSize` double(255,3) DEFAULT NULL,
  `MkdirRate` double(255,3) DEFAULT NULL,
  `MkdirResp` double(255,3) DEFAULT NULL,
  `RmdirRate` double(255,3) DEFAULT NULL,
  `RmdirResp` double(255,3) DEFAULT NULL,
  `CreateRate` double(255,3) DEFAULT NULL,
  `CreateResp` double(255,3) DEFAULT NULL,
  `OpenRate` double(255,3) DEFAULT NULL,
  `OpenResp` double(255,3) DEFAULT NULL,
  `CloseRate` double(255,3) DEFAULT NULL,
  `CloseResp` double(255,3) DEFAULT NULL,
  `DeleteRate` double(255,3) DEFAULT NULL,
  `DeleteResp` double(255,3) DEFAULT NULL,
  `Operationtabledate` datetime DEFAULT NULL,
  `Testcase` varchar(255) DEFAULT NULL,
  `Client_number` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1260893 DEFAULT CHARSET=utf8;
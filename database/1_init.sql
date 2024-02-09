create database tododb;
CREATE USER 'tododbuser'@'%' IDENTIFIED BY 'GHcn=zKhNr5o';
GRANT ALL PRIVILEGES ON tododb.* TO 'tododbuser'@'%';
/*grant all privileges on tododb.* to 'tododbuser'@'%' identified by 'GHcn=zKhNr5o';*/
flush privileges;

SET GLOBAL log_bin_trust_function_creators = 1;
SET GLOBAL log_bin_trust_function_creators = 0;
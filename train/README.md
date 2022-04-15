## 运行方式？

./train start|stop|em-down|em-up
#start 启动EM部署的所有服务
#stop 停止EM部署的所有服务
#em-down 停止em后台服务
#em-up 启动em后台服务

## config.yml配置说明

# databases
mysqldb.host: 172.16.10.221 #em 服务端ip
mysqldb.port: 3306
mysqldb.user: root
mysqldb.password: dtstack
mysqldb.dbname: dtagent

#agent EM服务端IP
agent.host: 172.16.10.221:8889

#em docker-compose路径
#docker-compose需要在系统PATH下
em.docker-compose: /opt/dtstack/easymanager/docker-compose.yml

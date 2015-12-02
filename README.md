# AirportCode-go
Building
========

go build

Running
=======

* Make sure you have an etcd environment set up and running on 127.0.0.1:2379
* Make sure you have mysql-server set up and running
* Add a user with credentials picasso:picasso who has full mysql permissions
* Add mysql to etcd
** etcdctl set /picasso/registry/mysql <addressOnly>
* ./AirportCode

API
===

GET /airport
GET /airport/<IATA>
GET /airport/id/<Id>
GET /airport/city/<City>
GET /airport/country/<Country>

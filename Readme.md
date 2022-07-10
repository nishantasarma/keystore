# Keystore

This project is now online

To get the value of the key abc-1

`curl 34.160.143.94/get/abc-1`

To search for keys with prefix xyz

`curl 34.160.143.94/prefix?xyz`


To search keys with suffix -1


`curl 34.160.143.94?suffix=-1`


To set a key and value use

`curl -ivk -trace 34.160.143.94/set  -H 'Content-Type:application/json' -X POST -d '{"abc" : "efg"}'`



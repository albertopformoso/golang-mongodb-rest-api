# Login to MongoDB with created User & Database by using

```
mongo -u [MONGO_INITDB_ROOT_USERNAME] -p [MONGO_INITDB_ROOT_PASSWORD] --authenticationDatabase [MONGO_INITDB_DATABASE]
```

or

```
mongo -u [MONGO_INITDB_ROOT_USERNAME] --authenticationDatabase [MONGO_INITDB_DATABASE]
```

# Connect to your database by using the URL
```
mongodb://[MONGO_INITDB_ROOT_USERNAME]:[MONGO_INITDB_ROOTPASSWORD]@127.0.0.1:27017/your-database-name
```
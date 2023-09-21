# blog
a blog project
frontend: vue3 + dayjs + elementPlus + md-editor-v3
backend:  gin + gorm + casbin + minio + zap
db: mysql + redis
deploy: github runner + docker-compose + nginx
server: ali-cloud ec2

![topo](https://github.com/ChocolateAceCream/blog/blob/master/topu.drawio.png)

demo: http://blog.easydelivery.ltd/
</br>
guestLogin: 'guest'
</br>
guestPassword: '123qwe!@#QWE'

### File structure
    minio               (docker-compose file)
    server
    ├── api
    │   └── v1
    ├── config          (config struct files that used to map config from config.yml)
    ├── db              (store db migrations)
    ├── docs            (swagger docs)
    ├── global          (global instance)
    ├── initialize      (initializer for all data models)
    ├── library         ()
    ├── log             (store server log files)
    ├── middleware
    ├── model           ()
    │   ├── dbTable     (data models)
    │   ├── request     (api request data structs)
    │   └── response    (api response data structs )
    ├── router          ()
    ├── service         (service layer)
    ├── unitTest         ()
    ├── source          ()
    └── utils           ()
        └── upload      (oss interface)

    web
    │   ├── router                 -- static routes and route guards
    │   ├── utils                  --
    │   │   ├── apiAxios           -- repackaged axios
    │   │   ├── date.js            -- time formatting
    │   │   ├── validate.js        -- form validation functions
    │   │   └── tree.js            -- convert list to tree
    |   ├── view                   -- pages
    |   |   ├── admin              --
    |   |   |   ├── endpoint       -- api management page
    |   |   |   ├── menu           -- menu management page
    |   |   |   ├── role           -- role management page
    |   |   ├── article
    |   |   |   ├── contentManagement       -- manage self posted/drafted articles
    |   |   |   ├── draft          -- article writing page
    |   |   |   ├── preview        -- article preview page
    |   |   ├── auth               -- login,register,password reset etc...
    |   |   ├── error              -- 404 page
    |   |   └── layout             -- footer, header, sidebar etc...
    └── vite.config.js             --

### conventions
1. for trees, backend only return array list and frontend will convert list to tree.

TOOD:
<!-- - [x] ~~~add README~~ -->
- [x] add README
- [x] Article preview/editing
- [x] Add comments add, list and delete
- [x] Add comments like count function
- [x] Add reply section
- [x] Add reply likes logic
- [X] Add notification system
- [ ] add loading more reply/comments button instead of infinite scroll
- [ ] rank most popular blogs (top viewed blogs) using go-routing to update data every 5 mins, then save the cached data to redis


Gotcha
1. no dao layer since gorm take in action

notice:
1. signature is calculated md5(APP_ID, nonce, timestamp, SIGNATURE_SECRET), turn off by default, change config to turn on this function
2. after adding a new model, remember to add it to /server/db/table_register.go
3. after adding a new route, remember to check if it's necessary also add it to casbin, endpoint list, initializer etc...
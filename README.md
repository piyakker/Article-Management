# Article Management Service
This is a simple application using go-kit used to build up a basic understanding with go-kit, microservice and clean architecture. With the help of three-layer developing pattern (service-endpoint-transport), I can easily decouple the business logic with communication method. 
I also use repository pattern to make the way of data persistence more flexible.Data are stored in memory currently. Last but not Least, I add basic mothod logging and error logging as middleware.

## Run
```
go run ./cmd/publishing
```

## APIs(RESTful API)
1. GetAllArticle
```
curl -GET localhost:8080/Articles
```

2. GetArticleByID
```
curl -GET localhost:8080/Articles/{id}
```

3. CreateArticle
```
curl -XPOST -d '{"title": "fun", "content": "funny", "author": "piyakker"}' localhost:8080/Articles
```

4. UpdateArticle
```
curl -XPUT -d '{"title": "not fun", "content": "not funny", "author": "piyakker"}' localhost:8080/Articles/{id}
```

5. DeleteArticle
```
curl -XDELETE localhost:8080/Articles/{id}
```
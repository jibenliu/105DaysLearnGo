#refer:[https://mp.weixin.qq.com/s/lakTSfqUOEYrE25oa2K6MQ](https://mp.weixin.qq.com/s/lakTSfqUOEYrE25oa2K6MQ)



### install
```
go get github.com/golang/mock/mockgen
```

### 初始化
```
├── mock
├── person
│   └── male.go
└── user
    ├── user.go
    └── user_test.go
```
```
mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock
```

##### 测试
```
go test ./user
// 查看测试覆盖率
go test -cover ./user
```

##### 可视化界面
```
1.生成测试覆盖率的 profile 文件
go test ./... -coverprofile=cover.out
2.利用 profile 文件生成可视化界面
go tool cover -html=cover.out
```
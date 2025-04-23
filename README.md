# SQLC Demo 项目

这是一个使用sqlc生成类型安全的Go代码与PostgreSQL交互的示例项目。

## 项目结构

```
sqlc-demo/
├── cmd/              // 命令行工具目录
│   └── main.go       // 主程序入口
├── config/           // 配置文件目录
│   ├── config.go     // 配置加载代码
│   └── config.yml    // 配置文件
├── internal/         // 内部代码
│   └── db/           // 数据库相关代码(由sqlc生成)
├── sqlc/             // SQLC源文件
│   └── sources/      // SQL源文件
│       ├── query.sql // 查询定义
│       └── schema.sql // 数据库架构
├── go.mod            // Go模块文件
├── go.sum            // Go依赖校验文件
├── docker-compose.yml // Docker配置文件
└── README.md         // 本文件
```

## 准备工作

### 安装依赖

1. 安装Go（1.16或更高版本）
2. 安装PostgreSQL（本地或使用Docker）
3. 安装sqlc工具

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### 安装项目依赖

```bash
go mod tidy
```

## 运行项目

使用Docker

1. 启动PostgreSQL容器

```bash
docker-compose up -d
```

2. 创建数据库并导入数据

```bash
docker exec -i my_postgres psql -U postgres -c "CREATE DATABASE testdb;"
cat sqlc/sources/schema.sql | docker exec -i my_postgres psql -U postgres -d testdb
cat testdata.sql | docker exec -i my_postgres psql -U postgres -d testdb
```

3. 更新配置文件（如需要）

编辑`config/config.yml`文件，确保数据库连接配置正确

4. 运行程序

```bash
go run cmd/main.go
```

5. 停止PostgreSQL容器

```bash
docker-compose down
```

## 修改和重新生成代码

如果你修改了`sqlc/sources`目录下的SQL文件，需要重新生成Go代码：

```bash
sqlc generate
```

## 配置说明

配置文件位于`config/config.yml`

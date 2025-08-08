# Nginx 配置文件管理工具

这是一个用 Go 语言编写的 Nginx 配置文件管理工具，可以根据 JSON 格式的配置文件自动生成 Nginx 配置文件。

## 功能特性

- 递归扫描目录查找 `nm-*.json` 文件
- 根据 JSON 配置生成对应的 Nginx `.conf` 配置文件
- 支持默认配置文件
- 使用 Cobra 命令行框架
- 使用 slog 日志记录

## 安装

```bash
go install
```

或者直接编译：

```bash
go build -o nginx-manager
```

## 使用方法

### 基本用法

```bash
# 递归扫描目录并生成配置文件
./nginx-manager trans /path/to/dir

# 使用指定的默认配置文件
./nginx-manager trans --default /path/to/default.json /path/to/dir
```

### 命令行参数

```bash
# 查看帮助信息
./nginx-manager --help

# 查看 trans 命令的帮助信息
./nginx-manager trans --help
```

## 配置文件格式

### 默认配置文件 (default.json)

默认配置文件包含服务器块和第一个 location 块的默认配置项。

```json
{
    "server_items": [
        "listen 443 ssl;",
        "listen [::]:443 ssl;",
        "http2 on;",
        "",
        "client_max_body_size 90000m;",
        "ssl_certificate /etc/nginx/ssl/certs/*.gmk.xdwanj.top.crt;",
        "ssl_certificate_key /etc/nginx/ssl/certs/*.gmk.xdwanj.top.key;",
        "",
        "# HSTS (ngx_http_headers_module is required) (63072000 seconds)",
        "add_header Strict-Transport-Security \"max-age=63072000\" always;",
        "",
        "# intermediate configuration",
        "ssl_protocols TLSv1.2 TLSv1.3;",
        "ssl_ecdh_curve X25519:prime256v1:secp384r1;",
        "ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:AES128-GCM-SHA256:AES256-GCM-SHA384:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK;",
        "",
        "# see also ssl_session_ticket_key alternative to stateful session cache",
        "ssl_session_timeout 1d;",
        "ssl_session_cache shared:MozSSL:10m; # about 40000 sessions"
    ],
    "location_first_items": [
        "# 设置是否允许 cookie 传输",
        "add_header Access-Control-Allow-Credentials true;",
        "# 先省略程序自带",
        "proxy_hide_header Access-Control-Allow-Origin;",
        "# 允许请求地址跨域 * 做为通配符",
        "add_header Access-Control-Allow-Origin * always;",
        "# 允许跨域的请求方法",
        "add_header Access-Control-Allow-Methods 'GET, POST, PUT, DELETE, OPTIONS';",
        "# 请求头",
        "add_header Access-Control-Allow-Headers 'DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization';",
        "if ($request_method = 'OPTIONS') {",
        "    return 204;",
        "}"
    ]
}
```

### 服务配置文件 (nm-*.json)

服务配置文件定义特定服务的配置信息。

```json
{
    "server_name": "example.com",
    "locations": [
        {
            "location": "/",
            "proxy_pass": "http://localhost:8080"
        },
        {
            "location": "/api",
            "proxy_pass": "http://localhost:8081",
            "items": [
                "proxy_set_header X-Real-IP $remote_addr;",
                "proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;"
            ]
        }
    ]
}
```

## 项目结构

```
nginx-manager/
├── cmd/                 # 命令行接口
│   ├── root.go          # 根命令
│   └── trans.go         # trans 命令
├── config/              # 配置文件处理
│   ├── types.go         # 配置结构体定义
│   └── loader.go        # 配置加载器
├── internal/            # 内部包
│   ├── generator/       # 配置文件生成器
│   │   └── generator.go # Nginx 配置生成器
│   └── scanner/         # 文件扫描器
│       └── scanner.go   # 目录扫描器
├── main.go              # 程序入口
├── go.mod               # Go 模块定义
├── README.md            # 说明文档
├── default.json         # 默认配置文件示例
├── sample.json          # 服务配置文件示例
└── sample.conf          # 生成的配置文件示例
```

## 开发

### 依赖

- Go 1.24.5 或更高版本
- Cobra 命令行框架

### 构建

```bash
go build -o nginx-manager
```

### 测试

```bash
go test ./...
```

## 许可证

MIT License
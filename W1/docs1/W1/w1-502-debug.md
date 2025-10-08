# Incident: 502 Bad Gateway（Nginx 反代端口配置错误）

## 现象
- `curl -v http://localhost` 返回 `HTTP/1.1 502 Bad Gateway`
- 浏览器访问显示 502 页面

## 定位步骤
1) 查看反代配置：
   grep -n 'proxy_pass' /etc/nginx/sites-enabled/demo
   -> 配置为 127.0.0.1:8082
2) 查看后端监听端口：
   ss -tnlp | grep 808
   -> Flask 实际监听 127.0.0.1:8081
3) 查看 Nginx 错误日志：
   journalctl -u nginx -n 30 --no-pager
   -> 日志提示 upstream 连接被拒绝 (Connection refused)

## 根因
- Nginx 配置文件中的 `proxy_pass` 端口写错（应为 8081，实际写成 8082）

## 修复
- 将 `proxy_pass http://127.0.0.1:8082;` 改为 `...8081;`
- `sudo nginx -t && sudo systemctl reload nginx`
- 复测 `curl http://localhost` 返回 200

## 预防
- 新增上线前检查清单：反代端口与后端监听端口一致性校验
- 在 CI 中加入简单的端口连通性 smoke test（curl 预检）
- 记录排障三板斧：配置 → 实际监听 → 日志

#!/bin/bash
# diagnose.sh - 简易排障三板斧

echo "=== 1. 当前 proxy_pass 配置 ==="
grep -n 'proxy_pass' /etc/nginx/sites-enabled/* || echo "未找到 proxy_pass 配置"

echo ""
echo "=== 2. 后端监听端口 (ss) ==="
ss -tnlp | grep 808 || echo "未发现 808 端口监听"

echo ""
echo "=== 3. Nginx 最近报错日志 ==="
journalctl -u nginx -n 10 --no-pager

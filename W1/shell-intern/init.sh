#!/bin/bash
set -e   # 遇到错误就停止执行

echo "=== Init Script Started ==="

# 1. 显示系统信息
echo "当前用户: $(whoami)"
echo "当前日期: $(date)"
echo "系统版本:"
lsb_release -a

# 2. 创建新用户 devops（如果不存在）
USERNAME="devops"
if ! id "$USERNAME" &>/dev/null; then
  sudo useradd -m -s /bin/bash "$USERNAME"
  echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/$USERNAME
  echo "✅ 用户 $USERNAME 已创建并授予 sudo 权限"
else
  echo "ℹ️ 用户 $USERNAME 已存在"
fi

# 3. 设置时区为 Asia/Shanghai
sudo timedatectl set-timezone Asia/Shanghai
echo "✅ 时区已设置为 Asia/Shanghai"

# 4. 安装常用工具
sudo apt update -y
sudo apt install -y curl wget git htop vim
echo "✅ 常用工具已安装"

echo "=== Init Script Finished ==="

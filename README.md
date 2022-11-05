# 执行如下命令

# 初始化module
go mod init "github.com/chenglh/common"

# 获取资源
go mod tidy

# 初始化成仓库
git init

# 把远程仓库地址加进来
git remote add origin https://github.com/chenglh/common.git

# 设置密钥
git remote set-url origin https://xxxxxxxxxxxxxxxxxxxxxx@github.com/chenglh/common.git

# 添加项目所有变动文件
git add .

# 提交 -m 是版本内容说明
git commit -m "init"

# 推送到 main 分支
git push -u origin master

# 显示远程仓库
git remote -v

# 移除远程仓库
git remote remove origin
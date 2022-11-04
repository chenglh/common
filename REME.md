# 执行如下命令
go mod init "github.com/chenglh/common"
go mod tidy

# 初始化成仓库
git init

# 把远程仓库地址加进来
git remote add origin https://github.com/chenglh/common.git

# 添加项目所有变动文件
git add .

# 提交 -m 是版本内容说明
git commit -m "init"

# 推送到 master 分支
git push origin master
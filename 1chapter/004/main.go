/*
 * @example: 用Sublime Text3打开文件
 * @author: cody.guo
 */

1> 使用Sublime Text3编写Go代码
2> 安装git工具
3> 在Sublime上安装插件，输入Ctrl + Shift + P，然后输出install
4> 安装支持Go语言的插件：GoSublime
5> 安装支持git的插件：SublimeGit
6> 在命令行中输入git config --global credential.helper store
7> git config --global user.name "CodyGuo"
8> git config --global user.email "guoxiao219@126.com"
9> git config --global push.default "matching"
10> 先在cmd中输入git push 把用户名、密码保存，然后在Sublime 中就可以畅快的提交了。


C:\Users\Administrator>git config --list
core.symlinks=false
core.autocrlf=true
color.diff=auto
color.status=auto
color.branch=auto
color.interactive=true
pack.packsizelimit=2g
help.format=html
http.sslcainfo=/bin/curl-ca-bundle.crt
sendemail.smtpserver=/bin/msmtp.exe
diff.astextplain.textconv=astextplain
rebase.autosquash=true
user.name=CodyGuo
user.email=guoxiao219@126.com
push.default=matching
gui.recentrepo=D:/go/cpath/src/282example
credential.helper=store
http.sslverify=false
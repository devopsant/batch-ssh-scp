[English](README-Chinese.md)

在windows电脑上批量执行shell命令或者上传文件。



# 背景

程序员或者测试员需要在一批linux上批量执行命令或者上传文件，如果使用linux操作机，可以使用pssh命令。

如果使用windows，可以使用本工具，加速工作效率。



# bscp 批量上传文件工具

bscp.exe 会读取properties文件，确定执行的主机范围。执行把本地的文件上传到一批服务器的指定目录。

**使用方式**

1. bscp.exe 和 xxx.properties文件放在一起；
2. 直接启动bscp.exe 双击exe或者命令启动均可
3. 提示输入启用的properties文件名，默认是host，可以执行输入xxx，启用xxx.properties
4. 输入本地文件的完整路径，比如是完整路径。例如：c:\xxx\xx.sh 只支持单文件
5. 输入目标主机地址，例如：/tmp/ 必须是一个目录地址，这里不支持重命名文件
6. 回车确认。所有的服务器会依次执行，如果中间有错误，会直接中断
7. 执行结束，不会自动退出，需要ctrl+c 退出。防止窗口自动消失，无法看到结果。

# bssh 批量执行命令工具
bssh.exe 会读取properties文件，确定执行的主机范围。执行时，会用当前登录的用户身份，执行shell命令

** 使用方式 **

1. bssh.exe 和 xxx.properties文件放在一起；
2. 直接启动bssh.exe 双击exe或者命令启动均可
3. 提示输入启用的properties文件名，默认是host，可以执行输入xxx，启用xxx.properties
4. 输入执行的命令，回车开始执行。
5. 所有的服务器会依次执行，如果中间有错误，会直接中断
6. 执行结束，不会自动退出，需要ctrl+c 退出。防止窗口自动消失，无法看到结果。

# properties文件
1. bssh 和 bscp 使用的是同一个properties文件文件，可以按照不同的业务进行区分，用英文命名文件名
2. 使用 # 开头的为注释行，不起作用

# 源码
1. 本工具修改来自于 https://github.com/tianshiyeben/wgcloud-scp
2. 如果有不满足的地方，可以直接修改wgcloud-scp的源码，或者在本源码的基础上，继续修改。
3. 编译方式，进入对应的目录执行
```bat
# 编译源码
cd bat;
.\build.bat
```

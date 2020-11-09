#!/usr/bin/expect
# 仅限 Linux & MacOS 系统使用
# 运行脚本必须使用 ./make_release.sh 或者 expect -f ./make_release.sh
# 设置超时时间
set timeout 300

spawn make release

expect {
        "(yes/no)?"
        {send "yes\n";exp_continue}
        "Password:"
        {send "Grandeep@123\n"}
	" password:"
	{send "Grandeep@123\n"}
}
# 执行上面的操作后,保持交互状态,把控制权交给终端
interact

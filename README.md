[中文介绍](README-Chinese.md)



batch execute command(shell) or upload file on Windows.


# Background

Programer or QA **tester** need to execute shell command on selected linux server or upload same file to selected linux server.

If you use linux, 'pssh' is useful, But a lot of people use Windows.So it's a  trouble to waste time.

If you use Windows,you cloud use this tool to work faster.



# bscp.exe batch scp upload file

 bscp.exe reads the properties file to determine the range of hosts to execute. When execute bscp.exe, this program will upload local files to a specified directory on a number of servers.



**opration step**



1. put bscp.exe and xxx.properties file together;

2. Start bscp.exe. Double-click the exe or use cmd to run the command;

3. A message is displayed prompting you to enter the name of the properties file to be enabled. The default value is "host", will use host.properties;

4. Enter the full path of the local file, For example, c:\xxx\xx.sh supports only a single file;

5. Enter the destination dir full path,only support dir full path. For example, '/tmp/' must be a dir path because renaming files are not supported;

6. Press Enter to confirm. All servers will execute in sequence, and if there is an error in the middle, it will interrupt directly;

7. After the command is executed, the system does not automatically exit. You need to press ctrl+c to exit. Prevents the window from automatically disappearing without being able to see the results.

 

# bssh.exe batch ssh execute command
 

bssh.exe reads the properties file to determine the range of hosts to execute. 



**opration step**



1. put bssh.exe and xxx.properties together.

2. Start bssh.exe. Double-click the exe or use cmd to run the command;

3. A message is displayed prompting you to enter the name of the properties file to be enabled. The default value is "host", will use host.properties;

4. Type the command and press Enter to start the command.

5. All servers execute the command in sequence. If there is an error, the command is interrupted directly

6. After the command is executed, the system does not automatically exit. You need to press ctrl+c to exit. Prevents the window from automatically disappearing without being able to see the results.

 



# properties file
1. bssh.exe and bscp.exe use the same properties file, you can edit many properties to choose different server;
2. Lines that start with # are comment lines and do not work

# 

# source code

1. souce code from  https://github.com/tianshiyeben/wgcloud-scp , change code to load different properties file and other modify.
2. If you are not satisfied, you can modify the source code of wgcloud-scp project, or modify this project.
3. How to compile.
```bat
# compire code
cd bat;
.\build.bat
```

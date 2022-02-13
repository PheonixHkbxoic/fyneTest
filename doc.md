

## go常用命令及参数
https://www.cnblogs.com/embedded-linux/p/11616183.html

```text
执行命令go mod init在当前目录下生成一个go.mod文件，执行这条命令时，当前目录不能存在go.mod文件。如果之前生成过，要先删除；
如果你工程中存在一些不能确定版本的包，那么生成的go.mod文件可能就不完整，因此继续执行下面的命令；
执行go mod tidy命令，它会添加缺失的模块以及移除不需要的模块。执行后会生成go.sum文件(模块下载条目)。添加参数-v，例如go mod tidy -v可以将执行的信息，即删除和添加的包打印到命令行；
执行命令go mod verify来检查当前模块的依赖是否全部下载下来，是否下载下来被修改过。如果所有的模块都没有被修改过，那么执行这条命令之后，会打印all modules verified。
执行命令go mod vendor生成vendor文件夹，该文件夹下将会放置你go.mod文件描述的依赖包，文件夹下同时还有一个文件modules.txt，它是你整个工程的所有模块。在执行这条命令之前，如果你工程之前有vendor目录，应该先进行删除。同理go mod vendor -v会将添加到vendor中的模块打印出来；


go build -ldflags="-s -w -H windowsgui" maim.go -o main.exe
1.-s strip 去掉无用的符号
2.-w DWARF 去掉DWARF调试信息，得到的可执行程序不可用调试器调试
3.-H windowsgui 生成带GUI界面的程序时，可去掉dos黑框
 
以上为3个常用的参数，此外-ldflags '-extldflags "-static"' 为静态编译
如果，想更加清楚的看到编译过程可加-x 参数，如bulid -x ......
```


##  go exe icon
1.安装
> go get github.com/akavel/rsrc

2.创建manifest文件, 命名：main.exe.manifest ：
```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="1.0.0.0"
    processorArchitecture="x86"
    name="controls"
    type="win32"
></assemblyIdentity>
<dependency>
    <dependentAssembly>
        <assemblyIdentity
            type="win32"
            name="Microsoft.Windows.Common-Controls"
            version="6.0.0.0"
            processorArchitecture="*"
            publicKeyToken="6595b64144ccf1df"
            language="*"
        ></assemblyIdentity>
    </dependentAssembly>
</dependency>
</assembly>
```

3. 生成syso文件
> rsrc -manifest main.exe.manifest -ico rc.ico -o main.syso


4. 将生成的main.syso文件拷贝到main.go同级目录

5. 编译生成main.exe
> go build -o main.exe


## fyne bundle 中文字体打包
[fyne字体打包](https://www.jianshu.com/p/3bb11c475a8b)

注意包名生成的为main，则bundle.go,theme.go需要放在根目录下

也可自己修改包名为theme，统一自己放到theme包下即可








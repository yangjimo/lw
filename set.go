package lwapi

/*
乐玩插件接口说明

BindWindow


函数简介:

绑定指定的窗口,并指定这个窗口的屏幕颜色获取方式,鼠标仿真模式,键盘仿真模式,以及模式设定，注意：在没用使用此函数绑定窗口之前，对象的所有键盘鼠标命令都无效。

参数定义:

hwnd 整形数: 指定的窗口句柄，-1表示绑定桌面窗口。

display 整数型: 屏幕颜色获取方式　该参数的缺省值为0取值有以下几种

0 : 正常模式,平常我们用的前台截屏模式
1: gdi模式,用于窗口采用GDI方式刷新时. 此模式占用CPU较大.

2 : gdi2模式,此模式兼容性较强,但是速度比gdi模式要慢许多,如果gdi模式发现后台不刷新时,可以考虑用gdi2模式.
3: dx2模式,用于窗口采用dx模式刷新,如果dx方式会出现窗口所在进程崩溃的状况,可以考虑采用这种.采用这种方式要保证窗口有一部分在屏幕外.vista以上系统不需要移动也可后台.此模式占用CPU较大.

4: dx模式,注意此模式需要管理员权限

5:opengl，用于窗口采用Opengl模式刷新。注意此模式需要管理员权限。

mouse 整数型: 鼠标仿真模式 取值有以下几种   该参数的缺省值为0

0: 正常模式,平常我们用的前台鼠标模式
1：Windows模式,采取模拟windows消息方式 同按键自带后台插件.

2，Windows3模式，采取模拟windows消息方式,可以支持有多个子窗口的窗口后台.


3，dx模式,采用模拟dx后台鼠标模式。有些窗口在此模式下绑定时，需要先激活窗口再绑定(或者绑定以后激活)，否则可能会出现绑定后鼠标无效的情况注意此模式需要管理员权限

4:dx2模式,同DX，如果DX无效，可以尝试下这个。



keypad 整数型: 键盘仿真模式  该参数的缺省值为0取值有以下几种

0: 正常模式,平常我们用的前台键盘模式
1: Windows模式,采取模拟windows消息方式 同按键的后台插件.
2: Windows3模式,采取模拟windows消息方式 同按键的后台插件，如果上一个模式不能正常按组合键，可能尝试此模式。

3:dx模式,采用模拟dx后台键盘模式。有些窗口在此模式下绑定时，需要先激活窗口再绑定(或者绑定以后激活)，否则可能会出现绑定后键盘无效的情况.

4:dx2模式,同DX，如果DX无效，可以尝试下这个。



added 整数型：附加信息，该参数的缺省值为0取值有以下几种，各值可以相加

1：锁定前台窗口，让目标窗口一直以为自己是前台窗口; 注意，部分窗口在此模式下会耗费大量资源慎用.


2：锁定窗口激活状态，保持目标窗口一直处于激活状态; 注意，部分窗口在此模式下会耗费大量资源慎用.

4：锁定鼠标位置，如果鼠标不能正常点击，可以尝试加上这个;

8：锁定键盘鼠标按键状态，如果不能正常后台后按键或者鼠标点击，可以尝试加上这个;

16：需要取后台鼠标特征码;

32：需要用输入法输入文本;

64：增强型锁定按键状态(如果ctrl,shift,alt键无效可尝试加上这个。) ；

128：增强型降CPU,需要在绑定图像模式不为，dx,opengl,opengl2模式时，仍要降低目标进程的CPU占用；

256：增强最小化截图，此模式主要是针对采用DX模式刷新的程序，如果DX图色模式可以正常截图，但不能最小化截图，可以尝试加上这个。


512：增强鼠标。某些窗口的某些鼠标动作要开启这个才能正常执行。





mode 整形数: 模式。 取值有以几下种  该参数的缺省值为0

     0 : 推荐模式此模式比较通用，而且后台效果是最好的，并且会隐藏目标进程中的lw.dll.

     1 : 如果嫌用模式0绑定慢，或者模式0和模式2会造成目标进程崩溃可以用这个模式，这个模式绑定速度最快。

     2 : 同模式0,如果模式0或者模式1绑定失败，可以尝试这个模式。

     3：如果前几个模式都绑定不了，一定要试试这个（如果绑定不成功，可能要较长时间判断）。

     4：强力绑定模式，如果前面几个绑不上，一定要试试这个。

 返回值:

整形数:
0: 失败
1: 成功



示例:


// display: 前台 鼠标:前台键盘:前台 模式0
ret = lw.BindWindow(hwnd,0,0,0,0,0)

// display: gdi 鼠标:前台 键盘:前台 模式1
ret = lw.BindWindow(hwnd,1,0,0,1)

// display: dx 鼠标:windows 键盘:windows 模式0
ret = lw.BindWindow(hwnd,3,1,1,0,0)


注意:


绑定之后,所有的坐标都相对于窗口的客户区坐标(不包含窗口边框)
另外,绑定窗口后,必须加以下代码,以保证所有资源正常释放

这个函数的意思是在脚本结束时,会调用这个函数。需要注意的是，目前的按键版本对于这个函数的执行不是线程级别的，也就是说，这个函数只会在主线程执行，子线程绑定的大漠对象，不保证完全释放。
Sub OnScriptExit()
    ret = lw.UnBindWindow()
End Sub

另外 绑定dx会比较耗时间,请不要频繁调用此函数.

另外如果绑定的是dx,要注意不可连续操作dx,中间至少加个10MS的延时,否则可能会导致操作失败.比如绑定图色DX,那么不要连续取色等,键鼠也是一样.



还有一点特别要注意的是,有些窗口绑定之后必须加一定的延时,否则后台也无效.一般1秒到2秒的延时就足够.



发现绑定失败的几种可能(一般是需要管理员权限的模式才有可能会失败)

1.     系统登录的帐号必须有Administrators权限

2.     如果是vista以上，启动窗口进程必须用管理员模式启动，脚本也必须用管理员模式启动.

3.     一些防火墙会防止插件注入窗口所在进程，比如360防火墙等，必须把lw.dll设置为信任.

4.     还有一个比较弱智的可能性，那就是插件没有注册到系统中，这时CreateObject压根就是失败的. 检测对象是否创建成功

5.     窗口所在进程有保护.

*/
func (com *LwSoft) BindWindow(hwnd, display, mouse, keypad, added, mode int) int {
	var ret, _ = com.lw.CallMethod("BindWindow", hwnd, display, mouse, keypad, added, mode)
	return int(ret.Val)
}

/*
函数简介:
解除绑定窗口,并释放系统资源.一般在OnScriptExit调用

参数定义:
返回值:
整形数:
0: 失败
1: 成功

*/
func (com *LwSoft) UnBindWindow() int {
	var ret, _ = com.lw.CallMethod("UnBindWindow")
	return int(ret.Val)
}

/*
函数简介:
 强制解除绑定窗口,并释放系统资源.
 参数定义:
  hwnd 整形数: 需要强制解除绑定的窗口句柄.

  返回值:
  整形数:0: 失败  1: 成功
注: 此接口会绑除所有的对象的绑定，不管是不是本进程绑定的，注意如果仍然可能会有对象对目标窗口操作，那么就不应该使用此函数。
*/
func (com *LwSoft) ForceUnBindWindow(hwnd int) int {
	var ret, _ = com.lw.CallMethod("ForceUnBindWindow", hwnd)
	return int(ret.Val)
}

/*
函数简介:
设定图色的获取方式，默认是显示器或者后台窗口(具体参考BindWindow)

参数定义:
mode 字符串: 图色输入模式取值有以下几种
1.     "0" 这个是默认的模式，表示使用显示器或者后台窗口
2.     "filename" ,如（“C\1.bmp”）指定输入模式为指定的图片,如果使用了这个模式，则所有和图色相关的函数均视为对此图片进行处理，比如文字识别查找图片 颜色 等等一切图色函数.
3.     "addr" ,如("1234126")指定输入模式为指定的图片,此图片在内存当中. addr为图像内存地址.注意,地址是十进制的!
如果使用了这个模式，则所有和图色相关的函数,均视为对此图片进行处理.比如文字识别 查找图片 颜色 等等一切图色函数.

返回值:
整形数: 0: 失败  1: 成功

示例:
// 设定为默认的模式
ret = lw.SetDisplayInput("0")
// 设定为图片模式 图片采用相对路径模式 相对于SetPath的路径
ret = lw.SetDisplayInput("test.bmp")
// 设为图片模式 图片采用绝对路径模式
ret = lw.SetDisplayInput("d:\test\test.bmp")
// 设置为图片模式,图片从内存中获取
ret = lw.SetDisplayInput("123456")
注意，如果是缓冲区中的图像数据，不支持加密图像。
*/
func (com *LwSoft) SetDisplayInput(mode int) int {
	var ret, _ = com.lw.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}

/*
函数简介:
获取当前对象已经绑定的窗口句柄. 无绑定返回0

返回值:
整形数: 窗口句柄
*/
func (com *LwSoft) GetBindWindow() int {
	var ret, _ = com.lw.CallMethod("GetBindWindow")
	return int(ret.Val)
}

/*
函数简介:
在不解绑的情况下,切换绑定窗口.(必须是同进程窗口)

参数定义:
hwnd 整形数: 需要切换过去的窗口句柄

返回值:
整形数: 0: 失败  1: 成功

示例:
// 绑定为后台
lw.SwitchBindWindow(hwnd1)
注:此函数一般用在绑定以后，窗口句柄改变了的情况。如果必须不解绑，那么此函数就很有用了。
*/
func (com *LwSoft) SwitchBindWindow(hwnd int) int {
	var ret, _ = com.lw.CallMethod("SwitchBindWindow", hwnd)
	return int(ret.Val)
}

/*
函数简介:

禁止外部输入到绑定窗口

参数定义:

lock 整形数: 0关闭锁定
       1 开启锁定(键盘鼠标都锁定)
       2 只锁定鼠标
       3 只锁定键盘


返回值:

整形数:
0: 失败
1: 成功

注意:此接口仅在绑定图色模式为DX,OPENGL,OPNEGL2,或者键盘模式为DX,或者鼠标模式为DX,或者附加信息不为0的时候有效另外，锁定完以后，记得要解除锁定，否则外部永远都无法输入了，除非解除了窗口绑定.

*/
func (com *LwSoft) LockInput(lock int) int {
	var ret, _ = com.lw.CallMethod("LockInput", lock)
	return int(ret.Val)
}

/*
函数简介:

降低目标窗口所在进程的CPU占用

函数原型:

int  DownCpu(rate)

参数定义:

rate 整形数: 取值范围大于0  取值为0 表示关闭CPU优化. 这个值越大表示降低CPU效果越好.（取值范围0－100，0表示关闭降CPU功能）

返回值:

整形数:
0: 失败
1: 成功

注意: 此接口必须在绑定窗口成功以后调用，窗口可以支持dx或者opengl或者opengl2或者在绑定的时候附加信息里设置有"增强降CPU"选项,否则降低CPU无效.

因为降低CPU是通过降低窗口刷新速度或者在系统消息循环增加延时来实现，所以注意，开启此功能以后会导致窗口刷新速度变慢.


*/
func (com *LwSoft) DownCpu(rate int) int {
	var ret, _ = com.lw.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

/*
函数简介:

禁止绑定窗口位置发生改变。



函数原型:
int LockMove(lock)

参数定义:

lock 整形数: 0=允许移动 ;1=禁止移动.



返回值:

整形数:
0: 失败
1: 成功

注意:

1,此接口仅在绑定图色模式为DX,OPENGL,OPNEGL2,或者键盘模式为DX,或者鼠标模式为DX,或者附加信息不为0的时候有效另外，锁定完以后，记得要解除锁定，否则外部永远都无法输入了，除非解除了窗口绑定.

2,禁止以后，也无法通过代码移动绑定窗口，同时会禁止窗口最小化。

*/
func (com *LwSoft) LockMove(lock int) int {
	var ret, _ = com.lw.CallMethod("LockMove", lock)
	return int(ret.Val)
}

/*
函数简介:

禁止绑定窗口尺寸发生改变。

参数定义:

lock 整形数: 0=允许尺寸改变 ;1=禁止尺寸改变.



返回值:

整形数:
0: 失败
1: 成功

注意:

1,绑定图色模式为DX,OPENGL,OPNEGL2,或者键盘模式为DX,或者鼠标模式为DX,或者附加信息不为0的时候，使用命令也无法改变窗口的尺寸，否则只是禁止用户拉动边框。

*/
func (com *LwSoft) LockSize(lock int) int {
	var ret, _ = com.lw.CallMethod("LockSize", lock)
	return int(ret.Val)
}

/*
函数简介:

当绑定的图色模式为opengl或者dx时，开启或者关闭同步截图模式。

函数原型:



int  OpenSynscreenshot(isOpen)
参数定义:

isOpen 整形数: 0=关闭同步模式；1＝开启同步模式。

返回值:

整形数:
0: 失败
1: 成功

同步模式时，调用截图命令必须等目标窗口刷新一次时才会返回，好处是画面一定能同步，但在某些情况下，目标窗口是要很久才刷新一次的。

对同一个目标窗口而言，只要有一个对象调用了这个命令，其它的对象也会切换过来。

注意，仅对opengl,,DX模式有效。pengl默认为异步，DX默认为同步。
*/
func (com *LwSoft) OpenSynscreenshot(isOpen int) int {
	var ret, _ = com.lw.CallMethod("OpenSynscreenshot", isOpen)
	return int(ret.Val)
}

/*
函数简介:

开启或者关闭暴力刷新模式。

参数定义:

isOpen 整形数: 0=关闭暴力模式；1＝开启暴力刷新模式。

返回值:

整形数:
0: 失败
1: 成功

在某些情况下，截到的图片始终和目标窗口不同步（如雷电模拟器的某些页面），这时候可以尝试开启暴力刷新模式。注意，针对同一个窗口，只要有一个对象开启了，其它的对象也会开启。

注意，仅对opengl,opengl2,DX模式有效。


*/
func (com *LwSoft) OpenViolentRedraw(isOpen int) int {
	var ret, _ = com.lw.CallMethod("OpenViolentRedraw", isOpen)
	return int(ret.Val)
}

/*
函数简介:

允许或者禁止绑定前自动激活目标窗口。


参数定义:

lock 整形数: 0=允许绑定前激活 ;1=禁止绑定前激活.

返回值:

整形数:
0: 失败
1: 成功

有些模式，在绑定前会自动激活一下窗口，如果不想激活，可以调用这个命令禁止掉，但是，有些模式，如果绑定前不激活一下，可能会有些异常，请自行判断。

*/
func (com *LwSoft) EnableSwitchWindow(flag int) int {
	var ret, _ = com.lw.CallMethod("EnableSwitchWindow", flag)
	return int(ret.Val)
}

/*
函数简介:

判定指定窗口是否已经被后台绑定. (前台无法判定)

参数定义:

hwnd 整形数: 窗口句柄

返回值:

整形数:
0: 没绑定,或者窗口不存在.
1: 已经绑定.


*/
func (com *LwSoft) IsBind(hwnd int) int {
	var ret, _ = com.lw.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

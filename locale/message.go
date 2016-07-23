// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locale

// 各个语种的语言对照表，通过相应文件的 init() 作初始化这样
// 在删除文件是，就自动删除相应的语言文件，不需要手修改代码。
var locales = map[string]map[string]string{}

// 各个语言需要翻译的所有字符串
const (
	SyntaxError  = "在[%v:%v]出现语法错误[%v]"      // app/errors.go:23
	OptionsError = "配置文件[%v]中配置项[%v]错误:[%v]" // app/errors.go:27

	// 与 flag 包相关的处理
	FlagHUsage              = "显示帮助信息"                    // main.go:28
	FlagVUsage              = "显示版本信息"                    // main.go:29
	FlagLUsage              = "显示所有支持的语言"                 // main.go:30
	FlagGUsage              = "在当前目录下创建一个默认的配置文件"         // main.go:31
	FlagPprofUsage          = "指定一种调试输出类型，可以为 cpu 或是 mem" // main.go:32
	FlagVersionBuildWith    = "%v %v build with %v\n"     // main.go:41
	FlagSupportedLangs      = "目前支持以下语言 %v\n"             // main.go:44
	FlagConfigWritedSuccess = "配置内容成功写入 %v\n"             // main.go:56
	FlagPprofWritedSuccess  = "pprof 的相关数据已经写入到 %v\n"     // main.go:73
	FlagInvalidPprrof       = "无效的 pprof 参数\n"            // main.go:89

	VersionInCompatible = "当前程序与配置文件中指定的版本号不兼容\n" // main.go:131
	Complete            = "完成！文档保存在 %v，总用时 %v\n"  // main.go:160

	DebugPort     = "当前为模板调试模式，调试端口为：%v\n" // output/html.go:58
	DebugTemplate = "当前为模板调试模式，调试模板为：%v\n" // output/html.go:59

	// 错误信息，可能在地方用到
	ErrRequired              = "不能为空"
	ErrInvalidFormat         = "格式不正确"
	ErrDirNotExists          = "目录不存在"
	ErrInvalidOutputType     = "无效的输出类型"            // output/output.go
	ErrTemplateNotExists     = "模板不存在"              // output/output.go
	ErrMkdirError            = "创建目录时发生以下错误%v"      // output/output.go:51
	ErrInvalidBlockType      = "无效的 block.Type 值%v" // input/block
	ErrUnsupportedInputLang  = "无效的输入语言%v"          // input
	ErrNotFoundEndFlag       = "找不到结束符号"            // input
	ErrNotFoundSupportedLang = "该目录下没有支持的语言文件"      // input/lang.go
)

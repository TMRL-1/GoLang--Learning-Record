Go 有三种函数用于输出文本：
Print()
Println()
Printf()

Print() 函数以默认格式打印其参数。
如果我们想在字符串参数之间添加一个空格，我们需要使用 " "：
如果两个参数都不是字符串，则 Print() 会在参数之间插入一个空格：fmt.Print("age", 18) // 输出：age 18

Println() 函数与 Print() 类似，不同之处在于参数之间添加空格，并在末尾添加换行符：

Printf() 函数首先根据给定的格式化动词格式化其参数，然后打印它们。%v 参数值  %T 参数类型

fmt.Sprint()  /  fmt.Sprintln()  /  fmt.Sprintf() 
只拼接字符串不输出，返回字符串，不算输出函数。



if  支持前置简单语句 + 条件表达式的写法，格式固定为：
if 简单语句; 布尔条件 {
}
在 if 内部定义一个临时变量，只在 if 块里生效
*不支持三元操作符(三目运算符) "a > b ? a : b"。

switch 语句
可省略 break  或 条件表达式
表达式不限制整数，啥类型都能比
case 可以是表达式，不一定是常量
一个 case 可以写多个值
支持 type switch（类型判断）
switch v := i.(type) {   或switch x.(type){
case int:
case string:
}

switch var1 {
    case val1:
    case val2:
    default:
}
变量 var1 可以是任何类型，而 val1 和 val2必须是相同的类型

package main
import "fmt"

type Student struct {
	ID    int     // 学号
	Name  string  // 姓名
	Age   int     // 年龄
	Score float64 // 成绩
}

//切片存储数据
var students []Student

//初始化
func init() {
	// 初始化切片，长度0，容量10
	students = make([]Student, 0, 10)
	fmt.Println("=== 学生信息管理工具初始化完成 ===")
}

func addStudent(s *[]Student) {
	var id, age int
	var name string
	var score float64

	fmt.Print("输入学生ID：")
	fmt.Scan(&id)
	fmt.Print("输入学生姓名：")
	fmt.Scan(&name)
	fmt.Print("输入学生年龄：")
	fmt.Scan(&age)
	fmt.Print("输入学生成绩：")
	fmt.Scan(&score)

	stu := Student{ID: id, Name: name, Age: age, Score: score}

	*s = append(*s, stu)
	fmt.Println("添加成功！")
}

//查询学生
func findStudent(id int) (Student, bool) {
	// range遍历切片，_是匿名变量，丢弃索引
	for _, s := range students {
		if s.ID == id {
			return s, true // 找到，返回学生+true
		}
	}
	return Student{}, false // 没找到，返回空结构体+false
}

//展示
func showAllStudents() {
	// 判断切片是否为空
	if len(students) == 0 {
		fmt.Println("暂无学生信息")
		return
	}

	fmt.Println("\n===== 所有学生信息 =====")
	for _, s := range students {
		fmt.Printf("ID：%d 姓名：%s 年龄：%d 成绩：%.2f\n", s.ID, s.Name, s.Age, s.Score)
	}
}

func main() {
//菜单无限循环
	for {
		fmt.Println("\n===== 学生信息管理菜单 =====")
		fmt.Println("1. 添加学生")
		fmt.Println("2. 查询学生")
		fmt.Println("3. 展示所有学生")
		fmt.Println("4. 退出程序")
		fmt.Print("请输入功能编号：")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("输入错误！")
			continue
		}

		switch choice {
		case 1:
			addStudent(&students)

		case 2:
			var searchID int
			fmt.Print("请输入查询ID：")
			fmt.Scan(&searchID)
			stu, ok := findStudent(searchID)
			if ok {
				fmt.Printf("ID：%d 姓名：%s 年龄：%d 成绩：%.2f\n", stu.ID, stu.Name, stu.Age, stu.Score)
			} else {
				fmt.Println("未找到该学生")
			}

		case 3:
			showAllStudents()

		case 4:
			fmt.Println("退出程序")
			return // 退出main函数

		default:
			fmt.Println("输入无效")
		}
	}
}
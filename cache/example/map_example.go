package main

import (
	"fmt"
	"pkg_v1/cache"
)

// User 定义一个示例结构体
type User struct {
	ID   int
	Name string
	Age  int
}

// MapCacheExample 展示 Map 缓存的基本使用方法
func MapCacheExample() {
	fmt.Println("\n=== 基本缓存操作示例 ===")
	// 创建一个新的 Map 缓存实例
	cache := cache.NewMapCache[string]()

	// 设置缓存项
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	fmt.Println("已设置两个缓存项")

	// 获取缓存项
	if value, exists := cache.Get("key1"); exists {
		fmt.Printf("获取到的值: %v\n", value)
	}

	// 检查键是否存在
	if _, exists := cache.Get("key2"); exists {
		fmt.Println("key2 存在")
	}

	// 删除缓存项
	cache.Remove("key1")
	fmt.Println("已删除 key1")

	// 获取缓存大小
	size := cache.Count()
	fmt.Printf("当前缓存大小: %d\n", size)

	// 清空所有缓存
	cache.Clear()
	fmt.Println("已清空所有缓存")
}

// MapCacheWithStructExample 展示如何使用 Map 缓存存储结构体
func MapCacheWithStructExample() {
	fmt.Println("\n=== 结构体缓存示例 ===")
	cache := cache.NewMapCache[User]()

	// 存储结构体
	user := User{
		ID:   1,
		Name: "张三",
		Age:  25,
	}
	cache.Set("user1", user)
	fmt.Println("已存储用户信息")

	// 获取并类型断言
	if value, exists := cache.Get("user1"); exists {
		fmt.Printf("获取到的用户信息: ID=%d, Name=%s, Age=%d\n", value.ID, value.Name, value.Age)
	}

	// 展示缓存大小
	fmt.Printf("当前缓存大小: %d\n", cache.Count())
}

func main() {
	fmt.Println("开始运行 Map 缓存示例...")

	// 运行基本示例
	MapCacheExample()

	// 运行结构体示例
	MapCacheWithStructExample()

	fmt.Println("\n示例运行完成！")
}

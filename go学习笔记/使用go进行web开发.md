# 使用Go进行WEB开发

## 概述

特征（分发速度快、安全、丰富的适合web开发的工具）
Go被设计成允许开发者可以快速的开发一个可伸缩的、安全的WEB应用，
1）使用简单
2）安全
3）高性能
4）支持最新的web技术包括：从HTTP/2、mysql、mongodb、elasticsearch到最新的加密技术（包括TLS 1.3）

## 特性功能

1. 快速的跨平台部署项目（Deploy across platforms in record speed）

For enterprises, Go is preferred for providing rapid cross-platform deployment. With its goroutines, native compilation, and the URI-based package namespacing, Go code compiles to a single, small binary—with zero dependencies—making it very fast.
(对应企业来说，GO被用于进行快速的跨平台的开发。通过它的goroutines、本地化编译和基于URL地址的包命名空间，Go的代码将会被编译成一个单一的、小巧二进制文件，并且没有依赖的组成部分)

2. 利用Go的开箱即用的性能可以轻松的实现程序的伸缩
- 编译成一个单一的文件
“Using static linking, Go actually combining all dependency libraries and modules into one single binary file based on OS type and architecture.”
（通过静态连接，实际上Go通过适配不同的操作系统类型和cpu架构将所有的依赖库和模块都编译成一个单一的二进制文件）
- 静态类型系统
“Type system is really important for large scale applications.”
（类型系统对于大型的可扩展的系统十分的重要）
- 执行性能
“Go performed better because of its concurrency model and CPU scalability. Whenever we need to process some internal request, we are doing it with separate Goroutines which are 10x cheaper in resources than Python Threads.”
(Go的高性能是基于它的并发模型和CPU的可伸缩性。无论任何时候，我们处理一个内部请求，我们都会将它分发到不同的协程上，这中间使用的资源比Python使用的线程的消耗少了至少10倍)
- 不需要一个WEB（No need for a web framework）
- 良好的集成环境和调试机制(Great IDE support and debugging)




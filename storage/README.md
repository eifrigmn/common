## 数据库相关

整理数据库相关驱动，将主要的数据库驱动进行二次封装，方便项目中直接引用，目录结构大致为：

~~~shell
.
├── badger
│   ├── tools
│   ├── v1
│   └── v2
├── bench
│   ├── badger
│   └── common
├── common
├── data
│   └── badger
│       ├── v1
│       └── v2
├── mock
├── storage.go
└── storage_test.go
~~~

固定目录为：

`bench`：对封装后各个包的性能测试

`common`：公共的数据结构定义，以及共用的常量

`mock`：方便运行测试用例而固化的数据

### Getting Started

选择需要使用的驱动，以levelDB的`badger`为例：

~~~shell
go get -u github.com/eifrigmn/common/storage/badger/v1
~~~

+ 初始化DB实例

~~~go
// 引入包
import ldb github.com/eifrigmn/common/storage/badger/v1

// init
dataPath := "path/to/levelDB/data"
db := ldb.NewDatastore(dataPath)
~~~

### 使用

DB初始化后，即可使用`storage.go`中定义的相关方法。

### 运行测试用例

每个包附有对应的测试用例文件，如需使用，请先修改`storage/mock/mock.go`文件中`BaseDataPath`变量的值，该变量用于指定levelDB数据文件的根目录。


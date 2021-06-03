module example.com/hello

go 1.16

replace example.com/greetings => ../greetings  // 这个注释掉，会编译失败

//require example.com/greetings v0.0.0-00010101000000-000000000000
require example.com/greetings v1.1.0

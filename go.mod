module github.com/rocwg/ged

go 1.27

require (
	// 引入契约仓的根模块（不要带子路径，版本号用 v0.0.0 占位即可）
	github.com/rocwg/grpc-contracts v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.82.0
)

require (
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260706201446-f0a921348800 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

// 本地开发时启用，提交代码前注释掉或删除
// 本地联动指向：直接指向根目录级别
replace github.com/rocwg/grpc-contracts => ../../grpc-contracts

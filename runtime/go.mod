module github.com/rocwg/goro-edge/runtime

go 1.26.4

require google.golang.org/grpc v1.82.0

require (
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260706201446-f0a921348800 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

// 引入契约仓的根模块（不要带子路径，版本号用 v0.0.0 占位即可）
require github.com/rocwg/grpc-contracts v0.0.0-00010101000000-000000000000

// 本地开发时启用，提交代码前注释掉或删除
// 本地联动指向：直接指向根目录级别
replace github.com/rocwg/grpc-contracts => ../../grpc-contracts

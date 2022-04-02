protoc --proto_path=src --go_out=src --go_opt=paths=source_relative simple/simple.proto
protoc --proto_path=src --go_out=src --go_opt=paths=source_relative enums/enum_example.proto
protoc --proto_path=src --go_out=src --go_opt=paths=source_relative complex/complex.proto

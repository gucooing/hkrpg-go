cd  ./hkrpg
protoc --proto_path=. --go_out=. *.proto
cd  ../server
protoc --proto_path=. --go_out=. *.proto

package config


type configStruct struct{
	Port         int
	Sqlite3_path string
}
var Conf configStruct
func NewConfig(){
	Conf = configStruct{
		Port:	50051,
		// Sqlite3_path: os.Getenv("SQLITE3_PATH"),
		Sqlite3_path: "/Users/iba/IT/06_gRPC_graqhql/sample_local_go_grpc_graphql/article/mydb.sqlite3",
	}
}

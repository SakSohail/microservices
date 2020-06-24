package main
)
const(
	port = ":50051"
)

type repository interface{
   Create(*pb.Con)
}
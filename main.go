package main

import { 
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
}
type MongoField struct{
	Fieldstr string 'json: "Field Str"'
	FieldInt int 'json: "Field Int"'
}

func main() {
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	clientOptions:= options.Client().ApplyURI("mongodb;//localhost:27017")
fmt.Println("ClientOptions TYPE:", reflect.Typeof(clientOptions), "\n")
client, err := mongo.Connect(context.TODO(),clientOptions)
if err != nil {
	fmt.Println("Mongo.connect() ERROR: ",err)
	os.Exit(1)
}
ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
col := client.Database("First_Document").collection("First Collection")
fmt.Println("Collection type: ", reflect.TypeOf(col), "\n")

oneDoc := MongoFields(
	FieldStr: " Upload Your Instagram picture here",
	FieldInt: 253648,
)
fmt.Println("oneDoc Type: ", reflect.Typeof(oneDoc), "\n")
result, insertErr := col.InsertOne(ctx, oneDoc)
if insertErr != nil {
	fmt.Println("InsertOne ERROR: ", insertErr)
	os.Exit(1)
}
else{
	fmt.Println("InsertOne() result Type: ", reflect.TypeOf(result))
	fmt.Println("InsertOne() api result Type: ", result)

	newID := result.InsertedID
	fmt.Println("InsertedOne(), newID" , newID)
	fmt.Println("InsertedOne(), newID type: " , reflect.TypeOf(newID))
}

}


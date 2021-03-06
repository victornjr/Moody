package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "encoding/json"
	  "net/http"

    // "github.com/mongodb/mongo-driver/bson"
    // "github.com/mongodb/mongo-driver/bson/primitive"
    // "github.com/mongodb/mongo-driver/mongo"
    // "github.com/mongodb/mongo-driver/mongo/options"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var pinguinos_collection *mongo.Collection
var leones_collection *mongo.Collection
var elefantes_collection *mongo.Collection
var ctx context.Context


type Pinguino struct {
  Caminando           string    `default:"-", json:caminando`
  De_pie              string    `default:"-", json:de_pie`
  De_panza            string    `default:"-", json:de_panza`
  Limpiandose         string    `default:"-", json:limpiandose`
  Limpiandose_en_agua string    `default:"-", json:limpiandose_en_agua`
  Nadando             string    `default:"-", json:nadando`
  Interactuando       string    `default:"-", json:interactuando`
  Peleando            string    `default:"-", json:peleando`
  Cargando_piedras    string    `default:"-", json:cargando_piedras`
  Haciendo_nidos      string    `default:"-", json:haciendo_nidos`
  Echados_en_nidos    string    `default:"-", json:echados_en_nidos`
  Alimentando_crias   string    `default:"-", json:alimentando_crias`
}
type Leon struct {
  Dormido             string    `default:"-", json:dormido`
  Jugando             string    `default:"-", json:jugando`
  Lamiendo            string    `default:"-", json:lamiendo`
  Nadando             string    `default:"-", json:nadando`
  Rugiendo            string    `default:"-", json:rugiendo`
}
type Elefante struct {
  Abanicando          string    `default:"-", json:abanicando`
  Banando             string    `default:"-", json:banando`
  Durmiendo           string    `default:"-", json:durmiendo`
  Comiendo            string    `default:"-", json:comiendo`
  Jugando             string    `default:"-", json:jugando`
}


func connect_to_db() {
  db_uri := os.Getenv("db_uri")
  ctx = context.Background()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_uri))
  if err != nil { log.Fatal(err) }
  err = client.Ping(ctx, nil)
  if err != nil { log.Fatal(err) }
  fmt.Println("Connected to MongoDB!")

  db := client.Database("animals")
  pinguinos_collection = db.Collection("pinguinos")
  leones_collection = db.Collection("leones")
  elefantes_collection = db.Collection("elefantes")
}


func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}


func add_pinguino(caminando string, de_pie string, de_panza string,
limpiandose string, limpiandose_en_agua string, nadando string,
interactuando string, peleando string, cargando_piedras string,
haciendo_nidos string, echados_en_nidos string, alimentando_crias string) {
  context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  res, err := pinguinos_collection.InsertOne(context, bson.M{
    "caminando": caminando, "de_pie": de_pie, "de_panza": de_panza,
    "limpiandose": limpiandose, "limpiandose_en_agua": limpiandose_en_agua,
    "nadando": nadando, "interactuando": interactuando, "peleando": peleando,
    "cargando_piedras": cargando_piedras, "haciendo_nidos": haciendo_nidos,
    "echados_en_nidos": echados_en_nidos, "alimentando_crias": alimentando_crias})
  if err != nil { log.Fatal(err) }
  fmt.Printf( "new pinguino created with id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
}

func add_lion(dormido string, jugando string, lamiendo string,
  nadando string, rugiendo string) {
  context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  res, err := leones_collection.InsertOne(context, bson.M{
    "dormido": dormido, "jugando": jugando, "lamiendo": lamiendo,
    "nadando": nadando, "rugiendo": rugiendo})
  if err != nil { log.Fatal(err) }
  fmt.Printf( "new leon created with id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
}

func add_elephant(abanicando string, banando string, durmiendo string,
  comiendo string, jugando string) {
  context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  res, err := elefantes_collection.InsertOne(context, bson.M{
    "abanicando": abanicando, "banando": banando, "durmiendo": durmiendo,
    "comiendo": comiendo, "jugando": jugando})
  if err != nil { log.Fatal(err) }
  fmt.Printf( "new elephant created with id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
}


func obtain_all_pinguinos(collection *mongo.Collection, ctx context.Context) []Pinguino {
  var results []Pinguino
  filter := bson.M{}
  findOptions := options.Find()

  cursor, err := collection.Find(ctx, filter, findOptions)
  if err != nil { log.Fatal(err) }
  for cursor.Next(ctx) {
      var p Pinguino
      if err := cursor.Decode(&p); err != nil { log.Fatal(err) }
      // fmt.Printf("p: %+v\n", p)
      results = append(results, p)
  }
  if err := cursor.Err(); err != nil { log.Fatal(err) }
  cursor.Close(ctx)
  return results
}

func obtain_all_leones(collection *mongo.Collection, ctx context.Context) []Leon {
  var results []Leon
  filter := bson.M{}
  findOptions := options.Find()

  cursor, err := collection.Find(ctx, filter, findOptions)
  if err != nil { log.Fatal(err) }
  for cursor.Next(ctx) {
      var p Leon
      if err := cursor.Decode(&p); err != nil { log.Fatal(err) }
      // fmt.Printf("p: %+v\n", p)
      results = append(results, p)
  }
  if err := cursor.Err(); err != nil { log.Fatal(err) }
  cursor.Close(ctx)
  return results
}

func obtain_all_elefantes(collection *mongo.Collection, ctx context.Context) []Elefante {
  var results []Elefante
  filter := bson.M{}
  findOptions := options.Find()

  cursor, err := collection.Find(ctx, filter, findOptions)
  if err != nil { log.Fatal(err) }
  for cursor.Next(ctx) {
      var p Elefante
      if err := cursor.Decode(&p); err != nil { log.Fatal(err) }
      // fmt.Printf("p: %+v\n", p)
      results = append(results, p)
  }
  if err := cursor.Err(); err != nil { log.Fatal(err) }
  cursor.Close(ctx)
  return results
}


func getjson_pinguinos(respuesta http.ResponseWriter, solicitud *http.Request) {
  results := obtain_all_pinguinos(pinguinos_collection, ctx)
  fmt.Printf("Found multiple pinginos: %+v\n\n", results)

  myjson, _ := json.Marshal(results)
  respuesta.Header().Set("Content-Type:", "aplication/json")
  respuesta.Write(myjson)
}

func getjson_leones(respuesta http.ResponseWriter, solicitud *http.Request) {
  results := obtain_all_leones(leones_collection, ctx)
  fmt.Printf("Found multiple leones: %+v\n\n", results)

  myjson, _ := json.Marshal(results)
  respuesta.Header().Set("Content-Type:", "aplication/json")
  respuesta.Write(myjson)
}

func getjson_elefantes(respuesta http.ResponseWriter, solicitud *http.Request) {
  results := obtain_all_elefantes(elefantes_collection, ctx)
  fmt.Printf("Found multiple elefantes: %+v\n\n", results)

  myjson, _ := json.Marshal(results)
  respuesta.Header().Set("Content-Type:", "aplication/json")
  respuesta.Write(myjson)
}


func penguin(respuesta http.ResponseWriter, solicitud *http.Request){
  switch solicitud.Method {
  case "GET":
    http.ServeFile(respuesta, solicitud, "visitors/penguin.html")
  case "POST":
    if err := solicitud.ParseForm(); err != nil { log.Fatal(err) }
    fmt.Printf("postForm: %+v\n", solicitud.PostForm)
    // fmt.Printf("caminando: %+v\n", solicitud.PostForm["caminando"][0:2])
    // for _, elem := range solicitud.PostForm["caminando"] {
    //   fmt.Printf("val caminando: %+v\n", elem)
    // }
    add_pinguino(
      solicitud.PostForm["caminando"][1], solicitud.PostForm["de_pie"][1],
      solicitud.PostForm["de_panza"][1],
      // solicitud.PostForm["limpiandose"][1],
      "-",
      // solicitud.PostForm["limpiandose_en_agua"][1],
      "-",
      solicitud.PostForm["nadando"][1],
      solicitud.PostForm["interactuando"][1], solicitud.PostForm["peleando"][1],
      solicitud.PostForm["cargando_piedras"][1], solicitud.PostForm["haciendo_nidos"][1],
      solicitud.PostForm["echados_en_nidos"][1], solicitud.PostForm["alimentando_crias"][1])
    http.ServeFile(respuesta, solicitud, "visitors/success.html")
  }
}


func lion(respuesta http.ResponseWriter, solicitud *http.Request){
  switch solicitud.Method {
  case "GET":
    http.ServeFile(respuesta, solicitud, "visitors/lion.html")
  case "POST":
    if err := solicitud.ParseForm(); err != nil { log.Fatal(err) }
    fmt.Printf("postForm: %+v\n", solicitud.PostForm)
    // fmt.Printf("caminando: %+v\n", solicitud.PostForm["caminando"][0:2])
    // for _, elem := range solicitud.PostForm["caminando"] {
    //   fmt.Printf("val caminando: %+v\n", elem)
    // }
    add_lion(
      solicitud.PostForm["dormido"][1], solicitud.PostForm["jugando"][1],
      solicitud.PostForm["lamiendo"][1], solicitud.PostForm["nadando"][1],
      solicitud.PostForm["rugiendo"][1])
    http.ServeFile(respuesta, solicitud, "visitors/success.html")
  }
}

func elephant(respuesta http.ResponseWriter, solicitud *http.Request){
  switch solicitud.Method {
  case "GET":
    http.ServeFile(respuesta, solicitud, "visitors/elephant.html")
  case "POST":
    if err := solicitud.ParseForm(); err != nil { log.Fatal(err) }
    fmt.Printf("postForm: %+v\n", solicitud.PostForm)
    // fmt.Printf("caminando: %+v\n", solicitud.PostForm["caminando"][0:2])
    // for _, elem := range solicitud.PostForm["caminando"] {
    //   fmt.Printf("val caminando: %+v\n", elem)
    // }
    add_elephant(
      solicitud.PostForm["abanicando"][1], solicitud.PostForm["banando"][1],
      solicitud.PostForm["durmiendo"][1], solicitud.PostForm["comiendo"][1],
      solicitud.PostForm["jugando"][1])
    http.ServeFile(respuesta, solicitud, "visitors/success.html")
  }
}


func home(respuesta http.ResponseWriter, solicitud *http.Request){
  // http.ServeFile(respuesta, solicitud, solicitud.URL.Path[1:])
  http.ServeFile(respuesta, solicitud, "index.html")
}
func done(respuesta http.ResponseWriter, solicitud *http.Request){
  http.ServeFile(respuesta, solicitud, "visitors/done.html")
}
func dashboard(respuesta http.ResponseWriter, solicitud *http.Request){
  http.ServeFile(respuesta, solicitud, "dashboard.html")
}
func demo1(respuesta http.ResponseWriter, solicitud *http.Request){
  http.ServeFile(respuesta, solicitud, "demo1.html")
}
func demo2(respuesta http.ResponseWriter, solicitud *http.Request){
  http.ServeFile(respuesta, solicitud, "demo2.html")
}
func demo3(respuesta http.ResponseWriter, solicitud *http.Request){
  http.ServeFile(respuesta, solicitud, "demo3.html")
}
func primerParcial(respuesta http.ResponseWriter, solicitud *http.Request){
  http.ServeFile(respuesta, solicitud, "primerParcial.html")
}


func main() {
  log.SetOutput(os.Stderr)
  log.SetOutput(os.Stdout)

  connect_to_db()
  http.HandleFunc("/", home)
  http.HandleFunc("/getjson_pinguinos", getjson_pinguinos)
  http.HandleFunc("/getjson_leones", getjson_leones)
  http.HandleFunc("/getjson_elefantes", getjson_elefantes)
  http.HandleFunc("/visitors/penguin", penguin)
  http.HandleFunc("/visitors/lion", lion)
  http.HandleFunc("/visitors/elephant", elephant)
  http.HandleFunc("/visitors/done", done)
  http.HandleFunc("/dashboard", dashboard)
  http.HandleFunc("/demo1", demo1)
  http.HandleFunc("/demo2", demo2)
  http.HandleFunc("/demo3", demo3)
  http.HandleFunc("/primerParcial", primerParcial)
  http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

  addr, error := determineListenAddress()
  if error != nil { log.Fatal(error) }
  fmt.Println("Listening on "+addr+" ...")
	http.ListenAndServe(addr, nil)
}

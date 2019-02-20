const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');
const mongoose = require('mongoose');
const Schema = mongoose.Schema;

mongoose.connect("mongodb://localhost/animals");

var pinguinoSchemaJSON = {
  caminando: String,
  de_pie: String,
  de_panza: String,
  limpiandose: String,
  nadando: String,
  interactuando: String,
  peleando: String,
  cargando_piedras: String,
  haciendo_nidos: String,
  alimentando_crias: String
};
var pinguino_schema = new Schema(pinguinoSchemaJSON);
var Pinguino = mongoose.model("Pinguino", pinguino_schema);


const app = express();
app.set('port', process.env.PORT || 3000);
app.disable('x-powered-by');
app.set('view engine', 'html');
// app.use(express.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(express.static(__dirname));  // TODO: change to something like app.use(express.static("public"));

app.use(function(req, res, next){
  console.log("Looking for URL: " + req.url);
  next();
});

app.get('/', function(req,res){
  res.send('ok');
});

app.get('/visitors/penguin',function(req, res){
  console.log("penguin...");
  res.sendFile(
    'visitors/penguin.html'
    ,{ root: path.join(__dirname, './') }
  );
});

app.post('/visitors/penguin',function(req, res){
  console.log("enviar...");
  // console.log(req.body);
  var nuevo_registro = new Pinguino({
    caminando: true,
    de_pie: true,
    de_panza: false,
    limpiandose: false,
    nadando: false,
    interactuando: false,
    peleando: false,
    cargando_piedras: false,
    haciendo_nidos: false,
    alimentando_crias: false
  });
  nuevo_registro.save(function(){
    Pinguino.find(function(err, doc){
      res.send("datos guardados.");
      // console.log(doc)
      for (entry in doc){
        console.log(doc[entry]['_id']);
      }
    });
  });
});

app.listen(app.get('port'), function(){
  console.log('Express started on http://localhost:' + app.get('port') + ' press Ctrl-C to terminate');
});

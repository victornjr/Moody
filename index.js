const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');
const mongoose = require('mongoose');
const Schema = mongoose.Schema;
const db_uri = process.env.db_uri;

mongoose.connect(db_uri)
        .then(db => console.log('db connected'))
        .catch(err => console.log('ERROR: ' + err));

var pinguinoSchemaJSON = {
  caminando: String,
  de_pie: String,
  de_panza: String,
  limpiandose: String,
  limpiandose_en_agua: String,
  nadando: String,
  interactuando: String,
  peleando: String,
  cargando_piedras: String,
  haciendo_nidos: String,
  echados_en_nidos: String,
  alimentando_crias: String
};
var pinguino_schema = new Schema(pinguinoSchemaJSON);
var Pinguino = mongoose.model("Pinguino", pinguino_schema);

const app = express();
app.set('port', process.env.PORT || 3000);
app.disable('x-powered-by');
app.set('view engine', 'html');
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
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
  var body = req.body;
  console.log(body);
  var nuevo_registro = new Pinguino({
    caminando: body['caminando'][1],
    de_pie: body['de_pie'][1],
    de_panza: body['de_panza'][1],
    limpiandose: body['limpiandose'][1],
    limpiandose_en_agua: body['limpiandose_en_agua'][1],
    nadando: body['nadando'][1],
    interactuando: body['interactuando'][1],
    peleando: body['peleando'][1],
    cargando_piedras: body['cargando_piedras'][1],
    haciendo_nidos: body['haciendo_nidos'][1],
    echados_en_nidos: body['echados_en_nidos'][1],
    alimentando_crias: body['alimentando_crias'][1]
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

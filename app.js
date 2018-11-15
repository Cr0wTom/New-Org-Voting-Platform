var express = require("express"); 
var app = express();
var vote = require("./vote.js");


app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  res.header("Content-Type: text/html");
  next();
});

app.get('/query', function(req, res) { 
      var queryres = vote.query(
		        function back(param){
                res.send(param);
			      }
        );
});

app.get('/query_votings', function(req, res) { 
  var queryres = vote.query_votings(
        function back(param){
            res.send(param);
        }
    );
});

app.get('/graph', function(req, res) { 
  var queryres = vote.graph(
        function back(param){
            res.send(param);
        }
    );
});

app.get('/index.html', function(req,res) {
  res.sendfile("./index.html");
});

app.get('/voting/:id/:voting', function (req, res, next) {
    console.log('ID:', req.params.id);
    console.log('VOTING:', req.params.voting);
    next();
  }, function (req, res, next) {
      vote.createvoting(req.params.id,req.params.voting,
            function back(param){ res.send(param); }
      );
  });

app.get('/voter/:id/:vote/:voting', function (req, res, next) {
  console.log('ID:', req.params.id);
  console.log('VOTE:', req.params.vote);
  console.log('VOTING:', req.params.voting);
  next();
}, function (req, res, next) {
    vote.dovote(req.params.id,req.params.vote,req.params.voting,
          function back(param){ res.send(param); }
    );
});

app.listen(8080); 

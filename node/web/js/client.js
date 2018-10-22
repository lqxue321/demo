var http = require('http');

var options = {
    host:'localhost',
    port:'8080',
    path:'../html/index.html'
}

var callback = function (res){
    var boby = '';
    res.on('data',function(data){
        boby+=data;
    });
    res.on('end',function(daat){
        console.log(boby);
    })
}
var req = http.request(options,callback);
req.end();
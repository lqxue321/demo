var http = require('http');
var fs = require('fs');
var url = require('url');
var port = require('./port');

http.createServer(function(req,res){
    //解析请求，包括文件名
    var pathname = url.parse(req.url).pathname;
     //输出请求的文件名
    console.log('request'+pathname+'received');
    fs.readFile(pathname.substr(1),function(err,data){
        if(err){
            console.log(err);
            res.writeHead(404, {'Content-Type': 'text/html'});
        }else{
            res.writeHead(20,{'Content-type':'text/html'});
            res.write(data.toString());
        }
        res.end();
    });
}).listen(port.ServicePort);
console.log('hello');


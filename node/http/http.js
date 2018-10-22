var http= require('http');
var url=require('url');


var server = http.createServer(function(req,res){
    console.log(req.method+':'+req.url);
    res.writeHead(200,
        {'Content-Type':'text-html'}
    );
    res.end('<h1>Hello World</h1>');
});

server.listen(8080);
//console.log('Server is running at localhost:8080/');
console.log(url.parse('https://user:pass@host.com:8080/path/to/file?query=string#hash'));
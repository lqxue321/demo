var fs = require("fs");
var request = require("request");


//获取订单数据 key(data)
function getData(){
    var options = { method: 'GET',
    url: 'https://couchdb-cloud.sparkpad-dev.com//b1/004dd5b7-a78a-11e8-b621-784f439359e1',
    headers: 
        {Authorization: 'Basic dGVzdF9hZG1pbjo4ODg4ODg=' }
     };

request(options, function (error, response, body) {  
  console.log((JSON.parse(body).data)); 
});  
}


//获取 oc token
function getToken(){
var options = { method: 'POST',
  url: 'https://de.sparkpad-dev.com/oauth/token',
  headers: 
   {
       'Content-Type': 'application/json','Content-Type': 'application/x-www-form-urlencoded' },
  form: 
   { grant_type: 'password',
     client_secret: 'KCv9dwvBgTxB',
     client_id: '20cd8f0e-7745-4c04-97dd-c052c6ce3e42',
     username: 'restws.gtdx',
     password: 'restws.gtdx' }
    };

request(options, function (error, response, body) {
  if (error) throw new Error(error);
  console.log(body);
});
}

getData();

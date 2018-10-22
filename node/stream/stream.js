var fs= require('fs');

//异步读文件
function con1(){
fs.readFile('input.txt','utf-8',function(err,body){
    console.log(body);
});
}

//写文件
function wri1(){
    var data= " my name is sakura1";
    fs.writeFile('input1.txt',data,function(err){
        if(err){
            console.log(err);
        }     
    });
    
}

//同步读文件
function con2(){
    var data = fs.readFileSync('input2.txt','utf-8');
    console.log(data);
}



//
function wri2(){
    var data= " my name is sakura2";
    fs.writeFileSync('input2.txt',data);
}


//stat
function statDemo(){
    fs.stat('input1.txt',function(err,stat){
        if(err){
            console.log(err);
        }else{
            console.log(stat.isFile());
        }
    });
}

//读
function readStream(){
    var stream = fs.createReadStream('input1.txt','utf-8');
    stream.on('data',function(chunk){
        console.log('data');
        console.log(chunk);
    });

    stream.on('end',function(){
        console.log('end');
    });

    stream.on('err',function(err){
        console.log(err);
    });
}


//写
function writeStream1(){
    var stream= fs.createWriteStream('input1.txt','utf-8');
    stream.write(' my age is 10');
    stream.write('end');
    stream.end();
}


function writeStream2(){
    var stream = fs.createWriteStream('input1.txt','utf-8');
    stream.write(new Buffer('使用Stream写入二进制数据...', 'utf-8'));
 //   stream.write(new Buffer('end','utf-8'));
    stream.end();
}


//pipe
function pipeDemo(){
    var readStream = fs.createReadStream('input1.txt');
    var writeStream = fs.createWriteStream('input2.txt');
    readStream.pipe(writeStream);
}



//writeStream2();
pipeDemo();

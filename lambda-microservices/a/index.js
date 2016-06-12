/* The MIT License (MIT)
* 
* Copyright (c) 2016 Claudemiro
* 
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
* 
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/

var AWS = require('aws-sdk');

AWS.config.accessKeyId = 'ACCESS_KEY_ID';
AWS.config.secretAccessKey = 'SECRET_ACCESS_KEY';
AWS.config.region = 'us-east-1';

var lambda = new AWS.Lambda();

exports.handler = function(event, context, callback) {
    var params = {
        FunctionName: 'FunctionB',
        InvocationType: 'RequestResponse',
        Payload: JSON.stringify({name: 'Claudemiro'})
    };

    lambda.invoke(params, function(err, data) {
        if (err) {
            callback(err, err.stack); 
        } else { 
            callback(null, data);
        }
    });
}

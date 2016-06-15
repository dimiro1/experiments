const fs = require('fs');
const wkhtmltopdf = require('wkhtmltopdf');
 
// URL 
wkhtmltopdf('http://google.com/', { pageSize: 'letter' })
  .pipe(fs.createWriteStream('out.pdf'));
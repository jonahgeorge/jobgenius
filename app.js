var express = require('express')
  , app = express();

/* Express config */ 
require('./config')(app)

/* Bootstrap routes */
require('./routes/routes')(app)

app.listen(3000);

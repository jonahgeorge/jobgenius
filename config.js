
/* Module dependencies */
var express = require('express');
var RedisStore = require('connect-redis')(express);
var gzippo = require('gzippo');
var dust = require('dustjs-linkedin');
var cons = require('consolidate');

module.exports = function (app) {
	app.set('showStackError', true)

	// set views path, template engine and default layout
	app.engine('dust', cons.dust);
	app.set('views', './views');
	app.set('view engine', 'dust');

	// cookieParser should be above session
	app.use(express.cookieParser())

	// bodyParser should be above methodOverride
	app.use(express.bodyParser())
	app.use(express.methodOverride())

	// Redis session storage
	app.use(express.session({ 
		store: new RedisStore({
			host: 'localhost',
			port: 6379,
			db: 2,
			pass: ''
		}),
		secret : "keyboard cat"
	}))

	app.use(app.router)

	// Favicon
	app.use(express.favicon())

	// Compression via Gzip
	app.use(gzippo.staticGzip('./public', { maxAge : 86400000 }))
	app.use(gzippo.compress());
}

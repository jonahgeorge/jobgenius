/* Module Dependencies */
var users = require('../controllers/users');
var articles = require('../controllers/articles');
var interviews = require('../controllers/interviews');
var statics = require('../controllers/statics');
var middleware = require('./middleware');

/* Expose routes */
module.exports = function (app) {

	app.get('/', statics.main);

	app.get('/privacy', statics.privacy);
	app.get('/terms', statics.terms);
	app.get('/about', statics.about);

	/* Article Routes */
	app.get('/a', articles.index);
	app.post('/a', middleware.requiresLogin, articles.create);
	app.get('/a/add', middleware.requiresLogin, articles.form);
	app.get('/a/:aid/edit', articles.edit);
	app.get('/a/:aid', articles.show);
	app.post('/a/:aid', articles.update);
	app.delete('/a/:aid', articles.delete);

	/* Interview Routes */
	app.get('/i', interviews.index);
	app.get('/i/add', interviews.form);
	app.post('/i', middleware.requiresLogin, interviews.create);
	app.get('/i/:id/edit', interviews.edit);
	app.get('/i/:id', interviews.show);
	app.post('/i/:id', interviews.update);

	/* Auth Routes */
	app.get('/auth/signout', users.signout);
	app.get('/auth/signin', users.lform);
	app.get('/auth/signup', users.sform);

	app.post('/api/signin', users.signin);
	app.post('/api/signup', users.signup);
}

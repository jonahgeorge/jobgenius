/* Module Dependencies */
var bcrypt = require('bcrypt');
var request = require('request');
var pool = require('../database').pool;

/* Render Sign In Form */
exports.lform = function (req, res) {
	res.render('users/signin', { 
		title : "Sign In" 
	});
};

/* Render Sign Up Form */
exports.sform = function (req, res) {
	res.render('users/signup', { 
		title : "Sign Up" 
	});
};

/* Process User Sign Up */
exports.signup = function (req,res) {
	pool.getConnection(function(err, connection) {

		// Check if user exists already exists
		connection.query("SELECT * FROM D_USERS WHERE email = ?", req.body.email, function (err, rows) {

			// if so, throw error
			if (rows.length > 0) { 
				res.json({ 
					"status" : "error", 
					"message" : "Email already in use." 
				});
			} 
 
			// if not, insert user
			else { 
				bcrypt.genSalt(10, function(error, salt) {
					bcrypt.hash(req.body.password, salt, function (err, hash) {

						var user = {
							email : req.body.email,
							password : hash,
							timestamp : new Date()
						}

						connection.query("INSERT INTO D_USERS SET ?", user, function (err) {
							if (err) throw err;
							var entity = { email : req.body.email };
							req.session.entity = entity;
							res.json({ "status" : "success" });
						});
					});
				});
			}
		});

		connection.release();
	});
};

/* Process User Sign Out */
exports.signout = function (req, res) {
	req.session.destroy();	
	res.redirect('/');
};

/* Process User Sign In */
exports.signin = function (req, res) {
	pool.getConnection(function(err, connection) {
		connection.query("SELECT * FROM D_USERS WHERE email = ?", req.body.email, function (err, user) {
			bcrypt.compare(req.body.password, user[0].password, function(err, response) {

				if (err) throw err;

				if (response) {	

					var entity = { 
						uid : user[0].uid,
						email : user[0].email, 
						role : user[0].role,
						pseudo_username : user[0].email.split("@")[0]
					};

					req.session.entity = entity;
					return res.send(200);
				}

				res.send(401);
			});
		});
		connection.release();
	});
};

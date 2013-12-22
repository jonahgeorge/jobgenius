var pool = require('../database').pool;

exports.about = function (req, res) {
	res.render('statics/about', {
		title : "About Job Genius",
		entity : req.session.entity
	});
};

exports.privacy = function (req, res) {
	res.render('statics/privacy', {
		title : "Privacy Policy",
		entity : req.session.entity
	});
};

exports.terms = function (req, res) {
	res.render('statics/terms', {
		title : "Terms of Service",
		entity : req.session.entity
	});
};

exports.main = function (req, res) {

	var aq = "select * from C_ARTICLE as A left join C_USER as U ON A.uid = U.uid where A.published = 1";
	var iq = "select * from C_INTERVIEW as I left join C_USER as U ON I.uid = U.uid where I.published = 1";

	pool.getConnection(function (err, conn) {
		conn.query(aq + ";" + iq, function (err, results) {
			res.render('statics/main', {
				title : "Dashboard",
				articles : results[0],
				interviews : results[1],
				entity : req.session.entity
			});
		});
		conn.release();
	});
};

var pool = require('../database').pool;

exports.index = function (req, res) {
	pool.getConnection(function (err, conn) {

		conn.query("select * from D_ARTICLES as A left join D_USERS as U ON A.uid = U.uid where A.published = 1", function (err, articles) {
			if (err) console.log(err);
			res.render('articles/index', {
				title : "Articles",
				articles : articles,
				entity : req.session.entity
			});
		});

		conn.release();
	});
};

exports.show = function (req, res) {
	pool.getConnection(function (err, conn) {
		conn.query("select * from D_ARTICLES where aid = ?", req.params.aid, function (err, article) {
			
			if (err) console.log(err);
			res.render('articles/show', {
				article : article[0],
				entity : req.session.entity
			});

		});
		conn.release();
	});
};

exports.form = function (req, res) {
	res.render('articles/form', {
		entity : req.session.entity
	});
};

exports.edit = function (req, res) {
	pool.getConnection(function (err, conn) {
		conn.query("select * from D_ARTICLES where aid = ?", req.params.aid, function (err, article) {
			res.render('articles/form', {
				article : article[0],
				entity : req.session.entity
			});
		});
	});
};

exports.create = function (req, res) {
	pool.getConnection(function (err, conn) {

		var article = {
			title : req.body.title,
			subtitle : req.body.subtitle,
			body : req.body.body,
			uid : req.session.entity.uid,
			published : 1,
			timestamp : new Date()
		};

		conn.query("insert into D_ARTICLES set ?", article, function (err) {
			if (err) console.log(err);
			res.redirect('/a');
		});
        
		conn.release();
	});
};

exports.update = function (req, res) {
	pool.getConnection(function (err, conn) {

		var article = {
			title : req.body.title,
			subtitle : req.body.subtitle,
			body : req.body.body,
			timestamp : new Date()
		};

		conn.query("update D_ARTICLES set ? where aid = " + req.params.aid, article, function (err) {
			if (err) console.log(err);
			res.redirect( "/a/" + req.params.aid );
		});
	});
};

exports.delete = function (req, res) {
   	pool.getConnection(function (err, conn) {
	    conn.query("select * from D_ARTICLES where aid = ?", req.params.aid, function (err, article) {
	    	if (article.uid == req.session.entity.uid || req.session.entity.role == 0) {
			    conn.query("update D_ARTICLES set published = 0 where aid = ? ", req.params.aid, function (err) {
			        if (err) {
			        	console.log(err);
			        	res.json({ "status" : "Error" });
			        }
			        res.json({ "status" : "Success" });
			    });
			}
		});
		conn.release();
	});
};

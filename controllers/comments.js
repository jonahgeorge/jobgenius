var pool = require('../database').pool;

exports.create = function (req, res) {

	var comment = {
		iid: req.body.id,
		uid: req.session.entity.uid,
		value: req.body.value,
		published: 1
	};

	pool.getConnection(function (err, conn) {
		conn.query("insert into F_COMMENT set ?", comment, function (err) {
			if (err) {
				console.log(err);
				res.send(500);
			}
			res.send(200);
		});
		conn.release();
	});
};

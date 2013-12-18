var pool = require('../database').pool;

exports.index = function (req, res) {
	pool.getConnection(function (err, conn) {

		conn.query("select * from D_INTERVIEWS as I left join D_USERS as U ON I.uid = U.uid where I.published = 1", function (err, interviews) {
			if (err) console.log(err);
			res.render('interviews/index', {
				title : "Interviews",
				interviews : interviews,
				entity : req.session.entity
			});
		});

		conn.release();
	});
};

exports.show = function (req, res) {
	pool.getConnection(function (err, conn) {

		var q = "select di.iid, di.name, di.position, ft.value as type, fs.value as sector, "
		      + "fi.value as industry, fex.value as experience, fen.value as environment, "
		      + "fsa.value as salary, fh.value as hours_per_week, fw.value as weekends_worked, "
		      + "fo.value as overnight_travel, fc.value as certs, fss.value as soft_skills "
			  + "from D_INTERVIEWS as di "
			  + "left join F_TYPES as ft on ft.fid = di.type "
			  + "left join F_SECTORS as fs on fs.fid = di.sector "
			  + "left join F_INDUSTRIES as fi on fi.fid = di.industry "
  			  + "left join F_EXPERIENCE as fex on fex.fid = di.experience "
			  + "left join F_ENVIRONMENT as fen on fen.fid = di.environment "
			  + "left join F_SALARY as fsa on fsa.fid = di.salary "
			  + "left join F_HOURS_PER_WEEK as fh on fh.fid = di.hours_per_week "
			  + "left join F_WEEKENDS_WORKED as fw on fw.fid = di.weekends_worked "
			  + "left join F_OVERNIGHT_TRAVEL as fo on fo.fid = di.overnight_travel "
			  + "left join F_CERTS as fc on fc.fid = di.certs "
			  + "left join F_SOFT_SKILLS as fss on fss.fid = di.soft_skills "
			  + "where iid = " + conn.escape(req.params.iid);

		conn.query(q, function (err, interviews) {
			
			if (err) console.log(err);
			res.render('interviews/show', {
				interview : interviews[0],
				entity : req.session.entity
			});

		});

		conn.release();
	});
};

exports.form = function (req, res) {
	res.render('interviews/form', {
		entity : req.session.entity
	});
};

exports.edit = function (req, res) {
	pool.getConnection(function (err, conn) {

		var q0 = "select * from D_INTERVIEWS where iid = " + conn.escape(req.params.iid) + ";";

		// Field Queries
		var f1 = "select * from F_TYPES;";
		var f2 = "select * from F_SECTORS;";
		var f3 = "select * from F_INDUSTRIES;";
		var f4 = "select * from F_EXPERIENCE;";
		var f5 = "select * from F_ENVIRONMENT;";
		var f6 = "select * from F_SALARY;";
		var f7 = "select * from F_HOURS_PER_WEEK;";
		var f8 = "select * from F_WEEKENDS_WORKED;";
		var f9 = "select * from F_OVERNIGHT_TRAVEL;";
		var f10 = "select * from F_CERTS;";
		var f11 = "select * from F_SOFT_SKILLS;";

		var q = q0 + f1 + f2 + f3 + f4 + f5 + f6 + f7 + f8 + f9 + f10 + f11;

		conn.query(q, function (err, results) {
			res.render('interviews/form', {
				interview 			: results[0][0],
				types 				: results[1],
				sectors 			: results[2],
				industries 			: results[3],
				experience 			: results[4],
				environments 		: results[5],
				salary 				: results[6],
				hours_per_week 		: results[7],
				weekends_worked 	: results[8],
				overnight_travel 	: results[9],
				certs 				: results[10],
				soft_skills 		: results[11],
				entity 				: req.session.entity
			});
		});
	});
};

exports.create = function (req, res) {
	pool.getConnection(function (err, conn) {

		var interview = {
			name : req.body.name,
			position : req.body.position,
			uid : req.session.entity.uid,
			published : 1,
			timestamp : new Date()
		};

		conn.query("insert into D_INTERVIEWS set ?", interview, function (err) {
			if (err) console.log(err);
			res.redirect('/i');
		});
        
		conn.release();
	});
};

exports.update = function (req, res) {
	pool.getConnection(function (err, conn) {

		var article = {
			name 				: req.body.name,
			position 			: req.body.position,
			type 				: req.body.type,
			sector 				: req.body.sector,
			industry 			: req.body.industry,
			experience 			: req.body.experience,
			environment 		: req.body.environment,
			salary 				: req.body.salary,
			hours_per_week 		: req.body.hours_per_week,
			weekends_worked 	: req.body.weekends_worked,
			overnight_travel 	: req.body.overnight_travel,
			certs				: req.body.certs,
			soft_skills			: req.body.soft_skills,
			timestamp 			: new Date()
		};

		conn.query("update D_INTERVIEWS set ? where iid = " + req.params.iid, article, function (err) {
			if (err) console.log(err);
			res.redirect( "/i/" + req.params.iid );
		});
	});
};

exports.delete = function (req, res) {

};

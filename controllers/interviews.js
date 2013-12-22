var pool = require('../database').pool;

exports.index = function (req, res) {
	pool.getConnection(function (err, conn) {

		conn.query("select * from C_INTERVIEW as I left join C_USER as U ON I.uid = U.uid where I.published = 1", function (err, interviews) {
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

        var q = "select C_INTERVIEW.*, "
            + "L_TYPE.value as type, "
            + "L_SECTOR.value as sector, "
            + "industry.value as industry, "
            + "L_EXPERIENCE.value as experience, "
            + "environment.value as environment, "
            + "L_SALARY.value as salary, "
            + "L_HOURS_PER_WEEK.value as hours_per_week, "
            + "L_WEEKENDS_WORKED.value as weekends_worked, "
            + "L_OVERNIGHT_TRAVEL.value as overnight_travel, "
            + "group_concat(L_CERTIFICATION.value separator ', ') as certs, "
            + "soft_skills.value as soft_skills, "
			+ "education.value as education "
            
            + "from C_INTERVIEW "
            
            // Types
            + "left join F_TYPE on F_TYPE.iid = C_INTERVIEW.id "
            + "left join L_TYPE on L_TYPE.id = F_TYPE.vid "
            
            // Sectors
            + "left join F_SECTOR on F_SECTOR.iid = C_INTERVIEW.id "
            + "left join L_SECTOR on L_SECTOR.id = F_SECTOR.vid "
            
            // Industries
            + "left join ( "
                + "select F_INDUSTRY.iid, group_concat(L_INDUSTRY.value separator ', ') as value "
                + "from F_INDUSTRY "
                + "left join L_INDUSTRY on L_INDUSTRY.id = F_INDUSTRY.vid "
                + "group by F_INDUSTRY.iid "
            + ") as industry on industry.iid = C_INTERVIEW.id "
            
            // Experience
            + "left join F_EXPERIENCE on F_EXPERIENCE.iid = C_INTERVIEW.id "
            + "left join L_EXPERIENCE on L_EXPERIENCE.id = F_EXPERIENCE.vid "
            
            // Environment
            + "left join ( "
                + "select F_ENVIRONMENT.iid, group_concat(L_ENVIRONMENT.value separator ', ') as value "
                + "from F_ENVIRONMENT "
                + "left join L_ENVIRONMENT on L_ENVIRONMENT.id = F_ENVIRONMENT.vid "
                + "group by F_ENVIRONMENT.iid "
            + ") as environment on environment.iid = C_INTERVIEW.id "
            
            // Salary
            + "left join F_SALARY on F_SALARY.iid = C_INTERVIEW.id "
            + "left join L_SALARY on L_SALARY.id = F_SALARY.vid "
            
            // Hours per Week
            + "left join F_HOURS_PER_WEEK on F_HOURS_PER_WEEK.iid = C_INTERVIEW.id "
            + "left join L_HOURS_PER_WEEK on L_HOURS_PER_WEEK.id = F_HOURS_PER_WEEK.vid "
            
            // Weekends Worked
            + "left join F_WEEKENDS_WORKED on F_WEEKENDS_WORKED.iid = C_INTERVIEW.id "
            + "left join L_WEEKENDS_WORKED on L_WEEKENDS_WORKED.id = F_WEEKENDS_WORKED.vid "
            
            // Overnight Travel
            + "left join F_OVERNIGHT_TRAVEL on F_OVERNIGHT_TRAVEL.iid = C_INTERVIEW.id "
            + "left join L_OVERNIGHT_TRAVEL on L_OVERNIGHT_TRAVEL.id = F_OVERNIGHT_TRAVEL.vid "
            
            // Certifications
            + "left join F_CERTIFICATION on F_CERTIFICATION.iid = C_INTERVIEW.id "
            + "left join L_CERTIFICATION on L_CERTIFICATION.id = F_CERTIFICATION.vid "
            
            // Soft Skills
            + "left join ( "
                + "select F_SOFT_SKILL.iid, group_concat(L_SOFT_SKILL.value separator ', ') as value "
                + "from F_SOFT_SKILL "
                + "left join L_SOFT_SKILL on L_SOFT_SKILL.id = F_SOFT_SKILL.vid "
                + "group by F_SOFT_SKILL.iid "
            + ") as soft_skills on soft_skills.iid = C_INTERVIEW.id "
           
			// Education
			+ "left join ( "
				+ "select tmp.iid, group_concat(tmp.value separator ';') as value "
				+ "from ( "
					+ "select F_EDUCATION.iid, concat_ws(',', L_DEGREE.value, L_MAJOR.value, L_UNIVERSITY.value, F_EDUCATION.year) as value "
					+ "from F_EDUCATION "
					+ "left join L_UNIVERSITY on L_UNIVERSITY.id = F_EDUCATION.university "
					+ "left join L_DEGREE on F_EDUCATION.degree = L_DEGREE.id "
					+ "left join L_MAJOR on F_EDUCATION.major = L_MAJOR.id "
				+ ") as tmp"
			+ ") as education on education.iid = C_INTERVIEW.id "

            + "where C_INTERVIEW.id = " + conn.escape(req.params.id) + ";";

		conn.query(q, function (err, interviews) {
			if (err) console.log(err);
			var interview = interviews[0];

			console.log(interview);

			// Prepare Industry Data
			if (interview.industry != null) {
				var industries = interview.industry.split(",");
				interview.industry = new Array();
				industries.forEach( function (i) {
					interview.industry.push({ value : i });
				});
			}	

			// Prepare Environment Data
			if (interview.environment != null) {
				var environments = interview.environment.split(",");
				interview.environment = new Array();
				environments.forEach( function (i) {
					interview.environment.push({ value : i });
				});
			}

			// Prepare Certifications Data
			if (interview.certs != null) {
				var certifications = interview.certs.split(",");
				interview.certs = new Array();
				certifications.forEach( function (i) {
					interview.certs.push({ value : i });
				});
			}	

			// Prepare Soft Skills Data
			if (interview.soft_skills) {
				var soft_skill = interview.soft_skills.split(",");
				interview.soft_skills = new Array();
				soft_skill.forEach( function (i) {
					interview.soft_skills.push({ value : i });
				});
			}

			// Prepare Education Data
			if (interview.education != null) {
				var degrees = interview.education.split(";");
				interview.education = new Array();
				degrees.forEach(function (i) {
					var arr = i.split(",");
					
					var degree = {
						degree: 	arr[0],
						major: 		arr[1],
						university: arr[2],
						year: 		arr[3]
					};

					interview.education.push(degree);
				});
			}

			console.log(interview);
			
			res.render('interviews/show', {
				interview : interview,
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

		var iq = "select * from C_INTERVIEW where id = " + conn.escape(req.params.id) + ";";

		var fq = "select * from L_TYPE;"
			   + "select * from L_SECTOR;"
			   + "select * from L_INDUSTRY;"
			   + "select * from L_EXPERIENCE;"
			   + "select * from L_ENVIRONMENT;"
			   + "select * from L_SALARY;"
			   + "select * from L_HOURS_PER_WEEK;"
			   + "select * from L_WEEKENDS_WORKED;"
			   + "select * from L_OVERNIGHT_TRAVEL;"
			   + "select * from L_CERTIFICATION;"
			   + "select * from L_SOFT_SKILL;"
			   + "select * from L_DEGREE;"
			   + "select * from L_MAJOR;"
			   + "select * from L_UNIVERSITY;"

		conn.query(fq + iq, function (err, results) {
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
				degrees 			: results[12],
				majors 				: results[13],
				universities 		: results[14],
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

		conn.query("insert into C_INTERVIEW set ?", interview, function (err) {
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

		conn.query("update C_INTERVIEW set ? where id = " + req.params.id, article, function (err) {
			if (err) console.log(err);
			res.redirect( "/i/" + req.params.id );
		});
	});
};

exports.delete = function (req, res) {

};

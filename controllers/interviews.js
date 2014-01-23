
/**
 * Module Dependencies
 */

var pool = require('../database').pool;
var interview = require('../models/interview');


/**
 * List interviews
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.index = function (req, res) {
    interview.list(null, function (err, interviews) {
        if (err) console.log(err);

        res.render('interviews/index', {
            title      : "Interviews",
            interviews : interviews,
            entity     : req.session.entity
        });
    });
};


/**
 * Retrieve interview
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.show = function (req, res) {
    interview.retrieve(req.params.id, function (err, results) {
        if (err) console.log(err);

        var interview = results[0][0];
        var comments  = results[1];

        comments.forEach(function (i) {
            i.email = i.email.split("@")[0];
        });

        interview.industry    = splitAndPush(interview.industry, ",");
        interview.environment = splitAndPush(interview.environment, ",");
        interview.certs       = splitAndPush(interview.certs, ",");
        interview.soft_skills = splitAndPush(interview.soft_skills, ",");
        interview.tools       = splitAndPush(interview.tools, ",");
        interview.skills      = splitAndPush(interview.skills, ",");

        if (interview.solo_group) {
            var obj = {
                solo : interview.solo_group,
                group : (100 - interview.solo_group)
            };
            interview.solo_group = new Object();
            interview.solo_group = obj;
            delete obj;
        }

        if (interview.work && interview.life) {
            var obj = {
                work : interview.work,
                life : interview.life
            };
            interview.fulfillment = new Object();
            interview.fulfillment = obj;
            delete obj;
        }

        // Prepare Education Data
        if (interview.education != null) {
            var degrees = interview.education.split(";");
            interview.education = new Array();
            degrees.forEach(function (i) {
                var arr = i.split(",");

                var degree = {
                    degree     : arr[0],
                    major      : arr[1],
                    university : arr[2],
                    year       : arr[3]
                };

                interview.education.push(degree);
            });
        }

        // Prepare Daily Tasks Data
        // { label: "Programming", data: 5 },
        if (interview.tasks != null) {
            var activity = interview.tasks.split(";");
            interview.tasks = new String();
            activity.forEach(function (element, index, array) {
                var arr = element.split(",");
                interview.tasks += "{ label: '" + arr[0] + "', data: " + arr[1] + " }";
                if (index < array.length - 1) interview.tasks += ", ";
            });
        }

        console.log(interview);

        console.log(interview.tasks);

        res.render('interviews/show', {
            interview: interview,
            comments: comments,
            entity: req.session.entity
        });
    });
};


/**
 * Render interview form
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.form = function (req, res) {

    var arr = [];
    //arr[arr.length] = "SELECT * FROM C_INTERVIEW WHERE id = " + pool.escape(req.params.id) + ";";
    arr[arr.length] = "SELECT * FROM L_TYPE;";
    arr[arr.length] = "SELECT * FROM L_SECTOR;";
    arr[arr.length] = "SELECT * FROM L_INDUSTRY;";
    arr[arr.length] = "SELECT * FROM L_EXPERIENCE;";
    arr[arr.length] = "SELECT * FROM L_ENVIRONMENT;";
    arr[arr.length] = "SELECT * FROM L_SALARY;";
    arr[arr.length] = "SELECT * FROM L_HOURS_PER_WEEK;";
    arr[arr.length] = "SELECT * FROM L_WEEKENDS_WORKED;";
    arr[arr.length] = "SELECT * FROM L_OVERNIGHT_TRAVEL;";
    arr[arr.length] = "SELECT * FROM L_CERTIFICATION;";
    arr[arr.length] = "SELECT * FROM L_SOFT_SKILL;";
    arr[arr.length] = "SELECT * FROM L_DEGREE;";
    arr[arr.length] = "SELECT * FROM L_MAJOR;";
    arr[arr.length] = "SELECT * FROM L_UNIVERSITY;";

    // Construct query
    var query = "";
    arr.forEach(function (str) {
        query += str;
    });

    pool.getConnection(function (err, conn) {
        conn.query(query, function (err, results) {
            res.render('interviews/form', {
                types            : results[0],
                sectors          : results[1],
                industries       : results[2],
                experience       : results[3],
                environments     : results[4],
                salary           : results[5],
                hours_per_week   : results[6],
                weekends_worked  : results[7],
                overnight_travel : results[8],
                certs            : results[9],
                soft_skills      : results[10],
                degrees          : results[11],
                majors           : results[12],
                universities     : results[13],
                entity           : req.session.entity
            });
        });
    });
};


/**
 * Edit interview
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.edit = function (req, res) {

    var arr = [];
    arr[arr.length] = "SELECT * FROM C_INTERVIEW WHERE id = " + pool.escape(req.params.id) + ";";
    arr[arr.length] = "SELECT * FROM L_TYPE;";
    arr[arr.length] = "SELECT * FROM L_SECTOR;";
    arr[arr.length] = "SELECT * FROM L_INDUSTRY;";
    arr[arr.length] = "SELECT * FROM L_EXPERIENCE;";
    arr[arr.length] = "SELECT * FROM L_ENVIRONMENT;";
    arr[arr.length] = "SELECT * FROM L_SALARY;";
    arr[arr.length] = "SELECT * FROM L_HOURS_PER_WEEK;";
    arr[arr.length] = "SELECT * FROM L_WEEKENDS_WORKED;";
    arr[arr.length] = "SELECT * FROM L_OVERNIGHT_TRAVEL;";
    arr[arr.length] = "SELECT * FROM L_CERTIFICATION;";
    arr[arr.length] = "SELECT * FROM L_SOFT_SKILL;";
    arr[arr.length] = "SELECT * FROM L_DEGREE;";
    arr[arr.length] = "SELECT * FROM L_MAJOR;";
    arr[arr.length] = "SELECT * FROM L_UNIVERSITY;";

    // Construct query
    var query = "";
    arr.forEach(function (str) {
        query += str;
    });

    pool.getConnection(function (err, conn) {
        conn.query(query, function (err, results) {
            res.render('interviews/form', {
                interview        : results[0][0],
                types            : results[1],
                sectors          : results[2],
                industries       : results[3],
                experience       : results[4],
                environments     : results[5],
                salary           : results[6],
                hours_per_week   : results[7],
                weekends_worked  : results[8],
                overnight_travel : results[9],
                certs            : results[10],
                soft_skills      : results[11],
                degrees          : results[12],
                majors           : results[13],
                universities     : results[14],
                entity           : req.session.entity
            });
        });
    });
};


/**
 * Create interview
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.create = function (req, res) {

    var fulfillment = {
        life : req.body.life,
        work : req.body.work
    };

    // Construct educational background
    var education = {
      degree        : req.body.degree,
      year          : req.body.year,
      concentration : req.body.concentration,
      university    : req.body.university 
    };

    // Construct activities
    var activities = [];
    req.body.tasks.forEach(function (e, i) {
        activities.push({
            task : req.body.tasks[i],
            time : req.body.hours[i]
        });
    });

    var payload = {
        name      : req.body.name,
        position  : req.body.position,
        uid       : req.session.entity.uid,
        published : 1,
        timestamp : new Date(),

        // Company?
        // 
        
        type             : req.body.type,
        sector           : req.body.sector,
        industry         : req.body.industry,
        experience       : req.body.experience,
        environment      : req.body.environment,
        salary           : req.body.salary,
        hours_per_week   : req.body.hours_per_week,
        weekends_worked  : req.body.weekends_worked,
        overnight_travel : req.body.overnight_travel,
        certs            : req.body.certs,
        soft_skills      : req.body.soft_skills,

        // Educational Background
        education        : education,
        
        // Others
        skills           : req.body.skills,
        tools            : req.body.tools,

        // Graphs
        activities       : activities,
        solo_group       : req.body.solo_group,
        fulfillment      : fulfillment
    };

    console.log(payload);

    interview.create(payload, function (err, results) {
        if (err) console.log(err);
        res.redirect('/i');
    });
};


/**
 * Update interview
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.update = function (req, res) {

    var article = {
        name             : req.body.name,
        position         : req.body.position,
        type             : req.body.type,
        sector           : req.body.sector,
        industry         : req.body.industry,
        experience       : req.body.experience,
        environment      : req.body.environment,
        salary           : req.body.salary,
        hours_per_week   : req.body.hours_per_week,
        weekends_worked  : req.body.weekends_worked,
        overnight_travel : req.body.overnight_travel,
        certs            : req.body.certs,
        soft_skills      : req.body.soft_skills,
        timestamp        : new Date()
    };

    pool.getConnection(function (err, conn) {
        conn.query("update C_INTERVIEW set ? where id = " + req.params.id, article, function (err) {
            if (err) console.log(err);
            res.redirect("/i/" + req.params.id);
        });
    });
};


/**
 * Delete interview
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.delete = function (req, res) {

};


/**
 * Helper Functions
 * 
 * @param {Object} value
 * @param {Object} separator
 * @api private
 */

function splitAndPush (value, separator) {
    if (value != null) {
        var tmp = value.split(separator);
        value = new Array();
        tmp.forEach(function (i) {
            value.push({ value : i });
        });
    }
    return value;
}


/**
 * Module Dependencies
 */

var pool = require('../database').pool;

/**
 * List Interviews
 *
 * @param {Object} options
 * @param {Function} cb
 * @api private
 */

exports.list = function (options, cb) {

  var query = "select * "
            + "from C_INTERVIEW as I "
            + "left join C_USER as U ON I.uid = U.uid "
            + "where I.published = 1";

  pool.getConnection(function (err, conn) {
    conn.query(query, function (err, interviews) {
      return cb(err, interviews);
    });
    conn.release();
  });
}

/**
 * Create Interview
 *
 * @param {Object} data
 * @param {Function} cb
 * @api private
 */

exports.create = function (data, cb) {

  sql = 'INSERT INTO C_INTERVIEW (name, position, uid, created, updated, published) VALUES ('+pool.escape(data.name)+','+pool.escape(data.position)+','+pool.escape(data.uid)+',"'+new Date()+'","null","1"); SELECT LAST_INSERT_ID();';

  pool.getConnection(function (err, conn) {
    conn.query(sql, function (err, results, fields) {
      var iid = results[1][0]['LAST_INSERT_ID()'];

      var arr = [];

      // Construct queries
      if (data.type) arr[arr.length] = 'INSERT INTO F_TYPE (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.type)+');';
      if (data.sector) arr[arr.length] = 'INSERT INTO F_SECTOR (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.sector)+');';
      if (data.industry) {
        data.industry.forEach(function (i) {
          arr[arr.length] = 'INSERT INTO F_INDUSTRY (iid, vid) VALUES ("'+iid+'",'+pool.escape(i)+');';
        });
      }
      if (data.experience) arr[arr.length] = 'INSERT INTO F_EXPERIENCE (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.experience)+');';
      if (data.environment) {
        data.environment.forEach(function (i) {
          arr[arr.length] = 'INSERT INTO F_ENVIRONMENT (iid, vid) VALUES ("'+iid+'",'+pool.escape(i)+');';
        });
      }
      if (data.salary) arr[arr.length] = 'INSERT INTO F_SALARY (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.salary)+');';
      if (data.hours_per_week) arr[arr.length] = 'INSERT INTO F_HOURS_PER_WEEK (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.hours_per_week)+');';
      if (data.weekends_worked) arr[arr.length] = 'INSERT INTO F_WEEKENDS_WORKED (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.weekends_worked)+');';
      if (data.overnight_travel) arr[arr.length] = 'INSERT INTO F_OVERNIGHT_TRAVEL (iid, vid) VALUES ("'+iid+'",'+pool.escape(data.overnight_travel)+');';
      if (data.certs) {
        data.certs.forEach(function (i) {
          arr[arr.length] = 'INSERT INTO F_CERTIFICATION (iid, vid) VALUES ("'+iid+'",'+pool.escape(i)+');';
        });
      }
      if (data.soft_skills) {
        data.soft_skills.forEach(function (i) {
          arr[arr.length] = 'INSERT INTO F_SOFT_SKILL (iid, vid) VALUES ("'+iid+'",'+pool.escape(i)+');';
        });
      }
      if (data.skills) {
        data.skills.forEach(function (i) {
          arr[arr.length] = 'INSERT INTO F_SKILL (iid, value) VALUES ("'+iid+'",'+pool.escape(i)+');';
        });
      }
      if (data.tools) {
        data.tools.forEach(function (i) {
          arr[arr.length] = 'INSERT INTO F_TOOL (iid, value) VALUES ("'+iid+'",'+pool.escape(i)+');';
        });
      }

      arr.forEach(function (query) {
        conn.query(query, function (err) {
          if (err) console.log(err); 
        });
      });
    });

    conn.release();
    return cb(err);

  });
}

/**
 * Retrieve Interview
 *
 * @param {Object} id
 * @param {Function} cb
 * @api private
 */

exports.retrieve = function (id, cb) {
  var query = "select C_INTERVIEW.*, "
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
            + "education.value as education, "
            + "tools.value as tools, "
            + "skills.value as skills "
            
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

            // Tools
            + "left join ( "
              + "select F_TOOL.iid, group_concat(F_TOOL.value separator ',') as value "
              + "from F_TOOL "
              + "group by F_TOOL.iid "
            + ") as tools on tools.iid = C_INTERVIEW.id "

            // Skills
            + "left join ( "
              + "select F_SKILL.iid, group_concat(F_SKILL.value separator ',') as value "
              + "from F_SKILL "
              + "group by F_SKILL.iid "
            + ") as skills on skills.iid = C_INTERVIEW.id "

            + "where C_INTERVIEW.id = " + pool.escape(id) + ";";

  var comment_query = "select F_COMMENT.id, C_USER.email, F_COMMENT.value "
                    + "from F_COMMENT "
                      + "left join C_USER on C_USER.uid = F_COMMENT.uid "
                    + "where F_COMMENT.iid = " + pool.escape(id) + " "
                      + "and F_COMMENT.published = 1;";

  pool.getConnection(function (err, conn) {
    conn.query(query + comment_query, function (err, results) {
      return cb(err, results);
    });
    conn.release();
  });
}
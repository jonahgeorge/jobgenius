var pool = require('../database').pool;

exports.index = function (req, res) {
    pool.getConnection(function (err, conn) {

        conn.query("select * from C_ARTICLE as A left join C_USER as U ON A.uid = U.uid where A.published = 1", function (err, articles) {
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
        conn.query("select * from C_ARTICLE where aid = ?", req.params.aid, function (err, article) {
            
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
        conn.query("select * from C_ARTICLE where aid = ?", req.params.aid, function (err, article) {
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

        conn.query("insert into C_ARTICLE set ?", article, function (err) {
            if (err) console.log(err);
            res.redirect('/a');
        });
        
        conn.release();
    });
};

exports.update = function (req, res) {
  pool.getConnection(function (err, conn) {

    var sql = "update C_ARTICLE set "
            + "title = "+pool.escape(req.body.title)+", "
            + "subtitle = "+pool.escape(req.body.subtitle)+", "
            + "body = "+pool.escape(req.body.body)+", "
            + "timestamp = "+pool.escape(new Date())+" "
            + "where aid = " + pool.escape(req.params.aid) + ";"

    console.log(sql);

    conn.query(sql, function (err) {
      if (err) console.log(err);
      res.redirect( "/a/" + req.params.aid );
    });
  });
};

exports.delete = function (req, res) {
  pool.getConnection(function (err, conn) {
    
    conn.query("select * from C_ARTICLE where aid = ?", req.params.aid, function (err, article) {
      
      if (article.uid == req.session.entity.uid || req.session.entity.role == 0) {

        conn.query("update C_ARTICLE set published = 0 where aid = ? ", req.params.aid, function (err) {
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

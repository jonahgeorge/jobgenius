/* Generic require login routing middleware */
exports.requiresLogin = function (req, res, next) {
	if (!req.session.entity) {
		return res.redirect('/auth/signin')
	}
	next();
}

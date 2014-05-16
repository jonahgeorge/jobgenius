	

$.ajax({
	url: "/api/charts/groupwork?i="+$("body").data("interview-id"),
	type: "GET",
	dataType: "json"
}).done(function( msg ) {

	var arr = [];

	var obj = [];
	obj.push("Label");
	obj.push("Value");
	arr.push(obj);

	var obj = [];
	obj.push("Solo");
	obj.push(msg.Solo);
	arr.push(obj);

	var obj = [];
	obj.push("Group");
	obj.push(msg.Group);
	arr.push(obj);

	google.setOnLoadCallback( drawChart(arr) );
});
	 
function drawChart(arr) {
	var data = google.visualization.arrayToDataTable(arr);

	var options = {
		width: '100%',
		pieHole: 0.4
	};

	// Instantiate and draw our chart, passing in some options.
	var chart = new google.visualization.PieChart(document.getElementById('solovgroup'));
	chart.draw(data, options);

	function resizeHandler () {
		chart.draw(data, options);
	}

	if (window.addEventListener) window.addEventListener('resize', resizeHandler, false);
	else if (window.attachEvent) window.attachEvent('onresize', resizeHandler);
}

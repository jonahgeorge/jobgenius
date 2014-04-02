
$.ajax({
	url: "/api/charts/breakdown?i="+$("body").data("interview-id"),
	type: "GET",
	dataType: "json"
}).done(function( msg ) {

	arr  = [];

	// Format data
	msg.forEach(function(i) {
		var obj = [];
		obj.push(i.Title.String);
		obj.push(i.Value.Int64);
		arr.push(obj);
	})

	google.setOnLoadCallback( drawPieChart(arr) );
});

function drawPieChart(arr) {

	var data = new google.visualization.DataTable();
	data.addColumn('string', 'Activity');
	data.addColumn('number', 'Hours');
	data.addRows(arr);

	// Set chart options
	var options = { 'width':'100%' };

	// Instantiate and draw our chart, passing in some options.
	var chart = new google.visualization.PieChart(document.getElementById('breakdown'));
	chart.draw(data, options);

	function resizeHandler () {
		chart.draw(data, options);
	}

	if (window.addEventListener) window.addEventListener('resize', resizeHandler, false);
	else if (window.attachEvent) window.attachEvent('onresize', resizeHandler);
}
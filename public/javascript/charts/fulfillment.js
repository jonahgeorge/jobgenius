
$.ajax({
	url: "/api/charts/fulfillment?i="+$("body").data("interview-id"),
	type: "GET",
	dataType: "json"
}).done(function( msg ) {

	arr  = [];

	var obj = [];
	obj.push("Label");
	obj.push("Edward Gonzalez's value");
	obj.push("Industry average");
	arr.push(obj);

	var obj = [];
	obj.push("Development");
	obj.push(msg.Development.Int64);
	obj.push(5); // Industry Average
	arr.push(obj);

	var obj = [];
	obj.push("Independence");
	obj.push(msg.Independence.Int64);
	obj.push(5); // Industry Average
	arr.push(obj);

	var obj = [];
	obj.push("Impact");
	obj.push(msg.Impact.Int64);
	obj.push(5); // Industry Average
	arr.push(obj);

	var obj = [];
	obj.push("Personal Life");
	obj.push(msg.Personal.Int64);
	obj.push(5); // Industry Average
	arr.push(obj);

	google.setOnLoadCallback( drawBarChart(arr) );
});
	
function drawBarChart(arr) {
	var data = google.visualization.arrayToDataTable(arr);

	var options = {
		legend: 'top',
		vAxis : {
			minValue : 0, 
			maxValue : 10 
		},
		isStacked: false
	};

	// Instantiate and draw our chart, passing in some options.
	var chart = new google.visualization.ColumnChart(document.getElementById('fulfillment'));
	chart.draw(data, options);

	function resizeHandler () {
		chart.draw(data, options);
	}

	if (window.addEventListener) window.addEventListener('resize', resizeHandler, false);
	else if (window.attachEvent) window.attachEvent('onresize', resizeHandler);
}
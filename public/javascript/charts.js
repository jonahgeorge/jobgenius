
// Load the Visualization API and the piechart package.
google.load('visualization', '1.0', {'packages':['corechart', 'gauge']});

// Set a callback to run when the Google Visualization API is loaded.
google.setOnLoadCallback(drawPieChart);
google.setOnLoadCallback(drawBarChart);
google.setOnLoadCallback(drawChart);
	
/*
 *
 */
	 
function drawChart() {
	var data = google.visualization.arrayToDataTable([
		['Label', 'Value'],
		['Solo',  60],
		['Group', 40]
	]);

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

/*
 *
 */
	
function drawBarChart() {
	var data = google.visualization.arrayToDataTable([
		['Label', "Edward Gonzalez's value", 'Industry Average'],
		['Development', 8, 5],
		['Independence', 9, 5],
		['Impact', 10, 5],
		['Personal Life', 4, 5]
	]);

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

/*
 *
 */

function drawPieChart() {
	var data = new google.visualization.DataTable();
	data.addColumn('string', 'Activity');
	data.addColumn('number', 'Hours');
	data.addRows([
		["Interviews",1],
		["Email",2],
		["Phone Calls",3],
		["Working with Developers",1],
		["Research",3]
	]);

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

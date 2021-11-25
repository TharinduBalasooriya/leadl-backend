package service

import (

	//Importing file storage utility

	//"strings"

	"fmt"
	"log"
	"strconv"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
)

func GetReportTemplate(report datamodels.Report) string {

	var templateStart = `<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Report 2</title>
  <link rel="preconnect" href="//fonts.gstatic.com" />
  <link
    href="https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,200;0,300;0,400;0,500;0,700;0,800;1,300&family=Martel:wght@900&display=swap"
    rel="stylesheet"
  />
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
  <link
    rel="stylesheet"
    type="text/css"
    href="//fonts.googleapis.com/icon?family=Material+Icons"
  />
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <!DOCTYPE html>
  <!-- Load d3.js -->
  <script src="https://d3js.org/d3.v4.js"></script>
</head>`

	var templateStyleStart = `<style>
  *,
  body {
    margin: 0;
  }
  html {
    font-family: 'Poppins', sans-serif;
    margin: 0;
  }
  .report-page {
    padding: 48px;
  }
  .report-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .head-txt h1 {
    font-size: 1.5rem;
    text-transform: uppercase;
    font-weight: 500;
    padding-top: 40px;
    color: #000000;
  }
  .logo-qr {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
  .logo-qr img.company-logo {
    width: 45%;
  }
  .logo-qr img.company-qr {
    width: 15%;
    filter: brightness(0);
  }
  .logo-qr span {
    font-style: italic;
    color: #000000;
  }
  .right-align {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
  }
  .sub-header {
    margin: 20px 0;
  }
  .hr {
    height: 7px;
    background: black;
  }
  .mrbtm {
    margin-bottom: 5px;
    text-align: left;
  }
  .sub-header p.description {
    font-size: 1rem;
    color: #000000;
    padding: 0 0;
  }
  .table-container {
    display: flex;
    margin: 0 0 36px;
    width: 80%;
  }
  .vertical-col {
    padding: 16px 0;
  }
  .tc-vertical {
    text-align: right;
    font-size: 15px;
  }
  .tableMainHeader {
    display: flex;
    align-items: center;
    padding: 6px 0;
    justify-content: space-between;
  }
  .tableMainHeader p {
    display: flex;
    align-items: center;
    font-size: 15px;
    color: #000000;
  }
  table.reportTable {
    border: 1px solid #ffffff;
    width: 100%;
    height: auto;
    text-align: center;
    border-collapse: collapse;
  }
  table.reportTable th {
    border-top: 2px solid #000000;
    border-bottom: 1px solid #000000;
    padding: 16px;
  }
  table.reportTable td {
    border: none;
    padding: 16px 0px 16px 16px;
  }
  table.reportTable tbody td {
    font-size: 18px;
    color: #000000;
  }
  table.reportTable tr td {
    background: #ffffff;
  }
  table.reportTable thead {
    background: #ebebeb;
  }
  table.reportTable  th {
    font-size: 20px;
    font-weight: bold;
    color: #000000;
    text-align: center;
  }
  table.reportTable tfoot td {
    font-size: 14px;
  }
  .footer {
    display: flex;
    align-items: center;
    padding: 24px 0;
    justify-content: space-between;
  }
  .footer p {
    display: flex;
    align-items: center;
    font-size: 15px;
    color: #000000;
  }
  .subtablehr {
    border-bottom: 3px solid #ebebeb;
  }
  .page-no-txt {
    background-color: #ffffff;
    color: rgb(0, 0, 0);
    padding: 8px 16px;
    text-align: right;
    font-size: 18px;
  }
  table.reportsubTable {
    border: 3px solid #ebebeb;
    width: 42%;
    height: auto;
    text-align: center;
    border-collapse: collapse;
    margin-left: auto;
    margin-right: 0px;
  }
  table.reportsubTable2 {
    border: 3px solid #ebebeb;
    width: 100%;
    height: auto;
    text-align: center;
    border-collapse: collapse;
    margin-left: auto;
    margin-right: 0px;
  }
  table.reportsubTable tbody tr th {
    border: none;
    font-size: 12px;
    font-weight: 350;
  }
  table.reportsubTable2 tbody tr th {
    border: none;
    font-size: 16px;
    font-weight: 350;
    padding: 5px;
  }
  table.reportsubTable2 thead th {
    background: #ebebeb;
    border: none;
    font-size: 18px;
    color: #000000;
    text-align: center;
    padding: 5px;
  }
  table.reportsubTable thead th {
    background: #ebebeb;
    border: none;
    font-size: 17px;
    color: #000000;
    text-align: center;
  }
  .axis .domain {
    display: none;
  }`

	var templateStyleEnd = `</Style>`

	var templateBodyStart = `  <body>
    <div class="report-page">
      <div class="report-header">
        <div class="head-txt"><h1>{subject}</h1></div>
        <div class="logo-qr">
          <div class="right-align">
            <img
              class="company-logo"
              src="https://tracified-report-images.s3.ap-south-1.amazonaws.com/tracified.png"
              alt="tracified"
              border="0"
            />
            <span>Trust through Traceability</span>
          </div>
          <img
            class="company-qr"
            src="https://tracified-report-images.s3.ap-south-1.amazonaws.com/qr.png"
            alt="frame-Copy-removebg-preview"
            border="0"
          />
        </div>
      </div>
      <hr style="height: 2px; background-color: #000000" />
      <div style="padding: 5px">
        <div class="sub-header"><p class="description">{subjectdesc}</p></div>
        <div class="tableMainHeader">
          <p class="left-text">
            <span
              ><b>Report Ordered by</b><br />
              {tenantName}</span>
          </p>
          <p class="right-text">
            <span><b>Report ID</b><br />{reportID}</span>
          </p>
        </div>
        <div class="tableMainHeader">
          <p class="left-text">
            <span
              ><b>Farm location</b><br />
              {FarmLocation}</span>
          </p>
          <p class="right-text">
            <span align="right"><b>Period</b><br />{dateRange}</span>
          </p>
        </div>
      </div>
      <div class="chartContainer">
      `

	var templateBodyContent = ``

	var templateBodyEnd = `
	<div class="footer">
	  <p class="left-text">
		<span>Date and Time of issue: {createdAt}</span>
	  </p>
	  <p class="right-text"><span>Page 01</span></p>
	</div>
  </div> 
</body>
</html>`

	/**
	 * Container Style
	 **/
	templateStyleStart = templateStyleStart + getChartContainerStyle(report.Widgets)
	/*
	 * Add Widget Styles
	 **/
	tableCount := 0
	groupedBarChartCount := 0
	bubbleChartCount := 0
	for _, widget := range report.Widgets {

		if widget.Type == "table" {
			tableCount++
			templateStyleStart = templateStyleStart + getTableStyle(strconv.Itoa(tableCount), widget)
			templateBodyContent = templateBodyContent + getTableTemplate(strconv.Itoa(tableCount))
		}

		if widget.Type == "gbar" {
			log.Print("Grouped barchart")
			groupedBarChartCount++
			templateStyleStart = templateStyleStart + getGroupedBarChartStyle(strconv.Itoa(groupedBarChartCount), widget)
			templateBodyContent = templateBodyContent + getGroupedBarChartTemplate(strconv.Itoa(groupedBarChartCount))
		}

		if widget.Type == "bubble" {
			log.Print("Bubble Chart")
			bubbleChartCount++
			templateStyleStart = templateStyleStart + getBubbleChartStyle(strconv.Itoa(bubbleChartCount), widget)
			templateBodyContent = templateBodyContent + getBubbbleChartTemplate(strconv.Itoa(bubbleChartCount))

		}

	}

	template := templateStart + templateStyleStart + templateStyleEnd + templateBodyStart + templateBodyContent + templateBodyEnd
	//template = strings.ReplaceAll(template, "\n", "")
	//template = strings.ReplaceAll(template,`\"`,"&quot;")
	return template
}

func getChartContainerStyle(widgets []datamodels.Widget) string {

	yMax, err := strconv.ParseFloat(widgets[0].Y, 64)
	if err != nil {
		log.Println("Float parse failed")
	}
	for _, widget := range widgets {

		y, err := strconv.ParseFloat(widget.Y, 64)
		if err != nil {
			log.Println("Float parse failed")
		}
		if y > yMax {
			yMax = y
		}

	}

	s := fmt.Sprintf("%f", yMax+200)
	fmt.Println(s)

	style := `.chartContainer{
  position: relative;
  height:` + s + `px;
  }`

	return style

}
func getTableStyle(tableNumber string, widget datamodels.Widget) string {

	style := `
  .table` + tableNumber + `{
    position: absolute;
    top:` + widget.Y + `px;
    left:` + widget.X + `px;` + `}
    `

	return style

}

func getGroupedBarChartStyle(chartNumber string, widget datamodels.Widget) string {

	style := `
  .gbar` + chartNumber + `{
    position: absolute;
    top:` + widget.Y + `px;
    left:` + widget.X + `px;` + `}
    `

	return style

}

func getBubbleChartStyle(chartNumber string, widget datamodels.Widget) string {

	style := `
  .bubble` + chartNumber + `{
    position: absolute;
    top:` + widget.Y + `px;
    left:` + widget.X + `px;` + `}
    `

	return style

}

func getTableTemplate(tableNumber string) string {
	var tableTemplate string
	tableTemplate = `
  <!--Table start-->
  <div class="table-container` + ` table` + tableNumber + `">
	<table id="reportTable" class="table table-hover table-striped table-bordered tableReport">
  <!--Data section-->
  <thead>
  <tr>
    <th>Heading1</th>
    <th>Heading2</th>
    <th>Heading3</th>
  </tr>
  </thead>
  <tbody >
  <tr>
    <td>Data1</td>
    <td>Data2</td>
    <td>Data3</td>
  </tr>
  <tr>
  <td>Data1</td>
  <td>Data2</td>
  <td>Data3</td>
  </tr>
  <tr>
  <td>Data1</td>
  <td>Data2</td>
  <td>Data3</td>
  </tr>
  <tr>
  <td>Data1</td>
  <td>Data2</td>
  <td>Data3</td>
  </tr>
  <tr>
  <td>Data1</td>
  <td>Data2</td>
  <td>Data3</td>
  </tr>
  </tbody>

	</table>
  </div>
  <!--Table End -->
	
	<script> 
	var objC = JSON.parse('{query}');
	buildTable(objC.mainTable.mainTableRows);
	 function buildTable(data) {
	  //change table id
	   var table = document.getElementById('myTable')
	   for (var i = 0; i < data.length; i++) {
		 var row = ` + `<tr><td style="">${data[i].Product}</td> <td>${data[i].Location}</td> <td>${data[i].Accepted_QTY_kg}</td> <td>${data[i].Received_QTY_kg}</td> <td>${data[i].Percentage}</td></tr>` + `table.innerHTML += row
	   }
	 }
   </script>`
	return tableTemplate
}
func getBubbbleChartTemplate(chartNumber string) string {
	template := `<div id="bubble` + chartNumber + `"></div>
  <script>
      var reportId = 2

        var chartData = [
          {
            date: '2021-01-16T13:45:00.010Z',
            data: [{ fertilizer: 'Dert', fertilizerQTY: '7', harvestedQTY: '500' }],
          },
          {
            date: '2021-03-16T13:45:00.010Z',
            data: [{ fertilizer: 'LakPE', fertilizerQTY: '9', harvestedQTY: '800' }],
          },
          {
            date: '2021-07-16T13:45:00.010Z',
            data: [{ fertilizer: 'MK23', fertilizerQTY: '10', harvestedQTY: '800' }],
          },
          {
            date: '2021-09-16T13:45:00.010Z',
            data: [
              { fertilizer: 'Compo', fertilizerQTY: '5', harvestedQTY: '200' },
            ],
          },
        ]

        var fertilizers = [{name: 'Fertilizersstop'},{name:'EcoScraps'},{name:'Soil Meal Inc'},{name:'CNH Industrial NV'}]
      
      //Sample Data
        var newData = [{'date':'01/01/2016',  fertilizer: 'Fertilizersstop', fertilizerQTY: '6', harvestedQTY: '250'},
                {'date':'02/01/2016', fertilizer: 'EcoScraps', fertilizerQTY: '7', harvestedQTY: '800'},
                {'date':'02/01/2016', fertilizer: 'Fertilizersstop', fertilizerQTY: '9', harvestedQTY: '300'},
                {'date':'04/01/2016', fertilizer: 'EcoScraps', fertilizerQTY: '8', harvestedQTY: '400'},
                {'date':'05/01/2016', fertilizer: 'Fertilizersstop', fertilizerQTY: '10', harvestedQTY: '600'},
				        {'date':'06/01/2016', fertilizer: 'EcoScraps', fertilizerQTY: '2', harvestedQTY: '500' },
                {'date':'06/01/2016', fertilizer: 'Fertilizersstop', fertilizerQTY: '6', harvestedQTY: '700'},
                {'date':'08/01/2016', fertilizer: 'Soil Meal Inc', fertilizerQTY: '7', harvestedQTY: '400'},
                {'date':'09/01/2016', fertilizer: 'CNH Industrial NV', fertilizerQTY: '3', harvestedQTY: '600'},]


                                    

                          function genReport3(data,fnames) {


                            var fertilizers = d3.map(fnames, function(d){return(d.name)}).keys()
                            console.log("fertizers",fertilizers)
                                    // set the dimensions and margins of the graph
                                    var margin = {top: 30, right: 30, bottom: 40, left: 50},
                                width = 900 - margin.left - margin.right,
                                height = 420 - margin.top - margin.bottom;

                            // append the svg object to the body of the page
                            var svg = d3.select('#bubble` + chartNumber + `')
                              .append('svg')
                                .attr('width', width + margin.left + margin.right)
                                .attr('height', height + margin.top + margin.bottom)
                              .append('g')
                                .attr('transform',
                                      'translate(' + margin.left + ',' + margin.top + ')');

                                      

                            
                              console.log('Data', data);
                              var xFormat = '%d-%b-%Y';;
                              var parseTime = d3.timeParse('%d/%m/%Y');
                              // Add X axis
                              var x = d3.scaleTime()
                                .rangeRound([ 0, width ])
                                .domain(d3.extent(data, function(d) { return parseTime(d.date); }));
                              svg.append('g')
                                .attr('transform', 'translate(0,' + height + ')')
                                .call(d3.axisBottom(x).tickFormat(d3.timeFormat(xFormat)));

                                // text label for the x axis
                              svg.append('text')             
                                .attr('transform',
                                      'translate(' + (width/2) + ' ,' + 
                                                    (height + margin.top) + ')')
                                .style('font', '12px sans-serif')
                                .style('text-anchor', 'middle')
                                .text('date of application');

                              // Add Y axis
                              var y = d3.scaleLinear()
                                .domain([0,d3.max(data, function(d) { return +d.fertilizerQTY;} )])
                                .range([ height, 0]);
                              svg.append('g')
                                .call(d3.axisLeft(y));

                              // text label for the y axis
                              svg.append('text')
                                .attr('transform', 'rotate(-90)')
                                .attr('y', 0 - margin.left)
                                .attr('x',0 - (height / 2))
                                .attr('dy', '2em')
                                .style('font', '12px sans-serif')
                                .style('text-anchor', 'middle')
                                .text('quantity(kg)');  

                               // color palette = one color per subgroup
                              var color = d3.scaleOrdinal()
                                .domain(fertilizers)
                                .range(['#F89C30', '#51BEFC', '#88FC51','#e41a1c','#377eb8','#92c5de','#4daf4a'])

                              // Add a scale for bubble size
                              var z = d3.scaleLinear()
                                .domain([200, 1000])
                                .range([ 1, 40]);

                              var c = d3.scaleLinear()
                                .domain([200, 1000])
                                .range([ 1, 40]);

                              // Add dots
                              svg.append('g')
                                .selectAll('dot')
                                .data(data)
                                .enter()
                                .append('circle')
                                  .attr('cx', function(d) { return x(parseTime(d.date)); } )
                                  .attr('cy', function (d) {  return y(d.fertilizerQTY); } )
                                  .attr('r', function (d) {  return z(d.harvestedQTY); } )
                                  .style('fill',function(d) { return color(d.fertilizer); })
                                  .attr('stroke', 'black')

                                  var legend = svg
                                  .selectAll('.legend')
                                  .data(fertilizers)
                                  .enter()
                                  .append('g')
                                  .attr('class', 'legend')
                                  .attr('transform', function (d, i) {
                                    return 'translate(0,' + i * 20 + ')'
                                  })
                                  .style('opacity', '0')
                                legend
                                  .append('circle')
                                  .attr('cx', width)
                                  .attr('r', z(300) )
                                  .style('fill', function (d) {
                                    return color(d)
                                  })
                                  .attr('stroke', 'black')
                                  
                                legend
                                  .append('text')
                                  .attr('x', width - 15)
                                  .attr('y', 0)
                                  .attr('dy', '.35em')
                                  .style("font", "10px sans-serif")
                                  .style('text-anchor', 'end')
                                  .text(function (d) {
                                    return d
                                  })
                                legend
                                  .style('opacity', '1')
                       }
      genReport3(newData,fertilizers)
    </script>`
	return template
}
func getGroupedBarChartTemplate(chartNumber string) string {
	template := `
  <!--Grouped barchart-->
  <div align="center" style="margin-top: 50px" class="gbar` + chartNumber + `" id="gBar` + chartNumber + `"></div>
 <div id="chart_desc"></div>
 <script>
 var dataM = [
   {
     product: 'Sample 1',
     values: [
       { value: 50, handler: 'Handler 1' },
       { value: 4, handler: 'Handler 2' },
       { value: 12, handler: 'Handler 3' },
     ],
   },
   {
     product: 'Sample 2',
     values: [
       { value: 10, handler: 'Handler 1' },
       { value: 21, handler: 'Handler 2' },
       { value: 13, handler: 'Handler 3' },
     ],
   },
   
 ]
 var maintableRows = [
   {
     Product: 'Farmington Hills',
     Location: 'Texas',
     Accepted_QTY_kg: '450',
     Received_QTY_kg: '480',
     Percentage: '93.75%',
   },
   {
     Product: 'Farmington Hills',
     Location: 'Texas',
     Accepted_QTY_kg: '450',
     Received_QTY_kg: '480',
     Percentage: '93.75%',
   },
 ]


 var margin = { top: 50, right: 200, bottom: 65, left: 70 },
   width = 980- margin.left - margin.right,
   height = 500 - margin.top - margin.bottom
 var svg = d3
   .select('#gBar` + chartNumber + `')
   .append('svg')
   .attr('width', width + margin.left + margin.right)
   .attr('height', height + margin.top + margin.bottom)
   .append('g')
   .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')')
 function generate(data) {
   var subgroups = data[0].values.map(function (d) {
     return d.handler
   })
   var groups = d3
     .map(data, function (d) {
       return d.product
     })
     .keys()
   var x = d3.scaleBand().domain(groups).range([0, width]).padding([0.2])
   svg
     .append('g')
     .style('font', '14px sans-serif')
     .attr('transform', 'translate(0,' + height + ')')
     .call(d3.axisBottom(x).tickSize(2))
   svg
     .append('text')
     .attr(
       'transform',
       'translate(' + width / 2 + ' ,' + (height + margin.top) + ')'
     )
     .style('font', '20px times')
     .style('text-anchor', 'middle')
     .text('product')
   var y = d3.scaleLinear().domain([0, 100]).range([height, 0])
   svg.append('g').style('font', '14px sans-serif').call(d3.axisLeft(y))
   svg
     .append('text')
     .attr('transform', 'rotate(-90)')
     .attr('y', 0 - margin.left)
     .attr('x', 0 - height / 2)
     .attr('dy', '1em')
     .style('font', '20px times')
     .style('text-anchor', 'middle')
     .text('efficacy (%)')
   var xSubgroup = d3
     .scaleBand()
     .domain(subgroups)
     .range([0, x.bandwidth()])
     .padding([0.05])
   var color = d3
     .scaleOrdinal()
     .domain(subgroups)
     .range([
       '#F89C30',
       '#51BEFC',
       '#88FC51',
       '#92c5de',
       '#e41a1c',
       '#377eb8',
       '#4daf4a',
     ])
   svg
     .append('g')
     .selectAll('g')
     .data(data)
     .enter()
     .append('g')
     .attr('transform', function (d) {
       console.log('d val', d)
       return 'translate(' + x(d.product) + ',0)'
     })
     .selectAll('rect')
     .data(function (d) {
       return subgroups.map(function (key) {
         console.log(
           ' val',
           d.values.find((elem) => elem.handler == key)
         )
         return {
           key: key,
           value: d.values.find((elem) => elem.handler == key).value,
         }
       })
     })
     .enter()
     .append('rect')
     .attr('x', function (d) {
       return xSubgroup(d.key)
     })
     .attr('y', function (d) {
       return y(d.value)
     })
     .attr('width', xSubgroup.bandwidth())
     .attr('height', function (d) {
       return height - y(d.value)
     })
     .attr('fill', function (d) {
       return color(d.key)
     })
   var legend = svg
     .selectAll('.legend')
     .data(
       data[0].values.map(function (d) {
         return d.handler
       })
     )
     .enter()
     .append('g')
     .attr('class', 'legend')
     .attr('transform', function (d, i) {
       return 'translate(0,' + i * 20 + ')'
     })
     .style('opacity', '0')
   legend
     .append('rect')
     .attr('x', width - 13)
     .attr('width', 13)
     .attr('height', 13)
     .style('fill', function (d) {
       return color(d)
     })
   legend
     .append('text')
     .attr('x', width - 24)
     .attr('y', 7)
     .attr('dy', '.35em')
     .style('font', '15px sans-serif')
     .style('text-anchor', 'end')
     .text(function (d) {
       return d
     })
   legend.style('opacity', '1')
 }
 generate(dataM)
</script>
<!-- Grouped barchart End -->
 `
	return template
}

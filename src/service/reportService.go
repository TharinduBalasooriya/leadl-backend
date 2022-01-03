package service

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/types"
)

var (
	AuthHeaderName = ""
	documentStart  = ReadFromFile("util\\templates\\reportHeader.html")
	documentEnd    = `</html>`
 
	styleStart = "<style> \n"+ ReadFromFile("util\\templates\\reportStyles.css")

	styleEnd = `</style>`

	body = ReadFromFile("util\\templates\\reportBody.html")

	reportDetailJSONStart = `
	<!--Report detail script-->
	<script>
	const report =`

	reportDetailJSONEnd = `</script>
	<!--End Report Detail container-->`

	mainHandler = ReadFromFile("util\\templates\\reportScript.html")
)

/*
*	Generate complete  HTML report template (css,javascript)
**/
func GetReportTemplate(report datamodels.Report) string {

	reportScript := GetReportDetailJSON(report)
	reportJSON, err := json.MarshalIndent(reportScript, "", "\t")
	if err != nil {
		log.Println(err)
	}

	template := documentStart + styleStart + GetStyles(report) + styleEnd +GetHTTPostParams(report) +body + GetHTTPHeaderFunction(report) + reportDetailJSONStart + string(reportJSON) + reportDetailJSONEnd + mainHandler + documentEnd

	return template
}


/*
* Configure css style tags in report tempalte
*/
func GetStyles(report datamodels.Report) string {
	style := ``
	for index, widget := range report.Widgets {
		if widget.Type == "table" {
			style = style + `.widget` + strconv.Itoa(index) + `{
            position: absolute;
            display: block;
            left:` + widget.X + `px;
            top:` + widget.Y + `px;
			width:` + widget.Width + `px;
			height:` + widget.Height + `px;

            }`
		}

		if widget.Type == "gbar" {
			style = style + `.widget` + strconv.Itoa(index) + `{
            position: absolute;
            display: block;
            left:` + widget.X + `px;
            top:` + widget.Y + `px;
			width:` + widget.Width + `px;
			height:` + widget.Height + `px;
			
            }`
		}

		if widget.Type == "bar" {
			style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				width:` + widget.Width + `px;
			    height:` + widget.Height + `px;
                }`

		}

		if widget.Type == "bubble" {
			style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				width:` + widget.Width + `px;
		     	height:` + widget.Height + `px;
                }`

		}

		if widget.Type == "text" {

			if widget.Styles.BorderColor != "" {

				if widget.Styles.TextBackgroundColor != "" {
					style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				border:solid+1px` + widget.Styles.BorderColor + `;
				color:` + widget.Styles.Fill + `;
                background-color:` + widget.Styles.TextBackgroundColor + `;
				font-size:` + strconv.Itoa(int(widget.Styles.FontSize)) + `px;
				font-weight:` + widget.Styles.FontWeight + `;
                font-family:` + widget.Styles.FontFamily + `;
                font-style:` + widget.Styles.FontStyle + `;
                }`
				} else {
					style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				border:solid+1px` + widget.Styles.BorderColor + `;
				color:` + widget.Styles.Fill + `;
				font-size:` + strconv.Itoa(int(widget.Styles.FontSize)) + `px;
				font-weight:` + widget.Styles.FontWeight + `;
                font-family:` + widget.Styles.FontFamily + `;
                font-style:` + widget.Styles.FontStyle + `;
                }`
				}
			} else {
				if widget.Styles.TextBackgroundColor != "" {
					style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				border-color:` + widget.Styles.BorderColor + `;
                background-color:` + widget.Styles.TextBackgroundColor + `;
				color:` + widget.Styles.Fill + `;
				font-size:` + strconv.Itoa(int(widget.Styles.FontSize)) + `px;
				font-weight:` + widget.Styles.FontWeight + `;
                font-family:` + widget.Styles.FontFamily + `;
                font-style:` + widget.Styles.FontStyle + `;
                }`
				} else {
					style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				border-color:` + widget.Styles.BorderColor + `;
				color:` + widget.Styles.Fill + `;
				font-size:` + strconv.Itoa(int(widget.Styles.FontSize)) + `px;
				font-weight:` + widget.Styles.FontWeight + `;
                font-family:` + widget.Styles.FontFamily + `;
                font-style:` + widget.Styles.FontStyle + `;
                }`
				}
			}

		}
		if widget.Type == "image" {

			style = style + `.widget` + strconv.Itoa(index) + `{
                position: absolute;
                display: block;
                left:` + widget.X + `px;
                top:` + widget.Y + `px;
				width:` + widget.Width + `px;
		    	height:` + widget.Height + `px;
                }`

		}
	}

	return style

}

/*
* Generate JSON which contain report details
**/
func GetReportDetailJSON(report datamodels.Report) types.ReportScript {
	reportScript := &types.ReportScript{
		ReportName:  report.ReportName,
		URL:         report.Url,
		AccessToken: report.AccessToken,
		Widgets:     []types.Widget{},
		FilterOptions :types.FilterOptions{},
	}


	reportScript.FilterOptions.FromDate = report.FilterOptions.FromDate
	reportScript.FilterOptions.ToDate = report.FilterOptions.ToDate
	reportScript.FilterOptions.ItemIds = report.FilterOptions.ItemIds



	for index, widget := range report.Widgets {

		log.Println(widget.Type)

		if widget.Type == "table" {
			tableWidget := &types.Widget{
				WidgetType: "table",
				WidgetName: `widget` + strconv.Itoa(index),
				Request:    types.Request{},
				Width:      widget.Width,
				Height:     widget.Height,
			}

			query := scriptrepo.GetLDALScripts(widget.ScriptId)

			decodedContent, err := base64.StdEncoding.DecodeString(query.Content)
			if err != nil {
				log.Println("decode error:", err)
			}
			tableWidget.Request.Query = string(decodedContent)

			//Get Tree
			tree := cusjsonrepo.GetCustomJson(query.BoundedId)
			// tableWidget.Request.Tree = tree.Content
			tableWidget.Request.Type = tree.JsonType

			reportScript.Widgets = append(reportScript.Widgets, *tableWidget)

		}
		if widget.Type == "gbar" {
			gbarWidget := &types.Widget{
				WidgetType:  "gbar",
				WidgetName:  `widget` + strconv.Itoa(index),
				Request:     types.Request{},
				ColorScheme: widget.BarCharts.Colors,
				XAxis:       widget.BarCharts.XAxis,
				YAxis:       widget.BarCharts.YAxis,
				Width:       widget.Width,
				Height:      widget.Height,
			}

			query := scriptrepo.GetLDALScripts(widget.ScriptId)

			decodedContent, err := base64.StdEncoding.DecodeString(query.Content)
			if err != nil {
				log.Println("decode error:", err)
			}
			gbarWidget.Request.Query = string(decodedContent)

			//Get Tree
			tree := cusjsonrepo.GetCustomJson(query.BoundedId)
			// gbarWidget.Request.Tree = tree.Content
			gbarWidget.Request.Type = tree.JsonType

			reportScript.Widgets = append(reportScript.Widgets, *gbarWidget)

		}

		if widget.Type == "bar" {

			barWidget := &types.Widget{
				WidgetType:  "bar",
				WidgetName:  `widget` + strconv.Itoa(index),
				Request:     types.Request{},
				ColorScheme: widget.BarCharts.Colors,
				XAxis:       widget.BarCharts.XAxis,
				YAxis:       widget.BarCharts.YAxis,
				Width:       widget.Width,
				Height:      widget.Height,
			}

			query := scriptrepo.GetLDALScripts(widget.ScriptId)

			decodedContent, err := base64.StdEncoding.DecodeString(query.Content)
			if err != nil {
				log.Println("decode error:", err)
			}
			barWidget.Request.Query = string(decodedContent)

			//Get Tree
			tree := cusjsonrepo.GetCustomJson(query.BoundedId)
			// barWidget.Request.Tree = tree.Content
			barWidget.Request.Type = tree.JsonType

			reportScript.Widgets = append(reportScript.Widgets, *barWidget)

		}

		if widget.Type == "bubble" {

			bubbleWidget := &types.Widget{
				WidgetType: "bubble",
				WidgetName: `widget` + strconv.Itoa(index),
				Request:    types.Request{},
				Width:      widget.Width,
				Height:     widget.Height,
			}

			query := scriptrepo.GetLDALScripts(widget.ScriptId)

			decodedContent, err := base64.StdEncoding.DecodeString(query.Content)
			if err != nil {
				log.Println("decode error:", err)
			}
			bubbleWidget.Request.Query = string(decodedContent)

			//Get Tree
			tree := cusjsonrepo.GetCustomJson(query.BoundedId)
			// bubbleWidget.Request.Tree = tree.Content
			bubbleWidget.Request.Type = tree.JsonType
			reportScript.Widgets = append(reportScript.Widgets, *bubbleWidget)

		}

		if widget.Type == "text" {

			textWidget := &types.Widget{
				WidgetType: "text",
				WidgetName: `widget` + strconv.Itoa(index),
				Request:    types.Request{},
				Value:      widget.Text,
			}

			reportScript.Widgets = append(reportScript.Widgets, *textWidget)

		}

		if widget.Type == "image" {

			imageWidget := &types.Widget{
				WidgetType: "image",
				Url:        widget.Url,
				WidgetName: `widget` + strconv.Itoa(index),
				Request:    types.Request{},
				Value:      widget.Text,
			}

			reportScript.Widgets = append(reportScript.Widgets, *imageWidget)

		}

	}
	return *reportScript
}

/**
* Generate JS script  tag which containt the logic for auth function
**/
func GetHTTPHeaderFunction(report datamodels.Report) string {
	result := `<script>`
	headers := report.Headers

	for _, header := range headers {
		if header.Name == "Authorization" {

			result = result + `let authorization = ` + header.JsFunction
		}
	}

	result = result + `
   </script>`
	return result

}

/**
* Generate JS script  tag which containt the logic for post parameters
**/
func GetHTTPostParams(report datamodels.Report) string {
	result := `<script>`
	params := report.POSTParams

	json := "let postRequest = {\n"
	for _, param := range params {
		json = json + param.Name + " : " + param.JsFunction +  ",\n"

	}
	json = json + "\n}"
	result = result + json + `
   </script>`
	return result

}

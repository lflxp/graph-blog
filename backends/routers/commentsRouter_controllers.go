package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"],
        beego.ControllerComments{
            Method: "GetDataByTablename",
            Router: `/:tablename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:tablename`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:tablename/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/kv`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"],
        beego.ControllerComments{
            Method: "AddTables",
            Router: `/tables`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:DbController"],
        beego.ControllerComments{
            Method: "GetAllTables",
            Router: `/tables`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:Neo4jController"] = append(beego.GlobalControllerRouter["github.com/lflxp/graph-blog/backends/controllers:Neo4jController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/search`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

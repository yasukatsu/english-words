webpackJsonp([1],{"5JDx":function(t,e){},I9Lo:function(t,e){},NHnr:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=n("7+uW"),r={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{attrs:{id:"app"}},[e("img",{staticClass:"header",attrs:{src:n("RVes")}}),this._v(" "),e("router-view")],1)},staticRenderFns:[]};var i=n("VU/8")({name:"App"},r,!1,function(t){n("I9Lo")},null,null).exports,a=n("/ocq"),o=n("mtWM"),c=n.n(o),u={name:"List",data:function(){return{words:[]}},created:function(){var t=this;c.a.get("http://0.0.0.0:8000/list").then(function(e){t.words=e.data}),console.log(this.words)}},_={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"list"},[n("table",[t._m(0),t._v(" "),n("tbody",t._l(t.words,function(e){return n("tr",{key:e},[n("td",[t._v(t._s(e.Spel))]),t._v(" "),n("td",[t._v(t._s(e.Define))]),t._v(" "),n("td",[t._v(t._s(e.Pos))]),t._v(" "),n("td",[t._v(t._s(e.Meaning))])])}),0)])])},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("thead",[e("tr",[e("th",[this._v("word")]),this._v(" "),e("th",[this._v("describe")]),this._v(" "),e("th",[this._v("pos")]),this._v(" "),e("th",[this._v("japanese")])])])}]};var v=n("VU/8")(u,_,!1,function(t){n("5JDx")},"data-v-62fcc6c4",null).exports;s.a.use(a.a);var p=new a.a({routes:[{path:"/",name:"List",component:v}]});s.a.config.productionTip=!1,new s.a({el:"#app",router:p,components:{App:i},template:"<App/>"})},RVes:function(t,e,n){t.exports=n.p+"static/img/ngsl.9223274.png"}},["NHnr"]);
//# sourceMappingURL=app.7d3362646b43a5c99a23.js.map
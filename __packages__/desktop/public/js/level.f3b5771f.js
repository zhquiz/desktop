(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["level"],{"51c0":function(e,t,n){},"73e0":function(e,t,n){"use strict";n.r(t);var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("section",[e.isLoading?n("b-loading",{attrs:{active:""}}):e._e(),Object.keys(e.allData).length>0?n("div",{staticClass:"LevelPage container"},[n("div",{staticClass:"field"},[n("label",{staticClass:"label"},[e._v("Filter")]),n("b-field",[n("b-radio-button",{attrs:{"native-value":"all",type:"is-success"},model:{value:e.whatToShow,callback:function(t){e.whatToShow=t},expression:"whatToShow"}},[e._v(" Show all ")]),n("b-radio-button",{attrs:{"native-value":"all-quiz",type:"is-info"},model:{value:e.whatToShow,callback:function(t){e.whatToShow=t},expression:"whatToShow"}},[e._v(" All quiz ")]),n("b-radio-button",{attrs:{"native-value":"learning",type:"is-warning"},model:{value:e.whatToShow,callback:function(t){e.whatToShow=t},expression:"whatToShow"}},[e._v(" Learning ")])],1)],1),n("b-table",{attrs:{data:e.currentData}},[n("b-table-column",{attrs:{field:"level",label:"Level",width:"40"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("span",{staticClass:"clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){e.selected=e.allData[t.row.level],e.$refs.context.open(n)}(n)}}},[e._v(" "+e._s(t.row.level)+" ")])]}}],null,!1,653076080)}),n("b-table-column",{attrs:{field:"entries",label:"Item"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("div",e._l(t.row.entries,(function(t){return n("span",{key:t,staticClass:"tag clickable",class:e.getTagClass(t),on:{contextmenu:function(n){return n.preventDefault(),function(n){e.selected=[t],e.$refs.context.open(n)}(n)}}},[e._v(" "+e._s(t)+" ")])})),0)]}}],null,!1,3147923365)})],1)],1):e._e(),n("ContextMenu",{ref:"context",attrs:{type:"vocab",entry:e.selected,pinyin:e.pinyinMap},on:{"quiz:added":function(t){return e.reload(t.entries)},"quiz:removed":function(t){return e.reload(t.entries)}}})],1)},r=[],i=(n("4de4"),n("a630"),n("d81d"),n("4fad"),n("d3b7"),n("25f0"),n("3ca3"),n("ddb0"),n("96cf"),n("1da1")),o=n("b85c"),s=n("3835"),l=n("d4ec"),u=n("bee2"),c=n("262e"),f=n("2caf"),h=n("9ab4"),v=n("1b40"),d=n("6825"),p=n("02ef"),b=n("5962"),w=n.n(b),y=function(e){Object(c["a"])(n,e);var t=Object(f["a"])(n);function n(){var e;return Object(l["a"])(this,n),e=t.apply(this,arguments),e.isLoading=!0,e.allData={},e.srsLevel={},e.selected=[],e.tagClassMap=[function(e){return e>2?"is-success":""},function(e){return e>0?"is-warning":""},function(e){return 0===e?"is-danger":""}],e.whatToShow="all",e.currentData=[],e.pinyinMap={},e}return Object(u["a"])(n,[{key:"setCurrentData",value:function(){var e=this;this.currentData=Object.entries(this.allData).map((function(t){var n=Object(s["a"])(t,2),a=n[0],r=n[1],i=parseInt(a);return{level:i,entries:Array.from(r).filter((function(t){return"all"===e.whatToShow||("learning"===e.whatToShow&&e.srsLevel[t]<=2||"all-quiz"===e.whatToShow&&"undefined"!==typeof e.srsLevel[t])})).sort()}})).filter((function(e){return e.entries.length>0})).sort((function(e,t){return e.level-t.level}))}},{key:"getTagClass",value:function(e){var t=this.srsLevel[e];if("undefined"!==typeof t){if(-1===t)return"is-info";var n,a=Object(o["a"])(this.tagClassMap);try{for(a.s();!(n=a.n()).done;){var r=n.value,i=r(t);if(i)return i}}catch(s){a.e(s)}finally{a.f()}}return"is-light"}},{key:"init",value:function(){var e=Object(i["a"])(regeneratorRuntime.mark((function e(){var t,n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,p["a"].get("/api/user",{params:{select:["settings.level.whatToShow"]}});case 2:return t=e.sent,n=t.data["settings.level.whatToShow"],a=void 0===n?null:n,a&&(this.whatToShow=a),e.next=8,this.reload([]);case 8:this.isLoading=!1;case 9:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"reload",value:function(){var e=Object(i["a"])(regeneratorRuntime.mark((function e(t){var n,a,r,i,o,s=this;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(0!==this.currentData.length){e.next=8;break}return e.next=3,p["a"].get("/api/vocab/level");case 3:n=e.sent,a=n.data.result,t=a.map((function(e){var t=e.entry,n=e.source,a=e.level;n&&"cedict"!==n&&(s.pinyinMap[t]=w()(t,{keepRest:!0,toneToNumber:!0}));var r=a.toString(),i=s.allData[r]||[];return i.push(t),s.allData[r]=i,t})),this.$set(this,"pinyinMap",this.pinyinMap),this.$set(this,"allData",this.allData);case 8:if(!(t.length>0)){e.next=18;break}return e.next=11,p["a"].post("/api/quiz/srsLevel",{entries:t,type:"vocab",select:["entry","srsLevel"]});case 11:r=e.sent,i=r.data.result,o=void 0===i?[]:i,t.map((function(e){delete s.srsLevel[e]})),o.map((function(e){var t=e.entry,n=e.srsLevel;s.srsLevel[t]="number"===typeof n?n:-1})),this.$set(this,"srsLevel",this.srsLevel),this.setCurrentData();case 18:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"onWhatToShowChanged",value:function(){var e=Object(i["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.setCurrentData(),e.next=3,p["a"].patch("/api/user",{"settings.level.whatToShow":this.whatToShow});case 3:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()}]),n}(v["d"]);Object(h["a"])([Object(v["c"])()],y.prototype,"context",void 0),y=Object(h["a"])([Object(v["a"])({components:{ContextMenu:d["a"]},created:function(){var e=window,t=e.frameElement;if(t){var n=parseInt(t.getAttribute("data-id")||"");window.parent.setName(n,"Level")}this.init()},watch:{whatToShow:function(){this.onWhatToShowChanged()}}})],y);var g=y,m=g,S=(n("8f1b"),n("2877")),T=Object(S["a"])(m,a,r,!1,null,"20fe6820",null);t["default"]=T.exports},"8f1b":function(e,t,n){"use strict";n("51c0")},b85c:function(e,t,n){"use strict";n.d(t,"a",(function(){return r}));n("a4d3"),n("e01a"),n("d28b"),n("d3b7"),n("3ca3"),n("ddb0");var a=n("06c5");function r(e,t){var n;if("undefined"===typeof Symbol||null==e[Symbol.iterator]){if(Array.isArray(e)||(n=Object(a["a"])(e))||t&&e&&"number"===typeof e.length){n&&(e=n);var r=0,i=function(){};return{s:i,n:function(){return r>=e.length?{done:!0}:{done:!1,value:e[r++]}},e:function(e){throw e},f:i}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var o,s=!0,l=!1;return{s:function(){n=e[Symbol.iterator]()},n:function(){var e=n.next();return s=e.done,e},e:function(e){l=!0,o=e},f:function(){try{s||null==n["return"]||n["return"]()}finally{if(l)throw o}}}}}}]);
//# sourceMappingURL=level.f3b5771f.js.map
(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["hanzi"],{"3c0e":function(t,e,n){},"4c53":function(t,e,n){"use strict";var a=n("23e7"),r=n("857a"),s=n("af03");a({target:"String",proto:!0,forced:s("sub")},{sub:function(){return r(this,"sub","","")}})},"664f":function(t,e,n){"use strict";var a=n("23e7"),r=n("857a"),s=n("af03");a({target:"String",proto:!0,forced:s("sup")},{sup:function(){return r(this,"sup","","")}})},"857a":function(t,e,n){var a=n("1d80"),r=/"/g;t.exports=function(t,e,n,s){var u=String(a(t)),i="<"+e;return""!==n&&(i+=" "+n+'="'+String(s).replace(r,"&quot;")+'"'),i+">"+u+"</"+e+">"}},af03:function(t,e,n){var a=n("d039");t.exports=function(t){return a((function(){var e=""[t]('"');return e!==e.toLowerCase()||e.split('"').length>3}))}},f0aa:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("section",[n("div",{staticClass:"HanziPage"},[n("form",{staticClass:"field",on:{submit:function(e){e.preventDefault(),t.q=t.q0}}},[n("div",{staticClass:"control"},[n("input",{directives:[{name:"model",rawName:"v-model",value:t.q0,expression:"q0"}],staticClass:"input",attrs:{type:"search",name:"q",placeholder:"Type here to search.","aria-label":"search"},domProps:{value:t.q0},on:{input:function(e){e.target.composing||(t.q0=e.target.value)}}})])]),n("div",{staticClass:"columns"},[n("div",{staticClass:"column is-6 entry-display"},[n("div",{staticClass:"hanzi-display clickable font-han",on:{contextmenu:function(e){return e.preventDefault(),function(e){return t.openContext(e,t.current,"hanzi")}(e)}}},[t._v(" "+t._s(t.current)+" ")]),n("div",{staticClass:"buttons has-addons"},[n("button",{staticClass:"button",attrs:{disabled:t.i<1},on:{click:function(e){t.i--},keypress:function(e){t.i--}}},[t._v(" Previous ")]),n("button",{staticClass:"button",attrs:{disabled:t.i>t.entries.length-2},on:{click:function(e){t.i++},keypress:function(e){t.i++}}},[t._v(" Next ")])])]),n("div",{staticClass:"column is-6"},[n("b-collapse",{staticClass:"card",attrs:{animation:"slide",open:!!t.sub.length},scopedSlots:t._u([{key:"trigger",fn:function(e){return n("div",{staticClass:"card-header",attrs:{role:"button"}},[n("h2",{staticClass:"card-header-title"},[t._v("Subcompositions")]),n("a",{staticClass:"card-header-icon",attrs:{role:"button"}},[n("fontawesome",{attrs:{icon:e.open?"caret-down":"caret-up"}})],1)])}}])},[n("div",{staticClass:"card-content"},t._l(t.sub,(function(e){return n("span",{key:e,staticClass:"font-han clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){return t.openContext(n,e,"hanzi")}(n)}}},[t._v(" "+t._s(e)+" ")])})),0)]),n("b-collapse",{staticClass:"card",attrs:{animation:"slide",open:!!t.sup.length},scopedSlots:t._u([{key:"trigger",fn:function(e){return n("div",{staticClass:"card-header",attrs:{role:"button"}},[n("h2",{staticClass:"card-header-title"},[t._v("Supercompositions")]),n("a",{staticClass:"card-header-icon",attrs:{role:"button"}},[n("fontawesome",{attrs:{icon:e.open?"caret-down":"caret-up"}})],1)])}}])},[n("div",{staticClass:"card-content"},t._l(t.sup,(function(e){return n("span",{key:e,staticClass:"font-han clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){return t.openContext(n,e,"hanzi")}(n)}}},[t._v(" "+t._s(e)+" ")])})),0)]),n("b-collapse",{staticClass:"card",attrs:{animation:"slide",open:!!t.variants.length},scopedSlots:t._u([{key:"trigger",fn:function(e){return n("div",{staticClass:"card-header",attrs:{role:"button"}},[n("h2",{staticClass:"card-header-title"},[t._v("Variants")]),n("a",{staticClass:"card-header-icon",attrs:{role:"button"}},[n("fontawesome",{attrs:{icon:e.open?"caret-down":"caret-up"}})],1)])}}])},[n("div",{staticClass:"card-content"},t._l(t.variants,(function(e){return n("span",{key:e,staticClass:"font-han clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){return t.openContext(n,e,"hanzi")}(n)}}},[t._v(" "+t._s(e)+" ")])})),0)]),n("b-collapse",{staticClass:"card",attrs:{animation:"slide",open:!!t.vocabs.length},scopedSlots:t._u([{key:"trigger",fn:function(e){return n("div",{staticClass:"card-header",attrs:{role:"button"}},[n("h2",{staticClass:"card-header-title"},[t._v("Vocabularies")]),n("a",{staticClass:"card-header-icon",attrs:{role:"button"}},[n("fontawesome",{attrs:{icon:e.open?"caret-down":"caret-up"}})],1)])}}])},[n("div",{staticClass:"card-content"},t._l(t.vocabs,(function(e,a){return n("div",{key:a,staticClass:"long-item"},[n("span",{staticClass:"clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){return t.openContext(n,e.simplified,"vocab")}(n)}}},[t._v(" "+t._s(e.simplified)+" ")]),e.traditional?n("span",{staticClass:"clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){return t.openContext(n,e.traditional,"vocab")}(n)}}},[t._v(" "+t._s(e.traditional)+" ")]):t._e(),n("span",{staticClass:"pinyin"},[t._v("["+t._s(e.pinyin)+"]")]),n("span",[t._v(t._s(e.english))])])})),0)]),n("b-collapse",{key:t.sentenceKey,staticClass:"card",attrs:{animation:"slide",open:!!t.sentences().length},scopedSlots:t._u([{key:"trigger",fn:function(e){return n("div",{staticClass:"card-header",attrs:{role:"button"}},[n("h2",{staticClass:"card-header-title"},[t._v("Sentences")]),n("a",{staticClass:"card-header-icon",attrs:{role:"button"}},[n("fontawesome",{attrs:{icon:e.open?"caret-down":"caret-up"}})],1)])}}])},[n("div",{staticClass:"card-content"},t._l(t.sentences(),(function(e,a){return n("div",{key:a,staticClass:"long-item"},[n("span",{staticClass:"clickable",on:{contextmenu:function(n){return n.preventDefault(),function(n){return t.openContext(n,e.chinese,"sentence")}(n)}}},[t._v(" "+t._s(e.chinese)+" ")]),n("span",[t._v(t._s(e.english))])])})),0)])],1)])]),n("ContextMenu",{ref:"context",attrs:{entry:t.selected.entry,type:t.selected.type,additional:t.additionalContext,pinyin:t.sentenceDef.pinyin,english:t.sentenceDef.english}})],1)},r=[],s=(n("4de4"),n("7db0"),n("c975"),n("d81d"),n("fb6a"),n("d3b7"),n("ac1f"),n("1276"),n("4c53"),n("664f"),n("ddb0"),n("ade3")),u=n("2909"),i=(n("96cf"),n("1da1")),c=n("d4ec"),o=n("bee2"),l=n("262e"),d=n("2caf"),p=n("9ab4"),f=n("1b40"),h=n("6825"),D=n("02ef"),v=n("d8bb"),b=function(t){Object(l["a"])(n,t);var e=Object(d["a"])(n);function n(){var t;return Object(c["a"])(this,n),t=e.apply(this,arguments),t.entries=[],t.i=0,t.sub=[],t.sup=[],t.variants=[],t.vocabs=[],t.selected={entry:"",type:""},t.q0="",t.sentenceKey=0,t}return Object(o["a"])(n,[{key:"sentences",value:function(){return v["c"].find({chinese:{$containsString:this.current}}).slice(0,10)}},{key:"openContext",value:function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:this.selected.entry,n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:this.selected.type;this.selected={entry:e,type:n},this.context.open(t)}},{key:"onQChange",value:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(e){var n,a,r,s,u;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(n=window,a=n.frameElement,a&&(r=parseInt(a.getAttribute("data-id")||""),window.parent.setName(r,(e?e+" - ":"")+"Hanzi")),!/(?:[\u2E80-\u2E99\u2E9B-\u2EF3\u2F00-\u2FD5\u3005\u3007\u3021-\u3029\u3038-\u303B\u3400-\u4DBF\u4E00-\u9FFC\uF900-\uFA6D\uFA70-\uFAD9]|\uD81B[\uDFF0\uDFF1]|[\uD840-\uD868\uD86A-\uD86C\uD86F-\uD872\uD874-\uD879\uD880-\uD883][\uDC00-\uDFFF]|\uD869[\uDC00-\uDEDD\uDF00-\uDFFF]|\uD86D[\uDC00-\uDF34\uDF40-\uDFFF]|\uD86E[\uDC00-\uDC1D\uDC20-\uDFFF]|\uD873[\uDC00-\uDEA1\uDEB0-\uDFFF]|\uD87A[\uDC00-\uDFE0]|\uD87E[\uDC00-\uDE1D]|\uD884[\uDC00-\uDF4A])/.test(e)){t.next=7;break}s=e.split("").filter((function(t){return/(?:[\u2E80-\u2E99\u2E9B-\u2EF3\u2F00-\u2FD5\u3005\u3007\u3021-\u3029\u3038-\u303B\u3400-\u4DBF\u4E00-\u9FFC\uF900-\uFA6D\uFA70-\uFAD9]|\uD81B[\uDFF0\uDFF1]|[\uD840-\uD868\uD86A-\uD86C\uD86F-\uD872\uD874-\uD879\uD880-\uD883][\uDC00-\uDFFF]|\uD869[\uDC00-\uDEDD\uDF00-\uDFFF]|\uD86D[\uDC00-\uDF34\uDF40-\uDFFF]|\uD86E[\uDC00-\uDC1D\uDC20-\uDFFF]|\uD873[\uDC00-\uDEA1\uDEB0-\uDFFF]|\uD87A[\uDC00-\uDFE0]|\uD87E[\uDC00-\uDE1D]|\uD884[\uDC00-\uDF4A])/.test(t)})),this.entries=s.filter((function(t,e){return s.indexOf(t)===e})),t.next=11;break;case 7:return t.next=9,D["a"].get("/api/hanzi/q",{params:{q:e}});case 9:u=t.sent,this.entries=u.data.result.map((function(t){var e=t.entry;return e}));case 11:this.i=0,this.load();case 13:case"end":return t.stop()}}),t,this)})));function e(e){return t.apply(this,arguments)}return e}()},{key:"load",value:function(){this.current?(this.loadHanzi(),this.loadVocab()):(this.sub=[],this.sup=[],this.variants=[],this.vocabs=[])}},{key:"loadHanzi",value:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,D["a"].get("/api/hanzi",{params:{entry:this.current}}).then((function(t){return t.data}));case 2:e=t.sent,this.sub=Object(u["a"])(e.sub),this.sup=Object(u["a"])(e.sup),this.variants=Object(u["a"])(e.variants);case 6:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()},{key:"loadVocab",value:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(){var e,n;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,D["a"].get("/api/vocab/q",{params:{q:this.current}});case 2:e=t.sent,n=e.data.result,this.vocabs=n;case 5:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()},{key:"loadSentences",value:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(v["a"])(this.current,10);case 2:if(!t.sent){t.next=4;break}this.sentenceKey=Math.random();case 4:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()},{key:"sentenceDef",get:function(){if("sentence"!==this.selected.type)return{};var t=v["c"].findOne({chinese:this.selected.entry});return t?{pinyin:Object(s["a"])({},this.selected.entry,t.pinyin),english:Object(s["a"])({},this.selected.entry,t.english)}:{}}},{key:"q",get:function(){var t=this.$route.query.q;return(Array.isArray(t)?t[0]:t)||""},set:function(t){this.$router.push({query:{q:t}})}},{key:"current",get:function(){return this.entries[this.i]}},{key:"additionalContext",get:function(){var t=this;return this.q?[]:[{name:"Reload",handler:function(){var e=Object(i["a"])(regeneratorRuntime.mark((function e(){var n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,D["a"].get("/api/hanzi/random");case 2:n=e.sent,a=n.data.result,t.q0=a;case 5:case"end":return e.stop()}}),e)})));function n(){return e.apply(this,arguments)}return n}()}]}}]),n}(f["d"]);Object(p["a"])([Object(f["c"])()],b.prototype,"context",void 0),Object(p["a"])([Object(f["e"])("q")],b.prototype,"onQChange",null),Object(p["a"])([Object(f["e"])("current")],b.prototype,"load",null),b=Object(p["a"])([Object(f["a"])({components:{ContextMenu:h["a"]},created:function(){var t=this;return Object(i["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(t.q0=t.q,!t.additionalContext[0]){e.next=4;break}return e.next=4,t.additionalContext[0].handler();case 4:t.onQChange(t.q0);case 5:case"end":return e.stop()}}),e)})))()}})],b);var C=b,F=C,y=(n("f589"),n("2877")),m=Object(y["a"])(F,a,r,!1,null,"10d59776",null);e["default"]=m.exports},f589:function(t,e,n){"use strict";n("3c0e")}}]);
//# sourceMappingURL=hanzi.a221c840.js.map
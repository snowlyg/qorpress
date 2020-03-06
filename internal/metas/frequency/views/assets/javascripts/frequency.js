"use strict";var _typeof="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e};!function(e){"function"==typeof define&&define.amd?define(["jquery"],e):"object"===("undefined"==typeof exports?"undefined":_typeof(exports))?e(require("jquery")):e(jQuery)}(function(o){var r="qor.metas.daterange",e="click."+r,t="change."+r,n="enable."+r,i=".qor-frequency__selector",s=".qor-frequency__week li";function f(e,t){this.$element=o(e),this.options=o.extend({},f.DEFAULTS,o.isPlainObject(t)&&t),this.init()}return f.prototype={constructor:f,init:function(){var e=this.$element;this.$selector=e.find(i),this.$month=e.find(".qor-frequency__monthly"),this.$week=e.find(".qor-frequency__weekly"),this.bind(),this.initData()},bind:function(){this.$element.on(t,i,this.change.bind(this)).on(e,s,this.selectDayOfWeek.bind(this))},unbind:function(){this.$element.off(t).off(e)},initData:function(){this.change()},change:function(){var e=this.$selector.val();this.$month.hide(),this.$week.hide(),"monthly"===e?this.$month.show():"weekly"===e&&this.$week.show()},selectDayOfWeek:function(e){var t=o(e.target),n=t.attr("value");this.$element.find(s).removeClass("selected"),t.addClass("selected"),this.$element.find(".qor-frequency__weekly--input").val(n)},destroy:function(){this.unbind(),this.$element.removeData(r)}},f.plugin=function(i){return this.each(function(){var e,t=o(this),n=t.data(r);if(!n){if(/destroy/.test(i))return;t.data(r,n=new f(this,i))}"string"==typeof i&&o.isFunction(e=n[i])&&e.apply(n)})},o(function(){var t='[data-toggle="qor.metas.frequency"]';o(document).on("disable.qor.metas.daterange",function(e){f.plugin.call(o(t,e.target),"destroy")}).on(n,function(e){f.plugin.call(o(t,e.target))}).triggerHandler(n)}),f});
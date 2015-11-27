
'use strict'
var app = angular.module('testApp', ['textAngular']);
console.log(app);

app.config(['$interpolateProvider', function ($interpolateProvider) {
    $interpolateProvider.startSymbol('[[');
    $interpolateProvider.endSymbol(']]');
}]);
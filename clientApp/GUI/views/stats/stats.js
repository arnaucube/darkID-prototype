'use strict';

angular.module('app.stats', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/stats', {
            templateUrl: 'views/stats/stats.html',
            controller: 'StatsCtrl'
        });
    }])

    .controller('StatsCtrl', function($scope, $rootScope, $http, $filter) {

        $rootScope.server = JSON.parse(localStorage.getItem("darkID_server"));

        $scope.generatingID = false;
        $scope.ids = [];
        $http.get(clientapi + 'ids')
            .then(function(data) {
                console.log('data success');
                console.log(data);
                $scope.ids = data.data;
                $scope.idsToChart();
            }, function(data) {
                console.log('data error');
            });

        $scope.newID = function() {
            $scope.generatingID = true;
            $http.get(clientapi + 'newid')
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;
                    $scope.generatingID = false;

                }, function(data) {
                    console.log('data error');
                });
        };

        $scope.blindAndSendToSign = function(id) {
            $http.get(clientapi + 'blindandsendtosign/' + id)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };
        $scope.verify = function(id) {
            $http.get(clientapi + 'verify/' + id)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };
        $scope.clientApp = function(route, param) {
            $http.get(clientapi + route + '/' + param)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };

        //chartjs
        $scope.chart1 = {
            colours: ['#4DD0E1', '#9575CD', '#F06292', '#FFF176'],
            labels: [],
            data: []
        };
        $scope.chart2 = {
            colours: ['#4DD0E1', '#9575CD', '#F06292', '#FFF176'],
            labels: [],
            data: []
        };
        $scope.idsToChart = function() {
            //chart1
            var dictionary = {};
            var ids = $scope.ids;
            for(var i=0; i<ids.length; i++) {
                var day = $filter('date')(ids[i].date, 'dd.MM.y, HH:mm');
                if(dictionary[day]==undefined) {
                    dictionary[day] = 1
                } else {
                    dictionary[day]++;
                }
            }
            console.log(dictionary);
            for(var key in dictionary) {
                $scope.chart1.labels.push(key);
                $scope.chart1.data.push(dictionary[key]);
            }


            //chart2
            var dictionary = {};
            for(var i=0; i<ids.length; i++) {
                if(ids[i].blockchainref) {
                    if(dictionary['in blockchain']==undefined) {
                        dictionary['in blockchain'] = 1
                    } else {
                        dictionary['in blockchain']++;
                    }
                } else if(ids[i].unblindedsig) {
                    if(dictionary['signed']==undefined) {
                        dictionary['signed'] = 1
                    } else {
                        dictionary['signed']++;
                    }
                } else {
                    if(dictionary['unsigned']==undefined) {
                        dictionary['unsigned'] = 1
                    } else {
                        dictionary['unsigned']++;
                    }
                }
            }
            console.log(dictionary);
            for(var key in dictionary) {
                $scope.chart2.labels.push(key);
                $scope.chart2.data.push(dictionary[key]);
            }

        };
    });

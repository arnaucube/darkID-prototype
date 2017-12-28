'use strict';

angular.module('app.main', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/main', {
            templateUrl: 'views/main/main.html',
            controller: 'MainCtrl'
        });
    }])

    .controller('MainCtrl', function($scope, $rootScope, $http) {

        $rootScope.server = JSON.parse(localStorage.getItem("darkID_server"));

        $scope.ids = [];
        $http.get(clientapi + 'ids')
            .then(function(data) {
                console.log('data success');
                console.log(data);
                $scope.ids = data.data;

            }, function(data) {
                console.log('data error');
            });

        $scope.newID = function() {
            $http.get(clientapi + 'newid')
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };

        $scope.blindAndSendToSign = function(pubK) {
            $http.get(clientapi + 'blindandsendtosign/' + pubK)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };
        $scope.verify = function(pubK) {
            $http.get(clientapi + 'verify/' + pubK)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };
    });

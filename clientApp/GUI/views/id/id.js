'use strict';

angular.module('app.id', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/id/:keyid', {
            templateUrl: 'views/id/id.html',
            controller: 'IdCtrl'
        });
    }])

    .controller('IdCtrl', function($scope, $rootScope, $http, $routeParams) {
        $scope.keyid = $routeParams.keyid;
        $scope.decryptData = {
            m:"",
            c:""
        };
        $scope.encryptData = {
            m:"",
            c:""
        };

        $rootScope.server = JSON.parse(localStorage.getItem("darkID_server"));

        $scope.id = {};
        $scope.clientApp = function(route, param) {
            $http.get(clientapi + route + '/' + param)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.id = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };
        $scope.clientApp('id', $routeParams.keyid);


        $scope.decrypt = function() {
            $http({
                    url: clientapi + 'decrypt/' + $routeParams.keyid,
                    method: "POST",
                    headers: {
                        "Content-Type": undefined
                    },
                    data: $scope.decryptData
                })
                .then(function(data) {
                        console.log("data: ");
                        console.log(data.data);
                        $scope.decryptData = data.data;
                    },
                    function(data) {
                        console.log(data);
                    });
        };
        $scope.encrypt = function() {
            $http({
                    url: clientapi + 'encrypt/' + $routeParams.keyid,
                    method: "POST",
                    headers: {
                        "Content-Type": undefined
                    },
                    data: $scope.encryptData
                })
                .then(function(data) {
                        console.log("data: ");
                        console.log(data.data);
                        $scope.encryptData = data.data;
                    },
                    function(data) {
                        console.log(data);
                    });
        };
    });

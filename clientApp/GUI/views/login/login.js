'use strict';

angular.module('app.login', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/login', {
            templateUrl: 'views/login/login.html',
            controller: 'LoginCtrl'
        });
    }])

    .controller('LoginCtrl', function($scope, $rootScope, $http, $routeParams, toastr) {
        $rootScope.server = ""
        $scope.user = {};
        //set server in goclient
        $http.get(clientapi + 'getserver')
            .then(function(data) {
                console.log("data: ");
                console.log(data.data);
                $rootScope.server = data.data;
                localStorage.setItem("darkID_server", JSON.stringify($rootScope.server));
                console.log("server", $rootScope.server);
            }, function(data) {
                console.log('data error');
            });

        $scope.login = function() {

            console.log('Doing login', $scope.user);
            console.log($rootScope.server + "login");



            $http({
                    url: $rootScope.server + 'login',
                    method: "POST",
                    headers: {
                        "Content-Type": undefined
                    },
                    data: $scope.user
                })
                .then(function(data) {
                        console.log("data: ");
                        console.log(data.data);
                        if (data.data.token) {
                            localStorage.setItem("darkID_token", data.data.token);
                            localStorage.setItem("darkID_user", JSON.stringify(data.data));
                            window.location.reload();
                        } else {
                            console.log("login failed");
                            toastr.error('Login failed, ' + data.data);
                        }


                    },
                    function(data) {
                        console.log(data);
                    });

        };
    });

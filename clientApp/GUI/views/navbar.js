'use strict';

angular.module('app.navbar', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/navbar', {
            templateUrl: 'views/navbar.html',
            controller: 'NavbarCtrl'
        });
    }])

    .controller('NavbarCtrl', function($scope, $rootScope, $http, $routeParams, $location) {
        $rootScope.server = JSON.parse(localStorage.getItem("darkID_server"));

        $scope.user = JSON.parse(localStorage.getItem("darkID_user"));

        $scope.logout = function() {
            localStorage.removeItem("darkID_token");
            localStorage.removeItem("darkID_user");
            localStorage.removeItem("darkID_server");
            window.location.reload();
        };

    });

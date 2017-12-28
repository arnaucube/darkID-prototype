'use strict';

angular.module('app.signup', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/signup', {
    templateUrl: 'views/signup/signup.html',
    controller: 'SignupCtrl'
  });
}])

.controller('SignupCtrl', function($scope, $http, $routeParams) {
    $scope.user = {};
    $scope.doSignup = function() {
      console.log('Doing signup', $scope.user);


    };
});

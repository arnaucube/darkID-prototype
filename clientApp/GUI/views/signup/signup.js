'use strict';

angular.module('app.signup', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/signup', {
    templateUrl: 'views/signup/signup.html',
    controller: 'SignupCtrl'
  });
}])

.controller('SignupCtrl', function($scope, $http, $routeParams, $rootScope) {
    $scope.user = {};
    $scope.signup = function() {
      console.log('Doing signup', $scope.user);
      $http({
              url: $rootScope.server + 'signup',
              method: "POST",
              headers: {
                  "Content-Type": undefined
              },
              data: $scope.user
          })
          .then(function(data) {
                  console.log("data: ");
                  console.log(data.data);
                  window.location="/";

              },
              function(data) {
                  console.log(data);
              });

    };
});

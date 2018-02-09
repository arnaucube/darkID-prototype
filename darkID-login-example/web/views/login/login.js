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
        $scope.proof = {
            publicKey: "",
            clear: "",
            question: "",
            answer: ""
        };
        $scope.getproof = function() {
            $http({
                    url: urlapi + 'getproof',
                    method: "POST",
                    headers: {
                        "Content-Type": undefined
                    },
                    data: $scope.proof
                })
                .then(function(data) {
                        console.log("data: ");
                        console.log(data.data);
                        $scope.proof = data.data;
                    },
                    function(data) {
                        console.log(data);
                        toastr.error("error: bad darkID PublicKey")
                    });

        };
        $scope.sendanswer = function() {
            $http({
                    url: urlapi + 'answerproof',
                    method: "POST",
                    headers: {
                        "Content-Type": undefined
                    },
                    data: $scope.proof
                })
                .then(function(data) {
                        console.log("data: ");
                        console.log(data.data);
                        if(data.data=="fail\n") {
                            toastr.error("Proof of darkID failed");
                        }else{
                            toastr.success("You are logged with darkID!");
                            window.location="#!/main";
                        }
                    },
                    function(data) {
                        console.log(data);
                    });

        };
    });

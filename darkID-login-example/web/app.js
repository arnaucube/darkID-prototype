'use strict';


var urlapi = "http://127.0.0.1:5010/";

// Declare app level module which depends on views, and components
angular.module('app', [
    'ngRoute',
    'ngMessages',
    'angularBootstrapMaterial',
    'ui.bootstrap',
    'toastr',
    'app.main',
    'app.login'
]).
config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
        $locationProvider.hashPrefix('!');
        $routeProvider.otherwise({
            redirectTo: '/login'
        });
    }])
    .config(function(toastrConfig) {
        angular.extend(toastrConfig, {
            autoDismiss: false,
            containerId: 'toast-container',
            maxOpened: 0,
            newestOnTop: true,
            positionClass: 'toast-bottom-right',
            preventDuplicates: false,
            preventOpenDuplicates: false,
            target: 'body'
        });
    })
    .factory('httpInterceptor', function httpInterceptor() {
        return {
            request: function(config) {
                return config;
            },

            requestError: function(config) {
                return config;
            },

            response: function(res) {
                return res;
            },

            responseError: function(res) {
                return res;
            }
        };
    })
    .factory('api', function($http) {
        return {
            init: function() {
                /*$http.defaults.headers.common['X-Access-Token'] = localStorage.getItem('block_webapp_token');
                $http.defaults.headers.post['X-Access-Token'] = localStorage.getItem('block_webapp_token');*/
            }
        };
    })
    .run(function(api) {
        api.init();
    });

'use strict';


var clientapi = "http://127.0.0.1:4100/";

// Declare app level module which depends on views, and components
angular.module('app', [
    'ngRoute',
    'ngMessages',
    'angularBootstrapMaterial',
    'ui.bootstrap',
    'toastr',
    'chart.js',
    'app.navbar',
    'app.main',
    'app.signup',
    'app.login',
    'app.id',
    'app.stats'
]).
config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
        $locationProvider.hashPrefix('!');

        if ((localStorage.getItem('darkID_token'))) {
            console.log(window.location.hash);
            if ((window.location.hash === '#!/login') || (window.location.hash === '#!/signup')) {
                window.location = '#!/main';
            }

            $routeProvider.otherwise({
                redirectTo: '/main'
            });
        } else {
            if ((window.location !== '#!/login') || (window.location !== '#!/signup')) {
                console.log('app, user no logged');

                localStorage.removeItem('darkID_token');
                localStorage.removeItem('darkID_userdata');
                window.location = '#!/login';
                $routeProvider.otherwise({
                    redirectTo: '/login'
                });
            }
        }
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
                /*$http.defaults.headers.common['X-Access-Token'] = localStorage.getItem('darkID_token');
                $http.defaults.headers.post['X-Access-Token'] = localStorage.getItem('darkID_token');*/
            }
        };
    })
    .run(function(api) {
        api.init();
    });



var domain = 'http://localhost:3000/api'






var backend = domain;
angular.module('app.config', [])

.factory('Configuration', function() {
    return {
        backend: backend
    }
});
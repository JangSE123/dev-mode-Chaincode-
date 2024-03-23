'use strict';

var app = angular.module('application', []);

app.controller('AppCtrl', function($scope, appFactory) {
    // 사용자 등록 함수
    $scope.registerUser = function() {
        var user = {
            userID: $scope.user.userID,
            name: $scope.user.name,
            buildingName: $scope.user.buildingName,
            description: $scope.user.description
        };
        appFactory.registerUser(user, function(data) {
            $scope.registerUserResult = data;
        });
    };

// 부동산 조회 함수 수정
$scope.viewBuilding = function() {
    appFactory.viewBuilding($scope.buildingName, function(data) {
        // 응답 데이터를 JSON 객체로 파싱합니다.
        $scope.buildingInfoParsed = JSON.parse(data);
    });
};

});

app.factory('appFactory', function($http) {
    var factory = {};

    // 사용자 등록을 위한 HTTP POST 요청 함수
    factory.registerUser = function(user, callback) {
        $http.post('/registerUser', user).then(function(response) {
            callback(response.data);
        }, function(error) {
            callback({ error: error });
        });
    };

    // 부동산 조회를 위한 HTTP GET 요청 함수
    factory.viewBuilding = function(buildingName, callback) {
        $http.get('/viewBuilding?buildingName=' + buildingName).then(function(response) {
            callback(response.data);
        }, function(error) {
            callback({ error: error });
        });
    };

    return factory;
});